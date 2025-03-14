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
	"os"
	"path/filepath"
	"testing"
)

func Test_installProtobuf(t *testing.T) {
	dir, err := os.MkdirTemp(".", "test*")
	if err != nil {
		t.Fatal(err)
	}

	if err := installProtobuf("21.12", dir); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(dir, "bin", protobufBinaryName())); os.IsNotExist(err) {
		t.Fatal("protobuf binary doesn't exist")
	}
}
