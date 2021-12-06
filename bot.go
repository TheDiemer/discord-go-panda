package main

import (
	//	"encoding/json"
	//	"flag"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"math/rand"
	"os"
	"os/signal"
	//"strings"
	"syscall"
	//	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"github.com/TheDiemer/discord-go-panda/commands"
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

func main() {
	// Read and Load up the config file
	conf, err := readFile("conf.yml", "", "yaml")
	// Drop out in case it fails to load the config
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return
	}
	// fmt.Println(conf.Discord.Token)

	// Create a new Discord session using our bot token
	dg, err := discordgo.New("Bot " + conf.Discord.Token)
	// Error checking!
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.

	//dg.AddHandler(messageCreate)
	dg.AddHandler(commands.CommandsHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
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
