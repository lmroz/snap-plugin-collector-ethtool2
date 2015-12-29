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
	"os/exec"
	"strings"
)

const command = "ethtool"

// Execute() executes ethtool on local machine and filters output to return
// list of non-empty lines.
func (self *LocalExecutor) Execute(option, iface string) ([]string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		path = "/sbin/ethtool"
	}
	outBytes, err := exec.Command(path, option, iface).Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(outBytes), "\n")
	output := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.Trim(line, "\r\n")
		if line != "" {
			output = append(output, line)
		}
	}

	return output, nil
}
