package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var users = []User{
	User{"Rodrigo Kochenburger", "rodrigo@example.com"},
	User{"John Smith", "john@example.com"},
	User{"Jane Doe", "jane@example.com"},
	User{"Mary Jane", "mary@example.com"},
}

type User struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	j := json.NewEncoder(w)
	if err := j.Encode(users); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/users", GetUsersHandler)
	log.Println("Listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
