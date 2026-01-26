package internal

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type SeverityErr interface {
	error
	Severity() Severity
}

type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
)

func HandleError(err error) {
	if err == nil {
		return
	}

	if severityErr, ok := err.(SeverityErr); ok {
		switch severityErr.Severity() {
		case SeverityWarning:
			fmt.Println(text.FgYellow.Sprint(severityErr))
		default:
			fmt.Println(text.FgRed.Sprint(severityErr))
		}
		os.Exit(1)
	}
	cobra.CheckErr(err)
}
