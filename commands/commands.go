package commands

import (
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
)

func CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore anything said by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Get the Arguments
	args := strings.Split(m.Content, " ")
	if len(args) < 1 {
		msg := errorMessage("Command missing", "For a list of commands type .help")
		s.ChannelMessageSend(m.ChannelID, msg.String())
		return
	}
	if string(args[0][0]) == "." {
		switch string(args[0][1:]) {
		case "beans":
			message := GimmeBeans(m.Author.Mention())
			_, err := s.ChannelMessageSend(m.ChannelID, message.String())
			if err != nil {
				fmt.Println(err)
			}
		case "yaface":
			_, err := s.ChannelMessageSend(m.ChannelID, "nahhh, definitely YOURS :stuck_out_tongue:")
			if err != nil {
				fmt.Println(err)
			}
		case "rollcall":
			var msg strings.Builder
			if len(args) < 2 {
				msg = errorMessage("Command missing", "For a list of commands type .help")
			} else {
				msg = Attendance(args[1])
			}
			_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
			if err != nil {
				fmt.Println(err)
			}

		default:
			msg := errorMessage("Invalid command", "For a list of commands type .help")
			_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
