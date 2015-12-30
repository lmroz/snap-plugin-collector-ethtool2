/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package etplugin

import (
	"errors"
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type cMock struct {
	mock.Mock
}

func (cm *cMock) GetNicStats(iface string) (map[string]string, error) {
	args := cm.Mock.Called(iface)
	var r0 map[string]string
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string]string)
	}
	return r0, args.Error(1)
}

func (cm *cMock) GetRegDump(iface string) (map[string]string, error) {
	args := cm.Mock.Called(iface)
	var r0 map[string]string
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string]string)
	}
	return r0, args.Error(1)
}

func (cm *cMock) GetDomStats(iface string) (map[string]string, error) {
	args := cm.Mock.Called(iface)
	var r0 map[string]string
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string]string)
	}
	return r0, args.Error(1)
}
func (cm *cMock) GetDriverInfo(iface string) (string, error) {
	args := cm.Mock.Called(iface)
	return args.Get(0).(string), args.Error(1)
}

func TestValidMetrics(t *testing.T) {

	Convey("ValidMetrics", t, func() {
		orgNetInterfaces := netInterfaces
		Reset(func() {
			netInterfaces = orgNetInterfaces
		})

		collector := &cMock{}
		sut := &metricCollectorImpl{Ethtool: collector}

		Convey("listing interfaces is unsuccessful", func() {
			netInterfaces = func() ([]net.Interface, error) {
				return nil, errors.New("x")
			}

			_, dutErr := sut.ValidMetrics()

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})
		})

		Convey("listing interfaces was successful", func() {
			collector.makeSomeData()

			Convey("getting driver name was unsuccessful", func() {
				netInterfaces = func() ([]net.Interface, error) {
					return []net.Interface{
						net.Interface{Name: "invalid0", Flags: net.FlagUp},
					}, nil
				}

				_, dutErr := sut.ValidMetrics()

				Convey("returns error", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})

			netInterfaces = func() ([]net.Interface, error) {
				return []net.Interface{
					net.Interface{Name: "valid1", Flags: net.FlagUp},
					net.Interface{Name: "valid2", Flags: net.FlagUp},
					net.Interface{Name: "invalid1", Flags: net.FlagUp},
					net.Interface{Name: "loopback3", Flags: net.FlagLoopback},
				}, nil
			}

			dut, dutErr := sut.ValidMetrics()

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})

			Convey("queries each non-loopback interface about metrics", func() {
				collector.AssertCalled(t, "GetDriverInfo", "valid1")
				collector.AssertCalled(t, "GetDriverInfo", "valid2")
				collector.AssertCalled(t, "GetDriverInfo", "invalid1")

				collector.AssertCalled(t, "GetNicStats", "valid1")
				collector.AssertCalled(t, "GetNicStats", "valid2")
				collector.AssertCalled(t, "GetNicStats", "invalid1")

				collector.AssertCalled(t, "GetRegDump", "valid1")
				collector.AssertCalled(t, "GetRegDump", "valid2")
				collector.AssertCalled(t, "GetRegDump", "invalid1")

				collector.AssertCalled(t, "GetDomStats", "valid1")
				collector.AssertCalled(t, "GetDomStats", "valid2")
				collector.AssertCalled(t, "GetDomStats", "invalid1")
				collector.AssertCalled(t, "GetDomStats", "invalid1")
			})

			Convey("does not query loopback interfaces", func() {
				collector.AssertNotCalled(t, "GetNicStats", "loopback3")

			})

			Convey("returned list of metrics", func() {

				Convey("contains metrics for interfaces supporting getting stats", func() {

					for src, metrics := range dut {
						So(src.kind, ShouldNotEqual, 0)
						So(src.device, ShouldNotBeEmpty)

						switch src.device {
						case "invalid1":
							// does not contain metrics from interfaces which not support getting stats
							So(metrics, ShouldBeEmpty)

						case "valid2":
							So(src.driver, ShouldEqual, "e1000e")
							if src.kind == COLLECT_NIC {
								Convey("contain metrics from valid2 interfaces which support getting stats", func() {
									So(metrics, ShouldContain, "m3")
								})
							} else {
								So(metrics, ShouldBeEmpty)
							}

						case "valid1":
							So(src.driver, ShouldEqual, "ixgbe")
							switch src.kind {
							case COLLECT_NIC:
								Convey("contains metrics from valid1 interfaces which support getting stats", func() {
									So(metrics, ShouldContain, "m1")
									So(metrics, ShouldContain, "m2")
								})

							case COLLECT_REGDUMP:
								Convey("contains metrics from valid1 interfaces which support getting register dump", func() {
									So(metrics, ShouldContain, "r1")
									So(metrics, ShouldContain, "r2")
								})

							case COLLECT_DOM:
								Convey("contains metrics from valid1 interfaces which support getting digital optical monitoring info", func() {
									So(metrics, ShouldContain, "s1")
									So(metrics, ShouldContain, "s2")
								})
							}
						}
					}
				})
			})
		})
	})

}

