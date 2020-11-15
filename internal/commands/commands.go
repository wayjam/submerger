package commands

import (
	"github.com/spf13/cobra"

	"github.com/wayjam/submerger/internal/version"
)

var (
	rootCmd = &cobra.Command{
		Use:   "submerger",
		Short: "A tool lets you customize clash config",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of submerger",
		Run: func(cmd *cobra.Command, args []string) {
			version.PrintVersion()
		},
	}
	rootCmd.AddCommand(versionCmd)
}
