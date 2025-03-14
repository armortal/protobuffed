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

// testPlugin used for testing purposes only so we can register this with the core.
type testPlugin struct{}

func (p *testPlugin) Name() string {
	return "testplugin"
}

func (p *testPlugin) Install(version string, dir string) error {
	return nil
}

// func TestPlugins(t *testing.T) {
// 	p := &testPlugin{}

// 	if _, ok := GetPlugin(p.Name()); ok {
// 		t.Fatal("plugin should not exist")
// 	}

// 	RegisterPlugin(p)

// 	v, ok := GetPlugin(p.Name())
// 	if !ok {
// 		t.Fatal("plugin should exist")
// 	}

// 	if v != p {
// 		t.Fatal("plugin mismatch")
// 	}
// }
