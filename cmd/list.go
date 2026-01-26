package cmd

import (
	"os"

	"github.com/PadBro/alias-bro/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Long: `Display all aliases currently managed by alias-bro.

Each alias shows its name and the command it executes.
The output is formatted in a neat table for readability.`,
	Run: func(cmd *cobra.Command, args []string) {
		aliases := internal.GetAliases()

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Alias", "Command"})
		i := 1
		for alias, command := range aliases {
			t.AppendRow(table.Row{i, alias, command})
			i++
		}
		t.Render()
	},
}
