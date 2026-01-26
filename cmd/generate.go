package cmd

import (
	"github.com/PadBro/alias-bro/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the shell aliases file",
	Long: `Generates a shell source file containing all configured aliases.
This file is sourced by your shell to make aliases available as commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.GenerateSourceFile()
	},
}
