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

type collectInfo struct {
	collect_s bool // execute "ethtool -S"; NIC statistics
	collect_d bool // execute "ethtool -d"; Register statistics (register dump)
}

// makeName creates metrics namespace includes device(contains driver info), type of metrics and metric name
func makeName(device, typeOf, metric string) []string {
	ns := append(namespacePrefix, strings.Split(device, "/")...)
	ns = append(ns, typeOf)
	ns = append(ns, strings.Split(metric, "/")...)
	return ns
}

// parseName performs reverse operation to make name, extracts driver, device, type of metric (eg. nic statistics or register dump)
// and metric name from namespace
func parseName(ns []string) (driver, device, typeOf, metric string) {
	if len(ns) < len(namespacePrefix)+4 {
		panic("Cannot parse metric namespace")
	}

	return ns[len(namespacePrefix)], ns[len(namespacePrefix)+1], ns[len(namespacePrefix)+2], joinNamespace(ns[len(namespacePrefix)+3:])
}

// joinNamespace concatenates the elements of `ns` to create a single string with slash as the separator between elements in the resulting string.
func joinNamespace(ns []string) string {
	return strings.Join(ns, "/")
}

// Plugin's main class.
type IXGBEPlugin struct {
	mc metricCollector
}

// setMetricType returns map of interfaces with information about what type of statistics are desired to be collected.
// That indicate ethtool command arguments to execute.
func setMetricType(mts []plugin.PluginMetricType) map[string]*collectInfo {

	isetMt := make(map[string]*collectInfo)

	for _, mt := range mts {

		// retrive info about stat type from metric namespace
		_, dev, typeOf, _ := parseName(mt.Namespace())

		// set items defaults to false if not initialized before
		if isetMt[dev] == nil {

			isetMt[dev] = &collectInfo{collect_s: false, collect_d: false}
		}

		switch typeOf {
		case statNicLabel: // type of stats indicates nic stats, set collect_s true
			isetMt[dev].collect_s = true
			break

		case statRegLabel: // type of stats indicates reg dump stats, set collect_d true
			isetMt[dev].collect_d = true
			break

		default: // unrecognize type of metric, exit function with nil map
			return nil
		}
	}

	return isetMt
}

// CollectMetrics retrieves values for given metrics
func (p *IXGBEPlugin) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {

	// with which interfaces and what type of statistics are desired to be collected; avoid unnecessary ethtool -d command execution
	iset := setMetricType(mts)

	if iset == nil {
		return nil, errors.New("Unrecognize type of metric and ethtool command options to execute")
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
		_, dev, typeOf, metric := parseName(mt.Namespace())
		ns := joinNamespace([]string{dev, typeOf, metric})
		val, ok := metrics[ns]
		if !ok {
			return nil, fmt.Errorf("unknown %s stat: %s on interface %s", typeOf, metric, dev)
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
	tags := map[string]string{}

	// metrics from command `ethtool -S <interface>`
	valid, err := p.mc.ValidMetrics()
	if err != nil {
		return nil, err
	}

	for dev, metrics := range valid {
		for _, metric := range metrics {
			ns := makeName(dev, statNicLabel, metric)
			tags["driver"], _, _, _ = parseName(ns)
			mts = append(mts, plugin.PluginMetricType{Namespace_: ns, Tags_: tags})
		}
	}

	// register dump raw metrics from command `ethtool -d <interface>`
	validRegDump, err := p.mc.ValidRegDumpMetrics()
	if err != nil {
		return nil, err
	}

	for dev, metrics := range validRegDump {
		for _, metric := range metrics {
			ns := makeName(dev, statRegLabel, metric)
			tags["driver"], _, _, _ = parseName(ns)
			mts = append(mts, plugin.PluginMetricType{Namespace_: ns, Tags_: tags})
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
