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

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func dndMessage(data string, name string) (message strings.Builder) {
	tmpSplit := strings.Split(data, " ")
	message.WriteString("\n")
	message.WriteString(name)
	message.WriteString(" for the ")
	message.WriteString(tmpSplit[0])
	message.WriteString(" campaign: ")
	message.WriteString(tmpSplit[1])
	return
}

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
		{
			Name:        "dnd",
			Description: "Commands for playing D&D!",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "notes",
					Description:  "Which campaign notes do you want?",
					Required:     false,
					Autocomplete: true,
				},
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "website",
					Description:  "Which campaign site do you want?",
					Required:     false,
					Autocomplete: true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "private",
					Description: "Only display the output to you.",
					Required:    false,
				},
			},
		},
	}
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"musiclink":  handleMusicLink,
		"rollcall":   handleRollcall,
		"randomsong": handleRandomSong,
		"dnd":        handleDnd,
	}
)

func handleDnd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		var private bool
		var message strings.Builder
		private = false
		switch tmp := len(data.Options); {
		case tmp == 1:
			switch data.Options[0].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[0].StringValue(), data.Options[0].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[0].BoolValue()
				message.WriteString("\n")
				tmpMsg := commands.ErrorMessage("Invalid Command", "\nYou gotta choose what data you want!")
				message.WriteString(tmpMsg.String())
			}
		case tmp == 2:
			switch data.Options[0].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[0].StringValue(), data.Options[0].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[0].BoolValue()
			}
			switch data.Options[1].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[1].StringValue(), data.Options[1].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[1].BoolValue()
			}
		case tmp == 3:
			switch data.Options[0].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[0].StringValue(), data.Options[0].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[0].BoolValue()
			}
			switch data.Options[1].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[1].StringValue(), data.Options[1].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[1].BoolValue()
			}
			switch data.Options[2].Name {
			case "notes", "website":
				temporary := dndMessage(data.Options[2].StringValue(), data.Options[2].Name)
				message.WriteString(temporary.String())
			case "private":
				private = data.Options[2].BoolValue()
			}
		}
		// we need to get rid of that very first newline...we just..don't need it
		var info string
		info = trimLeftChars(message.String(), 1)
		if private {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: info,
				},
			})
			if err != nil {
				panic(err)
			}
		} else {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: info,
				},
			})
			if err != nil {
				panic(err)
			}
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		var choices []*discordgo.ApplicationCommandOptionChoice
		notes := []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "noodles",
				Value: "Transmutation https://docs.google.com/document/d/1Gzav6DrOHzDGi_KgIzgDIxX9BkxnKToRwvY6amkowkA/edit",
			},
			{
				Name:  "Transmutation",
				Value: "Transmutation https://docs.google.com/document/d/1Gzav6DrOHzDGi_KgIzgDIxX9BkxnKToRwvY6amkowkA/edit",
			},
			{
				Name:  "jonesin",
				Value: "Mesegea https://docs.google.com/document/d/1mYEV4mF_Xxabe4AU1cEEjaH0HnDyFa882NiOoXXhW80/edit",
			},
			{
				Name:  "tropolis",
				Value: "Mesegea https://docs.google.com/document/d/1mYEV4mF_Xxabe4AU1cEEjaH0HnDyFa882NiOoXXhW80/edit",
			},
			{
				Name:  "mesegea",
				Value: "Mesegea https://docs.google.com/document/d/1mYEV4mF_Xxabe4AU1cEEjaH0HnDyFa882NiOoXXhW80/edit",
			},
		}
		sites := []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "noodles",
				Value: "Transmutation https://wiki.noodles.ninja/doku.php?id=start",
			},
			{
				Name:  "Transmutation",
				Value: "Transmutation https://wiki.noodles.ninja/doku.php?id=start",
			},
			{
				Name:  "jonesin",
				Value: "Mesegea https://ericdiemerjones.com/dnd/",
			},
			{
				Name:  "tropolis",
				Value: "Mesegea https://ericdiemerjones.com/dnd/",
			},
			{
				Name:  "mesegea",
				Value: "Mesegea https://ericdiemerjones.com/dnd/",
			},
		}
		switch tmp := len(data.Options); {
		case tmp == 1:
			switch data.Options[0].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
		case tmp == 2:
			switch data.Options[0].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
			switch data.Options[1].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
		case tmp == 3:
			switch data.Options[0].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
			switch data.Options[1].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
			switch data.Options[2].Name {
			case "notes":
				choices = notes
			case "website":
				choices = sites
			}
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
}

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
}
