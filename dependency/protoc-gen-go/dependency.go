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

package protocgengo

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
	// Get the platform.
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows.amd64.zip"
		case "arm64":
			platform = "windows.arm64.zip"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux.amd64.tar.gz"
		case "arm64":
			platform = "linux.arm64.tar.gz"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin.amd64.tar.gz"
		case "arm64":
			platform = "darwin.arm64.tar.gz"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	// Get the release and url.
	release := fmt.Sprintf("protoc-gen-go.%s.%s", version, platform)
	releasePath := dir.Join(release)

	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-go/releases/download/%s/%s", version, release)

	// Download and extract.
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := dir.Write(release, b, 0700); err != nil {
		return err
	}

	// Create the bin folder
	if err := os.Mkdir(dir.Join("bin"), 0700); err != nil {
		return err
	}

	// Decompress all .gz files.
	if filepath.Ext(release) == ".gz" {
		if err := util.DecompressGz(releasePath, dir.Join(strings.TrimSuffix(release, ".gz"))); err != nil {
			return err
		}
	}

	// Extract the archive.
	if err := util.Extract(dir.Join(strings.TrimSuffix(release, ".gz")), dir.Join("bin")); err != nil {
		return err
	}

	if err := os.Chmod(filepath.Join(dir.Path(), "bin", "protoc-gen-go"), 0700); err != nil {
		return err
	}

	return nil
}
