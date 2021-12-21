package commands

import (
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
)

const triggerChar = "."

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
	if string(args[0][0]) == triggerChar {
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
			// Lets make sure this is happening in #dnd or #bots
			switch m.ChannelID {
				// #dnd, #bots
			case "654446991912206346", "655092075389517894":
				// OKAY
				var msg strings.Builder
				var dm strings.Builder
				if len(args) < 2 {
					msg = errorMessage("Command missing", "For a list of commands type .help")
					dm = errorMessage("Command missing", "For a list of commands type .help")
				} else {
					msg, dm = Attendance(args[1], m.Author.Mention())
				}
				_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
				if err != nil {
					fmt.Println(err)
				} else {
					user, _ := s.UserChannelCreate(m.Author.ID)
					_, err := s.ChannelMessageSend(user.ID, dm.String())
					if err != nil {
						fmt.Println(err)
					}
				}
			default:
				msg := errorMessage("Invalid channel", "This command is only available in #dnd (use) or #bots (testing).")
				_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
				if err != nil {
					fmt.Println(err)
				}
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
