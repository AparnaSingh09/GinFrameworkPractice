package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var loginService *LoginService

func TestLoginService_Login(t *testing.T) {
	assert.Equal(t, true, loginService.Login("TW", "tw"))
}
