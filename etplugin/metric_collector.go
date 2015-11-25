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

	"github.com/intelsdi-x/snap-plugin-collector-ethtool/ethtool"
)

type metricCollector interface {
	CollectMetrics(iset map[string]bool) (map[string]string, error)
	ValidMetrics() (map[string][]string, error)
}

type metricCollectorImpl struct {
	Ethtool ethtool.Collector
}

// for mocking
var netInterfaces = func() ([]net.Interface, error) {
	return net.Interfaces()
}

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
		stats, err := mc.Ethtool.GetStats(iface.Name)
		if err != nil {
			continue
		}

		validList := make([]string, 0, len(stats))
		for stat, _ := range stats {
			validList = append(validList, stat)
		}
		result[iface.Name] = validList
	}

	return result, nil
}

func (mc *metricCollectorImpl) CollectMetrics(iset map[string]bool) (map[string]string, error) {
	result := map[string]string{}
	for k, _ := range iset {
		stats, err := mc.Ethtool.GetStats(k)
		if err != nil {
			return nil, fmt.Errorf("cant read stats from %s [%v]", k, err)
		}
		for stat, value := range stats {
			result[k+"/"+stat] = value
		}
	}

	return result, nil
}
