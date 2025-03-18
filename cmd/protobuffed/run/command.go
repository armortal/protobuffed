package run

import (
	"github.com/armortal/protobuffed/config"
	"github.com/armortal/protobuffed/operation"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a script",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			cfg, err := config.ReadFile(f)
			if err != nil {
				return err
			}

			return operation.Run(cmd.Context(), cfg, args[0])
		},
	}

	return cmd
}
