package cmd

import (
	"errors"
	"fmt"

	"github.com/PadBro/alias-bro/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "alias-bro",
	Version: internal.Version,
	Short:   "Manage shell command aliases from a single CLI",
	Long: `alias-bro lets you create, list, and remove shell command aliases
using a simple configuration file.

It generates sourceable shell functions so aliases work like native
commands in your terminal.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(internal.ConfigDir())
	viper.SetConfigType("yaml")
	viper.SetConfigName("aliases")

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			fmt.Println("Creating config")
			err = viper.SafeWriteConfig()
			cobra.CheckErr(err)
		} else {
			cobra.CheckErr(err)
		}
	}
}
