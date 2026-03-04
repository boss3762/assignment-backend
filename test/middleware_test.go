package test

import (
	"agnos/internal/middleware"
	"agnos/test/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupMiddlewareRouter(m *middleware.Middleware) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/protected", m.ValidateToken, func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"username": username})
	})
	return r
}

func TestMiddleware_MissingAuthHeader_Returns401(t *testing.T) {
	mockJWT := new(mocks.MockJWTService)
	m := middleware.NewMiddleware(mockJWT)
	r := setupMiddlewareRouter(m)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestMiddleware_InvalidToken_Returns401(t *testing.T) {
	mockJWT := new(mocks.MockJWTService)
	mockJWT.On("ValidateToken", "bad_token").Return("", errors.New("invalid token"))

	m := middleware.NewMiddleware(mockJWT)
	r := setupMiddlewareRouter(m)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer bad_token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockJWT.AssertExpectations(t)
}

func TestMiddleware_ValidToken_PassesThrough(t *testing.T) {
	mockJWT := new(mocks.MockJWTService)
	mockJWT.On("ValidateToken", "valid.jwt.token").Return("john_doe", nil)

	m := middleware.NewMiddleware(mockJWT)
	r := setupMiddlewareRouter(m)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid.jwt.token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "john_doe")
	mockJWT.AssertExpectations(t)
}
