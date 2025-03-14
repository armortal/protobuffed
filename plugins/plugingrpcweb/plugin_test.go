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
