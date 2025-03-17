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

package protocgenjs

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
			platform = "win64"
		case "arm64":
			platform = "win64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "osx-x86_64"
		case "arm64":
			platform = "osx-aarch_64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	release := fmt.Sprintf("protobuf-javascript-%s-%s.zip", strings.TrimPrefix(version, "v"), platform)
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-javascript/releases/download/%s/%s", version, release)

	r, err := http.Get(url)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := util.ExtractZipFromBytes(b, dir.Path()); err != nil {
		return err
	}

	return nil

}
