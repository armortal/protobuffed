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
