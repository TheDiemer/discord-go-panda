package slash

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/TheDiemer/discord-go-panda/commands"
	"github.com/TheDiemer/discord-go-panda/config"
	"github.com/TheDiemer/discord-go-panda/functions"
	"github.com/bwmarrin/discordgo"
)

var conf config.Config
var err error
var dndChannels []string
var transmutation []string
var mesegea []string
var rollCallReactions []string

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
	mesegea = []string{"Adelvir", "Ayayron", "Duvu", "Gisli", "Krasus", "Wrench"}
	rollCallReactions = []string{
		"yee:707604728724324382",
		"megusta:775469871454486549",
		"wolo:789952118739042334",
		"sus:808326972475834449",
		"nooo:846428536939741244",
	}
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
			Name:        "randomwiki",
			Description: "Get a random wiki entry.",
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
		{
			Name:        "alias",
			Description: "Setup an alias for a user.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "og-nick",
					Description: "What is the name you want to receive karma?",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "alias",
					Description: "What is the alias you want to apply to the og nick?",
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
			Name:        "quote",
			Description: "Get quotes!",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "What is the id of the quote you want to see?",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "quoted",
					Description: "Who do you want to get a quote from?",
					Required:    false,
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
			Name:        "addquote",
			Description: "Transcribe a memory into our history books.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "quote",
					Description: "What is the quote to be saved?",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "quoted",
					Description: "Who said the quote?",
					Required:    true,
				},
			},
		},
		{
			Name:        "poll",
			Description: "Ask your current channel a poll with configurable response options.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "question",
					Description: "Question or poll to ask the channel!",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "title",
					Description: "General Topic of your poll (default: POLL).",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "color",
					Description: "Decimal code for the Color on the side of the card (default: 15597568 (or EE0000 in hex)).",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "responses",
					Description: "Comma separated list of :emoji:=Meaning (default: :yee:=Positive, :wolo:=Neutral, :nooo:=Negative).",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "ping",
					Description: "Space separated list of users/roles to send a ping to.",
					Required:    false,
				},
			},
		},
		{
			Name:        "cheer",
			Description: "Send a gif to cheer someone on! Optionally Tag them if you want to highlight your target!",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "target",
					Description: "@ whomever you want to cheer on!",
					Required:    false,
				},
			},
		},
		{
			Name:        "blink",
			Description: "blink! Optionally blink At someone",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "target",
					Description: "@ Whomever you want to blink at...",
					Required:    false,
				},
			},
		},
	}
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"randomwiki": handleRandomWiki,
		"musiclink":  handleMusicLink,
		"rollcall":   handleRollcall,
		"randomsong": handleRandomSong,
		"dnd":        handleDnd,
		"alias":      handleAlias,
		"quote":      handleQuote,
		"addquote":   handleAddQuote,
		"poll":       handlePoll,
		"cheer":      handleCheer,
		"blink":      handleBlink,
	}
)

func handleAddQuote(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var quote string
	quote = data.Options[0].StringValue()
	var quoted string
	quoted = data.Options[1].StringValue()
	var quoter string
	quoter = i.Member.User.ID
	var channel string
	channel = i.ChannelID
	//var date string
	//date = time.Now().Format(time.RFC3339)
	//info, err := AddQuote(quote, quoted, quoter, channel, date)
	info, err := AddQuote(quote, quoted, quoter, channel)
	if err != nil {
		response := commands.ErrorMessage("Error saving quote", info.String())
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response.String(),
			},
		})
		//s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		//	Content: response.String(),
		//})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: info.String(),
			},
		})
		//s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		//	Content: info.String(),
		//})
	}

}

func handleBlink(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var target string
	for _, option := range data.Options {
		switch option.Name {
		case "target":
			target = option.StringValue()
		}
	}
	var response strings.Builder
	if target != "" {
		response.WriteString("Hey ")
		response.WriteString(target)
		response.WriteString(",\n")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response.String(),
			},
		})
	} else {
		response.WriteString("Hey ")
		response.WriteString(i.Member.User.Username)
		response.WriteString(",\n")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response.String(),
			},
		})
	}
	s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		Content: "https://cdn.discordapp.com/attachments/778744161317683283/1027334449039413389/blink1.gif",
	})

}

func handleCheer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var target string
	for _, option := range data.Options {
		switch option.Name {
		case "target":
			target = option.StringValue()
		}
	}
	var response strings.Builder
	gif := functions.GetGiph(conf)
	if target != "" {
		response.WriteString("Hey ")
		response.WriteString(target)
		response.WriteString(",\n")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response.String(),
			},
		})
	} else {
		response.WriteString("Here ")
		response.WriteString(i.Member.User.Username)
		response.WriteString(",\n")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response.String(),
			},
		})
	}
	s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
		Content: gif,
	})
}

