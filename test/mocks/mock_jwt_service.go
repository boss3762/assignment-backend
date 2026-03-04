package mocks

import "github.com/stretchr/testify/mock"

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(username string) string {
	args := m.Called(username)
	return args.String(0)
}

func (m *MockJWTService) ValidateToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}
