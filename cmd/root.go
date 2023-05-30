/*
Copyright © 2023 NAME HERE paradorn.chandeth@gmail.com
*/package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/FrameParadorn/fvcli/cmd/new"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fvcli",
	Short: "Command line interface tool for generate code of clean architecture",
	Long:  `Command line interface tool for generate code of clean architecture with dependency injection (DI) by uber-dig`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fvcli.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(versionCmd)

	newCommand := new.NewNew()
	rootCmd.AddCommand(newCommand.CreateCommand())

}
