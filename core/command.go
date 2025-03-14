// Copyright 2025 ARMORTAL TECHNOLOGIES PTY LTD

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"fmt"
	"os/exec"
)

func Command(config *Config, cache string) (*exec.Cmd, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	cmd := exec.Command(protobufBinaryPath(cache, config.Protobuf.Version))

	for _, plugin := range config.Plugins {

		cmd.Args = append(cmd.Args,
			fmt.Sprintf("--plugin=protoc-gen-%s=%s", plugin.Name, pluginBinaryPath(cache, plugin.Name, plugin.Version)))
		cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_out=%s", plugin.Name, plugin.Output))
		if plugin.Options != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_opt=%s", plugin.Name, plugin.Options))
		}

	}

	for _, i := range config.Imports {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--proto_path=%s", i))
	}

	cmd.Args = append(cmd.Args, config.Inputs...)

	return cmd, nil
}
