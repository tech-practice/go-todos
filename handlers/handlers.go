package handlers

import (
	"encoding/json"
	"go-todos/models"
	"io/ioutil"
	"net/http"
)

var URL = "https://jsonplaceholder.typicode.com/todos"

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []models.Todo{}

	body, err := httpGet(URL)
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

var httpGet = func(url string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}
