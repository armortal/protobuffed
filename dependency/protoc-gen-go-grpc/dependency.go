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

package protocgengogrpc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/util"
)

type Dependency struct{}

// Install installs the protoc-gen-go-grpc binary. This grpc-go repository doesn't provide the built binaries
// so we must download the source and we can run it directly.
func (d *Dependency) Install(ctx context.Context, dir *cache.Directory, version string) error {
	r, err := http.Get(fmt.Sprintf("https://github.com/grpc/grpc-go/archive/refs/tags/%s.zip", version))
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

	bin := filepath.Join(dir.Path(), "bin")
	if err := os.MkdirAll(bin, 0700); err != nil {
		return err
	}

	output := filepath.Join(bin, "protoc-gen-go-grpc")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", output, ".")
	// We join the filename twice because archives from git creates the same subfolder with its contents
	// The unzipped contents don't have the v prefix
	cmd.Dir = filepath.Join(dir.Path(), fmt.Sprintf("grpc-go-%s", strings.TrimPrefix(version, "v")), "cmd", "protoc-gen-go-grpc")
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := os.Chmod(output, 0700); err != nil {
		return err
	}

	return nil
}
