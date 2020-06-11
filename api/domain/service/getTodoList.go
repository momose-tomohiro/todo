package service

import (
	"todo/api/domain/model"
)

func (t *todoServiceStruct) GetTodoService() ([]model.Todo, error) {
	return t.todoRepo.GetTodoListAll()
}
