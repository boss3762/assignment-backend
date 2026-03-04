package mocks

import (
	"agnos/internal/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockPatientUsecase struct {
	mock.Mock
}

func (m *MockPatientUsecase) CreateNewPatient(ctx context.Context, staffname string, patient *domain.PatientInput) error {
	args := m.Called(ctx, staffname, patient)
	return args.Error(0)
}

func (m *MockPatientUsecase) FindPatient(ctx context.Context, hospitalName string, patient *domain.PatientSearchInput) ([]domain.Patient, error) {
	args := m.Called(ctx, hospitalName, patient)
	return args.Get(0).([]domain.Patient), args.Error(1)
}

func (m *MockPatientUsecase) FindPatientByID(ctx context.Context, id string) (*domain.Patient, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Patient), args.Error(1)
}
