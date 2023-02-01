// MIT License

// Copyright (c) 2023 Armortal Technologies Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package core

import (
	"fmt"
	"os/exec"
)

func Command(c *Config) (*exec.Cmd, error) {
	b := getProtobufExecutable(c.Version)
	cmd := exec.Command(b)

	for _, plugin := range c.Plugins {
		p, ok := GetPlugin(plugin.Name)
		if !ok {
			return nil, ErrPluginNotSupported(plugin.Name)
		}

		binary, err := p.Executable(plugin.Version, pluginPath(plugin.Name, plugin.Version))
		if err != nil {
			return nil, err
		}

		cmd.Args = append(cmd.Args, fmt.Sprintf("--plugin=protoc-gen-%s=%s", p.Name(), binary))
		cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_out=%s", plugin.Name, plugin.Output))
		if plugin.Options != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_opt=%s", plugin.Name, plugin.Options))
		}

	}

	for _, i := range c.Imports {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--proto_path=%s", i))
	}

	cmd.Args = append(cmd.Args, c.Inputs...)

	return cmd, nil
}
