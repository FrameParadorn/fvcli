package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of fvcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version v0.0.1")
	},
}