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

var plugins = map[string]Plugin{}

func GetPlugin(name string) (Plugin, bool) {
	p, ok := plugins[name]
	return p, ok
}

func RegisterPlugin(p Plugin) {
	plugins[p.Name()] = p
}

type Plugin interface {
	// Executable returns the path to the plugin executable, or an error if it exists.
	// The directory provided is the root of the plugin directory.
	Executable(version string, dir string) (string, error)

	// Install installs the plugin in the directory provided. Implementions should create
	// separate directories for each version inside this directory.
	Install(version string, dst string) error

	// Name returns the plugin name.
	Name() string
}
