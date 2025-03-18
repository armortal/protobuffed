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

package operation

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/armortal/protobuffed/config"
)

func Run(ctx context.Context, cfg *config.Config, script string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd", "/c")
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c")
	}

	s, ok := cfg.Scripts[script]
	if !ok {
		return errors.New("script not found")
	}

	cmd.Args = append(cmd.Args, s)

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
