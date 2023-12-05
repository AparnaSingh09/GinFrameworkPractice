package controller

import (
	"GinFrameworkPractice/entity"
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

var loginController *LoginController = NewLoginController()

func TestLoginController_Login(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Method = "POST"
	creds := entity.Credentials{
		Username: "TW",
		Password: "tw",
	}
	obj, _ := json.Marshal(creds)
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(obj))

	assert.NotEqual(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", loginController.Login(ctx))

}
