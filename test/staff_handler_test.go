package test

import (
	"agnos/internal/delivery/http"
	"agnos/internal/domain"
	"agnos/test/mocks"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupStaffRouter(mockUsecase *mocks.MockStaffUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := http.NewStaffHandler(mockUsecase)
	r.POST("/staff/register", h.CreateNewStaff)
	r.POST("/staff/login", h.LoginStaff)
	return r
}

func TestStaffHandler_CreateNewStaff_Success(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	input := domain.CreateStaffInput{
		Username:     "nurse01",
		Password:     "secret123",
		HospitalName: "โรงพยาบาลกรุงเทพ",
	}
	mockUC.On("CreateNewStaff", context.Background(), &input).Return(nil)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/staff/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "สร้าง staff สำเร็จ")
	mockUC.AssertExpectations(t)
}

func TestStaffHandler_CreateNewStaff_BadBody_Returns400(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	req := httptest.NewRequest(nethttp.MethodPost, "/staff/register", bytes.NewBufferString("not-json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusBadRequest, w.Code)
	mockUC.AssertNotCalled(t, "CreateNewStaff")
}

func TestStaffHandler_CreateNewStaff_UsecaseError_Returns500(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	input := domain.CreateStaffInput{
		Username:     "nurse01",
		Password:     "secret123",
		HospitalName: "โรงพยาบาลกรุงเทพ",
	}
	mockUC.On("CreateNewStaff", context.Background(), &input).Return(errors.New("db error"))

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/staff/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusInternalServerError, w.Code)
}

func TestStaffHandler_LoginStaff_Success(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	input := domain.CreateStaffInput{
		Username:     "nurse01",
		Password:     "secret123",
		HospitalName: "โรงพยาบาลกรุงเทพ",
	}
	mockUC.On("LoginStaff", context.Background(), &input).Return("jwt.token.string")

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/staff/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "jwt.token.string")
	mockUC.AssertExpectations(t)
}

func TestStaffHandler_LoginStaff_WrongCredentials_Returns500(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	input := domain.CreateStaffInput{Username: "nurse01", Password: "wrong", HospitalName: "โรงพยาบาลกรุงเทพ"}
	mockUC.On("LoginStaff", context.Background(), &input).Return("")

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/staff/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusInternalServerError, w.Code)
	mockUC.AssertExpectations(t)
}

func TestStaffHandler_LoginStaff_BadBody_Returns400(t *testing.T) {
	mockUC := new(mocks.MockStaffUsecase)
	r := setupStaffRouter(mockUC)

	req := httptest.NewRequest(nethttp.MethodPost, "/staff/login", bytes.NewBufferString("bad"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusBadRequest, w.Code)
}
