// MIT License

// Copyright (c) 2023 ARMORTAL TECHNOLOGIES PTY LTD

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
			return "", errRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return "", errRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "darwin":
		platform = "osx-universal_binary"
	default:
		return "", errRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
	}

	return fmt.Sprintf("protoc-%s-%s.zip", version, platform), nil
}
