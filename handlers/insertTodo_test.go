package handlers_test

import (
	"context"
	"encoding/json"
	"go-todos/config"
	"go-todos/database"
	"go-todos/handlers"
	"go-todos/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client database.TodoInterface

func init() {
	conf := config.MongoConfiguration{
		Server:     "mongodb://localhost:27017",
		Database:   "Mgo",
		Collection: "TodosTest",
	}
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf)
	collection := db.Collection(conf.Collection)

	client = &database.TodoClient{
		Col: collection,
		Ctx: ctx,
	}
}

func TestInsertTodo(t *testing.T) {
	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			payload:      `{"userId":1,"title":"learning golang","completed":false}`,
			expectedCode: 200,
			expected:     "learning golang",
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/todos", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(handlers.InsertTodo(client))
			h.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.Todo{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.expected, todo.Title)
				assert.NotNil(t, todo.ID)
				//cleanup
				_, _ = client.Delete(todo.ID.(string))
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}
