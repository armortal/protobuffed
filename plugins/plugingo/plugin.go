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

package plugingo

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

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
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-go/releases/download/v%s/%s", version, release)

	archive := filepath.Join(dst, release)
	if err := util.Download(url, archive); err != nil {
		return err
	}

	bin := filepath.Join(dst, "bin")
	// Create the bin folder
	if err := os.MkdirAll(bin, 0700); err != nil {
		return err
	}

	if strings.HasSuffix(release, ".zip") {
		if err := util.ExtractZip(archive, bin); err != nil {
			return err
		}
	} else {
		if err := util.ExtractTarGz(archive, bin); err != nil {
			return err
		}
	}
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
			return "", errRuntimeNotSupported(version)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux.amd64.tar.gz"
		case "arm64":
			platform = "linux.arm64.tar.gz"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin.amd64.tar.gz"
		case "arm64":
			platform = "darwin.arm64.tar.gz"
		default:
			return "", errRuntimeNotSupported(version)
		}
	default:
		return "", errRuntimeNotSupported(version)
	}

	return fmt.Sprintf("protoc-gen-go.v%s.%s", version, platform), nil
}

func errRuntimeNotSupported(version string) error {
	return core.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "go", version)
}
