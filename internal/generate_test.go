package internal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSourceFile_Success(t *testing.T) {
	SetupTestConfig(t)
	viper.Set("aliases", map[string]string{"ll": "ls -la"})

	tempDir := t.TempDir()

	orig := userConfigDir
	defer func() { userConfigDir = orig }()

	userConfigDir = func() (string, error) {
		return tempDir, nil
	}

	GenerateSourceFile()

	aliasesFile := filepath.Join(tempDir, "alias-bro", "aliases.sh")

	_, err := os.Stat(aliasesFile)
	assert.Nil(t, err)

	content, _ := os.ReadFile(aliasesFile)
	assert.True(t, strings.Contains(string(content), "ll()"))
}
