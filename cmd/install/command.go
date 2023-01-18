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
package install

import (
	"fmt"

	"github.com/armortal/protobuffed/core"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install binaries",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read in the config.
			f, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}
			cfg, err := core.ReadConfig(f)
			if err != nil {
				return err
			}

			return Exec(cfg)
		},
	}

	return cmd
}

func Exec(c *core.Config) error {
	// Install protobuf.
	fmt.Printf("Installing protobuf@%s\n", c.Version)
	if err := core.InstallProtobuf(c.Version); err != nil {
		return err
	}

	for _, plugin := range c.Plugins {
		fmt.Printf("Installing protoc-gen-%s@%s\n", plugin.Name, plugin.Version)
		if err := core.InstallPlugin(plugin); err != nil {
			return err
		}
	}
	return nil
}
