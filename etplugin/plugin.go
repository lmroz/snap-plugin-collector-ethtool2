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
	"os"
	"strconv"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"

	"github.com/intelsdi-x/snap-plugin-collector-ethtool/ethtool"
)

const (
	// Name of plugin
	Name = "ethtool"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

var namespacePrefix = []string{"intel", "net"}

func makeName(device, metric string) []string {
	result := []string{}
	result = append(result, namespacePrefix...)
	result = append(result, device, metric)
	return result

}

// performs reverse operation to make name, extracts device
// and metric from namespace
func parseName(ns []string) (device, metric string) {
	return ns[len(namespacePrefix)], ns[len(namespacePrefix)+1]

}

// Plugin's main class.
type IXGBEPlugin struct {
	mc metricCollector
}

// Retrieves values for given metrics.
// Asks each device only once about it's metrics.
func (p *IXGBEPlugin) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	iset := map[string]bool{}
	for _, mt := range mts {
		dev, _ := parseName(mt.Namespace())
		iset[dev] = true
	}
	metrics, err := p.mc.CollectMetrics(iset)
	if err != nil {
		return nil, err
	}

	// it's not worth to abort collection
	// when only os.Hostname() raised error
	host, _ := os.Hostname()
	t := time.Now()

	results := make([]plugin.PluginMetricType, len(mts))

	for i, mt := range mts {
		dev, metric := parseName(mt.Namespace())
		val, ok := metrics[dev+"/"+metric]
		if !ok {
			return nil, fmt.Errorf("unknown stat: %s on interface %s", metric, dev)
		}

		vInt, err := strconv.ParseInt(val, 10, 64)

		if err != nil {
			return nil, fmt.Errorf("incorrect metric value: %s = %s", metric, val)
		}

		results[i] = plugin.PluginMetricType{
			Namespace_: mt.Namespace(),
			Data_:      vInt,
			Source_:    host,
			Timestamp_: t,
		}

	}

	return results, nil
}

// Returns list of metrics.
// Metrics are put in namespaces {"intel", "net", INTERFACE, metrics...}
func (p *IXGBEPlugin) GetMetricTypes(_ plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mts := []plugin.PluginMetricType{}
	valid, err := p.mc.ValidMetrics()
	if err != nil {
		return nil, err
	}
	for dev, metrics := range valid {
		for _, metric := range metrics {
			ns := makeName(dev, metric)
			mts = append(mts, plugin.PluginMetricType{Namespace_: ns})
		}
	}

	return mts, nil
}

// GetConfigPolicy
func (p *IXGBEPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

// Creates new instance of plugin and returns pointer to initialized object.
func NewIXGBECollector() *IXGBEPlugin {
	exc := new(ethtool.LocalExecutor)
	col := &ethtool.ToolCollector{Tool: exc}
	mc := &metricCollectorImpl{Ethtool: col}

	return &IXGBEPlugin{mc: mc}
}

// Returns plugin's metadata
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}
