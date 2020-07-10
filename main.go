package main

import (
	"go-todos/apis"
	"go-todos/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	client := &apis.HttpClient{}
	r.HandleFunc("/todos", handlers.GetTodos(client))

	http.ListenAndServe(":8080", r)
}
