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

package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/armortal/protobuffed/util"
)

// Install will download the binary and extract it.
func InstallProtobuf(version string, dst string) error {
	// Check if it already exists
	if _, err := os.Stat(filepath.Join(dst, "bin", getProtobufBinaryName())); os.IsExist(err) {
		return nil
	}

	// Get the archive name
	archiveName, err := getProtobufArchiveName(version)
	if err != nil {
		return err
	}

	// We need to create the output directory
	// dir := filepath.Join(dst, version)
	// if err := os.Mkdir(dir, 0700); err != nil {
	// 	return err
	// }

	archive := filepath.Join(dst, archiveName)
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%s/%s", version, archiveName)
	if err := util.Download(url, archive); err != nil {
		return err
	}

	if err := util.ExtractZip(archive, dst); err != nil {
		return err
	}

	return nil
}

func getProtobufExecutable(version string) string {
	return filepath.Join(protobufPath(version), "bin", getProtobufBinaryName())
}

func getProtobufBinaryName() string {
	f := "protoc"
	if runtime.GOOS == "windows" {
		f += ".exe"
	}
	return f
}

func getProtobufArchiveName(version string) (string, error) {
	// Get the base filename.
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "win64"
		default:
			return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
		}
	case "darwin":
		platform = "osx-universal_binary"
	default:
		return "", ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH)
	}

	return fmt.Sprintf("protoc-%s-%s.zip", version, platform), nil
}
