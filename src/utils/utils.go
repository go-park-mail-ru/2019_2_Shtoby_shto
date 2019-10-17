package utils

import (
	"2019_2_Shtoby_shto/src/config"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

const (
	hexPattern = "^(urn\\:uuid\\:)?\\{?([a-z0-9]{8})-([a-z0-9]{4})-" +
		"([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})\\}?$"
)

func IsUUID(str string) bool {
	md := regexp.MustCompile(hexPattern).FindStringSubmatch(str)
	if md == nil {
		return false
	} else {
		return true
	}
}

func GenerateUUID() (uuid.UUID, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return uuid.Nil, err
	}
	return uuid.NewV4(), nil
}

func Join(args ...string) string {
	return strings.Join(args, "/")
}

var apiURL = ""

func init() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	config.InitConfig(logger)

	apiURL = config.GetInstance().FrontendURL

	logger.Println("Frontend URL:", apiURL)
}
