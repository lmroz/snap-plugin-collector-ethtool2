// +build unit

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
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type cMock struct {
	mock.Mock
}

func (self *cMock) GetStats(iface string) (map[string]string, error) {
	args := self.Mock.Called(iface)
	var r0 map[string]string = nil
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string]string)
	}
	return r0, args.Error(1)
}

func (self *cMock) makeSomeData() {
	self.On("GetStats", "valid1").Return(map[string]string{"m1": "1",
		"m2": "2"}, nil)
	self.On("GetStats", "invalid0").Return(nil, errors.New("virtual one!"))
	self.On("GetStats", "valid2").Return(map[string]string{"m3": "3"}, nil)
	self.On("GetStats", "loopback3").Return(map[string]string{"m4": "4"}, nil)

}

func TestValidMetrics(t *testing.T) {
	Convey("ValidMetrics", t, func() {

		orgNetInterfaces := netInterfaces
		Reset(func() {
			netInterfaces = orgNetInterfaces
		})

		collector := &cMock{}
		sut := &metricCollectorImpl{Ethtool: collector}

		Convey("when listing interfaces is unsuccessful", func() {

			netInterfaces = func() ([]net.Interface, error) {
				return nil, errors.New("x")
			}

			_, dut_err := sut.ValidMetrics()

			Convey("returns error", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("when listing interfaces was successful", func() {

			collector.makeSomeData()

			netInterfaces = func() ([]net.Interface, error) {
				return []net.Interface{
					net.Interface{Name: "valid1", Flags: net.FlagUp},
					net.Interface{Name: "invalid0", Flags: net.FlagUp},
					net.Interface{Name: "valid2", Flags: 0},
					net.Interface{Name: "loopback3", Flags: net.FlagLoopback},
				}, nil
			}

			dut, _ := sut.ValidMetrics()

			Convey("queries each non-loopback interface about metrics", func() {

				collector.AssertCalled(t, "GetStats", "valid1")
				collector.AssertCalled(t, "GetStats", "valid2")
				collector.AssertCalled(t, "GetStats", "invalid0")
			})

			Convey("does not query loopback interfaces", func() {

				collector.AssertNotCalled(t, "GetStats", "loopback3")

			})

			Convey("returned list of metrics", func() {

				Convey("contains metrics for interfaces supporting getting stats", func() {

					data, ok := dut["valid1"]
					So(ok, ShouldBeTrue)
					if ok {
						So(data, ShouldContain, "m1")
						So(data, ShouldContain, "m2")
					}

					data, ok = dut["valid2"]
					So(ok, ShouldBeTrue)
					if ok {
						So(data, ShouldContain, "m3")
					}

				})

				Convey("does not contain metrics from interfaces which not support getting stats", func() {

					_, ok := dut["invalid0"]
					So(ok, ShouldBeFalse)

				})

			})

		})

	})
}

func TestCollectMetrics(t *testing.T) {
	Convey("CollectMetrics", t, func() {

		collector := &cMock{}
		sut := &metricCollectorImpl{Ethtool: collector}

		metrics := map[string]bool{"valid1": true, "valid2": true}

		Convey("queries each given interface about metrics", func() {

			collector.On("GetStats", mock.AnythingOfType("string")).Return(map[string]string{"m": "0"}, nil)

			sut.CollectMetrics(metrics)

			collector.AssertCalled(t, "GetStats", "valid1")
			collector.AssertCalled(t, "GetStats", "valid2")

		})

		Convey("when querying interface fails", func() {

			collector.On("GetStats", mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dut_err := sut.CollectMetrics(metrics)

			Convey("error is returned", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("returned list of metrics", func() {

			collector.makeSomeData()

			dut, _ := sut.CollectMetrics(metrics)

			Convey("contains metrics from each interface", func() {

				t.Log(dut)

				got := map[string]bool{}
				for k, _ := range dut {
					got[strings.Split(k, "/")[0]] = true
				}

				So(got["valid1"], ShouldBeTrue)
				So(got["valid2"], ShouldBeTrue)

			})

			Convey("is complete", func() {

				So(dut["valid1/m1"], ShouldEqual, "1")
				So(dut["valid1/m2"], ShouldEqual, "2")
				So(dut["valid2/m3"], ShouldEqual, "3")

			})

			Convey("does not contain unsupported metrics", func() {

				So(len(dut), ShouldEqual, 3)
			})

		})

	})
}
