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
	"os"
	"regexp"
	"strings"
)

var (
	// regular expression for retrieving NIC stats ("ethtool -S" output)
	nicRegex = regexp.MustCompile(`^\s*(\w+)\s*:\s*(\w+)\s*$`)

	// regular expression for retrieving driver info ("ethtool -i" output)
	driverRegex = regexp.MustCompile(`^\s*driver\s*:\s*(\w+)\s*$`)

	// regular expression for retrieving register dump info ("ethtool -d" output),
	// e.g. line:	"	0x040D0: tpr (Total Packets Received) 0x077A9D0E "
	regDumpRegex = regexp.MustCompile(`^\s*0x(\w+)\s*:\s*(\w+)\s*[(]*(.*)*[)]\s*(\w+)\s*$`)

	// universal regular expression,
	// e.g. line:	" Link Speed : 10G "
	uniRegex = regexp.MustCompile(`^\s*(.*)\s+:\s+(.*)\s*$`) // universal regex
)

// GetDriverInfo returns name of net driver that can be gathered from ethtool using -i switch.
func (tc *ToolCollector) GetDriverInfo(iface string) (string, error) {
	lines, err := tc.Tool.Execute("-i", iface)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get driver information")
		return "", err
	}

	for _, line := range lines {
		s := strings.ToLower(line)
		if strings.Contains(s, "driver:") {
			match := driverRegex.FindStringSubmatch(s)
			return match[len(match)-1], nil
		}
	}

	return "", errors.New("no driver info available")
}

// GetNicStats returns NIC statistics that can be gathered from ethtool using -S switch.
func (tc *ToolCollector) GetNicStats(iface string) (map[string]string, error) {
	lines, err := tc.Tool.Execute("-S", iface)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get NIC statistics")
		return nil, err
	}

	results := map[string]string{}
	for i, line := range lines {
		if i == 0 {
			if !strings.HasPrefix(line, "NIC statistics:") {
				return nil, fmt.Errorf("invalid output: first line = %s", line)
			}
			continue
		}

		match := nicRegex.FindStringSubmatch(line)
		if match == nil {
			return nil, fmt.Errorf("invalid output: %s", line)
		}

		results[match[1]] = match[2]
	}

	return results, nil
}

// GetRegDump returns register dump that can be gathered from ethtool using -d switch.
func (tc *ToolCollector) GetRegDump(iface string) (map[string]string, error) {

	lines, err := tc.Tool.Execute("-d", iface)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get register dump information")
		return nil, err
	}

	results := map[string]string{}

	for _, line := range lines {

		// line with metric has to contain char ":", skip if not
		if !strings.ContainsAny(line, ":") {
			continue
		}

		match := regDumpRegex.FindStringSubmatch(line)

		if len(match) > 3 {
			results[formatMetricName(match[2], true)] = match[4]
			continue
		}

		// it might be global info, use universal regex
		match = uniRegex.FindStringSubmatch(line)
		if len(match) > 2 {
			// make metric names in this case to lower
			results[formatMetricName(strings.ToLower(match[1]), true)] = match[2]
		} else {
			continue
		}
	}

	return results, nil
}

// GetDomStats returns statistics that can be gathered from ethtool using -m option.
func (tc *ToolCollector) GetDomStats(iface string) (map[string]string, error) {

	lines, err := tc.Tool.Execute("-m", iface)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get digital optical monitoring information")
		return nil, err
	}

	results := map[string]string{}

	for _, line := range lines {

		// line with metric has to contain char ":", skip if not
		if !strings.ContainsAny(line, ":") {
			continue
		}
		// use universal regex
		match := uniRegex.FindStringSubmatch(line)

		if match == nil {
			return nil, fmt.Errorf("invalid output: %s", line)
		}

		if len(match) > 2 {
			// make metric names in this case to lower
			metric := formatMetricName(strings.ToLower(match[1]), false)

			if len(results[metric]) != 0 {
				// add comma, combined metric's value (e.g. for dom/transceiver_type)
				results[metric] += ", "
			}
			results[metric] += match[2]

		} else {
			continue
		}
	}

	return results, nil
}

// formatMetricName prepares metric name what includes replacing space with char "_", removing brackets
// and also removing text in brackets if `removeTextInBrackets` is set to true
func formatMetricName(name string, removeTextInBrackets bool) string {
	// check if opening bracket occures in name
	index := strings.Index(name, "(")
	if index > 0 {
		if removeTextInBrackets {
			// remove text in brackets
			name = name[:index]
		} else {
			// remove brackets, leave text
			name = strings.Replace(name, "(", "", -1)
			name = strings.Replace(name, ")", "", -1)
		}
	}
	// trim white spaces in name
	name = strings.TrimSpace(name)
	// replace spaces, commas and slashes to "_"
	name = strings.Replace(strings.Replace(strings.Replace(name, "/", "_", -1), ",", "_", -1), " ", "_", -1)

	// remove double underline
	return strings.Replace(name, "__", "_", -1)
}
