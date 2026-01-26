package cmd

import (
	"testing"

	"github.com/PadBro/alias-bro/internal"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRm_Success(t *testing.T) {
	internal.SetupTestConfig(t)
	viper.Set("aliases", map[string]string{"ll": "ls -la"})

	assert.Nil(t, rm("ll"))

	aliases := internal.GetAliases()
	assert.Equal(t, "", aliases["ll"])
}

func TestAdd_NotExisting(t *testing.T) {
	internal.SetupTestConfig(t)

	assert.IsType(t, rm("ll"), &errAliasNotFound{})
}
