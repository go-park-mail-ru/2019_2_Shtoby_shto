package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	defaultFileSettingsName = "trello-local-settings.json"
)

type Config struct {
	ImagePath     string `json:"trello.service.image.path"`
	DbConfig      string `json:"trello.service.db.config"`
	RedisConfig   string `json:"trello.service.redis.config"`
	RedisPass     string `json:"trello.service.redis.password"`
	RedisDbNumber int    `json:"trello.service.redis.db.number"`
}

var ToolConfig *Config

func readConfig(fileName string) {
	configFile, err := os.Open(fileName)
	if err != nil {
		log.Println("Can't open properties file: " + err.Error())
	}
	defer configFile.Close()
	if err = json.NewDecoder(configFile).Decode(ToolConfig); err != nil {
		log.Println("Can't parsing properties file: " + err.Error())
	}
}

func InitConfig(workPath, settingsFileName string) {
	if settingsFileName == "" {
		settingsFileName = defaultFileSettingsName
	}
	configFileName := path.Join(workPath, settingsFileName)
	ToolConfig = new(Config)
	readConfig(configFileName)
}

func GetInstance() *Config {
	return ToolConfig
}
