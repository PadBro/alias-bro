package main

import (
	"github.com/PadBro/alias-bro/cmd"
	"github.com/spf13/cobra"
)

func main() {
	err := cmd.Execute()
	cobra.CheckErr(err)
}
