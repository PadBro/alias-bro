package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateSourceFile() {
	aliases := GetAliases()

	var output []string
	for alias, command := range aliases {
		output = append(output, fmt.Sprintf("%s() {\n\texecute %s \"$@\"\n}\n", alias, command))
	}
	output = append(output, `execute() {
    local GREEN="\e[32m"
    local RESET="\e[0m"

    echo -e "${GREEN}Executing:${RESET} $@"
    "$@"
}
`)

	fileName := filepath.Join(ConfigDir(), "aliases.sh")
	err := os.WriteFile(fileName, []byte(strings.Join(output, "\n")), 0644)
	if err != nil {
		fmt.Println("unable to write file:", err)
	}
}
