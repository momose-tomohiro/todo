package main

import (
	"todo/api/application/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	handler.DI()

	e.GET("/todos", handler.GetTodoList)
	e.POST("/todos", handler.RegisterTodo)
	e.DELETE("/todos/:id", handler.RemoveTodo)

	e.Logger.Fatal(e.Start(":8888"))
}
