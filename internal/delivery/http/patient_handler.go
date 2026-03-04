package http

import (
	"agnos/internal/domain"
	"net/http"

	// "fmt"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientUsecase domain.PatientUsecase
}

func NewPatientHandler(patientUsecase domain.PatientUsecase) *PatientHandler {
	return &PatientHandler{patientUsecase: patientUsecase}
}

// CreateNewPatient godoc
// @Summary      สร้างข้อมูลผู้ป่วยใหม่
// @Description  สร้างผู้ป่วยใหม่ภายใต้โรงพยาบาลของ staff ที่ login อยู่
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      domain.PatientInput  true  "ข้อมูลผู้ป่วย"
// @Success      200    {object}  map[string]string    "Patient created successfully"
// @Failure      400    {object}  map[string]string    "ข้อมูลไม่ถูกต้อง"
// @Failure      500    {object}  map[string]string    "เกิดข้อผิดพลาดภายใน"
// @Router       /patient/create [post]
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

// FindPatient godoc
// @Summary      ค้นหาผู้ป่วย
// @Description  ค้นหาผู้ป่วยภายใต้โรงพยาบาลของ staff ผ่าน filter ต่างๆ (ทุก field เป็น optional)
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      domain.PatientSearchInput  true  "เงื่อนไขการค้นหา (ทุก field optional)"
// @Success      200    {object}  map[string]interface{}     "รายการผู้ป่วยที่พบ"
// @Failure      400    {object}  map[string]string          "ข้อมูลไม่ถูกต้อง"
// @Failure      500    {object}  map[string]string          "เกิดข้อผิดพลาดภายใน"
// @Router       /patient/search [post]
func (p *PatientHandler) FindPatient(c *gin.Context) {
	username := c.MustGet("username").(string)
	input := domain.PatientSearchInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := p.patientUsecase.FindPatient(c.Request.Context(), username, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

// FindPatientByID godoc
// @Summary      ค้นหาผู้ป่วยด้วย ID
// @Description  ค้นหาผู้ป่วยด้วย National ID หรือ Passport ID
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string                 true  "National ID หรือ Passport ID"
// @Success      200  {object}  map[string]interface{} "ข้อมูลผู้ป่วย"
// @Failure      400  {object}  map[string]string      "ไม่ได้ระบุ ID"
// @Failure      500  {object}  map[string]string      "ไม่พบข้อมูล"
// @Router       /patient/search/{id} [get]
func (p *PatientHandler) FindPatientByID(c *gin.Context) {
	id := c.Param("id")
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
