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

	"github.com/armortal/protobuffed/cmd/generate"
	cmdinit "github.com/armortal/protobuffed/cmd/init"
	"github.com/armortal/protobuffed/cmd/install"
	"github.com/armortal/protobuffed/cmd/print"
	"github.com/armortal/protobuffed/core"
	"github.com/armortal/protobuffed/plugins/plugingo"
	"github.com/armortal/protobuffed/plugins/plugingogrpc"
	"github.com/armortal/protobuffed/plugins/plugingrpcgateway"
	"github.com/armortal/protobuffed/plugins/plugingrpcweb"
	"github.com/armortal/protobuffed/plugins/pluginjs"
	"github.com/spf13/cobra"
)

func init() {
	core.RegisterPlugin(plugingo.New())
	core.RegisterPlugin(plugingogrpc.New())
	core.RegisterPlugin(plugingrpcgateway.New())
	core.RegisterPlugin(plugingrpcweb.New())
	core.RegisterPlugin(pluginjs.New())
}

func main() {
	cmd := &cobra.Command{
		Use:   "protobuffed",
		Short: "Protocol buffers buffed up. A lightweight tool for managing your protobuf projects.",
	}

	cmd.PersistentFlags().StringP("cache", "c", ".protobuffed", "Path where binaries will be installed and executed from")
	cmd.PersistentFlags().StringP("file", "f", "protobuffed.json", "Path of the configuration file")

	cmd.AddCommand(generate.Command())
	cmd.AddCommand(cmdinit.Command())
	cmd.AddCommand(install.Command())
	cmd.AddCommand(print.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
