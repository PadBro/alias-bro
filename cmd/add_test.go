package cmd

import (
	"testing"

	"github.com/PadBro/alias-bro/internal"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAdd_Success(t *testing.T) {
	internal.SetupTestConfig(t)

	assert.Nil(t, add("ll", "ls -la"))

	aliases := internal.GetAliases()
	assert.Equal(t, "ls -la", aliases["ll"])
}

func TestAdd_Duplicate(t *testing.T) {
	internal.SetupTestConfig(t)
	viper.Set("aliases", map[string]string{"ll": "ls -la"})

	assert.IsType(t, add("ll", "ls -la"), &errAliasExists{})
}
