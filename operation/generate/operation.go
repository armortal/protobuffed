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

package generate

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"

	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/config"
)

func Execute(ctx context.Context, cfg *config.Config, cache *cache.Cache) error {
	return new(Operation).Execute(ctx, cfg, cache)
}

type Operation struct{}

// Execute will generate the source code using the given configuration file.
// Calling this function will assume the binaries have been installed.
// If an error occurs, it will be returned.
func (o *Operation) Execute(ctx context.Context, cfg *config.Config, cache *cache.Cache) error {
	cmd := exec.Command(".protobuffed/protoc/bin/protoc")

	for _, plugin := range cfg.Plugins {
		cmd.Args = append(cmd.Args,
			fmt.Sprintf("--plugin=%s", fmt.Sprintf(".protobuffed/%s/bin/%s", plugin.Name, plugin.Name)))
		cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_out=%s", plugin.Name, plugin.Output))
		if plugin.Options != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_opt=%s", plugin.Name, plugin.Options))
		}

	}

	for _, i := range cfg.Imports {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--proto_path=%s", i))
	}

	cmd.Args = append(cmd.Args, cfg.Inputs...)

	// return cmd, nil
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	outScanner := bufio.NewScanner(stdout)
	go func() {
		for outScanner.Scan() {
			fmt.Printf("%s\n", outScanner.Text())
		}
	}()

	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			fmt.Printf("%s\n", errScanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func command(config *config.Config, cache string) (*exec.Cmd, error) {
	// if err := config.validate(); err != nil {
	// 	return nil, err
	// }

	// cmd := exec.Command(protobufBinaryPath(cache, config.Protobuf.Version))

	// for _, plugin := range config.Plugins {

	// 	cmd.Args = append(cmd.Args,
	// 		fmt.Sprintf("--plugin=protoc-gen-%s=%s", plugin.Name, pluginBinaryPath(cache, plugin.Name, plugin.Version)))
	// 	cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_out=%s", plugin.Name, plugin.Output))
	// 	if plugin.Options != "" {
	// 		cmd.Args = append(cmd.Args, fmt.Sprintf("--%s_opt=%s", plugin.Name, plugin.Options))
	// 	}

	// }

	// for _, i := range config.Imports {
	// 	cmd.Args = append(cmd.Args, fmt.Sprintf("--proto_path=%s", i))
	// }

	// cmd.Args = append(cmd.Args, config.Inputs...)

	// return cmd, nil
	return nil, nil
}
