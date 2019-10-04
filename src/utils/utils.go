package utils

import (
	"2019_2_Shtoby_shto/src/config"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"log"
	"os"
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

func ConvertUriToQueueName(uriFull string) string {
	res := ""
	uri := uriFull
	if strings.Index(uri, "?") > 0 {
		uri = uri[:strings.Index(uri, "?")]
	}
	uriItems := strings.Split(uri, "/")
	if len(uriItems) == 1 {
		res = uri
	} else {
		for u := range uriItems {
			if uriItems[u] != "" {
				if IsUUID(uriItems[u]) {
					uriItems[u] = ":uuid"
				}
				res = res + "/" + uriItems[u]
			} else if u == len(uriItems)-1 {
				res = res + "/"
			}
		}
	}
	return res
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

func SetHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", apiURL)
	// (*w).Header().Set("Access-Control-Allow-Origin", "https://fmonster2.aleshka2012.now.sh")
	// "http://localhost:3000" https://fmonster2.aleshka2012.now.sh, https://20192shtobyshto-git-mergetest.ivanshport98.now.sh/")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
