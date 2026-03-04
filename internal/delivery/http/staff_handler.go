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

// CreateNewStaff godoc
// @Summary      สร้าง staff ใหม่
// @Description  ลงทะเบียน staff พร้อมระบุโรงพยาบาล (ถ้าโรงพยาบาลยังไม่มีจะสร้างให้อัตโนมัติ)
// @Tags         Staff
// @Accept       json
// @Produce      json
// @Param        input  body      domain.CreateStaffInput  true  "ข้อมูล Staff"
// @Success      200    {object}  map[string]string        "สร้าง staff สำเร็จ"
// @Failure      400    {object}  map[string]string        "ข้อมูลไม่ถูกต้อง"
// @Failure      500    {object}  map[string]string        "เกิดข้อผิดพลาดภายใน"
// @Router       /staff/create [post]
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

// LoginStaff godoc
// @Summary      เข้าสู่ระบบ
// @Description  ตรวจสอบ username/password แล้วคืน JWT access token
// @Tags         Staff
// @Accept       json
// @Produce      json
// @Param        input  body      domain.CreateStaffInput  true  "ข้อมูล Login"
// @Success      200    {object}  map[string]string        "access_token"
// @Failure      400    {object}  map[string]string        "ข้อมูลไม่ถูกต้อง"
// @Failure      500    {object}  map[string]string        "username หรือ password ไม่ถูกต้อง"
// @Router       /staff/login [post]
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
