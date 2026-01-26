package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetAliases_Success(t *testing.T) {
	SetupTestConfig(t)
	viper.Set("aliases", map[string]string{"ll": "ls -la"})
	aliases := GetAliases()

	assert.Equal(t, "ls -la", aliases["ll"])
}

func TestConfigDir_Success(t *testing.T) {
	tempDir := t.TempDir()

	orig := userConfigDir
	defer func() { userConfigDir = orig }()

	userConfigDir = func() (string, error) {
		return tempDir, nil
	}

	dir := ConfigDir()

	assert.NotEqual(t, "", dir)

	info, err := os.Stat(dir)
	assert.Nil(t, err)

	assert.True(t, info.IsDir(), "ConfigDir path is not a directory: %s", dir)

	assert.Equal(t, "alias-bro", filepath.Base(dir), "expected 'alias-bro', got %s", filepath.Base(dir))
}
