package test

import (
	"agnos/internal/domain"
	"agnos/internal/usecase"
	"agnos/test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewPatient_Success(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	hospitalID := uuid.New()
	staff := &domain.Staff{Username: "nurse01", HospitalID: hospitalID}
	input := &domain.PatientInput{
		FirstNameTH: "สมชาย", LastNameTH: "ใจดี",
		FirstNameEN: "Somchai", LastNameEN: "Jaidee",
		PatientHN: "HN001", NationalID: "1234567890123",
	}

	mockStaffRepo.On("FindByUsername", context.Background(), "nurse01").Return(staff, nil)
	mockPatientRepo.On("Create", &domain.Patient{
		HospitalID:  hospitalID,
		FirstNameTH: "สมชาย", LastNameTH: "ใจดี",
		FirstNameEN: "Somchai", LastNameEN: "Jaidee",
		PatientHN: "HN001", NationalID: "1234567890123",
	}).Return(nil)

	err := uc.CreateNewPatient(context.Background(), "nurse01", input)

	assert.NoError(t, err)
	mockStaffRepo.AssertExpectations(t)
	mockPatientRepo.AssertExpectations(t)
}

func TestCreateNewPatient_StaffNotFound_ReturnsError(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	mockStaffRepo.On("FindByUsername", context.Background(), "unknown").
		Return(nil, errors.New("staff not found"))

	err := uc.CreateNewPatient(context.Background(), "unknown", &domain.PatientInput{})

	assert.Error(t, err)
	assert.EqualError(t, err, "staff not found")
	mockStaffRepo.AssertExpectations(t)
	mockPatientRepo.AssertNotCalled(t, "Create")
}

func TestFindPatient_Success(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	hospitalID := uuid.New()
	staff := &domain.Staff{Username: "nurse01", HospitalID: hospitalID}
	searchInput := &domain.PatientSearchInput{}
	expected := []domain.Patient{{FirstNameTH: "สมชาย"}}

	mockStaffRepo.On("FindByUsername", context.Background(), "nurse01").Return(staff, nil)
	mockPatientRepo.On("FindPatientRepo", context.Background(), hospitalID, searchInput).
		Return(expected, nil)

	result, err := uc.FindPatient(context.Background(), "nurse01", searchInput)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockStaffRepo.AssertExpectations(t)
	mockPatientRepo.AssertExpectations(t)
}

func TestFindPatient_StaffNotFound_ReturnsError(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	mockStaffRepo.On("FindByUsername", context.Background(), "nobody").
		Return(nil, errors.New("staff not found"))

	result, err := uc.FindPatient(context.Background(), "nobody", &domain.PatientSearchInput{})

	assert.Error(t, err)
	assert.Nil(t, result)
	mockPatientRepo.AssertNotCalled(t, "FindPatientRepo")
}

func TestFindPatientByID_Success(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	expected := &domain.Patient{NationalID: "1234567890123"}
	mockPatientRepo.On("FindPatientByIDRepo", context.Background(), "1234567890123").
		Return(expected, nil)

	result, err := uc.FindPatientByID(context.Background(), "1234567890123")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockPatientRepo.AssertExpectations(t)
}

func TestFindPatientByID_NotFound_ReturnsError(t *testing.T) {
	mockPatientRepo := new(mocks.MockPatientRepository)
	mockStaffRepo := new(mocks.MockStaffRepository)
	uc := usecase.NewPatientUsecase(mockPatientRepo, mockStaffRepo)

	mockPatientRepo.On("FindPatientByIDRepo", context.Background(), "XXXX").
		Return(nil, errors.New("record not found"))

	result, err := uc.FindPatientByID(context.Background(), "XXXX")

	assert.Error(t, err)
	assert.Nil(t, result)
}
