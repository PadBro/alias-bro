package cmd

import (
	"fmt"

	"github.com/PadBro/alias-bro/internal"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type errAliasExists struct {
	alias   string
	command string
}

func (err *errAliasExists) Error() string {
	return fmt.Sprintf("The alias `%s` already exists with the command: `%s`", err.alias, err.command)
}

func (e *errAliasExists) Severity() internal.Severity {
	return internal.SeverityError
}

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [alias] [command]",
	Short: "Adds an alias [alias] [command]",
	Long: `Add a new alias to alias-bro.

The first argument is the alias name (e.g., "ll"),
and the second argument is the command it should execute (e.g., "ls -la").

This command updates your alias-bro configuration and regenerates
the shell source file, so the alias is available in your shell
(after sourcing your aliases file).`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		command := args[1]

		err := add(alias, command)
		internal.HandleError(err)

		internal.GenerateSourceFile()

		fmt.Println(text.FgGreen.Sprintf("The alias %s was added", alias))
		fmt.Println(text.FgHiBlack.Sprintf(
			"Run `source %s/aliases.sh` or restart your shell to apply the changes.",
			internal.ConfigDir(),
		))
	},
}

func add(alias string, command string) error {
	aliases := internal.GetAliases()

	if command, exists := aliases[alias]; exists {
		return &errAliasExists{alias, command}
	}

	aliases[alias] = command

	viper.Set("aliases", aliases)
	err := viper.WriteConfig()
	cobra.CheckErr(err)

	return nil
}
