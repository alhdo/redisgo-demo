package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"server"`
	Redis struct {
		Host string `mapstructure:"address"`
		Port string `mapstrucure:"port"`
	}
}

func New(confFile string) *Config {
	viper.SetConfigName(confFile)
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/simple-go-worker/")
	viper.AddConfigPath("$HOME/.simple-go-worker")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error: %w \n", err))
	}
	viper.SetDefault("server.address", "127.0.0.1:8080")
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", "6379")
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &config
}
