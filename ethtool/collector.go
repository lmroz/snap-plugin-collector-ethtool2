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

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// regular expression for retrieving stats ("ethtool -S" output) and  driver info ("ethtool -i" output)
var statRegex = regexp.MustCompile(`^\s*(\w+)\s*:\s*(\w+)\s*$`)
var driverRegex = regexp.MustCompile(`^\s*driver\s*:\s*(\w+)\s*$`)

// ********		Regular expression for Register Dump, where
// ***			- examplary reg dump line contained register info:	0x040D0: tpr (Total Packets Received) 0x077A9D0E
// ***			- examplary reg dump line contained global info:	Link Speed: 10G
///////
var regDumpRegex = regexp.MustCompile(`^\s*0x(\w+)\s*:\s*(\w+)\s*[(]*(.*)*[)]\s*(\w+)\s*$`)
var regDumpRegexGlobal = regexp.MustCompile(`^\s*(.*)\s*:\s*(.*)\s*$`)

// GetDriverInfo returns name of network driver that can be gathered from ethtool using -i switch.
func (self *ToolCollector) GetDriverInfo(iface string) (string, error) {
	lines, err := self.Tool.Execute("-i", iface)

	if err != nil {
		return "", err
	}

	if len(lines) < 1 {
		return "", errors.New("no info output")
	}

	for _, line := range lines {
		s := strings.ToLower(line)
		if strings.Contains(s, "driver:") {
			match := driverRegex.FindStringSubmatch(s)
			return match[len(match)-1], nil
		}
	}

	return "", errors.New("no driver info")
}

// GetStats returns statistics that can be gathered from ethtool using -S switch.
func (self *ToolCollector) GetStats(iface string) (map[string]string, error) {
	lines, err := self.Tool.Execute("-S", iface)
	if err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, errors.New("no output")
	}

	results := map[string]string{}
	for i, line := range lines {
		if i == 0 {
			if !strings.HasPrefix(line, "NIC statistics:") {
				return nil, fmt.Errorf("invalid output: first line = %s", line)
			}
			continue
		}

		match := statRegex.FindStringSubmatch(line)
		if match == nil {
			return nil, fmt.Errorf("invalid output: %s", line)
		}

		results[match[1]] = match[2]
	}

	return results, nil
}

// GetRegDump returns register dump that can be gathered from ethtool using -d switch.
func (self *ToolCollector) GetRegDump(iface string) (map[string]string, error) {

	lines, err := self.Tool.Execute("-d", iface)
	if err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, errors.New("no output")
	}

	results := map[string]string{}

	for _, line := range lines {

		// line with metric has to contain char ":", skip if not
		if !strings.ContainsAny(line, ":") {
			continue
		}

		match := regDumpRegex.FindStringSubmatch(line)
		if len(match) > 3 {
			results[formatMetricName(match[2])] = match[4]
			continue
		}

		// it might be global info
		match = regDumpRegexGlobal.FindStringSubmatch(line)
		if len(match) > 2 {
			// make metric names in this case to lower
			results[formatMetricName(strings.ToLower(match[1]))] = match[2]
		} else {
			continue
		}
	}

	return results, nil
}

// formatMetricName prepares metric name what includes removing text in brackets and replacing space with char "_"
func formatMetricName(name string) string {
	// check if brackets occure in name
	index := strings.Index(name, "(")
	if index > 0 {
		// remove text in brackets
		name = name[:index]
	}

	// replace space to "_" and trim white spaces in name
	return strings.Replace(strings.TrimSpace(name), " ", "_", -1)
}
