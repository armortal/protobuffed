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
package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

var home string

func init() {
	if h := os.Getenv("PROTOBUFFED_HOME"); h != "" {
		home = h
	} else {
		h, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		home = filepath.Join(h, ".protobuffed")
	}

	if err := create(); err != nil {
		panic(err)
	}
}

func create() error {
	if err := os.MkdirAll(filepath.Join(home, "plugins"), 0700); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(home, "protobuf"), 0700); err != nil {
		return err
	}
	return nil
}

func Protobuf() string {
	return filepath.Join(home, "protobuf")
}

func Plugins() string {
	return filepath.Join(home, "plugins")
}

func Plugin(name string, version string) string {
	return filepath.Join(Plugins(), fmt.Sprintf("protoc-gen-%s", name), version)
}
