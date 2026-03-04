package http

import (
	"agnos/internal/domain"
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientUsecase domain.PatientUsecase
}

func NewPatientHandler(patientUsecase domain.PatientUsecase) *PatientHandler {
	return &PatientHandler{patientUsecase: patientUsecase}
}

func (p *PatientHandler) CreateNewPatient(c *gin.Context) {
	staffname := c.MustGet("username").(string)
	var input domain.PatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.patientUsecase.CreateNewPatient(c.Request.Context(), staffname, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully"})
}

func (p *PatientHandler) FindPatient(c *gin.Context) {
	username := c.MustGet("username").(string)
	input := domain.PatientSearchInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient,err := p.patientUsecase.FindPatient(c.Request.Context(), username, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

func (p *PatientHandler) FindPatientByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	patient, err := p.patientUsecase.FindPatientByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}
