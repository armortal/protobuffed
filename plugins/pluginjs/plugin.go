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

package pluginjs

import (
	"fmt"
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
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-javascript/releases/download/v%s/%s", version, release)

	archive := filepath.Join(dst, release)
	if err := util.Download(url, archive); err != nil {
		return err
	}

	if err := util.ExtractZip(archive, dst); err != nil {
		return err
	}

	return nil

}

func release(version string) (string, error) {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "win64"
		case "arm64":
			platform = "win64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "osx-x86_64"
		case "arm64":
			platform = "osx-aarch_64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	default:
		return "", errRuntimeNotSupported(version)
	}

	return fmt.Sprintf("protobuf-javascript-%s-%s.zip", version, platform), nil
}

func (p *Plugin) Name() string {
	return "js"
}

func errRuntimeNotSupported(version string) error {
	return core.ErrRuntimeNotSupported(runtime.GOOS, runtime.GOARCH, "go", version)
}
