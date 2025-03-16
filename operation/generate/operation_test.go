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

package generate

import (
	"testing"
)

func TestCommand(t *testing.T) {
	// config := &Config{
	// 	Protobuf: &ProtobufConfig{
	// 		Version: "21.12",
	// 	},
	// 	Imports: []string{"test"},
	// 	Inputs:  []string{"test.proto"},
	// 	Plugins: []*PluginConfig{
	// 		{
	// 			Name:    "testplugin",
	// 			Version: "1.0.0",
	// 			Options: "paths=source_relative",
	// 			Output:  "test",
	// 		},
	// 	},
	// }

	// // register the test plugin
	// RegisterPlugin(&testPlugin{})

	// act, err := Command(config, "test")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// exp := fmt.Sprintf("test/protobuf/21.12/bin/%s --plugin=protoc-gen-testplugin=test/plugins/testplugin/1.0.0/bin/%s --testplugin_out=test --testplugin_opt=paths=source_relative --proto_path=test test.proto", protobufBinaryName(), pluginBinaryName("testplugin"))

	// if act.String() != exp {
	// 	t.Fatal()
	// }
}
