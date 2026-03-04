package http

import (
	"agnos/internal/domain"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	uc domain.StaffUsecase
}

func NewStaffHandler(uc domain.StaffUsecase) *StaffHandler {
	return &StaffHandler{uc: uc}
}

func (h *StaffHandler) CreateNewStaff(c *gin.Context) {
	var input domain.CreateStaffInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateNewStaff(c.Request.Context(), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "สร้าง staff สำเร็จ"})
}

func (h *StaffHandler) LoginStaff(c *gin.Context) {
	var input domain.CreateStaffInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := h.uc.LoginStaff(c.Request.Context(), &input)
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการเชื่อมต่อหรือบันทึกข้อมูล"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "เข้าสู่ระบบสำเร็จ", "access_token": token})
}
