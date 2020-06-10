package handler

import (
	"todo/domain/service"
	"todo/infrastructure"
)

var todoService service.TodoServiceInterface

func DI() {
	todoRepo := infrastructure.NewTodoDB(infrastructure.DB)
	todoService = service.NewTodoService(todoRepo)
}
