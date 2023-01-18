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
package plugingo

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/armortal/protobuffed/core/errors"
	"github.com/armortal/protobuffed/util"
)

type Plugin struct{}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Install(version string, dst string) error {
	release, err := release(version)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-go/releases/download/%s/%s", version, release)

	archive := filepath.Join(dst, release)
	if err := util.Download(url, archive); err != nil {
		return err
	}

	if strings.HasSuffix(release, ".zip") {
		if err := util.ExtractZip(archive, dst); err != nil {
			return err
		}
	} else {
		if err := util.ExtractTarGz(archive, dst); err != nil {
			return err
		}
	}
	return nil
}

func (p *Plugin) Executable(version string, dst string) (string, error) {
	return filepath.Join(dst, "protoc-gen-go"), nil
}

func (p *Plugin) Extract(version string, dst string) error {
	return nil
}

func (p *Plugin) Name() string {
	return "go"
}

func release(version string) (string, error) {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows.amd64.zip"
		case "arm64":
			platform = "windows.arm64.zip"
		default:
			return "", errors.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux.amd64.tar.gz"
		case "arm64":
			platform = "linux.arm64.tar.gz"
		default:
			return "", errors.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin.amd64.tar.gz"
		case "arm64":
			platform = "darwin.arm64.tar.gz"
		default:
			return "", errors.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	default:
		return "", errors.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
	}

	return fmt.Sprintf("protoc-gen-go.%s.%s", version, platform), nil
}
