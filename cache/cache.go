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

package cache

import (
	"errors"
	"os"
	"path/filepath"
)

type Cache struct {
	path string
}

func (c *Cache) RemoveAll() error {
	return os.RemoveAll(c.path)
}

// func (c *Cache) load() error {
// 	if _, err := os.Stat(c.path); os.IsNotExist(err) {
// 		if err := os.MkdirAll(c.path, 0700); err != nil {
// 			return err
// 		}
// 		return nil
// 	}

// 	entries, err := os.ReadDir(c.path)
// 	if err != nil {
// 		return err
// 	}

// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			if entry.Name() == "bin" {

// 			}
// 			if entry.Name() == "dependencies" {
// 			}
// 		}
// 	}
// 	return nil
// }

func New() (*Cache, error) {
	path, err := filepath.Abs(".protobuffed")
	if err != nil {
		return nil, err
	}

	if err := os.RemoveAll(path); err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0700); err != nil {
			return nil, err
		}
	}

	return &Cache{
		path: path,
	}, nil
}

func Load() (*Cache, error) {
	path, err := filepath.Abs(".protobuffed")
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("run protobuffed install")
	}

	return &Cache{
		path: path,
	}, nil
}

func (c *Cache) Directory(name string) *Directory {
	path := filepath.Join(c.path, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0700); err != nil {
			panic(err)
		}
	}

	return &Directory{
		name: name,
		path: path,
	}
}

type Directory struct {
	name string
	path string
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Path() string {
	return d.path
}

func (d *Directory) Write(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filepath.Join(d.path, name), data, perm)
}

func (d *Directory) Exists(name string) (bool, error) {
	entries, err := os.ReadDir(d.path)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if entry.Name() == name {
			return true, nil
		}
	}
	return false, nil
}
