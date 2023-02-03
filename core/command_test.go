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
	"testing"
)

func TestCommand(t *testing.T) {
	config := &Config{
		Protobuf: &ProtobufConfig{
			Version: "21.12",
		},
		Imports: []string{"test"},
		Inputs:  []string{"test.proto"},
		Plugins: []*PluginConfig{
			{
				Name:    "testplugin",
				Version: "1.0.0",
				Options: "paths=source_relative",
				Output:  "test",
			},
		},
	}

	// register the test plugin
	RegisterPlugin(&testPlugin{})

	act, err := Command(config, "test")
	if err != nil {
		t.Fatal(err)
	}

	exp := fmt.Sprintf("test/protobuf/21.12/bin/%s --plugin=protoc-gen-testplugin=test/plugins/testplugin/1.0.0/bin/%s --testplugin_out=test --testplugin_opt=paths=source_relative --proto_path=test test.proto", protobufBinaryName(), pluginBinaryName("testplugin"))

	if act.String() != exp {
		t.Fatal()
	}
}
