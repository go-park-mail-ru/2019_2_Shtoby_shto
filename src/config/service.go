package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	localSettingsFile  = "trello-local-settings.json"
	remoteSettingsFile = "trello-settings.json"
	deployEnvVar       = "DEPLOYAPI"
	awsAccessKeyID     = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

//easyjson:json
type Config struct {
	Port             int    `json:"trello.service.port"`
	FrontendURL      string `json:"trello.service.frontend.url"`
	SecurityURL      string `json:"trello.service.security.url"`
	ImagePath        string `json:"trello.service.image.path"`
	StorageAccessKey string `json:"trello.service.storage.access.key"`
	StorageSecretKey string `json:"trello.service.storage.secret.key"`
	StorageRegion    string `json:"trello.service.storage.region"`
	StorageEndpoint  string `json:"trello.service.storage.endpoint"`
	StorageBucket    string `json:"trello.service.storage.bucket"`
	DbConfig         string `json:"trello.service.db.config"`
	RedisConfig      string `json:"trello.service.redis.config"`
	RedisPass        string `json:"trello.service.redis.password"`
	RedisDbNumber    int    `json:"trello.service.redis.db.number"`
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
	// s3 set env
	if err = os.Setenv(awsAccessKeyID, ToolConfig.StorageAccessKey); err != nil {
		return err
	}
	if err = os.Setenv(awsSecretAccessKey, ToolConfig.StorageSecretKey); err != nil {
		return err
	}
	return nil
}
