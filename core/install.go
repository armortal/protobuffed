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
	"os"
	"path/filepath"

	"github.com/armortal/protobuffed/util"
)

// Install will install protobuf and plugins using the given configuration and store it
// at destination provided.
func Install(config *Config, cache string) error {
	// validate the config
	if err := config.validate(); err != nil {
		return err
	}

	// install protobuf
	pb := protobufVersionPath(cache, config.Version)
	if _, err := os.Stat(pb); os.IsNotExist(err) {
		if err := os.MkdirAll(pb, 0700); err != nil {
			return err
		}
		if err := installProtobuf(config.Version, pb); err != nil {
			return err
		}
	}

	// install the plugins
	for _, p := range config.Plugins {
		plugin, ok := GetPlugin(p.Name)
		if !ok {
			return errPluginNotSupported(p.Name)
		}

		// create the folder and download the plugin if it doesn't already exist
		dir := pluginVersionPath(cache, p.Name, p.Version)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0700); err != nil {
				return err
			}

			if err := plugin.Install(p.Version, dir); err != nil {
				return err
			}
		}
	}

	return nil
}

// Download and install protobuf for the specific version at the directory provided.
func installProtobuf(version string, dst string) error {
	// Get the archive name
	archiveName, err := protobufArchiveName(version)
	if err != nil {
		return err
	}

	archive := filepath.Join(dst, archiveName)
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%s/%s", version, archiveName)
	if err := util.Download(url, archive); err != nil {
		return err
	}

	if err := util.ExtractZip(archive, dst); err != nil {
		return err
	}

	return nil
}
