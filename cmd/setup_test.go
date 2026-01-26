package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellName(t *testing.T) {
	orig := os.Getenv("SHELL")
	defer func() {
		_ = os.Setenv("SHELL", orig)
	}()
	err := os.Setenv("SHELL", "/bin/zsh")
	assert.Nil(t, err)
	shell, exists := shellName()
	assert.True(t, exists)
	assert.Equal(t, "zsh", shell)

	err = os.Unsetenv("SHELL")
	assert.Nil(t, err)
	shell, exists = shellName()
	assert.False(t, exists)
	assert.Equal(t, "", shell)
}

func TestRcFileName(t *testing.T) {
	name, found := rcFileName("bash")
	assert.True(t, found)
	assert.Equal(t, ".bashrc", name)
	name, found = rcFileName("zsh")
	assert.True(t, found)
	assert.Equal(t, ".zshrc", name)
	name, found = rcFileName("none")
	assert.False(t, found)
	assert.Equal(t, "", name)
}
