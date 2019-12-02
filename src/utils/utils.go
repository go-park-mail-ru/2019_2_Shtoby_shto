package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	hexPattern = "^(urn\\:uuid\\:)?\\{?([a-z0-9]{8})-([a-z0-9]{4})-" +
		"([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})\\}?$"
)

type RequestParams struct {
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Sort   []string `json:"sort"`
}

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

func ParseRequestParams(params url.URL) RequestParams {
	requestParams := RequestParams{}
	lim := params.Query().Get("limit")
	limit, err := strconv.Atoi(lim)
	if err != nil {
		limit = 100
	}
	requestParams.Limit = limit
	off := params.Query().Get("offset")
	offset, err := strconv.Atoi(off)
	if err != nil {
		offset = 0
	}
	requestParams.Offset = offset
	s := params.Query().Get("sort")
	sort := strings.Split(s, ",")
	requestParams.Sort = sort
	return requestParams
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
