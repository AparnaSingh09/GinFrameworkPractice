package server

import (
	"GinFrameworkPractice/entity"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server = SetUpServer()

func TestSetUpServer(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "endpoint working", w.Body.String())
}

func TestFindAll(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	server.ServeHTTP(w, req)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	expected := "[{\"id\":1,\"title\":\"The Bee Sting\",\"author\":\"Paul Murray\",\"price\":\"700\"},{\"id\":2,\"title\":\"North Woods\",\"author\":\"Daniel Mason\",\"price\":\"900\"}]"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), string(actual))

}

func TestFindById(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	server.ServeHTTP(w, req)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	expected := "{\"id\":1,\"title\":\"The Bee Sting\",\"author\":\"Paul Murray\",\"price\":\"700\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), string(actual))

}

func TestAddBook(t *testing.T) {
	w := httptest.NewRecorder()
	testBook := entity.Book{
		Id:     4,
		Title:  "abc",
		Author: "jnv",
		Price:  "600",
	}
	obj, _ := json.Marshal(testBook)
	req, _ := http.NewRequest("POST", "/books", io.NopCloser(bytes.NewBuffer(obj)))
	server.ServeHTTP(w, req)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	expected := "{\"id\":4,\"title\":\"abc\",\"author\":\"jnv\",\"price\":\"600\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), string(actual))
}

func TestRemoveBook(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books?id=1", nil)
	server.ServeHTTP(w, req)
	actual, _ := ioutil.ReadAll(w.Result().Body)
	expected := "{\"id\":1,\"title\":\"The Bee Sting\",\"author\":\"Paul Murray\",\"price\":\"700\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), string(actual))
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	creds := entity.Credentials{
		Username: "TW",
		Password: "tw",
	}
	obj, _ := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/login", io.NopCloser(bytes.NewBuffer(obj)))
	server.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
