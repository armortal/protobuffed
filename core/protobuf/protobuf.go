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

package protobuf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	err "github.com/armortal/protobuffed/core/errors"
	"github.com/armortal/protobuffed/util"
)

// Exists checks whether the given version exists.
func Exists(version string, dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, version)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Executable will return the path of the executable file or an error if there is one.
func Executable(version string, dir string) (string, error) {
	if !Exists(version, dir) {
		return "", errors.New("plugin does not exist")
	}
	return filepath.Join(dir, version, "bin", binary()), nil
}

// Install will download the binary and extract it.
func Install(version string, dst string) error {
	// Get the archive name
	release, err := release(version)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/%s/v%s", version, release)

	// We need to create the output directory
	dir := filepath.Join(dst, version)
	if err := os.Mkdir(dir, 0700); err != nil {
		return err
	}

	zip := filepath.Join(dir, "release.zip")
	if err := util.Download(url, zip); err != nil {
		return err
	}

	unzip := filepath.Join(dir)
	if err := util.ExtractZip(zip, unzip); err != nil {
		return err
	}

	if err := os.RemoveAll(zip); err != nil {
		return err
	}

	return nil
}

func binary() string {
	f := "protoc"
	if runtime.GOOS == "windows" {
		f += ".exe"
	}
	return f
}

func release(version string) (string, error) {
	// Get the base filename.
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "win64"
		default:
			return "", err.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return "", err.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "darwin":
		platform = "osx-universal_binary"
	default:
		return "", err.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
	}

	// The protoc binaries don't have a 'v' prefix for the version.
	v := version
	if strings.HasPrefix(version, "v") {
		v = strings.Split(version, "v")[1]
	}

	return fmt.Sprintf("protoc-%s-%s.zip", v, platform), nil
}
