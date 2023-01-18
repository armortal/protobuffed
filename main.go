package main

import (
	"fmt"
	"os"

	"github.com/armortal/protobuffed/cmd/generate"
	"github.com/armortal/protobuffed/cmd/install"
	"github.com/armortal/protobuffed/cmd/print"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.PersistentFlags().StringP("file", "f", "protobuffed.json", "The configuration file")

	cmd.AddCommand(generate.Command())
	cmd.AddCommand(install.Command())
	cmd.AddCommand(print.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
