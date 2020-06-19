package model

type Todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"time_limit"`
}
