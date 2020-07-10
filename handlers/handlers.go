package handlers

import (
	"encoding/json"
	"go-todos/apis"
	"go-todos/models"
	"net/http"
)

var URL = "https://jsonplaceholder.typicode.com/todos"

func GetTodos(client apis.HttpInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := []models.Todo{}

		body, err := client.Get(URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		err = json.Unmarshal(body, &todos)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}
