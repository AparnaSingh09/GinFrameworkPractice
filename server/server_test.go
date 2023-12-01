package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetUpServer(t *testing.T) {
	server := SetUpServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "endpoint working", w.Body.String())
}
