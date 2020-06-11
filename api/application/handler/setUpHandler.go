package handler

import (
	"todo/api/domain/service"
	"todo/api/infrastructure"
)

var todoService service.TodoServiceInterface

func DI() {
	todoRepo := infrastructure.NewTodoDB(infrastructure.DB)
	todoService = service.NewTodoService(todoRepo)
}
