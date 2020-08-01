package handlers_test

import (
	"go-todos/database"
	"go-todos/handlers"
	"go-todos/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestUpdateTodo(t *testing.T) {
	client := &database.MockTodoClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		payload      string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			payload:      `{"completed": true}`,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Update", test.id, mock.Anything).Return(models.TodoUpdate{}, nil)
			}
			req, _ := http.NewRequest("PATCH", "/todos/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.PATCH("/todos/:id", handlers.UpdateTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Update")
			}
		})
	}
}
