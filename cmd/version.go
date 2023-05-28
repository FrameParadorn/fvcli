package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getVersion() string {
	return "version v0.0.1"
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of fvcli",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(getVersion())
	},
}
