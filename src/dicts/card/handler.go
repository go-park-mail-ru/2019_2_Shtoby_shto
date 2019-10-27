package card

import (
	"2019_2_Shtoby_shto/src/customType"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/task"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type Handler struct {
	userService      user.HandlerUserService
	cardService      HandlerCardService
	cardUsersService сardUsers.HandlerCardUsersService
	taskService      task.HandlerTaskService
	securityService  security.HandlerSecurity
	handle.HandlerImpl
}

func NewCardHandler(e *echo.Echo, userService user.HandlerUserService, cardService HandlerCardService, cardUsersService сardUsers.HandlerCardUsersService, taskService task.HandlerTaskService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:      userService,
		cardService:      cardService,
		cardUsersService: cardUsersService,
		taskService:      taskService,
		securityService:  securityService,
	}
	e.GET("/cards/:id", handler.Get)
	e.GET("/cards", handler.Fetch)
	//e.POST( "/cards/board", handler.PostCardsBoard)
	e.POST("/cards/user", handler.FetchUserCards)
	e.POST("/cards", handler.Post)
	e.PUT("/cards/:id", handler.Put)
	e.DELETE("/cards/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	card, err := h.cardService.FindCardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetCardById error", http.StatusBadRequest, err)
		return err
	}
	//h.taskService.
	if err := h.cardService.FillLookupFields(card); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Lookup fields add error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, card)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	users, err := h.cardService.FetchCards(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, users)
}

func (h Handler) FetchUserCards(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	cardUsers, err := h.cardUsersService.FindCardUsersByUserID(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	resultCardUsersIDs := make([]string, 0)
	for _, cardUser := range cardUsers {
		resultCardUsersIDs = append(resultCardUsersIDs, cardUser.CardID.String())
	}
	cards, err := h.cardService.FetchCardsByIDs(resultCardUsersIDs)
	return ctx.JSON(http.StatusOK, cards)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	newCard, err := h.cardService.CreateCard(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("get user_id failed")
	}
	cardUser, err := h.cardUsersService.CreateCardUsers(userID, newCard.ID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	newCard.CardUserID = cardUser.ID
	return ctx.JSON(http.StatusOK, newCard)
}

func (h Handler) Put(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	board, err := h.cardService.UpdateCard(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update card error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.cardService.DeleteCard(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete card error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
