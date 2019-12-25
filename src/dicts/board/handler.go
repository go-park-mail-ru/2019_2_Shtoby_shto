package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	"2019_2_Shtoby_shto/src/dicts/card"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
	"2019_2_Shtoby_shto/src/dicts/cardTags"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/checkList"
	"2019_2_Shtoby_shto/src/dicts/comment"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/dicts/tag"
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
	userService       user.HandlerUserService
	boardService      HandlerBoardService
	cardService       card.HandlerCardService
	cardUsersService  сardUsers.HandlerCardUsersService
	cardGroupService  cardGroup.HandlerCardGroupService
	tagService        tag.HandlerTagService
	cardTagsService   cardTags.HandlerCardTagsService
	commentService    comment.HandlerCommentService
	checkListService  checkList.HandlerCheckListService
	boardUsersService boardUsers.HandlerBoardUsersService
	securityService   security.HandlerSecurity
	handle.HandlerImpl
}

func NewBoardHandler(e *echo.Echo, userService user.HandlerUserService,
	boardService HandlerBoardService,
	boardUsersService boardUsers.HandlerBoardUsersService,
	cardService card.HandlerCardService,
	cardUsersService сardUsers.HandlerCardUsersService,
	cardGroupService cardGroup.HandlerCardGroupService,
	tagService tag.HandlerTagService,
	cardTagService cardTags.HandlerCardTagsService,
	commentService comment.HandlerCommentService,
	checkListService checkList.HandlerCheckListService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:       userService,
		boardService:      boardService,
		boardUsersService: boardUsersService,
		cardService:       cardService,
		cardUsersService:  cardUsersService,
		cardGroupService:  cardGroupService,
		tagService:        tagService,
		cardTagsService:   cardTagService,
		commentService:    commentService,
		checkListService:  checkListService,
		securityService:   securityService,
	}
	e.GET("/board/:id", handler.Get)
	e.GET("/board", handler.Fetch)
	e.GET("/board/user/:id", handler.FetchUserBoards)
	e.POST("/board", handler.Post)
	e.GET("/:short-url", handler.GetShortURL)
	e.POST("/board/user/attach", handler.AttachUserToBoard)
	e.POST("/board/user/detach", handler.DetachUserToBoard)
	e.PUT("/board/:id", handler.Put)
	e.DELETE("/board/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	board, err := h.boardService.FindBoardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
		return err
	}
	if err := h.fillBoardFields(board); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillBoardFields error", http.StatusInternalServerError, err)
		return err
	}
	if err := h.fillBoardUsers(board); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillBoardUsers error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	boards, err := h.boardService.FetchBoards(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	for i := range boards {
		if err := h.fillBoardUsers(&boards[i]); err != nil {
			errorsLib.ErrorHandler(ctx.Response(), "fillBoardUsers error ", http.StatusBadRequest, err)
			ctx.Logger().Error(err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, boards)
}

func (h Handler) FetchUserBoards(ctx echo.Context) error {
	userID := customType.StringUUID(ctx.Param("id"))
	curUserID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("get user_id failed")
	}
	if curUserID != userID {
		return ctx.String(http.StatusUnauthorized, "It is not your data, man")
	}
	if !userID.IsUUID() {
		return ctx.JSON(http.StatusBadRequest, errors.New("Not valid userID"))
	}
	boardUsers, err := h.boardUsersService.FetchBoardUsersByUserID(userID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FetchBoardUsersByUserID error", http.StatusInternalServerError, err)
		return err
	}
	resultBoardIDs := make([]string, 0)
	for _, boardUser := range boardUsers {
		resultBoardIDs = append(resultBoardIDs, boardUser.BoardID.String())
	}
	boards, err := h.boardService.FetchBoardsByIDs(resultBoardIDs)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FetchBoardUsersByUserID error", http.StatusInternalServerError, err)
		return err
	}
	for i := range boards {
		if err := h.fillBoardUsers(&boards[i]); err != nil {
			errorsLib.ErrorHandler(ctx.Response(), "fillBoardUsers error ", http.StatusBadRequest, err)
			ctx.Logger().Error(err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, boards)
}

func (h Handler) AttachUserToBoard(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	attachRequest := &models.BoardsUserAttachRequest{}
	if err := attachRequest.UnmarshalJSON(body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UnmarshalJSON body error", http.StatusInternalServerError, err)
		return err
	}
	count, err := h.boardUsersService.FindBoardUsersByIDs(attachRequest.UserID, attachRequest.BoardID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FindBoardUsersByIDs error", http.StatusInternalServerError, err)
		return err
	}
	if count == 0 {
		_, err = h.boardUsersService.CreateBoardUsers("", attachRequest.UserID, attachRequest.BoardID)
		if err != nil {
			ctx.Logger().Error(err)
			errorsLib.ErrorHandler(ctx.Response(), "CreateBoardUsers error", http.StatusInternalServerError, err)
			return err
		}
	}
	return ctx.JSON(http.StatusOK, attachRequest)
}

func (h Handler) DetachUserToBoard(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	attachRequest := &models.BoardsUserAttachRequest{}
	if err := attachRequest.UnmarshalJSON(body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UnmarshalJSON body error", http.StatusInternalServerError, err)
		return err
	}
	err = h.boardUsersService.DeleteBoardUsersByIDs(attachRequest.UserID, attachRequest.BoardID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "DeleteCardUsersByIDs error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, attachRequest)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}

	newBoardUsersID, err := utils.GenerateUUID()
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GenerateUUID error", http.StatusInternalServerError, err)
		return err
	}
	newBoard, err := h.boardService.CreateBoard(body, customType.StringUUID(newBoardUsersID.String()))
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

	boardUser, err := h.boardUsersService.CreateBoardUsers(customType.StringUUID(newBoardUsersID.String()), userID, newBoard.ID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}

	newBoard.BoardUsersID = boardUser.ID

	newBoard.Users = append(newBoard.Users, userID.String())

	//event := &models.EventMessage{
	//	Event: "event",
	//}
	//
	//e, err := event.MarshalJSON()
	//if err != nil {
	//	errorsLib.ErrorHandler(ctx.Response(), "event.MarshalJSON", http.StatusInternalServerError, err)
	//	ctx.Logger().Error(err)
	//	return err
	//}
	//
	//h.hub.Broadcast <- e

	return ctx.JSON(http.StatusOK, newBoard)
}

func (h Handler) GetShortURL(ctx echo.Context) error {
	shortURL := ctx.Param("short-url")
	board, err := h.boardService.FindBoardByURL(shortURL)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "FindBoardByURL error", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}

	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("get user_id failed")
	}

	count, _ := h.boardUsersService.FindBoardUsersByIDs(userID, board.ID)
	if count != 0 {
		return ctx.JSON(http.StatusOK, "already exist")
	}

	userBoard, err := h.boardUsersService.CreateBoardUsers("", userID, board.ID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, userBoard)
}

