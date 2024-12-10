/*
Copyright © 2024 Nabhdeep
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "CLI for Advent of Code",
	Long:  `A command-line tool to download daily puzzles, create files with inputs, and run solutions for Advent of Code. Ensure you are in the project's root directory when running aoc-cli.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
