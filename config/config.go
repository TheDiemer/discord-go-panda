package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Discord struct {
		Token     string `yaml:"token"`
		Send_time string `yaml:"send_time"`
		Owner     string `yaml:"owner"`
		Guild     string `yaml:"guild"`
		Self      string `yaml:"self"`
		Trigger   string `yaml:"trigger"`
	} `yaml:"discord"`
	Database struct {
		IP          string `yaml:"ip"`
		DB_Username string `yaml:"db_username"`
		DB_Password string `yaml:"db_password"`
	} `yaml:"database"`
}

func ReadFile(conf_name, conf_path, conf_type string) (config Config, err error) {
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
