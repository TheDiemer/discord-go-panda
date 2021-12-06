package commands

import (
	"strings"
)

func errorMessage(title string, message string) (err strings.Builder) {
	err.WriteString("❌  **")
	err.WriteString(title)
	err.WriteString("**\n")
	err.WriteString(message)
	return err
}
