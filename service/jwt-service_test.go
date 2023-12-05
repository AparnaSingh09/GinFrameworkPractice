package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var jwtService *JWTService = NewJWTService()

func TestJWTService_GenerateToken(t *testing.T) {
	assert.NotEqual(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", jwtService.GenerateToken("TW"))
}
