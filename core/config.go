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
