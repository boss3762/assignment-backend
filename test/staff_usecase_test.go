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
	"golang.org/x/crypto/bcrypt"
)

func TestLoginStaff_Success_ReturnsToken(t *testing.T) {
	mockStaffRepo := new(mocks.MockStaffRepository)
	mockJWT := new(mocks.MockJWTService)
	uc := usecase.NewStaffUsecase(mockStaffRepo, mockJWT)

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte("secret123"), bcrypt.DefaultCost)
	input := &domain.CreateStaffInput{
		Username:     "nurse01",
		Password:     "secret123",
		HospitalName: "",
	}
	staff := &domain.Staff{
		ID:         uuid.New(),
		Username:   "nurse01",
		Password:   string(hashedPw),
		HospitalID: uuid.New(),
	}

	mockStaffRepo.On("FindByUsernameHospitalname", context.Background(), input).Return(staff, nil)
	mockJWT.On("GenerateToken", "nurse01").Return("jwt.token.here")

	token := uc.LoginStaff(context.Background(), input)

	assert.Equal(t, "jwt.token.here", token)
	mockStaffRepo.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestLoginStaff_WrongPassword_ReturnsEmpty(t *testing.T) {
	mockStaffRepo := new(mocks.MockStaffRepository)
	mockJWT := new(mocks.MockJWTService)
	uc := usecase.NewStaffUsecase(mockStaffRepo, mockJWT)

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte("correct_password"), bcrypt.DefaultCost)
	staff := &domain.Staff{Username: "nurse01", Password: string(hashedPw)}
	input := &domain.CreateStaffInput{
		Username:     "nurse01",
		Password:     "wrong_password",
		HospitalName: "",
	}
	mockStaffRepo.On("FindByUsernameHospitalname", context.Background(), input).Return(staff, nil)

	token := uc.LoginStaff(context.Background(), input)

	assert.Empty(t, token)
	mockJWT.AssertNotCalled(t, "GenerateToken")
}

func TestLoginStaff_UserNotFound_ReturnsEmpty(t *testing.T) {
	mockStaffRepo := new(mocks.MockStaffRepository)
	mockJWT := new(mocks.MockJWTService)
	uc := usecase.NewStaffUsecase(mockStaffRepo, mockJWT)

	input := &domain.CreateStaffInput{
		Username:     "nobody",
		Password:     "any",
		HospitalName: "",
	}
	mockStaffRepo.On("FindByUsernameHospitalname", context.Background(), input).
		Return(nil, errors.New("record not found"))

	token := uc.LoginStaff(context.Background(), input)

	assert.Empty(t, token)
	mockJWT.AssertNotCalled(t, "GenerateToken")
}