func handlePoll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Ack! Generating and sending your poll!",
		},
	})
	data := i.ApplicationCommandData()
	var question string
	var title string
	var color string
	var responses string
	var ping string
	for _, option := range data.Options {
		switch option.Name {
		case "question":
			question = option.StringValue()
		case "title":
			title = option.StringValue()
		case "color":
			color = option.StringValue()
		case "responses":
			responses = option.StringValue()
		case "ping":
			ping = option.StringValue()
		}
	}
	if title == "" {
		title = "POLL"
	}
	if color == "" {
		color = "15597568"
	}
	finalColor, _ := strconv.Atoi(color)
	if responses == "" {
		responses = "<:yee:707604728724324382>=Positive Response, <:wolo:789952118739042334>=Neutral Response, <:nooo:846428536939741244>=Negative Response"
	}
	tmpResponses := strings.Split(responses, ", ")
	responses = strings.Join(tmpResponses, "\n")
	if ping != "" {
		var pinged strings.Builder
		pinged.WriteString("Hey, ")
		pinged.WriteString(ping)
		pinged.WriteString(", <@")
		id := i.Member.User.ID
		if id == "" {
			id := i.User.ID
			fmt.Println(id)
		}
		fmt.Println(id)
		pinged.WriteString(id)
		pinged.WriteString("> is asking a poll! Please respond as you're able!")
		_, err := s.ChannelMessageSend(i.ChannelID, pinged.String())
		if err != nil {
			fmt.Println(err)
		}
	}
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       finalColor,
		Description: "Taking a poll",
		Type:        "rich",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   question,
				Value:  responses,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     title,
	}
	s.ChannelMessageSendEmbed(i.ChannelID, embed)
}

func handleQuote(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var id string
	var quoted string
	var private bool
	private = false
	for _, option := range data.Options {
		switch option.Name {
		case "id":
			id = option.StringValue()
		case "quoted":
			quoted = option.StringValue()
		case "private":
			private = option.BoolValue()
		}
	}
	info, err := GetQuote(id, quoted, conf)
	var response strings.Builder
	if err != nil {
		response = commands.ErrorMessage("Error getting quote", "No quote to be found.")
		if private {
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Flags:   1 << 6,
				Content: response.String(),
			})
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: response.String(),
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response.String(),
				},
			})
			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: response.String(),
			})
		}
	} else {
		var answer strings.Builder
		answer.WriteString("```")
		answer.WriteString(info.quote)
		answer.WriteString("```\n -- ")
		answer.WriteString(info.quoted)
		answer.WriteString(", ")
		answer.WriteString(info.date)
		answer.WriteString(" [")
		answer.WriteString(strconv.FormatInt(info.id, 10))
		answer.WriteString("]")
		response = commands.SuccessMessage("Quote Collected", answer.String())
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x317F43, //Signal Green
			Description: info.quote,
			Type:        "rich",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Original Speaker:",
					Value:  info.quoted,
					Inline: true,
				},
			},
			Timestamp: info.date,
			Title:     "Quote #" + strconv.FormatInt(info.id, 10),
		}
		var embeds []*discordgo.MessageEmbed
		embeds = append(embeds, embed)

		if private {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embeds,
					Flags:  1 << 6,
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embeds,
				},
			})
		}
	}
}

func handleRandomWiki(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var private bool
	if len(i.ApplicationCommandData().Options) > 0 {
		private = i.ApplicationCommandData().Options[0].BoolValue()
	} else {
		private = false
	}
	wiki, err := GetWiki()
	if err != nil {
		fmt.Println("error is: ", err.Error())
		message := commands.ErrorMessage("Random Wiki Failed", err.Error())
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message.String(),
				Flags:   1 << 6,
			},
		})
	} else {
		fmt.Println("wiki entry is: ", wiki)
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xCD7F32, //Wiki bronze
			Description: wiki.Extract,
			Type:        "rich",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "url",
					Value:  wiki.ContentURLs.Desktop.Page,
					Inline: true,
				},
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     wiki.Title,
		}

		var embeds []*discordgo.MessageEmbed
		embeds = append(embeds, embed)

		if private {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embeds,
					Flags:  1 << 6,
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: embeds,
				},
			})
		}
	}
}

