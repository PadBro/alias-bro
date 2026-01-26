package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PadBro/alias-bro/internal"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type ErrShellNotFound struct{}

func (err *ErrShellNotFound) Error() string {
	return "No shell found!"
}
func (e *ErrShellNotFound) Severity() internal.Severity {
	return internal.SeverityError
}

type ErrUnsupportedShell struct {
	shell string
}

func (err *ErrUnsupportedShell) Error() string {
	return fmt.Sprintf("Unsupported shell: `%s`", err.shell)
}
func (e *ErrUnsupportedShell) Severity() internal.Severity {
	return internal.SeverityError
}

type ErrAlreadySourced struct {
	rcFile string
}

func (err *ErrAlreadySourced) Error() string {
	return fmt.Sprintf("Aliases are already sourced in: `%s`", err.rcFile)
}
func (e *ErrAlreadySourced) Severity() internal.Severity {
	return internal.SeverityWarning
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure your shell to load alias-bro aliases",
	Long: `Detects your shell and updates its rc file to source
alias-bro's generated aliases automatically.`,
	Run: func(cmd *cobra.Command, args []string) {

		rcFile, aliasesFile, err := setup()
		internal.HandleError(err)

		fmt.Println(text.FgGreen.Sprintf("`%s` has been sourced in `%s`", aliasesFile, rcFile))
	},
}

func setup() (string, string, error) {
	shellName, shellFound := shellName()
	if !shellFound {
		return "", "", &ErrShellNotFound{}
	}

	rcFile, supported := rcFile(shellName)
	if !supported {
		return "", "", &ErrUnsupportedShell{shellName}
	}

	aliasesFile := filepath.Join(internal.ConfigDir(), "aliases.sh")

	// check if rc file exists
	_, err := os.Stat(rcFile)
	cobra.CheckErr(err)

	if isSourced := isSourced(rcFile); isSourced {
		return "", "", &ErrAlreadySourced{rcFile}
	}

	addSource(rcFile, aliasesFile)

	return rcFile, aliasesFile, nil
}

func shellName() (string, bool) {
	shell, exists := os.LookupEnv("SHELL")
	if !exists {
		return "", false
	}
	shellName := filepath.Base(shell)
	return shellName, true
}

func rcFile(shellName string) (string, bool) {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	rcFileName, found := rcFileName(shellName)
	if !found {
		return "", false
	}

	return filepath.Join(home, rcFileName), true
}

func rcFileName(shellName string) (string, bool) {
	switch shellName {
	case "bash":
		return ".bashrc", true
	case "zsh":
		return ".zshrc", true
	default:
		return "", false
	}
}

func isSourced(rcFile string) bool {
	// read the file
	content, err := os.ReadFile(rcFile)
	cobra.CheckErr(err)

	// check if already sourced
	if strings.Contains(string(content), "alias-bro/aliases.sh") {
		return true
	}

	return false
}

func addSource(rcFile string, aliasesFile string) bool {
	// opens rc file to add source
	f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_WRONLY, 0644)
	cobra.CheckErr(err)
	defer func() {
		_ = f.Close()
	}()

	// add source to rc file
	_, err = fmt.Fprintf(f, "\nsource %s\n", aliasesFile)
	cobra.CheckErr(err)

	return false
}
