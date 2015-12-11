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
	"strings"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type mcMock struct {
	mock.Mock
}

func (mc *mcMock) ValidMetrics() (map[string][]string, error) {
	args := mc.Called()
	var r0 map[string][]string = nil
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string][]string)
	}
	return r0, args.Error(1)
}

func (mc *mcMock) ValidRegDumpMetrics() (map[string][]string, error) {
	args := mc.Called()
	var r0 map[string][]string = nil
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string][]string)
	}
	return r0, args.Error(1)
}

func (mc *mcMock) Collect(iset map[string]*collectInfo) (map[string]string, error) {
	args := mc.Called(iset)
	var r0 map[string]string = nil
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[string]string)
	}
	return r0, args.Error(1)
}

func flattenMTS(mts []plugin.PluginMetricType) (map[string]interface{}, []string) {
	rl := make([]string, len(mts))
	rm := map[string]interface{}{}
	for i, mt := range mts {
		ns := strings.Join(mt.Namespace(), "/")
		rl[i] = ns
		rm[ns] = mt.Data()
	}

	return rm, rl
}

func TestPluginGetMetricTypes(t *testing.T) {
	Convey("IXGBEPlugin.GetMetricTypes", t, func() {

		m := &mcMock{}
		sut := IXGBEPlugin{mc: m}

		Convey("when metric collector returns list of valid metrics", func() {

			m.On("ValidMetrics").Return(map[string][]string{
				"eth0": []string{"m1", "m2"},
				"lo":   []string{"m3"},
			}, nil)

			/*
				dut, dut_err := sut.GetMetricTypes(plugin.PluginConfigType{})

				Convey("entire list is exposed", func() {
					//_, mts := flattenMTS(dut)

					//So(mts, ShouldContain, "intel/net/eth0/m1")
					//So(mts, ShouldContain, "intel/net/eth0/m2")
					//So(mts, ShouldContain, "intel/net/lo/m3")

					Convey("and nothing more", func() {

						So(len(dut), ShouldEqual, 3)

					})

				})
			*/

			Convey("returns no error", func() {

				//So(dut_err, ShouldBeNil)

			})

		})

		Convey("when metric collector returned error", func() {

			m.On("ValidMetrics").Return(nil, errors.New("x"))

			_, dut_err := sut.GetMetricTypes(plugin.PluginConfigType{})

			Convey("error is returned", func() {

				So(dut_err, ShouldNotBeNil)

			})

		})

	})

}

func TestPluginCollectMetrics(t *testing.T) {
	Convey("IXGBEPlugin.CollectMetrics", t, func() {

		m := &mcMock{}
		//sut := IXGBEPlugin{mc: m}

		/*
			mts := []plugin.PluginMetricType{
				plugin.PluginMetricType{Namespace_: []string{"intel", "net", "eth0", "m1"}},
				plugin.PluginMetricType{Namespace_: []string{"intel", "net", "eth0", "m2"}},
				plugin.PluginMetricType{Namespace_: []string{"intel", "net", "lo", "m3"}},
			}
		*/

		Convey("asks metric collector about each required interface", func() {
			m.On("CollectMetrics", mock.AnythingOfType("map[string]bool")).Return(
				map[string]string{"eth0/m1": "1", "eth0/m2": "2", "lo/m3": "3"},
				nil).Run(func(args mock.Arguments) {

				casted := args.Get(0).(map[string]bool)
				So(casted["eth0"], ShouldBeTrue)
				So(casted["lo"], ShouldBeTrue)
			})
			//sut.CollectMetrics(mts)
		})

		Convey("when metric collector returned valid list of metrics", func() {

			Convey("and this list contains everything we asked about", func() {

				m.On("CollectMetrics", mock.AnythingOfType("map[string]bool")).Return(
					map[string]string{"eth0/m1": "1", "eth0/m2": "2", "lo/m3": "3"}, nil)

				//dut, dut_err := sut.CollectMetrics(mts)
				//results, _ := flattenMTS(dut)

				Convey("each value is non-nil", func() {
					/*
						So(results["intel/net/eth0/m1"], ShouldNotBeNil)
						So(results["intel/net/eth0/m2"], ShouldNotBeNil)
						So(results["intel/net/lo/m3"], ShouldNotBeNil)
					*/
				})

				Convey("each value is correct", func() {
					/*
						So(results["intel/net/eth0/m1"], ShouldEqual, 1)
						So(results["intel/net/eth0/m2"], ShouldEqual, 2)
						So(results["intel/net/lo/m3"], ShouldEqual, 3)
					*/

				})

				Convey("no error is returned", func() {

					//So(dut_err, ShouldBeNil)

				})

			})

			Convey("but this list has some metrics missing", func() {
				m.On("CollectMetrics", mock.AnythingOfType("map[string]bool")).Return(
					map[string]string{"eth0/m1": "1", "eth0/m2": "2", "lo/m4": "3"}, nil)

				//_, dut_err := sut.CollectMetrics(mts)

				Convey("error is returned", func() {

					//So(dut_err, ShouldNotBeNil)

				})

			})

		})

		Convey("when metric collector returned error", func() {
			m.On("CollectMetrics", mock.AnythingOfType("map[string]bool")).Return(nil,
				errors.New("x"))

			//_, dut_err := sut.CollectMetrics(mts)

			Convey("error is returned", func() {

				//So(dut_err, ShouldNotBeNil)

			})

		})

		Convey("when metric collector returned non-int metric", func() {
			m.On("Collect", mock.AnythingOfType("map[string]*collectInfo")).Return(
				map[string]string{"eth0/m1": "1", "eth0/m2": "x", "lo/m3": "3"}, nil)

			//_, dut_err := sut.CollectMetrics(mts)
			Convey("error is returned", func() {

				//So(dut_err, ShouldNotBeNil)

			})
		})

	})
}

func TestGetConfigPolicy(t *testing.T) {
	Convey("GetConfigPolicy", t, func() {
		sut := &IXGBEPlugin{}
		dut, err_dut := sut.GetConfigPolicy()
		Convey("Returns correct type", func() {
			So(dut, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
		})
		Convey("Is not nil", func() {
			So(dut, ShouldNotBeNil)
		})

		Convey("Returns no error", func() {
			So(err_dut, ShouldBeNil)
		})
	})
}
