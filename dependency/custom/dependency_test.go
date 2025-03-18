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

package custom

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/armortal/protobuffed/cache"
)

func TestDependency_Install(t *testing.T) {
	dep := Dependency{}

	c, err := cache.New()
	if err != nil {
		t.Fatal(err)
	}

	dir := c.Directory("protobuffed")

	if err := dep.Install(context.Background(), dir, "https://github.com/armortal/protobuffed/archive/refs/heads/main.zip"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(dir.Path(), "main.zip")); os.IsNotExist(err) {
		t.Fatal("zip doesn't exist")
	}
}
