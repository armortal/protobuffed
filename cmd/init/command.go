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

package init

import (
	"encoding/json"
	"os"

	"github.com/armortal/protobuffed/core"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := &core.Config{
				Protobuf: &core.ProtobufConfig{
					Version: "latest",
				},
				Imports: make([]string, 0),
				Inputs:  make([]string, 0),
				Plugins: make([]*core.PluginConfig, 0),
			}

			b, err := json.MarshalIndent(config, "", "    ")
			if err != nil {
				return err
			}

			f, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			if err := os.WriteFile(f, b, 0700); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