func (h Handler) Put(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}

	board, err := h.boardService.UpdateBoard(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}

	if err := h.fillBoardFields(board); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillBoardFields error", http.StatusInternalServerError, err)
		return err
	}

	if err := h.fillBoardUsers(board); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "fillBoardUsers error", http.StatusInternalServerError, err)
		return err
	}

	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Delete(ctx echo.Context) error {
	boardID := customType.StringUUID(ctx.Param("id"))
	if err := h.boardService.DeleteBoard(boardID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}

func (h Handler) fillBoardFields(board *models.Board) error {
	cardGroups, err := h.cardGroupService.FetchCardGroupsByBoardIDs([]string{board.ID.String()})
	if err != nil {
		return err
	}
	for i, group := range cardGroups {
		cards, err := h.cardService.FetchCardsByCardGroupIDs([]string{group.ID.String()})
		if err != nil {
			return err
		}
		for j, card := range cards {
			comments, err := h.commentService.FetchCommentsByCardID(card.ID.String())
			if err != nil {
				return err
			}
			cards[j].Comments = comments
			checkLists, err := h.checkListService.FetchCheckListsByCardID(card.ID.String())
			if err != nil {
				return err
			}
			cards[j].CheckLists = checkLists
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
			cards[j].Tags = tags
			cUsers, err := h.cardUsersService.FetchCardUsersByCardID(cards[j].ID)
			usersResult := make([]string, 0)
			for _, value := range cUsers {
				usersResult = append(usersResult, value.UserID.String())
			}
			cards[j].Users = usersResult
		}
		cardGroups[i].Cards = cards
	}
	board.CardGroups = cardGroups
	return nil
}

func (h Handler) fillBoardUsers(board *models.Board) error {
	bUsers, err := h.boardUsersService.FetchBoardUsersByBoardID(board.ID)
	if err != nil {
		return err
	}
	usersResult := make([]string, 0)
	for _, value := range bUsers {
		usersResult = append(usersResult, value.UserID.String())
	}
	board.Users = usersResult
	return nil
}
