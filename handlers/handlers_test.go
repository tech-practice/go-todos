package handlers

import (
	"go-todos/apis"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	apis.Client = &apis.MockHttpClient{}

	apis.MockGet = func(string) ([]byte, error) {
		return []byte(`[{"title":"return this title"}]`), nil
	}
	req, _ := http.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(GetTodos)
	h.ServeHTTP(rec, req)
	assert.Contains(t, rec.Body.String(), `"title":"return this title"`)
}
