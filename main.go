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

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/armortal/protobuffed/cmd/command"
	"github.com/armortal/protobuffed/cmd/generate"
	"github.com/armortal/protobuffed/cmd/install"
	"github.com/armortal/protobuffed/core"
	"github.com/armortal/protobuffed/plugins/plugingo"
	"github.com/armortal/protobuffed/plugins/plugingogrpc"
	"github.com/spf13/cobra"
)

func init() {
	core.RegisterPlugin(plugingo.New())
	core.RegisterPlugin(plugingogrpc.New())
}

func main() {
	cmd := &cobra.Command{
		Use:   "protobuffed",
		Short: "Protocol buffers buffed up. Making it easier to work with protobuf files and binaries",
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cmd.PersistentFlags().StringP("cache", "c", filepath.Join(home, ".protobuffed"), "The directory where binaries will be installed.")
	cmd.PersistentFlags().StringP("file", "f", "protobuffed.json", "The configuration file")

	cmd.AddCommand(generate.Command())
	cmd.AddCommand(install.Command())
	cmd.AddCommand(command.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
