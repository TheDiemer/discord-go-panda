package commands

import (
	"math/rand"
	"strings"
	"time"
)

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func SuccessMessage(title string, message string) (success strings.Builder) {
	// Define our optional types of successes and the sorts of things the bot should say
	// General means the exact cause of the issue is unknown or is TOO specific to get its own category
	general := []string{
		"Congrats! That worked! ",
	}

	success.WriteString("✅  **")
	success.WriteString(title)
	success.WriteString("**\n")
	switch title {
	default:
		success.WriteString(general[rand.Intn(len(general))])
		success.WriteString(message)
	}
	return
}

func ErrorMessage(title string, message string) (err strings.Builder) {
	// Define our optional types of errors and the sorts of things the bot should say
	// invalid means they specified a command that isn't in our list
	invalid := []string{
		"HEY, you! You typed it in wrong! ",
	}
	// Channel means the channel they ran the command in is not an approved one
	channel := []string{
		"HEY! This command can't be run in this channel! ",
	}
	// General means the exact cause of the issue is unknown or is TOO specific to get its own category
	general := []string{
		"HEY! What'd you do that for? ",
	}

	err.WriteString("❌  **")
	err.WriteString(title)
	err.WriteString("**\n")
	switch title {
	case "Invalid Command":
		err.WriteString(invalid[rand.Intn(len(invalid))])
		err.WriteString(message)
	case "Invalid Channel":
		err.WriteString(channel[rand.Intn(len(channel))])
		err.WriteString(message)
	default:
		err.WriteString(general[rand.Intn(len(general))])
		err.WriteString(message)
	}
	return
}
