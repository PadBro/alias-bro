package internal

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func SetupTestConfig(t *testing.T) {
	t.Helper()

	tempDir := t.TempDir()
	viper.Reset()

	viper.AddConfigPath(tempDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("aliases")

	err := viper.SafeWriteConfig()
	assert.Nil(t, err)
}
