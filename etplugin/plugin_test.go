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

func (mc *mcMock) ValidMetrics() (map[metricSource][]string, error) {
	args := mc.Called()
	var r0 map[metricSource][]string
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[metricSource][]string)
	}
	return r0, args.Error(1)
}

func (mc *mcMock) Collect(iset map[metricSource]bool) (map[metricSource]map[string]string, error) {
	args := mc.Called(iset)
	var r0 map[metricSource]map[string]string
	if args.Get(0) != nil {
		r0 = args.Get(0).(map[metricSource]map[string]string)
	}

	return r0, args.Error(1)
}

func flattenMts(mts []plugin.PluginMetricType) (map[string]interface{}, []string) {
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
	Convey("GetMetricTypes", t, func() {

		m := &mcMock{}
		sut := NetPlugin{mc: m}

		Convey("when metric collector returns list of valid metrics", func() {

			// mock metrics source
			srcValid1Nic := metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_NIC}
			srcValid1Reg := metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_REGDUMP}
			srcValid1Dom := metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_DOM}
			srcValid2Nic := metricSource{device: "eth1", driver: "e1000e", kind: COLLECT_NIC}

			m.On("ValidMetrics").Return(map[metricSource][]string{
				srcValid1Nic: []string{"m1", "m2"},
				srcValid2Nic: []string{"m3", "m4"},
				srcValid1Reg: []string{"r1", "r2"},
				srcValid1Dom: []string{"s1", "s2"},
			}, nil)

			dut, dutErr := sut.GetMetricTypes(plugin.PluginConfigType{})

			Convey("returns no error", func() {
				So(dutErr, ShouldBeNil)
			})

			Convey("entire list is exposed", func() {
				_, mts := flattenMts(dut)

				So(mts, ShouldContain, "intel/net/ixgbe/eth0/nic/m1")
				So(mts, ShouldContain, "intel/net/ixgbe/eth0/nic/m2")

				So(mts, ShouldContain, "intel/net/e1000e/eth1/nic/m3")
				So(mts, ShouldContain, "intel/net/e1000e/eth1/nic/m4")

				So(mts, ShouldContain, "intel/net/ixgbe/eth0/reg/r1")
				So(mts, ShouldContain, "intel/net/ixgbe/eth0/reg/r2")

				So(mts, ShouldContain, "intel/net/ixgbe/eth0/dom/s1")
				So(mts, ShouldContain, "intel/net/ixgbe/eth0/dom/s2")

				Convey("and nothing more", func() {
					// 8 metrics exposed
					So(len(dut), ShouldEqual, 8)

				})

			})
		})

		Convey("when metric collector returned error", func() {

			m.On("ValidMetrics").Return(nil, errors.New("x"))

			dut, dutErr := sut.GetMetricTypes(plugin.PluginConfigType{})

			So(dutErr, ShouldNotBeNil)
			So(dut, ShouldBeNil)
		})

	})

}

func TestPluginCollectMetrics(t *testing.T) {
	Convey("CollectMetrics", t, func() {

		m := &mcMock{}
		sut := NetPlugin{mc: m}

		mts := []plugin.PluginMetricType{
			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "nic", "m1"}},
			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "nic", "m2"}},

			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "reg", "r1"}},
			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "reg", "r2"}},

			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "dom", "s1"}},
			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "ixgbe", "eth0", "dom", "s2"}},

			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "e1000e", "eth1", "nic", "m3"}},
			plugin.PluginMetricType{Namespace_: []string{"intel", "net", "e1000e", "eth1", "nic", "m4"}},
		}

		Convey("asks metric collector about each required interface", func() {

			Convey("when metric collector returned valid list of metrics", func() {

				m.On("Collect", mock.AnythingOfType("map[etplugin.metricSource]bool")).Return(
					map[metricSource]map[string]string{
						metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_NIC}:     map[string]string{"m1": "1", "m2": "2"},
						metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_REGDUMP}: map[string]string{"r1": "1", "r2": "2"},
						metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_DOM}:     map[string]string{"s1": "1", "s2": "2"},
						metricSource{device: "eth1", driver: "e1000e", kind: COLLECT_NIC}:    map[string]string{"m3": "3", "m4": "4"},
					}, nil)

				dut, dutErr := sut.CollectMetrics(mts)

				Convey("no error is returned", func() {
					So(dutErr, ShouldBeNil)
				})

				Convey("only desired metrics are returned", func() {
					So(len(dut), ShouldEqual, len(mts))
				})

				Convey("entire list is exposed", func() {
					results, _ := flattenMts(dut)

					So(results["intel/net/ixgbe/eth0/nic/m1"], ShouldEqual, "1")
					So(results["intel/net/ixgbe/eth0/nic/m2"], ShouldEqual, "2")

					So(results["intel/net/ixgbe/eth0/reg/r1"], ShouldEqual, "1")
					So(results["intel/net/ixgbe/eth0/reg/r2"], ShouldEqual, "2")

					So(results["intel/net/ixgbe/eth0/dom/s1"], ShouldEqual, "1")
					So(results["intel/net/ixgbe/eth0/dom/s2"], ShouldEqual, "2")

					So(results["intel/net/e1000e/eth1/nic/m3"], ShouldEqual, "3")
					So(results["intel/net/e1000e/eth1/nic/m4"], ShouldEqual, "4")

					Convey("and nothing more", func() {
						So(len(dut), ShouldEqual, len(mts))
					})

				})

			})

			Convey("when this list has some metrics missing", func() {
				m.On("Collect", mock.AnythingOfType("map[etplugin.metricSource]bool")).Return(
					map[metricSource]map[string]string{
						metricSource{device: "eth0", driver: "ixgbe", kind: COLLECT_NIC}: map[string]string{"m1": "1", "m2": "2"},
					}, nil)

				_, dutErr := sut.CollectMetrics(mts)

				Convey("error is returned", func() {
					So(dutErr, ShouldNotBeNil)
				})

			})

			Convey("when metric collector returned error", func() {
				m.On("Collect", mock.AnythingOfType("map[etplugin.metricSource]bool")).Return(nil, errors.New("x"))

				_, dutErr := sut.CollectMetrics(mts)

				Convey("error is returned", func() {
					So(dutErr, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestGetConfigPolicy(t *testing.T) {
	Convey("GetConfigPolicy", t, func() {
		sut := &NetPlugin{}
		dut, dutErr := sut.GetConfigPolicy()
		Convey("Returns correct type", func() {
			So(dut, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
		})
		Convey("Is not nil", func() {
			So(dut, ShouldNotBeNil)
		})

		Convey("Returns no error", func() {
			So(dutErr, ShouldBeNil)
		})
	})
}
