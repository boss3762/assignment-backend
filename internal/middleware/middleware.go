package middleware

import (
	"agnos/internal/auth"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	authService auth.JWTService
}

func NewMiddleware(authService auth.JWTService) *Middleware {
	return &Middleware{authService: authService}
}

func (m *Middleware) ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	token = token[7:]
	username, err := m.authService.ValidateToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("username", username)
	c.Next()
}
