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

package plugingrpcweb

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const testVersion = "1.4.2"

func testDirectory(version string) (string, error) {
	wd, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, version), nil
}

func TestPlugin_Install(t *testing.T) {
	p := New()

	dir, err := testDirectory(testVersion)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}

	if err := p.Install(testVersion, dir); err != nil {
		t.Fatal(err)
	}

	binary := "protoc-gen-grpc-web"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	if _, err := os.Stat(filepath.Join(dir, "bin", binary)); os.IsNotExist(err) {
		t.Fatal("binary doesn't exist")
	}
}

func TestPlugin_Name(t *testing.T) {
	p := New()
	if p.Name() != "grpc-web" {
		t.Fatal()
	}
}
