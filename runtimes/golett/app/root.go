package app

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCommand := &cobra.Command{
		Use: "golett",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return rootCommand
}