func TestCollectMetrics(t *testing.T) {

	Convey("CollectMetrics", t, func() {
		collector := &cMock{}
		sut := &metricCollectorImpl{Ethtool: collector}

		// mock metrics source
		srcValid1Nic := metricSource{device: "valid1", driver: "ixgbe", kind: COLLECT_NIC}
		srcValid1Reg := metricSource{device: "valid1", driver: "ixgbe", kind: COLLECT_REGDUMP}
		srcValid1Dom := metricSource{device: "valid1", driver: "ixgbe", kind: COLLECT_DOM}
		srcDisabled := metricSource{device: "valid2", driver: "e1000e", kind: COLLECT_NIC}

		mockIset := map[metricSource]bool{
			srcValid1Nic: true,
			srcValid1Reg: true,
			srcValid1Dom: true,
			srcDisabled:  false,
		}

		Convey("returned list of metrics", func() {
			collector.makeSomeData()
			dut, dutErr := sut.Collect(mockIset)

			Convey("no error is returned", func() {
				So(dutErr, ShouldBeNil)
			})

			Convey("does not contain unsupported metrics", func() {
				So(dut[srcDisabled], ShouldBeNil)
			})

			Convey("contains metrics from each interface", func() {
				So(dut[srcValid1Nic]["m1"], ShouldEqual, "1")
				So(dut[srcValid1Nic]["m2"], ShouldEqual, "2")

				So(dut[srcValid1Reg]["r1"], ShouldEqual, "1")
				So(dut[srcValid1Reg]["r2"], ShouldEqual, "2")

				So(dut[srcValid1Dom]["s1"], ShouldEqual, "1")
				So(dut[srcValid1Dom]["s2"], ShouldEqual, "2")
			})
		})

		Convey("when querying interface fails", func() {

			Convey("GetNicStats querying interface fails", func() {
				collector.On("GetNicStats", mock.AnythingOfType("string")).Return(nil, errors.New("x"))
				collector.makeSomeData()

				_, dutErr := sut.Collect(mockIset)

				Convey("error is returned", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})

			Convey("GetRegDump querying interface fails", func() {
				collector.On("GetRegDump", mock.AnythingOfType("string")).Return(nil, errors.New("x"))
				collector.makeSomeData()
				_, dutErr := sut.Collect(mockIset)

				Convey("error is returned", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})

			Convey("GetDomStats querying interface fails", func() {
				collector.On("GetDomStats", mock.AnythingOfType("string")).Return(nil, errors.New("x"))
				collector.makeSomeData()
				_, dutErr := sut.Collect(mockIset)

				Convey("error is returned", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})
		})
	})

}

func (cm *cMock) makeSomeData() {
	// mock for GetDriverInfo
	cm.On("GetDriverInfo", "valid1").Return("ixgbe", nil)
	cm.On("GetDriverInfo", "valid2").Return("e1000e", nil)
	cm.On("GetDriverInfo", "invalid0").Return("", errors.New("Cannot get driver information: Operation not supported"))
	cm.On("GetDriverInfo", "invalid1").Return("ixgbe", nil)
	cm.On("GetDriverInfo", "loopback3").Return("test", nil)

	// mock for GetNicStats
	cm.On("GetNicStats", "valid1").Return(map[string]string{"m1": "1", "m2": "2"}, nil)
	cm.On("GetNicStats", "valid2").Return(map[string]string{"m3": "3"}, nil)
	cm.On("GetNicStats", "invalid0").Return(nil, errors.New("virtual one"))
	cm.On("GetNicStats", "invalid1").Return(nil, errors.New("virtual one"))
	cm.On("GetNicStats", "loopback3").Return(map[string]string{"m4": "4"}, nil)

	// mock for GetRegDump
	cm.On("GetRegDump", "valid1").Return(map[string]string{"r1": "1", "r2": "2"}, nil)
	cm.On("GetRegDump", "valid2").Return(nil, errors.New("Cannot get register dump: Operation not supported"))
	cm.On("GetRegDump", "invalid0").Return(nil, errors.New("virtual one"))
	cm.On("GetRegDump", "invalid1").Return(nil, errors.New("virtual one"))
	cm.On("GetRegDump", "loopback3").Return(map[string]string{"m3": "3"}, nil)

	// mock for GetDomStats
	cm.On("GetDomStats", "valid1").Return(map[string]string{"s1": "1", "s2": "2"}, nil)
	cm.On("GetDomStats", "valid2").Return(nil, errors.New("Cannot get module EEPROM information: Operation not supported"))
	cm.On("GetDomStats", "invalid0").Return(nil, errors.New("virtual one"))
	cm.On("GetDomStats", "invalid1").Return(nil, errors.New("virtual one"))
	cm.On("GetDomStats", "loopback3").Return(map[string]string{"m3": "3"}, nil)
}
