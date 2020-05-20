package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type todoList struct {
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"timeLimit"`
}

func main() {
	db, err := sql.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/register", register)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	var form todoList

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	fmt.Println(string(b))
	json.Unmarshal(b, &form)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	schedule := form.Schedule
	priority := form.Priority
	timeLimit := form.TimeLimit

	db, err := sql.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO trn_todo(schedule, priority, time_limit) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	defer stmt.Close()

	result, err := stmt.Exec(schedule, priority, timeLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	fmt.Println(id)

}
