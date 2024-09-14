// MIT License

// Copyright (c) 2024 ARMORTAL TECHNOLOGIES PTY LTD

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

package plugingrpcgateway

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/armortal/protobuffed/core"
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

	// /v2.19.1/protoc-gen-grpc-gateway-v2.19.1-darwin-arm64
	url := fmt.Sprintf("https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v%s/%s", version, release)
	fmt.Println(url)
	bin := filepath.Join(dst, "bin")
	// Create the bin folder
	if err := os.MkdirAll(bin, 0700); err != nil {
		return err
	}

	output := filepath.Join(bin, "protoc-gen-grpc-gateway")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	if err := util.Download(url, output); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) Name() string {
	return "grpc-gateway"
}

func release(version string) (string, error) {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows-x86_64.exe"
		case "arm64":
			platform = "windows-arm64.exe"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-arm64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin-x86_64"
		case "arm64":
			platform = "darwin-arm64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	default:
		return "", errRuntimeNotSupported(version)
	}

	return fmt.Sprintf("protoc-gen-grpc-gateway-v%s-%s", version, platform), nil
}

func errRuntimeNotSupported(version string) error {
	return core.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "go", version)
}
