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

package protocgengrpcweb

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/armortal/protobuffed"
	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/util"
)

type Dependency struct{}

func (d *Dependency) Install(ctx context.Context, dir *cache.Directory, version string) error {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows-x86_64.exe"
		case "arm64":
			platform = "windows-aarch64.exe"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin-x86_64"
		case "arm64":
			platform = "darwin-aarch64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	v := strings.TrimPrefix(version, "v")

	release := fmt.Sprintf("protoc-gen-grpc-web-%s-%s", v, platform)

	url := fmt.Sprintf("https://github.com/grpc/grpc-web/releases/download/%s/%s", v, release)

	if err := os.MkdirAll(filepath.Join(dir.Path(), "bin"), 0700); err != nil {
		return err
	}

	output := filepath.Join(dir.Path(), "bin", "protoc-gen-grpc-web")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	if err := util.Download(url, output); err != nil {
		return err
	}

	return nil
}
