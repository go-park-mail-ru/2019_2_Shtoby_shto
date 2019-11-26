package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	localSettingsFile  = "session-local-settings.json"
	remoteSettingsFile = "session-settings.json"
	deployEnvVar       = "DEPLOYAPI"
)

//easyjson:json
type Config struct {
	Port          string `json:"trello.service.port"`
	RedisConfig   string `json:"trello.service.redis.config"`
	RedisPass     string `json:"trello.service.redis.password"`
	RedisDbNumber int    `json:"trello.service.redis.db.number"`
}

var ToolConfig *Config

func GetInstance() *Config {
	return ToolConfig
}

func readConfig(fileName string) error {
	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("Can't open properties file: " + err.Error())
	}
	if err = ToolConfig.UnmarshalJSON(configFile); err != nil {
		return errors.New("Can't parsing properties file: " + err.Error())
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
