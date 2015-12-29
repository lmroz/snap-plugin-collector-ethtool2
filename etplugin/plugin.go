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

	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	//"github.com/intelsdi-x/snap-plugin-utilities/config"

	"github.com/intelsdi-x/snap-plugin-collector-ethtool/ethtool"
)

const (
	// Name of plugin
	Name = "ethtool"
	// Version of plugin
	Version = 2
	// Type of plugin
	Type = plugin.CollectorPluginType
)

const (
	statNicLabel = "nic"
	statRegLabel = "reg"
)

var namespacePrefix = []string{"intel", "net"}
var prefixLength = len(namespacePrefix)
// makeName creates metrics namespace includes device(contains driver info), type of metrics and metric name
func makeName(source metricSource, metric string) []string {
	kind := statNicLabel
	if source.kind == COLLECT_REGDUMP {
		kind = statRegLabel
	}
	ns := append(namespacePrefix, source.driver)
	ns = append(ns, source.device)
	ns = append(ns, kind)
	ns = append(ns, strings.Split(metric, "/")...)
	return ns
}

// parseName performs reverse operation to make name, extracts driver, device, type of metric (eg. nic statistics or register dump)
// and metric name from namespace
func parseName(ns []string) (source metricSource, metric string) {

	if len(ns) > prefixLength+3 {
		source.driver = ns[prefixLength]
		source.device = ns[prefixLength+1]
		kind := ns[prefixLength+2]
		metric = strings.Join(ns[prefixLength+3:], "/")
	
		switch kind {
			case statNicLabel: 
				source.kind = COLLECT_STAT
			case statRegLabel:
				source.kind = COLLECT_REGDUMP
			default:
				source.kind = COLLECT_NONE
		}
	}
	return
}

// joinNamespace concatenates the elements of `ns` to create a single string with slash as the separator between elements in the resulting string.
func joinNamespace(ns []string) string {
	return strings.Join(ns, "/")
}

// Plugin's main class.
type IXGBEPlugin struct {
	mc metricCollector
}

// CollectMetrics retrieves values for given metrics
func (p *IXGBEPlugin) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {

	// with which interfaces and what type of statistics are desired to be collected; avoid unnecessary ethtool -d command execution
	iset := map[metricSource]bool{}
	
	for _, mt := range mts {
		src, _ := parseName(mt.Namespace())
		iset[src] = true
	}

	metrics, err := p.mc.Collect(iset)
	if err != nil {
		return nil, err
	}

	// it's not worth to abort collection when only os.Hostname() raised error
	host, _ := os.Hostname()
	t := time.Now()

	results := make([]plugin.PluginMetricType, len(mts))

	for i, mt := range mts {
		src, metric := parseName(mt.Namespace())
		val, ok := metrics[src][metric]
		if !ok {
			return nil, fmt.Errorf("unknown stat: %s on interface %s", metric, src.device)
		}

		results[i] = plugin.PluginMetricType{
			Namespace_: mt.Namespace(),
			Data_:      val,
			Source_:    host,
			Timestamp_: t,

		}
	}

	return results, nil
}

// GetMetricTypes returns list of metrics. Metrics are put in namespaces {"intel", "net", DRIVER, INTERFACE, metrics...}
func (p *IXGBEPlugin) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mts := []plugin.PluginMetricType{}

	// metrics from command `ethtool -S <interface>`
	valid, err := p.mc.ValidMetrics()
	if err != nil {
		return nil, err
	}

	for source, metrics := range valid {
		for _, metric := range metrics {
			ns := makeName(source, metric)
			mts = append(mts, plugin.PluginMetricType{Namespace_: ns})
		}
	}
	/*
		for dev, metrics := range valid {
			for _, metric := range metrics {
				ns := makeName(dev, statNicLabel, metric)
				tags["driver"], _, _, _ = parseName(ns)
				mts = append(mts, plugin.PluginMetricType{Namespace_: ns, Tags_: tags})
			}
		}

		for dev, metrics := range valid {
			for _, metric := range metrics {
				ns := makeName(dev, statNicLabel, metric)
				tags["driver"], _, _, _ = parseName(ns)
				mts = append(mts, plugin.PluginMetricType{Namespace_: ns, Tags_: tags})
			}
		}
	*/
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
