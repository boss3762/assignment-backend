package test

import (
	"agnos/internal/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken_ReturnsNonEmptyString(t *testing.T) {
	svc := auth.NewAuthService()
	token := svc.GenerateToken("john_doe")
	assert.NotEmpty(t, token, "token ที่ได้ต้องไม่ว่าง")
}

func TestValidateToken_ValidToken_ReturnsUsername(t *testing.T) {
	svc := auth.NewAuthService()
	token := svc.GenerateToken("john_doe")

	username, err := svc.ValidateToken(token)

	assert.NoError(t, err)
	assert.Equal(t, "john_doe", username)
}

func TestValidateToken_InvalidToken_ReturnsError(t *testing.T) {
	svc := auth.NewAuthService()

	username, err := svc.ValidateToken("this.is.not.a.valid.token")

	assert.Error(t, err)
	assert.Empty(t, username)
}

func TestValidateToken_TamperedToken_ReturnsError(t *testing.T) {
	svc := auth.NewAuthService()
	token := svc.GenerateToken("user1")
	// แก้ token ให้ผิด
	tampered := token + "x"

	_, err := svc.ValidateToken(tampered)

	assert.Error(t, err)
}
