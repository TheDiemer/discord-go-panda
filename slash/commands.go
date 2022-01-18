package slash

import (
	"fmt"
//	"github.com/TheDiemer/discord-go-panda/commands"
	"github.com/TheDiemer/discord-go-panda/config"
	"github.com/bwmarrin/discordgo"
//	"strconv"
	"strings"
)

var conf config.Config
var err error
var dndChannels []string

func init() {
	// Putting this here for now, but commented out till we actually need it
	// So we can get variables as necessary
	conf, err = config.ReadFile("config/conf.yml", "", "yaml")
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return
	}
	// This is the Channel ID for the #bots channel, where I test and develop the bot
	dndChannels = append(dndChannels, "655092075389517894")
	// This is #dnd, which is the PRIMARY intended target
	dndChannels = append(dndChannels, "654446991912206346")
}

func isRole(guild string, author string, role string, s *discordgo.Session) (is bool) {
	member, err := s.State.Member(guild, author)
	if err != nil {
		fmt.Println(err)
		if member, err = s.GuildMember(guild, author); err != nil {
			fmt.Println(err)
			is = false
		} else {
			fmt.Println(err)
			for _, roleID := range member.Roles {
				if roleID == role {
					is = true
				}
			}
		}
	} else {
		for _, roleID := range member.Roles {
			if roleID == role {
				is = true
			}
		}
	}
	return
}

func who(s *discordgo.Session, id string, name string) (mention string) {
	users, _ := s.GuildMembers(id, "0", 1000)
	for _, member := range users {
		mention = member.User.Username
		if mention == name {
			mention = member.Mention()
			return
		}
	}
	mention = name
	return
}

func handle(part1 string, part2 string) string {
	var combined strings.Builder
	tmp := strings.Fields(part1)
	tmp2 := strings.Fields(part2)
	// things
	for _, word := range tmp {
		combined.WriteString(word)
		combined.WriteString(" ")
	}
	for i, word := range tmp2 {
		combined.WriteString(word)
		if i < len(tmp) {
			combined.WriteString(" ")
		}
	}
	return combined.String()
}

func channelCheck(channel string, approvedList []string) (approved bool) {
	approved = false
	for _, v := range approvedList {
		if channel == v {
			approved = true
		}
	}
	return
}

var (
	Commands = []*discordgo.ApplicationCommand{
//		{
//			Name:        "schedule",
//			Description: "Schedule a D&D session.",
//			Type:        discordgo.ChatApplicationCommand,
//			Options: []*discordgo.ApplicationCommandOption{
//				{
//					Type:        discordgo.ApplicationCommandOptionString,
//					Name:        "date-time",
//					Description: "Date and Time (of any understandable format) for the session to occur.",
//					Required:    true,
//				},
//				{
//					Type:        discordgo.ApplicationCommandOptionString,
//					Name:        "topic",
//					Description: "Session topic, like `explore the ruins nearby`.",
//					Required:    true,
//				},
//			},
//		},
//		{
//			Name:        "check",
//			Description: "Check if any sessions are planned, or check on the status of a specified session.",
//			Type:        discordgo.ChatApplicationCommand,
//			Options: []*discordgo.ApplicationCommandOption{
//				{
//					Type:        discordgo.ApplicationCommandOptionInteger,
//					Name:        "session-number",
//					Description: "A session number to check specifically.",
//					Required:    false,
//				},
//				{
//					Type:        discordgo.ApplicationCommandOptionBoolean,
//					Name:        "private",
//					Description: "Only display the output to you.",
//					Required:    false,
//				},
//			},
//		},
//		{
//			Name:        "register",
//			Description: "Register that you want to play in a given session.",
//			Type:        discordgo.ChatApplicationCommand,
//			Options: []*discordgo.ApplicationCommandOption{
//				{
//					Type:        discordgo.ApplicationCommandOptionInteger,
//					Name:        "session-number",
//					Description: "A session number to check specifically.",
//					Required:    true,
//				},
//			},
//		},
//		{
//			Name:        "deregister",
//			Description: "Remove your name from the list of players or the queue for a given session.",
//			Type:        discordgo.ChatApplicationCommand,
//			Options: []*discordgo.ApplicationCommandOption{
//				{
//					Type:        discordgo.ApplicationCommandOptionInteger,
//					Name:        "session-number",
//					Description: "A session number to check specifically.",
//					Required:    true,
//				},
//			},
//		},
//		{
//			Name:        "delete",
//			Description: "DM ONLY: Delete sessions from the list of planned sessions.",
//			Type:        discordgo.ChatApplicationCommand,
//			Options: []*discordgo.ApplicationCommandOption{
//				{
//					Type:        discordgo.ApplicationCommandOptionInteger,
//					Name:        "session-number",
//					Description: "A session number to delete specifically.",
//					Required:    true,
//				},
//			},
//		},
		{
			Name:        "rollcall",
			Description: "Check in with the players of a given campaign. We assume a Wednesday playday, so this just asks about the week and lists the players.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "campaign",
					Description:  "Which campaign are we polling about?",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
	}
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"rollcall":   handleRollcall,
//		"schedule":   handleSchedule,
//		"check":      handleCheck,
//		"register":   handleRegister,
//		"deregister": handleDeregister,
//		"delete":     handleDelete,
	}
)

