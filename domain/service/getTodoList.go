package service

import (
	"todo/domain/model"
)

func (t *todoServiceStruct) GetTodoService() ([]model.Todo, error) {
	return t.todoRepo.GetTodoListAll()
}
