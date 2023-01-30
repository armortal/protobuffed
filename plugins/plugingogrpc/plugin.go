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

package plugingogrpc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func (p *Plugin) Executable(version string, dir string) (string, error) {
	return filepath.Join(dir, "bin", "protoc-gen-go-grpc"), nil
}

func (p *Plugin) Install(version string, dst string) error {
	// grpc-go don't provide the built binaries, we must download the source and we can run it directly.
	archive := filepath.Join(dst, fmt.Sprintf("%s.zip", release(version)))
	bin := filepath.Join(dst, "bin")

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
		if err := util.Download(fmt.Sprintf("https://github.com/grpc/grpc-go/archive/refs/tags/%s.zip", version), archive); err != nil {
			return err
		}

		if err := util.ExtractZip(archive, dst); err != nil {
			return err
		}

		cmd := exec.Command("go", "build", "-o", filepath.Join(bin, "protoc-gen-go-grpc"), ".")
		// We join the filename twice because archives from git creates the same subfolder with its contents
		// The unzipped contents don't have the v prefix
		cmd.Dir = filepath.Join(dst, release(strings.Split(version, "v")[1]), "cmd", "protoc-gen-go-grpc")
		if err := cmd.Run(); err != nil {
			return err
		}

		ex, err := p.Executable(version, dst)
		if err != nil {
			return err
		}
		if err := os.Chmod(ex, 0700); err != nil {
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