//		// if they are a DM, they are allowed to do this!
//		if isRole(conf.Discord.Guild, i.Member.User.ID, "918615545039970335", s) {

func handleRollcall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if channelCheck(i.ChannelID, dndChannels) {
			data := i.ApplicationCommandData()
			// Privately ack the input
			var msgformat strings.Builder
			msgformat.WriteString("I understood that you want to call rollcall for **")
			msgformat.WriteString(data.Options[0].StringValue())
			msgformat.WriteString("'s** campaign")
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: msgformat.String(),
				},
			})
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		//data := i.ApplicationCommandData()
		choices := []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "noodles",
				Value: "transmutation",
			},
			{
				Name:  "Transmutation",
				Value: "transmutation",
			},
			{
				Name:  "jonesin",
				Value: "mesegea",
			},
			{
				Name:  "tropolis",
				Value: "mesegea",
			},
			{
				Name:  "mesegea",
				Value: "mesegea",
			},
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{
				Choices: choices,
			},
		})
		if err != nil {
			panic(err)
		}
	}

		//} else {
		//	message := commands.ErrorMessage("Deletion Failed", info.String())
		//	s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		//		Flags:   1 << 6,
		//		Content: message.String(),
		//	})
		//}
}

// func handleDelete(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if channelCheck(i.ChannelID, scheduleChannels) {
// 		if isRole(conf.Discord.Guild, i.Member.User.ID, "918615545039970335", s) {
// 			session := strconv.Itoa(int(i.ApplicationCommandData().Options[0].IntValue()))
// 			// Privately ack the input
// 			var msgformat strings.Builder
// 		https://discord.com/blog/slash-commands-are-here	msgformat.WriteString("I understood that you want to delete the session numbered: **")
// 			msgformat.WriteString(session)
// 			msgformat.WriteString("**.")
// 			msgformat.WriteString("\nLemme go work on that...")
// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Flags:   1 << 6,
// 					Content: msgformat.String(),
// 				},
// 			})
// 			success, info, names := SlashDelete(session, conf)
// 			if success {
// 				message := commands.SuccessMessage("Player Deregistered", info.String())
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Flags:   1 << 6,
// 					Content: message.String(),
// 				})
// 				if len(names) > 0 {
// 					var tmpMsg strings.Builder
// 					tmpMsg.WriteString("Hey ")
// 					for i, user := range names {
// 						name := who(s, conf.Discord.Guild, user)
// 						tmpMsg.WriteString(name)
// 						if i < len(names)-1 {
// 							tmpMsg.WriteString(", ")
// 						}
// 					}
// 					tmpMsg.WriteString("!\nSession **")
// 					tmpMsg.WriteString(session)
// 					tmpMsg.WriteString("** has been deleted. Please followup with ")
// 					tmpMsg.WriteString(i.Member.Mention())
// 					tmpMsg.WriteString(" if you want more information.")
// 					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 						Content: tmpMsg.String(),
// 					})
// 				}
// 
// 			} else {
// 				message := commands.ErrorMessage("Deletion Failed", info.String())
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Flags:   1 << 6,
// 					Content: message.String(),
// 				})
// 			}
// 		} else {
// 			msg := commands.ErrorMessage("Invalid User", "The `delete` command can only be run by a `Dungeon Master`.")
// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Flags:   1 << 6,
// 					Content: msg.String(),
// 				},
// 			})
// 		}
// 	} else {
// 		msg := commands.ErrorMessage("Invalid Channel", "The `delete` command can only be run in the `#scheduling` channel.")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msg.String(),
// 			},
// 		})
// 	}
// }
// func handleDeregister(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if channelCheck(i.ChannelID, scheduleChannels) {
// 		session := strconv.Itoa(int(i.ApplicationCommandData().Options[0].IntValue()))
// 		// Privately ack the input
// 		var msgformat strings.Builder
// 		msgformat.WriteString("I understood that you want to deregister from the session numbered: **")
// 		msgformat.WriteString(session)
// 		msgformat.WriteString("**.")
// 		msgformat.WriteString("\nLemme go work on that...")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msgformat.String(),
// 			},
// 		})
// 		success, info, name, followup := SlashDeregister(session, conf, i.Member.User.String())
// 		if success {
// 			message := commands.SuccessMessage("Player Deregistered", info.String())
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 			if followup.String() != "" {
// 				name = who(s, conf.Discord.Guild, name)
// 				var tmpMsg strings.Builder
// 				tmpMsg.WriteString("Hey ")
// 				tmpMsg.WriteString(name)
// 				tmpMsg.WriteString(followup.String())
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Content: tmpMsg.String(),
// 				})
// 			}
// 
// 		} else {
// 			message := commands.ErrorMessage("Deregistration Failed", info.String())
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 		}
// 	} else {
// 		msg := commands.ErrorMessage("Invalid Channel", "The `deregister` command can only be run in the `#scheduling` channel.")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msg.String(),
// 			},
// 		})
// 	}
// }
// 
// func handleRegister(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if channelCheck(i.ChannelID, scheduleChannels) {
// 		session := strconv.Itoa(int(i.ApplicationCommandData().Options[0].IntValue()))
// 		// Privately ack the input
// 		var msgformat strings.Builder
// 		msgformat.WriteString("I understood that you want to register for the session numbered: **")
// 		msgformat.WriteString(session)
// 		msgformat.WriteString("**.")
// 		msgformat.WriteString("\nLemme go work on that...")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msgformat.String(),
// 			},
// 		})
// 		success, info := SlashRegister(session, conf, i.Member.User.String())
// 		if success {
// 			message := commands.SuccessMessage("Player Registered", info.String())
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 		} else {
// 			message := commands.ErrorMessage("Registration Failed", info.String())
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 		}
// 	} else {
// 		msg := commands.ErrorMessage("Invalid Channel", "The `register` command can only be run in the `#scheduling` channel.")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msg.String(),
// 			},
// 		})
// 	}
// }
// 
// func handleCheck(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if channelCheck(i.ChannelID, scheduleChannels) {
// 		var session string
// 		private := false
// 		if len(i.ApplicationCommandData().Options) == 1 {
// 			if i.ApplicationCommandData().Options[0].Name == "session-number" {
// 				session = strconv.Itoa(int(i.ApplicationCommandData().Options[0].IntValue()))
// 			} else {
// 				private = i.ApplicationCommandData().Options[0].BoolValue()
// 			}
// 
// 		} else if len(i.ApplicationCommandData().Options) == 2 {
// 			if i.ApplicationCommandData().Options[0].Name == "session-number" {
// 				session = strconv.Itoa(int(i.ApplicationCommandData().Options[0].IntValue()))
// 				private = i.ApplicationCommandData().Options[1].BoolValue()
// 			} else {
// 				private = i.ApplicationCommandData().Options[0].BoolValue()
// 				session = strconv.Itoa(int(i.ApplicationCommandData().Options[1].IntValue()))
// 			}
// 		}
// 		// Privately ack the input
// 		var msgformat strings.Builder
// 		msgformat.WriteString("I understood that you want to check on")
// 		if session != "" {
// 			msgformat.WriteString(" the session numbered: **")
// 			msgformat.WriteString(session)
// 			msgformat.WriteString("**.")
// 		} else {
// 			msgformat.WriteString(" all available sessions.")
// 		}
// 		if private {
// 			msgformat.WriteString("\nAnd you want the result to only be shown to you.")
// 		}
// 		msgformat.WriteString("\nLemme go look...")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msgformat.String(),
// 			},
// 		})
// 		success, info, sessions := CheckSlash(session, conf)
// 		// let's now check if it was successful!
// 		if success {
// 			message := commands.SuccessMessage("Session Data Collected", info.String())
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 
// 			for _, single := range sessions {
// 				var err error
// 				if private {
// 					_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 						Content: single.String(),
// 						Flags:   1 << 6,
// 					})
// 				} else {
// 					_, err = s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 						Content: single.String(),
// 					})
// 				}
// 				// precautionary step
// 				if err != nil {
// 					msg := commands.ErrorMessage("Message Failure", "The response on storing your session failed to send.")
// 					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 						Flags:   1 << 6,
// 						Content: msg.String(),
// 					})
// 					return
// 				}
// 			}
// 		} else {
// 			// it failed, so lets discretely tell them it failed for some reason
// 			tmp := commands.ErrorMessage("Check Failure", info.String())
// 			_, err := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: tmp.String(),
// 			})
// 			// precautionary step
// 			if err != nil {
// 				msg := commands.ErrorMessage("Message Failure", "The response on storing your session failed to send.")
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Flags:   1 << 6,
// 					Content: msg.String(),
// 				})
// 				return
// 			}
// 		}
// 	} else {
// 		msg := commands.ErrorMessage("Invalid Channel", "The `check` command can only be run in the `#scheduling` channel.")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msg.String(),
// 			},
// 		})
// 	}
// }
// 
// func handleSchedule(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if channelCheck(i.ChannelID, scheduleChannels) {
// 		margs := []interface{}{
// 			i.ApplicationCommandData().Options[0].StringValue(),
// 			i.ApplicationCommandData().Options[1].StringValue(),
// 		}
// 		var msgformat strings.Builder
// 		msgformat.WriteString("I understood that you want to schedule a session for:\n```")
// 		msgformat.WriteString(fmt.Sprintf("%s", margs[0]))
// 		msgformat.WriteString("```\nWith the following topic/focus:\n```")
// 		msgformat.WriteString(fmt.Sprintf("%s", margs[1]))
// 		msgformat.WriteString("```\nMoving to record this.")
// 
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msgformat.String(),
// 			},
// 		})
// 
// 		information := handle(fmt.Sprintf("%s", margs[0]), fmt.Sprintf("%s", margs[1]))
// 		success, info := ScheduleSlash(information, conf)
// 		// let's now check if it was successful!
// 		if success {
// 			// It was, so lets discretely tell the user that
// 			message := commands.SuccessMessage("Session Stored", "The session you submitted has been successfully stored!")
// 			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: message.String(),
// 			})
// 			// And then also tell the whole world that there is a new session to register for
// 			var tmp strings.Builder
// 			tmp.WriteString("Hello ")
// 			tmp.WriteString(i.Member.Mention())
// 			tmp.WriteString(",\nYour planned session: `")
// 			tmp.WriteString(information)
// 			tmp.WriteString("` is now ready to accept players!\nThe first SIX (6) players to register using the `register` command to specify this session, number **")
// 			tmp.WriteString(info)
// 			tmp.WriteString("**, will be added to the session!\nAny additional players past SIX (6) will be added to the Queue so they can be added to the list of players if someone drops out!")
// 			_, err := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Content: tmp.String(),
// 			})
// 			// precautionary step
// 			if err != nil {
// 				msg := commands.ErrorMessage("Message Failure", "The response on storing your session failed to send.")
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Flags:   1 << 6,
// 					Content: msg.String(),
// 				})
// 				return
// 			}
// 		} else {
// 			// it failed, so lets discretely tell them it failed for some reason
// 			var reason strings.Builder
// 			reason.WriteString("The scheduling tool failed to record your session:\n```\n")
// 			reason.WriteString(info)
// 			reason.WriteString("\n```")
// 			tmp := commands.ErrorMessage("Schedule Failure", reason.String())
// 			_, err := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 				Flags:   1 << 6,
// 				Content: tmp.String(),
// 			})
// 			// precautionary step
// 			if err != nil {
// 				msg := commands.ErrorMessage("Message Failure", "The response on storing your session failed to send.")
// 				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
// 					Flags:   1 << 6,
// 					Content: msg.String(),
// 				})
// 				return
// 			}
// 		}
// 	} else {
// 		msg := commands.ErrorMessage("Invalid Channel", "The `schedule` command can only be run in the `#scheduling` channel.")
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Flags:   1 << 6,
// 				Content: msg.String(),
// 			},
// 		})
// 	}
// }
