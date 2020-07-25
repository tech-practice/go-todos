package handlers_test

import (
	"encoding/json"
	"go-todos/handlers"
	"go-todos/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewTodo() string {
	todo := models.Todo{
		UserID:    1,
		Title:     "learning golang",
		Completed: false,
	}
	res, _ := client.Insert(todo)
	return res.ID.(primitive.ObjectID).Hex()
}

func TestUpdateTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id            string
		payload       string
		expectedCode  int
		modifiedCount int64
	}{
		"should return 200 and modified count 1": {
			id:            id,
			payload:       `{"completed":true}`,
			expectedCode:  200,
			modifiedCount: 1,
		},
		"should return 200 and modified count 0": {
			id:            id,
			payload:       `{"title":"learning golang"}`,
			expectedCode:  200,
			modifiedCount: 0,
		},
		"should return 400": {
			id:           "abc",
			expectedCode: 400,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("PATH", "/todos/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", handlers.UpdateTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.TodoUpdate{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.modifiedCount, todo.ModifiedCount)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	//cleanup
	_, _ = client.Delete(id)
}
