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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupPatientRouter(mockUsecase *mocks.MockPatientUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := http.NewPatientHandler(mockUsecase)
	// inject username เหมือน middleware ทำ
	r.Use(func(c *gin.Context) {
		c.Set("username", "nurse01")
		c.Next()
	})
	r.POST("/patient/create", h.CreateNewPatient)
	r.POST("/patient/search", h.FindPatient)
	r.GET("/patient/search/:id", h.FindPatientByID)
	return r
}

func TestPatientHandler_CreateNewPatient_Success(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	input := domain.PatientInput{
		FirstNameTH: "สมชาย", LastNameTH: "ใจดี",
		FirstNameEN: "Somchai", LastNameEN: "Jaidee",
		PatientHN: "HN001", NationalID: "1234567890123",
	}
	mockUC.On("CreateNewPatient", context.Background(), "nurse01", &input).Return(nil)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/patient/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Patient created successfully")
	mockUC.AssertExpectations(t)
}

func TestPatientHandler_CreateNewPatient_BadBody_Returns400(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	req := httptest.NewRequest(nethttp.MethodPost, "/patient/create", bytes.NewBufferString("not-json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusBadRequest, w.Code)
	mockUC.AssertNotCalled(t, "CreateNewPatient")
}

func TestPatientHandler_CreateNewPatient_UsecaseError_Returns500(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	input := domain.PatientInput{
		FirstNameTH: "สมชาย", LastNameTH: "ใจดี",
		FirstNameEN: "Somchai", LastNameEN: "Jaidee",
		PatientHN: "HN001", NationalID: "1234567890123",
	}
	mockUC.On("CreateNewPatient", context.Background(), "nurse01", &input).
		Return(errors.New("db error"))

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(nethttp.MethodPost, "/patient/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusInternalServerError, w.Code)
}

func TestPatientHandler_FindPatient_Success(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	searchInput := domain.PatientSearchInput{}
	patients := []domain.Patient{{FirstNameTH: "สมชาย", NationalID: "1234567890123"}}
	mockUC.On("FindPatient", context.Background(), "nurse01", &searchInput).
		Return(patients, nil)

	body, _ := json.Marshal(searchInput)
	req := httptest.NewRequest(nethttp.MethodPost, "/patient/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "สมชาย")
	mockUC.AssertExpectations(t)
}

func TestPatientHandler_FindPatient_UsecaseError_Returns500(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	searchInput := domain.PatientSearchInput{}
	mockUC.On("FindPatient", context.Background(), "nurse01", &searchInput).
		Return([]domain.Patient{}, errors.New("db error"))

	body, _ := json.Marshal(searchInput)
	req := httptest.NewRequest(nethttp.MethodPost, "/patient/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusInternalServerError, w.Code)
}

func TestPatientHandler_FindPatientByID_Success(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	patientID := uuid.New()
	patient := &domain.Patient{ID: patientID, NationalID: "1234567890123"}
	mockUC.On("FindPatientByID", context.Background(), "1234567890123").
		Return(patient, nil)

	req := httptest.NewRequest(nethttp.MethodGet, "/patient/search/1234567890123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "1234567890123")
	mockUC.AssertExpectations(t)
}

func TestPatientHandler_FindPatientByID_NotFound_Returns500(t *testing.T) {
	mockUC := new(mocks.MockPatientUsecase)
	r := setupPatientRouter(mockUC)

	mockUC.On("FindPatientByID", context.Background(), "XXXX").
		Return(nil, errors.New("record not found"))

	req := httptest.NewRequest(nethttp.MethodGet, "/patient/search/XXXX", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, nethttp.StatusInternalServerError, w.Code)
}
