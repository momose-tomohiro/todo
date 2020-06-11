package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterTodo(c echo.Context) error {
	err := todoService.RegisterTodoService(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
