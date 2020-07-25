package handlers_test

import (
	"encoding/json"
	"go-todos/handlers"
	"go-todos/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTodo(t *testing.T) {
	id := AddNewTodo()
	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200 - found": {
			payload:      `{"title": "learning golang"}`,
			expectedCode: 200,
			expected:     "learning golang",
		},
		"should return 200 - not found": {
			payload:      `{"title": "learning docker"}`,
			expectedCode: 200,
			expected:     "",
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/todos?q="+test.payload, nil)
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(handlers.SearchTodos(client))
			h.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todos := []models.Todo{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todos)
				for _, todo := range todos {
					assert.Equal(t, test.expected, todo.Title)
				}
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	//cleanup
	_, _ = client.Delete(id)
}
