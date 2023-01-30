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

import (
	"os"

	"github.com/armortal/protobuffed/core/errors"
	"github.com/armortal/protobuffed/core/protobuf"
	"github.com/armortal/protobuffed/core/storage"
)

func InstallProtobuf(version string) error {
	if !protobuf.Exists(version, storage.Protobuf()) {
		if err := protobuf.Install(version, storage.Protobuf()); err != nil {
			return err
		}
	}
	return nil
}

// Download the plugin.
func InstallPlugin(plugin *PluginConfig) error {
	p, ok := GetPlugin(plugin.Name)
	if !ok {
		return errors.ErrPluginNotSupported(plugin.Name)
	}
	// First let's check and see if the plugin directory already exists.
	// If not we'll create one.
	dir := storage.Plugin(plugin.Name, plugin.Version)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}
	return p.Install(plugin.Version, dir)
}
