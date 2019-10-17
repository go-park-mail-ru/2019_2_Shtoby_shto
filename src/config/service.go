package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
)

const (
	localSettingsFile  = "trello-local-settings.json"
	remoteSettingsFile = "trello-settings.json"
	deployEnvVar       = "DEPLOYAPI"
)

type Config struct {
	Port          int    `json:"trello.service.port"`
	FrontendURL   string `json:"trello.service.frontend.url"`
	ImagePath     string `json:"trello.service.image.path"`
	DbConfig      string `json:"trello.service.db.config"`
	RedisConfig   string `json:"trello.service.redis.config"`
	RedisPass     string `json:"trello.service.redis.password"`
	RedisDbNumber int    `json:"trello.service.redis.db.number"`
}

var ToolConfig *Config

func GetInstance() *Config {
	return ToolConfig
}

func readConfig(fileName string) error {
	configFile, err := os.Open(fileName)
	if err != nil {
		return errors.New("Can't open properties file: " + err.Error())
	}
	if err = json.NewDecoder(configFile).Decode(ToolConfig); err != nil {
		return errors.New("Can't parsing properties file: " + err.Error())
	}
	if err := configFile.Close(); err != nil {
		return err
	}
	return nil
}

func InitConfig() error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	deployVar := os.Getenv(deployEnvVar)

	var settingsFileName string

	if deployVar == "" {
		log.Printf(
			"%s not set, expecting requests from api on localhost deployment\n",
			deployEnvVar,
		)
		settingsFileName = localSettingsFile
	} else {
		settingsFileName = remoteSettingsFile
	}

	configFileName := path.Join(dir, settingsFileName)
	ToolConfig = new(Config)
	if err := readConfig(configFileName); err != nil {
		return err
	}
	return nil
}
