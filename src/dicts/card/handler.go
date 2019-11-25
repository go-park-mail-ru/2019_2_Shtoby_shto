package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/cardTags"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/comment"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/dicts/tag"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type Handler struct {
	userService      user.HandlerUserService
	cardService      HandlerCardService
	cardUsersService сardUsers.HandlerCardUsersService
	tagService       tag.HandlerTagService
	cardTagsService  cardTags.HandlerCardTagsService
	commentService   comment.HandlerCommentService
	securityService  security.HandlerSecurity
	handle.HandlerImpl
}

func NewCardHandler(e *echo.Echo, userService user.HandlerUserService,
	cardService HandlerCardService,
	cardUsersService сardUsers.HandlerCardUsersService,
	tagService tag.HandlerTagService,
	cardTagsService cardTags.HandlerCardTagsService,
	commentService comment.HandlerCommentService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:      userService,
		cardService:      cardService,
		cardUsersService: cardUsersService,
		tagService:       tagService,
		cardTagsService:  cardTagsService,
		commentService:   commentService,
		securityService:  securityService,
	}
	e.GET("/cards/:id", handler.Get)
	e.GET("/cards", handler.Fetch)
	//e.POST( "/cards/board", handler.PostCardsBoard)
	e.POST("/cards/user", handler.FetchUserCards)
	e.POST("/cards/user/attach", handler.AttachUserToCard)
	e.POST("/cards/user/detach", handler.DetachUserToCard)
	e.POST("/cards/:id/file/attach", handler.AttachFileToCard)
	e.GET("/cards/:id/file", handler.GetCardFile)
	e.POST("/cards", handler.Post)
	e.PUT("/cards/:id", handler.Put)
	e.DELETE("/cards/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	card, err := h.cardService.FindCardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FindCardByID error", http.StatusBadRequest, err)
		return err
	}
	if err := h.fillCardFields(card); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillCardFields error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, card)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	cards, err := h.cardService.FetchCards(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	for i := range cards {
		if err := h.fillCardFields(&cards[i]); err != nil {
			errorsLib.ErrorHandler(ctx.Response(), "fillBoardFields error ", http.StatusBadRequest, err)
			ctx.Logger().Error(err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, cards)
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
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "FetchCardsByIDs error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	for i := range cards {
		if err := h.fillCardFields(&cards[i]); err != nil {
			errorsLib.ErrorHandler(ctx.Response(), "fillBoardFields error ", http.StatusBadRequest, err)
			ctx.Logger().Error(err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, cards)
}

func (h Handler) DetachUserToCard(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	attachRequest := &models.CardsUserAttachRequest{}
	if err := attachRequest.UnmarshalJSON(body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UnmarshalJSON body error", http.StatusInternalServerError, err)
		return err
	}
	err = h.cardUsersService.DeleteCardUsersByIDs(attachRequest.UserID, attachRequest.CardID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "DeleteCardUsersByIDs error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, attachRequest)
}

func (h Handler) AttachUserToCard(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	attachRequest := &models.CardsUserAttachRequest{}
	if err := attachRequest.UnmarshalJSON(body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UnmarshalJSON body error", http.StatusInternalServerError, err)
		return err
	}
	count, err := h.cardUsersService.FindCardUsersByIDs(attachRequest.UserID, attachRequest.CardID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FindBoardUsersByIDs error", http.StatusInternalServerError, err)
		return err
	}
	if count == 0 {
		_, err = h.cardUsersService.CreateCardUsers(attachRequest.UserID, attachRequest.CardID)
		if err != nil {
			ctx.Logger().Error(err)
			errorsLib.ErrorHandler(ctx.Response(), "CreateCardUsers error", http.StatusInternalServerError, err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, attachRequest)
}

func (h Handler) AttachFileToCard(ctx echo.Context) error {
	rr := bufio.NewReader(ctx.Request().Body)
	card, err := h.cardService.DownloadFileToCard(rr, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, card)
}

func (h Handler) GetCardFile(ctx echo.Context) error {
	file, err := h.cardService.GetCardFile(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Add("Content-Type", "multipart/form-data")
	if _, err := ctx.Response().Write([]byte(file)); err != nil {
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, file)
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
	card, err := h.cardService.UpdateCard(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update card error", http.StatusInternalServerError, err)
		return err
	}
	err = h.fillCardFields(card)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillCardFields card error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, card)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.cardService.DeleteCard(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete card error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}

func (h Handler) fillCardFields(card *models.Card) error {
	comments, err := h.commentService.FetchCommentsByCardID(card.ID.String())
	if err != nil {
		return err
	}
	card.Comments = comments
	cardTags, err := h.cardTagsService.FindCardTagsByCardID(card.ID)
	if err != nil {
		return err
	}
	resultTagIDs := make([]string, 0)
	for _, ct := range cardTags {
		resultTagIDs = append(resultTagIDs, ct.TagID.String())
	}
	tags, err := h.tagService.FetchTagsByIDs(resultTagIDs)
	if err != nil {
		return err
	}
	card.Tags = tags
	err = h.fillCardUsers(card)
	if err != nil {
		return err
	}
	return nil
}

func (h Handler) fillCardUsers(card *models.Card) error {
	cUsers, err := h.cardUsersService.FetchCardUsersByCardID(card.ID)
	if err != nil {
		return err
	}
	usersResult := make([]string, 0)
	for _, value := range cUsers {
		usersResult = append(usersResult, value.UserID.String())
	}
	card.Users = usersResult
	return nil
}
