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
