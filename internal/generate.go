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
		commands := strings.Split(command, " && ")
		output = append(output, fmt.Sprintf("%s() {", alias))
		for _, command := range commands {

			newCommand := command
			if !strings.Contains(command, "\"$@\"") {
				newCommand = fmt.Sprintf("%s \"$@\"", command)
			}

			output = append(output, fmt.Sprintf("\texecute %s", newCommand))
		}
		output = append(output, "}\n")
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
