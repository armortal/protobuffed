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

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Name         string            `json:"name"`
	Dependencies map[string]string `json:"dependencies"`
	Imports      []string          `json:"imports"`
	Inputs       []string          `json:"inputs"`
	Plugins      []*PluginConfig   `json:"plugins"`
}

type PluginConfig struct {
	Name    string `json:"name"`
	Options string `json:"options"`
	Output  string `json:"output"`
}

func ReadFile(file string) (*Config, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(f, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %s", err.Error())
	}

	return config, nil
}

func Read() (*Config, error) {
	return ReadFile("protobuffed.json")
}
