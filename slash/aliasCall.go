package slash

import (
	"strings"
)

func Alias(nick string, alias string) (info strings.Builder, err error) {
	info.WriteString("\nNow, `")
	info.WriteString(nick)
	info.WriteString("` is known as `")
	info.WriteString(alias)
	info.WriteString("`! Thanks!")
	return
}
