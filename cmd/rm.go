package cmd

import (
	"fmt"

	"github.com/PadBro/alias-bro/internal"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type errAliasNotFound struct {
	alias string
}

func (err *errAliasNotFound) Error() string {
	return fmt.Sprintf("The alias %s does not exists", err.alias)
}
func (e *errAliasNotFound) Severity() internal.Severity {
	return internal.SeverityError
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm [alias]",
	Short: "Remove an existing alias",
	Long: `Remove an alias from alias-bro.

The specified alias will be deleted from your configuration
and the shell source file will be updated accordingly.

Example:
  alias-bro rm ll`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]

		err := rm(alias)
		internal.HandleError(err)

		internal.GenerateSourceFile()

		fmt.Println(text.FgGreen.Sprintf("The alias %s was removed", alias))
		fmt.Println(text.FgHiBlack.Sprintf(
			"Run `source %s/aliases.sh` or restart your shell to apply the changes.",
			internal.ConfigDir(),
		))
	},
}

func rm(alias string) error {
	aliases := internal.GetAliases()

	if _, exists := aliases[alias]; !exists {
		return &errAliasNotFound{alias}
	}

	delete(aliases, alias)

	viper.Set("aliases", aliases)
	err := viper.WriteConfig()
	cobra.CheckErr(err)

	return nil
}
