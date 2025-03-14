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
	"fmt"
	"path/filepath"
	"runtime"
)

var plugins = map[string]Plugin{}

// Plugin represents a protoc plugin (e.g. --go_out, --go-grpc_out). Plugins should implement
// this interface and be registered with RegisteredPlugin in order for it to be available
// in the configuration file.
type Plugin interface {
	// Install installs the plugin in the directory provided and returns an error if it exists. Implementations
	// must install the binary at dst/bin/protoc-gen-${PLUGIN_NAME}.
	Install(version string, dst string) error

	// Name returns the plugin name.
	Name() string
}

// GetPlugin returns the plugin associated with the name. Custom plugins may be registered with
// RegisterPlugin.
func GetPlugin(name string) (Plugin, bool) {
	p, ok := plugins[name]
	return p, ok
}

// RegisterPlugin registers a plugin.
func RegisterPlugin(p Plugin) {
	plugins[p.Name()] = p
}

func pluginBinaryName(name string) string {
	bin := fmt.Sprintf("protoc-gen-%s", name)
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	return bin
}

func pluginBinaryPath(cache string, name string, version string) string {
	return filepath.Join(pluginVersionPath(cache, name, version), "bin", pluginBinaryName(name))
}

func pluginVersionPath(cache string, name string, version string) string {
	return filepath.Join(cache, "plugins", name, version)
}
