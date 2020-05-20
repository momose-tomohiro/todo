package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to: ", version)

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/register", register)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	schedule := r.Form["schedule"][0]
	priority := r.Form["priority"][0]
	timeLimit := r.Form["time_limit"][0]

	db, err := sql.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO trn_todo(schedule, priority, time_limit) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(schedule, priority, timeLimit)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)

	http.Redirect(w, r, "/", 303)
}
