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

package core

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func protobufVersionPath(cache string, version string) string {
	return filepath.Join(cache, "protobuf", version)
}

func protobufBinaryName() string {
	f := "protoc"
	if runtime.GOOS == "windows" {
		f += ".exe"
	}
	return f
}
func protobufBinaryPath(cache string, version string) string {
	return filepath.Join(protobufVersionPath(cache, version), "bin", protobufBinaryName())
}

func protobufArchiveName(version string) (string, error) {
	// Get the base filename.
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "win64"
		default:
			return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "protobuf", version)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "protobuf", version)
		}
	case "darwin":
		platform = "osx-universal_binary"
	default:
		return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "protobuf", version)
	}

	return fmt.Sprintf("protoc-%s-%s.zip", version, platform), nil
}
