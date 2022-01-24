package slash

import (
	"fmt"
	"github.com/TheDiemer/discord-go-panda/commands"
	"github.com/TheDiemer/discord-go-panda/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var conf config.Config
var err error
var dndChannels []string
var transmutation []string
var mesegea []string

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
	transmutation = []string{"Tisi", "Ptrosk", "Baldrick", "Ikol", "Avu", "Red Staché"}
	mesegea = []string{"Adelvir", "Akta", "Ayayron", "Duvu", "Gisli", "Krasus", "Wrench"}
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
		{
			Name:        "musiclink",
			Description: "Turn a song link from most platforms into a generic link.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "link",
					Description: "Link for song to be converted.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "private",
					Description: "Only display the output to you.",
					Required:    false,
				},
			},
		},
		{
			Name:        "randomsong",
			Description: "Get a random song url.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "private",
					Description: "Only display the output to you.",
					Required:    false,
				},
			},
		},
		{
			Name:        "rollcall",
			Description: "Check in with the players of a given campaign.",
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
		"musiclink":  handleMusicLink,
		"rollcall":   handleRollcall,
		"randomsong": handleRandomSong,
	}
)

func handleRandomSong(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var msgformat strings.Builder
	msgformat.WriteString("I understood that you want a random song! Gimme a sec...")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: msgformat.String(),
		},
	})
	var private bool
	if len(i.ApplicationCommandData().Options) > 0 {
		private = i.ApplicationCommandData().Options[0].BoolValue()
	} else {
		private = false
	}
	song, err := GetSong()
	fmt.Println("song link is:", song)
	if err != nil {
		fmt.Println("error is:", err.Error())
	}
	if err != nil {
		message := commands.ErrorMessage("Random Song Failed", err.Error())
		s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Flags:   1 << 6,
			Content: message.String(),
		})
		// problems
	} else {
		message := commands.SuccessMessage("Random Song Chosen", song)
		fmt.Println(message.String())
		if private {
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Flags:   1 << 6,
				Content: message.String(),
			})
		} else {
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: message.String(),
			})
		}

	}
}

func handleMusicLink(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var link string
	var private bool
	if len(i.ApplicationCommandData().Options) > 1 {
		if i.ApplicationCommandData().Options[0].Name == "link" {
			link = i.ApplicationCommandData().Options[0].StringValue()
			private = i.ApplicationCommandData().Options[1].BoolValue()
		} else {
			private = i.ApplicationCommandData().Options[0].BoolValue()
			link = i.ApplicationCommandData().Options[1].StringValue()
		}
	} else {
		link = i.ApplicationCommandData().Options[0].StringValue()
		private = false
	}
	// Privately ack the input
	var msgformat strings.Builder
	msgformat.WriteString("I understood that you want to convert the following url: **")
	msgformat.WriteString(link)
	msgformat.WriteString("**\nOne sec...")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: msgformat.String(),
		},
	})
	success, info := GetSongLink(link)
	if success {
		message := commands.SuccessMessage("Song Converted", info.String())
		if private {
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Flags:   1 << 6,
				Content: message.String(),
			})
		} else {
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: message.String(),
			})
		}
	} else {
		message := commands.ErrorMessage("Song Conversion Failed", info.String())
		s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
			Flags:   1 << 6,
			Content: message.String(),
		})
	}
}

func handleRollcall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		var info strings.Builder
		if channelCheck(i.ChannelID, dndChannels) {
			if data.Options[0].StringValue() == "transmutation" {
				// Privately ack the input
				var msgformat strings.Builder
				msgformat.WriteString("I understood that you want to call rollcall for the **")
				msgformat.WriteString(data.Options[0].StringValue())
				msgformat.WriteString("** campaign. One sec...")
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   1 << 6,
						Content: msgformat.String(),
					},
				})
				info.WriteString("Hey, <@&654763072828997642>! Who of ")
				for i, person := range transmutation {
					info.WriteString(person)
					if i == len(transmutation)-2 {
						info.WriteString(", and ")
					} else if i < len(transmutation)-2 {
						info.WriteString(", ")
					}
				}
				info.WriteString(" will come to the call this week?")
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: info.String(),
				})
			} else if data.Options[0].StringValue() == "mesegea" {
				// Privately ack the input
				var msgformat strings.Builder
				msgformat.WriteString("I understood that you want to call rollcall for the **")
				msgformat.WriteString(data.Options[0].StringValue())
				msgformat.WriteString("** campaign! One sec...")
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   1 << 6,
						Content: msgformat.String(),
					},
				})
				info.WriteString("Hey, <@&654763072828997642>! Who of ")
				for i, person := range mesegea {
					info.WriteString(person)
					if i == len(mesegea)-2 {
						info.WriteString(", and ")
					} else if i < len(mesegea)-2 {
						info.WriteString(", ")
					}
				}
				info.WriteString(" will come to the call this week?")
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: info.String(),
				})
			} else {
				info.WriteString("\nUnfortunately, `")
				info.WriteString(data.Options[0].StringValue())
				info.WriteString("` is not a valid option for campaigns!\nI can't ping players I don't know about!")
				message := commands.ErrorMessage("Rollcall Error", info.String())
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   1 << 6,
						Content: message.String(),
					},
				})
			}
		} else {
			var info strings.Builder
			info.WriteString("\nUnfortunately, `rollcall` can only be run in `#dnd` to prevent unnecessary noise.")
			message := commands.ErrorMessage("Invalid Channel", info.String())
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: message.String(),
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
