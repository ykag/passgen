package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "passgen",
		Short: "A simple password generator that can generate passwords of different complexities",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
