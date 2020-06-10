package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetTodoList(c echo.Context) error {
	todoList, err := todoService.GetTodoService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todoList)
}
