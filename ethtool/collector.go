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

var statRegex = regexp.MustCompile(`^\s*(\w+)\s*:\s*(\w+)\s*$`)

// Returns statistics that can be gathered from ethtool using -S switch.
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
