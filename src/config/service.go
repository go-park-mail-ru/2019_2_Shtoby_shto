package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	localSettingsFile  = "trello-local-settings.json"
	remoteSettingsFile = "trello-settings.json"
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

const (
	deployEnvVar = "DEPLOYAPI"
)

func InitConfig(logger *log.Logger) {
	dir, err := os.Getwd()

	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
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
	readConfig(configFileName)
}

var ToolConfig *Config

func GetInstance() *Config {
	return ToolConfig
}
