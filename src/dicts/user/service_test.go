package user

import (
	_ "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestService_CreateUser(t *testing.T) {

	mokUser := &User{
		BaseInfo: dicts.BaseInfo{},
		Login:    "Ivan",
		Password: "123456",
		PhotoID:  nil,
	}

	data, _ := mokUser.MarshalJSON()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := NewMockHandlerUserService(ctrl)
	service.EXPECT().CreateUser(data).Return(mokUser, nil)

	handler := Handler{userService: service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/registration", strings.NewReader(string(data)))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/registration")

	body, err := handler.userService.CreateUser([]byte(fmt.Sprintf("%v", c)))
	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	comparable, _ := body.MarshalJSON()
	//body,_ := ioutil.ReadAll(rec.Body)
	//t.Log(body)

	if string(comparable) != string(data) {
		t.Errorf("Expected: %s , got: %s", string(data), string(comparable))
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mokUser, rec.Body.String())
	}
}
