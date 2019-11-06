package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	"2019_2_Shtoby_shto/src/dicts/card"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
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
	userService       user.HandlerUserService
	boardService      HandlerBoardService
	cardService       card.HandlerCardService
	cardGroupService  cardGroup.HandlerCardGroupService
	taskService       task.HandlerTaskService
	boardUsersService boardUsers.HandlerBoardUsersService
	securityService   security.HandlerSecurity
	handle.HandlerImpl
}

func NewBoardHandler(e *echo.Echo, userService user.HandlerUserService,
	boardService HandlerBoardService,
	boardUsersService boardUsers.HandlerBoardUsersService,
	cardService card.HandlerCardService,
	cardGroupService cardGroup.HandlerCardGroupService,
	taskService task.HandlerTaskService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:       userService,
		boardService:      boardService,
		boardUsersService: boardUsersService,
		cardService:       cardService,
		cardGroupService:  cardGroupService,
		taskService:       taskService,
		securityService:   securityService,
	}
	e.GET("/board/:id", handler.Get)
	e.GET("/board", handler.Fetch)
	e.GET("/board/user/:id", handler.FetchUserBoards)
	e.POST("/board", handler.Post)
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
	cardGroups, err := h.cardGroupService.FetchCardGroupsByBoardIDs([]string{board.ID.String()})
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "FetchCardGroupsByBoardIDs error", http.StatusBadRequest, err)
		return err
	}
	for i, group := range cardGroups {
		cards, err := h.cardService.FetchCardsByCardGroupIDs([]string{group.ID.String()})
		if err != nil {
			ctx.Logger().Error(err)
			errorsLib.ErrorHandler(ctx.Response(), "FetchCardsByCardGroupIDs error", http.StatusBadRequest, err)
			return err
		}
		for j, card := range cards {
			tasks, err := h.taskService.FetchTasksByCardIDs([]string{card.ID.String()})
			if err != nil {
				ctx.Logger().Error(err)
				errorsLib.ErrorHandler(ctx.Response(), "FetchCardsByCardGroupIDs error", http.StatusBadRequest, err)
				return err
			}
			cards[j].Tasks = tasks
		}
		cardGroups[i].Cards = cards
	}
	board.CardGroups = cardGroups
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	users, err := h.boardService.FetchBoards(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, users)
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
		return errors.New("It is not your data, man")
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
	return ctx.JSON(http.StatusOK, boards)
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
	return ctx.JSON(http.StatusOK, newBoard)
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
