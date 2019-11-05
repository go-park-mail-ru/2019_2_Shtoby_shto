package task

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService     user.HandlerUserService
	taskService     HandlerTaskService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewTaskHandler(e *echo.Echo, userService user.HandlerUserService,
	taskService HandlerTaskService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		taskService:     taskService,
		securityService: securityService,
	}
	e.GET("/tasks/:id", handler.Get)
	e.GET("/tasks", handler.Fetch)
	e.POST("/tasks", handler.Post)
	e.PUT("/tasks/:id", handler.Put)
	e.DELETE("/tasks/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.taskService.FindTaskByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetTaskById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	tasks, err := h.taskService.FetchTasks(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	responseData, err := h.taskService.CreateTask(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, responseData)
}

func (h Handler) Put(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	task, err := h.taskService.UpdateTask(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update task error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, task)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.taskService.DeleteTask(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete task error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
