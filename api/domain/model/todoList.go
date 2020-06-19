package model

type Todo struct {
	ID        int    `json:"id" xorm:"'id'"`
	Schedule  string `json:"schedule" xorm:"'schedule'"`
	Priority  string `json:"priority" xorm:"'priority'"`
	TimeLimit string `json:"time_limit" xorm:"'time_limit'"`
}
