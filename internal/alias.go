package internal

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userConfigDir = os.UserConfigDir

func GetAliases() map[string]string {
	aliases := map[string]string{}
	err := viper.UnmarshalKey("aliases", &aliases)
	cobra.CheckErr(err)
	return aliases
}

func ConfigDir() string {
	config, err := userConfigDir()
	cobra.CheckErr(err)

	configDir := filepath.Join(config, "alias-bro")
	err = os.MkdirAll(configDir, os.ModePerm)
	cobra.CheckErr(err)

	return configDir
}
