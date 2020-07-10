package handlers

import (
	"go-todos/apis"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	client := &apis.MockHttpClient{}
	client.On("Get", URL).Return([]byte(`[{"title":"return this title"}]`), nil)

	req, _ := http.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(GetTodos(client))
	h.ServeHTTP(rec, req)

	assert.Contains(t, rec.Body.String(), `"title":"return this title"`)
	client.AssertExpectations(t)
}
