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
	"fmt"
	"net"
	"os"

	"github.com/intelsdi-x/snap-plugin-collector-ethtool/ethtool"
)

type metricCollector interface {
	Collect(iset map[string]*collectInfo) (map[string]string, error)
	ValidMetrics() (map[string][]string, error)
	ValidRegDumpMetrics() (map[string][]string, error)
}

type metricCollectorImpl struct {
	Ethtool ethtool.Collector
}

// for mocking
var netInterfaces = func() ([]net.Interface, error) {
	return net.Interfaces()
}

// ValidMetrics returns map of net interfaces existing in system (includes driver name)
// and available metrics for them retrived from `ethtool -S`
func (mc *metricCollectorImpl) ValidMetrics() (map[string][]string, error) {
	ifaces, err := netInterfaces()
	if err != nil {
		return nil, err
	}

	result := map[string][]string{}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		driver, err := mc.Ethtool.GetDriverInfo(iface.Name)
		if err != nil {
			return nil, err
		}

		// include driver's name and net interface into metric namespace
		name := joinNamespace([]string{driver, iface.Name})

		stats, err := mc.Ethtool.GetStats(iface.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting NIC stats for %+v, %+v\n", iface.Name, err)
			continue
		}

		validList := make([]string, 0, len(stats))
		for stat, _ := range stats {
			validList = append(validList, stat)
		}
		result[name] = validList
	}

	return result, nil
}

// ValidRegDumpMetrics returns map of net interfaces existing in system (includes driver name)
// and available register dump metrics for them retrived from `ethtool -d`
func (mc *metricCollectorImpl) ValidRegDumpMetrics() (map[string][]string, error) {

	ifaces, err := netInterfaces()
	if err != nil {
		return nil, err
	}

	result := map[string][]string{}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		driver, err := mc.Ethtool.GetDriverInfo(iface.Name)
		if err != nil {
			return nil, err
		}

		// include driver's name and net interface into metric namespace
		name := joinNamespace([]string{driver, iface.Name})

		stats, err := mc.Ethtool.GetRegDump(iface.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting register dump for %+v, %+v\n", iface.Name, err)
			continue
		}

		validList := make([]string, 0, len(stats))
		for stat, _ := range stats {
			validList = append(validList, stat)
		}

		result[name] = validList
	}

	return result, nil
}

// CollectMetrics returns all desired net metrics defined in task manifest
func (mc *metricCollectorImpl) Collect(iset map[string]*collectInfo) (map[string]string, error) {
	result := map[string]string{}

	for dev, info := range iset {

		if info.collect_s {
			//collect output from ethtool -s (adapter statistics)
			stats, err := mc.Ethtool.GetStats(dev)
			if err != nil {
				return nil, fmt.Errorf("cant read stats from %s [%v]", dev, err)
			}
			for stat, value := range stats {
				ns := joinNamespace([]string{dev, statNicLabel, stat})
				result[ns] = value
			}
		}

		if info.collect_d {
			//collect output from ethtool -d (statictics from register dump)
			statsReg, errReg := mc.Ethtool.GetRegDump(dev)
			if errReg != nil {
				return nil, fmt.Errorf("cant read register dump raw stats from %s [%v]", dev, errReg)
			}
			for statReg, valueReg := range statsReg {
				ns := joinNamespace([]string{dev, statRegLabel, statReg})
				result[ns] = valueReg
			}
		}

	}

	return result, nil
}
