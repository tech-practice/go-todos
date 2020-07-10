package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	req, _ := http.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(GetTodos)
	h.ServeHTTP(rec, req)
	assert.Contains(t, rec.Body.String(), `"title":"delectus aut autem"`)
}
