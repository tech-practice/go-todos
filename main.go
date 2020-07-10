package main

import (
	"go-todos/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", handlers.GetTodos)

	http.ListenAndServe(":8080", r)
}
