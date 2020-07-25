package handlers_test

import (
	"encoding/json"
	"go-todos/handlers"
	"go-todos/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id           string
		expectedCode int
		deletedCount int64
	}{
		"should return 200 and deleted count 1": {
			id:           id,
			expectedCode: 200,
			deletedCount: 1,
		},
		"should return 200 and modified count 0": {
			id:           id,
			expectedCode: 200,
			deletedCount: 0,
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
			req, _ := http.NewRequest("DELETE", "/todos/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.TodoDelete{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.deletedCount, todo.DeletedCount)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	//cleanup
	_, _ = client.Delete(id)
}
