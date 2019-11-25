package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

const (
	remoteSettingsFile = "file-settings.json"
	deployEnvVar       = "DEPLOYAPI"
	awsAccessKeyID     = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

//easyjson:json
type Config struct {
	Port             string `json:"trello.service.port"`
	StorageAccessKey string `json:"trello.service.storage.access.key"`
	StorageSecretKey string `json:"trello.service.storage.secret.key"`
	StorageRegion    string `json:"trello.service.storage.region"`
	StorageEndpoint  string `json:"trello.service.storage.endpoint"`
	StorageBucket    string `json:"trello.service.storage.bucket"`
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

	configFileName := path.Join(dir, remoteSettingsFile)
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
