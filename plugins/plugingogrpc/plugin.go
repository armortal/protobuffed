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

package plugingogrpc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/armortal/protobuffed/util"
)

type Plugin struct{}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Exists(version string, dir string) bool {
	v := filepath.Join(dir, version)
	if _, err := os.Stat(v); os.IsNotExist(err) {
		return false
	} else {
		// Check if the binary if there
		if _, err := os.Stat(filepath.Join(v, "bin", "protoc-gen-go-grpc")); os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (p *Plugin) Install(version string, dst string) error {
	// grpc-go don't provide the built binaries, we must download the source and we can run it directly.
	archive := filepath.Join(dst, fmt.Sprintf("%s.zip", release(version)))
	bin := filepath.Join(dst, "bin")
	binAbs, err := filepath.Abs(bin)
	if err != nil {
		return err
	}

	// Check existence of binary and remove all directories if it doesn't exist.
	if !p.Exists(version, dst) {
		if err := os.RemoveAll(dst); err != nil {
			return err
		}
		// Create our directories
		if err := os.MkdirAll(bin, 0700); err != nil {
			return err
		}

		// Download the repo.
		if err := util.Download(fmt.Sprintf("https://github.com/grpc/grpc-go/archive/refs/tags/v%s.zip", version), archive); err != nil {
			return err
		}

		if err := util.ExtractZip(archive, dst); err != nil {
			return err
		}

		output := filepath.Join(binAbs, "protoc-gen-go-grpc")
		if runtime.GOOS == "windows" {
			output += ".exe"
		}

		cmd := exec.Command("go", "build", "-o", output, ".")
		// We join the filename twice because archives from git creates the same subfolder with its contents
		// The unzipped contents don't have the v prefix
		cmd.Dir = filepath.Join(dst, release(version), "cmd", "protoc-gen-go-grpc")
		if err := cmd.Run(); err != nil {
			return err
		}

		if err := os.Chmod(output, 0700); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) Name() string {
	return "go-grpc"
}

func release(version string) string {
	return fmt.Sprintf("grpc-go-%s", version)
}
