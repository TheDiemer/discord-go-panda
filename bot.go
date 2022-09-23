package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TheDiemer/discord-go-panda/config"
	"github.com/TheDiemer/discord-go-panda/slash"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session
var conf config.Config
var err error

func init() {
	// Read and Load up the config file
	conf, err = config.ReadFile("config/conf.yml", "", "yaml")
	// Drop out in case it fails to load the config
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return
	}
	// Create a new Discord session using our bot token
	dg, err = discordgo.New("Bot " + conf.Discord.Token)
	// Error checking!
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// First, lets make a handler that tells us when the bot is Actually up and ready
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { fmt.Println("Bot is up! Press CTRL-C to exit.") })
	// Next, lets add a new handler for Each of the slash commands
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := slash.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	// These intents are for letting us do more with the messages and access certains parts of discord
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open the connection to discord
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Close later, when we say so
	defer dg.Close()

	// Actually Create all of the slash commands in discord itself
	createdCommands, err := dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, conf.Discord.Guild, slash.Commands)
	if err != nil {
		fmt.Printf("cannot register commands: %v", err)
	}

	// Setup the trigger to kill the bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("\nStarting shut down.\nFirst, to delete our slash commands...")

	// Last thing before we disconnect from discord, lets delete our existing commands so we don't just leave old ones hanging around
	for _, cmd := range createdCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, conf.Discord.Guild, cmd.ID)
		if err != nil {
			fmt.Printf("cannot delete %q command: %v", cmd.Name, err)
		}
	}
	fmt.Println("Shutting down gracefully.")
}
