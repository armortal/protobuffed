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

package main

import (
	"fmt"
	"os"

	"github.com/armortal/protobuffed/cmd/protobuffed/generate"
	cmdinit "github.com/armortal/protobuffed/cmd/protobuffed/init"
	"github.com/armortal/protobuffed/cmd/protobuffed/install"
	"github.com/armortal/protobuffed/cmd/protobuffed/run"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "protobuffed",
		Short: "Protocol buffers buffed up. A lightweight tool for managing your protobuf projects.",
	}

	cmd.PersistentFlags().StringP("file", "f", "protobuffed.json", "Path of the configuration file")

	cmd.AddCommand(cmdinit.Command())
	cmd.AddCommand(generate.Command())
	cmd.AddCommand(install.Command())
	cmd.AddCommand(run.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
