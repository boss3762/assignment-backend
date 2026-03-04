package mocks

import (
	"agnos/internal/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockStaffRepository struct {
	mock.Mock
}

func (m *MockStaffRepository) Create(ctx context.Context, staff *domain.Staff) error {
	args := m.Called(ctx, staff)
	return args.Error(0)
}

func (m *MockStaffRepository) FindByUsername(ctx context.Context, username string) (*domain.Staff, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Staff), args.Error(1)
}
