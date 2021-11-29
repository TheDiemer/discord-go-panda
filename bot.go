package main

import (
	//	"encoding/json"
	//	"flag"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"math/rand"
	"log"
	"os"
	"os/signal"
	"strings"
	//	"syscall"
	//	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Building our config structure so we can use the unmarshalled config values
type Config struct {
	Discord struct {
		Token     string `yaml:"token"`
		Send_time string `yaml:"send_time"`
		Owner     string `yaml:"owner"`
		Guild     string `yaml:"guild"`
	} `yaml:"discord"`
	Database struct {
		IP          string `yaml:"ip"`
		DB_Username string `yaml:"db_username"`
		DB_Password string `yaml:"db_password"`
		DL_Username string `yaml:"dl_username"`
		DL_Password string `yaml:"dl_password"`
	} `yaml:"database"`
}

var s *discordgo.Session
var conf Config

func init() {
	// Read and Load up the config file
	conf, err := readFile("conf.yml", "", "yaml")
	// Drop out in case it fails to load the config
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return
	}
	fmt.Println(conf.Discord.Token)
//	var err error
	// Create a new Discord session using our bot token
	fmt.Println(conf.Discord.Token)
	s, err = discordgo.New("Bot " + conf.Discord.Token)
	// Error checking!
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "beans",
			Description: "Give the gift of beans",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"beans": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "test",//gimmeBeans(s.State.User.Mention()).String(),
				},
			})
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	// This is a placeholder to make sure the config gets loaded properly
	// fmt.Println(conf.Discord.Token)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("bot is up!")
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

// 	command := &discordgo.ApplicationCommand{
// 		Name: "beans",
// 		Type: discordgo.ChatApplicationCommand,
// 		Description: "give the gift of beans",
// 	}
// 	s.ApplicationCommandCreate(s.State.User.ID, conf.Discord.Guild, command)
	 // _, err := s.ApplicationCommandCreate(s.State.User.ID, conf.Discord.Guild, command)
	 //if err != nil {
	 //	log.Panicf("cannot create '%v' command %v", command, err)
	 //}
	 for _, v := range commands {
	 	_, err := s.ApplicationCommandCreate(s.State.User.ID, conf.Discord.Guild, v)
	 	if err != nil {
	 		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	 	}
	 }
	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("gracefully shutting down")
//	// Register the messageCreate func as a callback for MessageCreate events.
//	s.AddHandler(messageCreate)
//
//	s.Identify.Intents = discordgo.IntentsGuildMessages
//
//
//	fmt.Println("Bot is now running. Press CTRL-C to exit.")
//	sc := make(chan os.Signal, 1)
//	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
//	<-sc
//
//	defer s.Close()
}

func readFile(conf_name, conf_path, conf_type string) (config Config, err error) {
	viper.SetConfigName(conf_name)
	viper.SetConfigType(conf_type)
	if conf_path == "" {
		viper.AddConfigPath(".")
	}
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&config)
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages sent by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "yaface") {
		_, err := s.ChannelMessageSend(m.ChannelID, "nahhh, definitely yours :stuck_out_tongue:")
		if err != nil {
			fmt.Println(err)
		}
	}
	//// Beans Function, v2
	//if strings.Contains(m.Content, "beans") {
	//	// Get our message!
	//	message := gimmeBeans(m.Author.Mention())

	//	// Now lets Send our bean!
	//	_, err := s.ChannelMessageSend(m.ChannelID, message.String())
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
}