func handleAlias(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var nick string
	var alias string
	var private bool
	private = false
	for _, option := range data.Options {
		switch option.Name {
		case "og-nick":
			nick = option.StringValue()
		case "alias":
			alias = option.StringValue()
		case "private":
			private = option.BoolValue()
		}
	}
	info, err := Alias(nick, alias, conf)
	var response strings.Builder
	if err != nil {
		response = commands.ErrorMessage("Error Creating Alias", info.String())
		if private {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: response.String(),
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response.String(),
				},
			})
		}
	} else {
		response = commands.SuccessMessage("Alias Created", info.String())
		if private {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: response.String(),
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response.String(),
				},
			})
		}
	}
}

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
	data := i.ApplicationCommandData()
	for _, option := range data.Options {
		switch option.Name {
		case "link":
			link = option.StringValue()
		case "private":
			private = option.BoolValue()
		}
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
	var checkMsg strings.Builder
	checkMsg.WriteString("<:")
	checkMsg.WriteString(rollCallReactions[0])
	checkMsg.WriteString("> = Yee\n")
	checkMsg.WriteString("<:")
	checkMsg.WriteString(rollCallReactions[1])
	checkMsg.WriteString("> = ProbablYee\n")
	checkMsg.WriteString("<:")
	checkMsg.WriteString(rollCallReactions[2])
	checkMsg.WriteString("> = MabYee\n")
	checkMsg.WriteString("<:")
	checkMsg.WriteString(rollCallReactions[3])
	checkMsg.WriteString("> = ProbablNYee\n")
	checkMsg.WriteString("<:")
	checkMsg.WriteString(rollCallReactions[4])
	checkMsg.WriteString("> = NYee")
	var mydudes string
	tmp, _ := s.GuildRoles(conf.Discord.Guild)

	for _, role := range tmp {
		if role.ID == "654763072828997642" {
			mydudes = role.Mention()
		}
	}

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
				// This is to ping the role!
				var tmp strings.Builder
				tmp.WriteString("Hey, ")
				tmp.WriteString(mydudes)
				tmp.WriteString("!")
				_, err := s.ChannelMessageSend(i.ChannelID, tmp.String())
				if err != nil {
					fmt.Println(err)
				}
				// Now for the content of names!
				info.WriteString("Who of ")
				for i, person := range transmutation {
					info.WriteString(person)
					if i == len(transmutation)-2 {
						info.WriteString(", and ")
					} else if i < len(transmutation)-2 {
						info.WriteString(", ")
					}
				}
				info.WriteString(" will come to the call this week?")
				embed := &discordgo.MessageEmbed{
					Author:      &discordgo.MessageEmbedAuthor{},
					Color:       0xC73032, //DDB Red
					Description: info.String(),
					Type:        "rich",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Can you make it?",
							Value:  checkMsg.String(),
							Inline: true,
						},
					},
					Timestamp: time.Now().Format(time.RFC3339),
					Title:     "Roll Calllllll",
				}
				sent, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
				if err != nil {
					fmt.Println(err)
				}
				for _, emoji := range rollCallReactions {
					s.MessageReactionAdd(sent.ChannelID, sent.ID, emoji)
				}
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
				// This is to ping the role!
				var tmp strings.Builder
				tmp.WriteString("Hey, ")
				tmp.WriteString(mydudes)
				tmp.WriteString("!")
				_, err := s.ChannelMessageSend(i.ChannelID, tmp.String())
				if err != nil {
					fmt.Println(err)
				}
				// Now for the content of names!
				info.WriteString("Who of ")
				for i, person := range mesegea {
					info.WriteString(person)
					if i == len(mesegea)-2 {
						info.WriteString(", and ")
					} else if i < len(mesegea)-2 {
						info.WriteString(", ")
					}
				}
				info.WriteString(" will come to the call this week?")
				embed := &discordgo.MessageEmbed{
					Author:      &discordgo.MessageEmbedAuthor{},
					Color:       0xC73032, //DDB Red
					Description: info.String(),
					Type:        "rich",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Can you make it?",
							Value:  checkMsg.String(),
							Inline: true,
						},
						&discordgo.MessageEmbedField{
							Name:   "Recap:",
							Value:  "https://ericdiemerjones.com/dnd/last",
							Inline: false,
						},
					},
					Timestamp: time.Now().Format(time.RFC3339),
					Title:     "Roll Calllllll",
				}
				sent, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
				if err != nil {
					fmt.Println(err)
				}
				for _, emoji := range rollCallReactions {
					s.MessageReactionAdd(sent.ChannelID, sent.ID, emoji)
				}
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
