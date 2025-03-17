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
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

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
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "protoc")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "protoc")
	}

	// Set PATH variables
	path := ""
	// Add all dependencies' bin directories to the PATH
	for name := range cfg.Dependencies {
		p, err := filepath.Abs(fmt.Sprintf(".protobuffed/%s/bin", name))
		if err != nil {
			return err
		}
		if path != "" {
			path = fmt.Sprintf("%s:%s", path, p)
		} else {
			path = p
		}
	}

	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s:%s", path, os.Getenv("PATH")))

	// fmt.Println(cmd.Path)
	for _, plugin := range cfg.Plugins {
		args := cmd.Args[len(cmd.Args)-1]
		args = fmt.Sprintf("%s --%s_out=%s", args, plugin.Name, plugin.Output)
		if plugin.Options != "" {
			args = fmt.Sprintf("%s --%s_opt=%s", args, plugin.Name, plugin.Options)
		}
		cmd.Args[len(cmd.Args)-1] = args
	}

	for _, i := range cfg.Imports {
		args := cmd.Args[len(cmd.Args)-1]
		args = fmt.Sprintf("%s --proto_path=%s", args, i)
		cmd.Args[len(cmd.Args)-1] = args
	}

	for _, i := range cfg.Inputs {
		args := cmd.Args[len(cmd.Args)-1]
		args = fmt.Sprintf("%s %s", args, i)
		cmd.Args[len(cmd.Args)-1] = args
	}

	fmt.Printf("Executing -> %s\n", cmd.Args[2])

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
