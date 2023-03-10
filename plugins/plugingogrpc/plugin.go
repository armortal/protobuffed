// MIT License

// Copyright (c) 2023 ARMORTAL TECHNOLOGIES PTY LTD

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
