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

	todos := e.Group("/todos")
	todos.GET("", handler.GetTodoList)
	todos.POST("", handler.RegisterTodo)
	todos.DELETE("/:id", handler.RemoveTodo)

	e.Logger.Fatal(e.Start(":8888"))
}
