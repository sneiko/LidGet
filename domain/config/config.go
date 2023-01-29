package config

import (
	"encoding/json"
	"lidget/domain/pkg/logger"
	"os"
)

type (
	Config struct {
		Http               HttpConfig `json:"api"`
		DbConnectionString string     `json:"db_connection_string"`
		Tg                 TgConfig   `json:"telegram"`
	}

	HttpConfig struct {
		Port string `json:"port"`
	}

	TgConfig struct {
		Id          int    `json:"id"`
		Hash        string `json:"hash"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
)

func LoadConfig() (*Config, error) {
	var config Config

	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		config = Config{
			Http: HttpConfig{
				Port: "5000",
			},
		}
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &config, nil
}
