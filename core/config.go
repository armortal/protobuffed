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
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Protobuf *ProtobufConfig `json:"protobuf"`
	Imports  []string        `json:"imports"`
	Inputs   []string        `json:"inputs"`
	Plugins  []*PluginConfig `json:"plugins"`
}

type ProtobufConfig struct {
	Version string `json:"version"`
}

type PluginConfig struct {
	Name    string `json:"name"`
	Options string `json:"options"`
	Output  string `json:"output"`
	Version string `json:"version"`
}

func ReadConfig(file string) (*Config, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(f, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %s", err.Error())
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.Protobuf == nil {
		return fmt.Errorf("config: protobuf configuration not provided")
	}

	if c.Protobuf.Version == "" {
		return errors.New("config: protobuf version not provided")
	}

	for _, p := range c.Plugins {
		if p.Version == "" {
			return fmt.Errorf("config: %s plugin version not provided", p.Name)
		}
	}
	return nil
}
