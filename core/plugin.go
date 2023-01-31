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

import "os"

var plugins = map[string]Plugin{}

// Plugin represents a protoc plugin (e.g. --go_out, --go-grpc_out). Plugins should implement
// this interface and be registered with RegisteredPlugin in order for it to be available
// in the configuration file.
type Plugin interface {

	// Executable returns the path to the plugin executable, or an error if it exists.
	// The directory provided is the root of the plugin directory.
	Executable(version string, dir string) (string, error)

	// Install installs the plugin in the directory provided and returns an error if it exists.
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

// InstallPlugin installs the plugin using the given config.
func InstallPlugin(plugin *PluginConfig) error {
	p, ok := GetPlugin(plugin.Name)
	if !ok {
		return ErrPluginNotSupported(plugin.Name)
	}
	// First let's check and see if the plugin directory already exists.
	// If not we'll create one.
	dir := pluginPath(plugin.Name, plugin.Version)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}
	return p.Install(plugin.Version, dir)
}

// RegisterPlugin registers a plugin.
func RegisterPlugin(p Plugin) {
	plugins[p.Name()] = p
}
