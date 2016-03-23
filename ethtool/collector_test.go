/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

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

package ethtool

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type executorMock struct {
	mock.Mock
}

func (eMock *executorMock) Execute(option, iface string) ([]string, error) {
	args := eMock.Called(option, iface)
	r0 := []string{}
	if args.Get(0) != nil {
		r0 = args.Get(0).([]string)
	}

	return r0, args.Error(1)
}

func TestGetNicStats(t *testing.T) {

	Convey("GetNicStats", t, func() {

		executor := &executorMock{}
		sut := &ToolCollector{Tool: executor}

		Convey("succefully get NIC stats", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{"NIC statistics:",
				"a : b  ", "\tc:d", "   e   :f"}, nil)

			dut, dutErr := sut.GetNicStats("abc")

			Convey("returns parsed stats", func() {
				So(dut["a"], ShouldEqual, "b")
				So(dut["c"], ShouldEqual, "d")
				So(dut["e"], ShouldEqual, "f")
			})

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})
		})

		Convey("execution failed", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dutErr := sut.GetNicStats("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})
		})

		Convey("empty execution output ", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{}, errors.New("No output"))

			_, dutErr := sut.GetNicStats("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})
		})

		Convey("incorrect execution output, invalid first line", func() {

			Convey("invalid title line", func() {
				executor.On("Execute", mock.AnythingOfType("string"),
					mock.AnythingOfType("string")).Return([]string{"NIC statistics:",
					"a : b  ", "\tc:d", "   e invalid"}, nil)

				_, dutErr := sut.GetNicStats("abc")

				Convey("returns error", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})

			Convey("invali data format", func() {
				executor.On("Execute", mock.AnythingOfType("string"),
					mock.AnythingOfType("string")).Return([]string{"Unknown statistics",
					"a : b  ", "\tc:d", "   e   :f"}, nil)

				_, dutErr := sut.GetNicStats("abc")

				Convey("returns error", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestGetRegDump(t *testing.T) {

	Convey("GetRegDump", t, func() {
		executor := &executorMock{}
		sut := &ToolCollector{Tool: executor}

		Convey("succefully get register dump", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"\t0x040DC: ptc127  (Packets Tx (65-127B) Count)  0x0015F7CC",
				"Link Status(test)\t:  up",
				"Link Speed :\t10G"}, nil)

			dut, dutErr := sut.GetRegDump("abc")

			Convey("returns parsed stats", func() {
				So(dut["ptc127"], ShouldEqual, "0x0015F7CC")
				So(dut["link_status"], ShouldEqual, "up")
				So(dut["link_speed"], ShouldEqual, "10G")
			})

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})
		})

		Convey("execution failed", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dutErr := sut.GetRegDump("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})

		Convey("empty execution output ", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{}, errors.New("No output"))

			_, dutErr := sut.GetRegDump("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})

		Convey("incorrect execution output", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"ab", "c d", "e::f"}, nil)

			dut, _ := sut.GetRegDump("abc")

			Convey("returns empty dut", func() {
				So(dut, ShouldBeEmpty)
			})
		})
	})
}

func TestGetDomStats(t *testing.T) {

	Convey("GetDomStats", t, func() {
		executor := &executorMock{}
		sut := &ToolCollector{Tool: executor}

		Convey("succefully get plugable transceiver DOM info", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"Transceiver type : info1",
				"Transceiver type : info2",
				"BR, Nominal : 1000MBd",
				"Length (SMF) : 0m",
				"Module voltage : 1.0000 V"}, nil)

			dut, dutErr := sut.GetDomStats("abc")

			Convey("returns parsed stats", func() {
				So(dut["transceiver_type"], ShouldEqual, "info1, info2")
				So(dut["br_nominal"], ShouldEqual, "1000MBd")
				So(dut["length_smf"], ShouldEqual, "0m")
				So(dut["module_voltage"], ShouldEqual, "1.0000 V")
			})

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})
		})

		Convey("execution failed", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dutErr := sut.GetDomStats("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})
		})

		Convey("empty execution output ", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{}, errors.New("No output"))

			_, dutErr := sut.GetDomStats("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})

		Convey("incorrect execution output", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"a : b", "\tcd", "e::f"}, nil)

			_, dutErr := sut.GetDomStats("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})
	})
}

func TestGetDriverInfo(t *testing.T) {

	Convey("GetDriverInfo", t, func() {
		executor := &executorMock{}
		sut := &ToolCollector{Tool: executor}

		Convey("succefully get driver name", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"header",
				"   driver : drv0 ",
				"footer"}, nil)

			dut, dutErr := sut.GetDriverInfo("abc")

			Convey("returns parsed driver name", func() {
				So(dut, ShouldEqual, "drv0")
			})

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})
		})

		Convey("execution failed", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dutErr := sut.GetDriverInfo("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})
		})

		Convey("empty execution output", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{}, errors.New("No output"))

			_, dutErr := sut.GetDriverInfo("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})

		Convey("incorrect execution output", func() {
			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{
				"xxx"}, nil)

			_, dutErr := sut.GetDriverInfo("abc")

			Convey("returns error", func() {
				So(dutErr, ShouldNotBeNil)
			})

		})
	})
}
