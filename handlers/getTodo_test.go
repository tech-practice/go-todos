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

func TestGetTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id           string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
			expected:     "learning golang",
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/todos/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", handlers.GetTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.Todo{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.expected, todo.Title)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	//cleanup
	_, _ = client.Delete(id)
}
