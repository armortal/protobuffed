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

package install

import (
	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/config"
	"github.com/armortal/protobuffed/operation/install"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install binaries",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			cfg, err := config.ReadFile(f)
			if err != nil {
				return err
			}

			c, err := cache.New()
			if err != nil {
				return err
			}

			return install.Execute(cmd.Context(), cfg, c)
		},
	}

	return cmd
}
