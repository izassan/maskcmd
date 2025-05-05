package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.0.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("maskcmd %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
