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

func (self *executorMock) Execute(option, iface string) ([]string, error) {
	args := self.Called(option, iface)
	var r0 []string = nil
	if args.Get(0) != nil {
		r0 = args.Get(0).([]string)
	}

	return r0, args.Error(1)
}

func TestGetStats(t *testing.T) {
	Convey("GetStats", t, func() {

		executor := &executorMock{}
		sut := &ToolCollector{Tool: executor}

		Convey("when fed with correct output", func() {

			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{"NIC statistics:",
				"a : b  ", "\tc:d", "   e   :f"}, nil)

			dut, dut_err := sut.GetStats("abc")

			Convey("returns parsed stats", func() {

				So(dut["a"], ShouldEqual, "b")
				So(dut["c"], ShouldEqual, "d")
				So(dut["e"], ShouldEqual, "f")

			})

			Convey("returns no error", func() {

				So(dut_err, ShouldBeNil)

			})

			Convey("calls execute once with correct interface", func() {

				executor.AssertCalled(t, "Execute", mock.AnythingOfType("string"),
					"abc")
				executor.AssertNumberOfCalls(t, "Execute", 1)

			})

		})

		Convey("when fed with output with correct header but unknown row format", func() {

			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{"NIC statistics:", "xx"}, nil)

			_, dut_err := sut.GetStats("abc")

			Convey("returns error", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("when fed with output with incorrect header", func() {

			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{"x y z", "a : b"}, nil)

			_, dut_err := sut.GetStats("abc")

			Convey("returns error", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("when execution failed", func() {

			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return(nil, errors.New("x"))

			_, dut_err := sut.GetStats("abc")

			Convey("returns error", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("when execution returned empty output", func() {

			executor.On("Execute", mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).Return([]string{}, nil)

			_, dut_err := sut.GetStats("abc")

			Convey("returns error", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

	})
}
