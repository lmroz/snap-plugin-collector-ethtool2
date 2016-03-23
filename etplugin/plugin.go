/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2016 Intel Corporation
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
	"strings"
	"time"

	"github.com/intelsdi-x/snap-plugin-collector-ethtool/ethtool"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (
	// Name of plugin
	Name = "ethtool"
	// Version of plugin
	Version = 5
	// Type of plugin
	Type = plugin.CollectorPluginType
)

const (
	statNicLabel = "nic"
	statRegLabel = "reg"
	statDomLabel = "dom"
)

var namespacePrefix = []string{"intel", "net"}
var prefixLength = len(namespacePrefix)

// makeName creates metrics namespace includes device(contains driver info), type of metrics and metric name
func makeName(source metricSource, metric string) []string {
	kind := statNicLabel

	switch source.kind {
	case COLLECT_REGDUMP:
		kind = statRegLabel
	case COLLECT_DOM:
		kind = statDomLabel
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
			source.kind = COLLECT_NIC
		case statRegLabel:
			source.kind = COLLECT_REGDUMP
		case statDomLabel:
			source.kind = COLLECT_DOM
		default:
			source.kind = COLLECT_NONE
		}
	}
	return
}

// NetPlugin is the main class
type NetPlugin struct {
	mc metricCollector
}

// CollectMetrics retrieves values for given metrics
func (p *NetPlugin) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {

	// with which interfaces and what type of statistics are desired to be collected; avoid unnecessary command execution
	iset := map[metricSource]bool{}

	for _, mt := range mts {
		src, _ := parseName(mt.Namespace().Strings())
		iset[src] = true
	}

	metrics, err := p.mc.Collect(iset)
	if err != nil {
		return nil, err
	}

	t := time.Now()

	results := make([]plugin.MetricType, len(mts))

	for i, mt := range mts {
		src, metric := parseName(mt.Namespace().Strings())
		val, ok := metrics[src][metric]
		if !ok {
			return nil, fmt.Errorf("unknown stat: %s on interface %s", metric, src.device)
		}

		results[i] = plugin.MetricType{
			Namespace_: mt.Namespace(),
			Data_:      val,
			Timestamp_: t,
		}
	}

	return results, nil
}

// GetMetricTypes returns list of metrics. Metrics are put in namespaces {"intel", "net", DRIVER, INTERFACE, metrics...}
func (p *NetPlugin) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}

	valid, err := p.mc.ValidMetrics()
	if err != nil {
		return nil, err
	}

	for source, metrics := range valid {
		for _, metric := range metrics {
			ns := makeName(source, metric)
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace(ns...)})
		}
	}

	return mts, nil
}

// GetConfigPolicy returns config policy
func (p *NetPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

// NewNetCollector creates new instance of plugin and returns pointer to initialized object
func NewNetCollector() *NetPlugin {
	exc := new(ethtool.LocalExecutor)
	col := &ethtool.ToolCollector{Tool: exc}
	mc := &metricCollectorImpl{Ethtool: col}

	return &NetPlugin{mc: mc}
}

// Meta returns plugin's metadata
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}
