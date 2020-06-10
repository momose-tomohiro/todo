package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func RemoveTodo(c echo.Context) error {
	err := todoService.RemoveTodoService(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
