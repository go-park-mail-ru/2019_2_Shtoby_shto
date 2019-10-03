package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	PROPERTIESDEFAULT = "trello-settings.json"
)

type Config struct {
	ImagePath string `json:"trello.service.image.path"`
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

func InitConfig(workPath string) {
	configFileName := path.Join(workPath, PROPERTIESDEFAULT)
	ToolConfig = new(Config)
	readConfig(configFileName)
}

func GetInstance() *Config {
	return ToolConfig
}
