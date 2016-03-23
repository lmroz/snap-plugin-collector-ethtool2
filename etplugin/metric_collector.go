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

type metricKind int

const (
	COLLECT_NONE metricKind = 0
	COLLECT_NIC             = 1 << iota
	COLLECT_REGDUMP
	COLLECT_DOM
)

type metricSource struct {
	device string
	driver string
	kind   metricKind
}

type metricDescription struct {
	name, driver string
}

type metricCollector interface {
	Collect(iset map[metricSource]bool) (map[metricSource]map[string]string, error)
	ValidMetrics() (map[metricSource][]string, error)
}

type metricCollectorImpl struct {
	Ethtool ethtool.Collector
}

// for mocking
var netInterfaces = func() ([]net.Interface, error) {
	return net.Interfaces()
}

// getKeys returns metric keys as a slice
func getKeys(m map[string]string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}

	//transform map to slice
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s, nil
}

// ValidMetrics returns map of metrics for each existing net interfaces
func (mc *metricCollectorImpl) ValidMetrics() (map[metricSource][]string, error) {
	ifaces, err := netInterfaces()
	if err != nil {
		return nil, err
	}

	result := map[metricSource][]string{}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		driverName, err := mc.Ethtool.GetDriverInfo(iface.Name)
		if err != nil {
			continue
		}

		result[metricSource{driver: driverName, device: iface.Name, kind: COLLECT_NIC}], err = getKeys(mc.Ethtool.GetNicStats(iface.Name))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting NIC stats for %+v, %+v\n", iface.Name, err)
		}

		result[metricSource{driver: driverName, device: iface.Name, kind: COLLECT_REGDUMP}], err = getKeys(mc.Ethtool.GetRegDump(iface.Name))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting register dump for %+v, %+v\n", iface.Name, err)
		}

		result[metricSource{driver: driverName, device: iface.Name, kind: COLLECT_DOM}], err = getKeys(mc.Ethtool.GetDomStats(iface.Name))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting module EEPROM information for %+v, %+v\n", iface.Name, err)
		}
	}

	return result, nil
}

// Collect returns all available metrics exposed by metric source (defined as interface name and kind of ethtool command to execute)
func (mc *metricCollectorImpl) Collect(iset map[metricSource]bool) (map[metricSource]map[string]string, error) {
	result := map[metricSource]map[string]string{}

	for src, ok := range iset {

		if !ok {
			// skip this source
			continue
		}

		var err error
		stats := map[string]string{}

		if src.kind == COLLECT_NIC {
			//collect output from ethtool -s (adapter statistics)
			stats, err = mc.Ethtool.GetNicStats(src.device)
			if err != nil {
				return nil, fmt.Errorf("canot read NIC stats from %s [%v]", src.device, err)
			}
		}

		if src.kind == COLLECT_REGDUMP {
			//collect output from ethtool -d (statictics from register dump)
			stats, err = mc.Ethtool.GetRegDump(src.device)
			if err != nil {
				return nil, fmt.Errorf("cant read register dump raw stats from %s [%v]", src.device, err)
			}
		}

		if src.kind == COLLECT_DOM {
			//collect output from ethtool -m, which provides digital optical monitoring info (diagnostics data and alarms for optical transceivers)
			stats, err = mc.Ethtool.GetDomStats(src.device)
			if err != nil {
				return nil, fmt.Errorf("Cannot read digital optical monitoring information from %s [%v]", src.device, err)
			}
		}

		if len(stats) > 0 {
			result[src] = stats
		}

	}

	return result, nil
}
