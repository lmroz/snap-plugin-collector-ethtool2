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

package ethtool

// Executor executes ethtool with given option on given interface
type Executor interface {
	Execute(option, iface string) ([]string, error)
}

// Collector performs data collection from given interface
type Collector interface {
	GetNicStats(iface string) (map[string]string, error)
	GetRegDump(iface string) (map[string]string, error)
	GetDomStats(iface string) (map[string]string, error)
	GetDriverInfo(iface string) (string, error)
}

// LocalExecutor executes ethtool on local machine
type LocalExecutor struct {
}

// ToolCollector performs data collection from ethtool
type ToolCollector struct {
	Tool Executor
}
