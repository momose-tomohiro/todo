package handler

import (
	"todo/domain/model"
	"todo/domain/service"
)

type todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"time_limit"`
}

//一覧表示を返す
func Display() []model.Todo {
	return service.GetTodoList()
}
