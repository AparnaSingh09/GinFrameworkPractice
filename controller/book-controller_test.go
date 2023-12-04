package controller

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/repository"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var book = entity.Book{
	Id:     1,
	Title:  "The Bee Sting",
	Author: "Paul Murray",
	Price:  "700",
}
var bookController *BookController = New()

func TestBookController_FindAll(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookController.FindAll())
}

func TestBookController_FindBookById(t *testing.T) {
	assert.Equal(t, book, bookController.FindBookById(1))
}

func TestBookController_AddBook(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Method = "POST"
	testBook := entity.Book{
		Id:     4,
		Title:  "abc",
		Author: "jnv",
		Price:  "600",
	}
	obj, _ := json.Marshal(testBook)
	ctx.Request.Header.Set("Content-Type", "application/json")

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(obj))
	assert.Equal(t, testBook, bookController.AddBook(ctx))
}
