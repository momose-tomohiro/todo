package main

import (
	"todo/api/application/handler"
	"todo/api/domain/service"
	"todo/api/infrastructure"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	todoRepo := infrastructure.NewTodoDB(infrastructure.Engine)
	todoService := service.NewTodoService(todoRepo)
	handler := handler.NewTodoHandler(todoService)

	e.GET("/todos", handler.GetTodoList)
	e.POST("/todos", handler.RegisterTodo)
	e.DELETE("/todos/:id", handler.RemoveTodo)

	e.Logger.Fatal(e.Start(":8880"))
}
