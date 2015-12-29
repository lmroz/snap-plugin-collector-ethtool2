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
	COLLECT_STAT            = 1 << iota
	COLLECT_REGDUMP
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

func keysError(m map[string]string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}

	s := make([]string, 0, len(m))
	for k, _ := range m {
		s = append(s, k)
	}
	return s, nil
}

// ValidMetrics returns map of net interfaces existing in system (includes driver name)
// and available metrics for them retrived from `ethtool -S`
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
			return nil, err
		}

		result[metricSource{driver: driverName, device: iface.Name, kind: COLLECT_STAT}], err = keysError(mc.Ethtool.GetStats(iface.Name))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting NIC stats for %+v, %+v\n", iface.Name, err)
		}

		result[metricSource{driver: driverName, device: iface.Name, kind: COLLECT_REGDUMP}], err = keysError(mc.Ethtool.GetRegDump(iface.Name))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting register dump for %+v, %+v\n", iface.Name, err)
		}
	}

	return result, nil
}

// CollectMetrics returns all desired metrics defined in task manifest
func (mc *metricCollectorImpl) Collect(iset map[metricSource]bool) (map[metricSource]map[string]string, error) {
	result := map[metricSource]map[string]string{}

	for src, _ := range iset {

		var err error
		stats := map[string]string{}

		if src.kind == COLLECT_STAT {
			//collect output from ethtool -s (adapter statistics)
			stats, err = mc.Ethtool.GetStats(src.device)
			if err != nil {
				return nil, fmt.Errorf("cant read stats from %s [%v]", src.device, err)
			}
		}

		if src.kind == COLLECT_REGDUMP {
			//collect output from ethtool -d (statictics from register dump)
			stats, err = mc.Ethtool.GetRegDump(src.device)
			if err != nil {
				return nil, fmt.Errorf("cant read register dump raw stats from %s [%v]", src.device, err)
			}
		}
		if len(stats) > 0 {
			result[src] = stats
		}

	}

	return result, nil
}
