package mocks

import (
	"agnos/internal/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockStaffUsecase struct {
	mock.Mock
}

func (m *MockStaffUsecase) CreateNewStaff(ctx context.Context, input *domain.CreateStaffInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

func (m *MockStaffUsecase) LoginStaff(ctx context.Context, input *domain.CreateStaffInput) string {
	args := m.Called(ctx, input)
	return args.String(0)
}
