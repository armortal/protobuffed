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

	cache, err := filepath.Abs(cache)
	if err != nil {
		return err
	}

	// install protobuf
	pb := protobufVersionPath(cache, config.Protobuf.Version)
	if _, err := os.Stat(pb); os.IsNotExist(err) {
		if err := os.MkdirAll(pb, 0700); err != nil {
			return err
		}
		if err := installProtobuf(config.Protobuf.Version, pb); err != nil {
			return err
		}
	}

	// install the plugins
	for _, p := range config.Plugins {
		plugin, ok := GetPlugin(p.Name)
		if !ok {
			return ErrPluginNotSupported(p.Name)
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
