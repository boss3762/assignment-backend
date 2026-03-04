package mocks

import (
	"agnos/internal/domain"
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockPatientRepository struct {
	mock.Mock
}

func (m *MockPatientRepository) Create(patient *domain.Patient) error {
	args := m.Called(patient)
	return args.Error(0)
}

func (m *MockPatientRepository) FindPatientRepo(ctx context.Context, hospitalID uuid.UUID, patient *domain.PatientSearchInput) ([]domain.Patient, error) {
	args := m.Called(ctx, hospitalID, patient)
	return args.Get(0).([]domain.Patient), args.Error(1)
}

func (m *MockPatientRepository) FindPatientByIDRepo(ctx context.Context, id string) (*domain.Patient, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Patient), args.Error(1)
}
