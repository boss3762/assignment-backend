package usecase

import (
	"agnos/internal/auth"
	"agnos/internal/domain"

	"agnos/config"
	"context"

	"golang.org/x/crypto/bcrypt"
	// "github.com/gin-gonic/gin"
)

type staffUsecase struct {
	repo domain.StaffRepository
	jwtService auth.JWTService
}

func NewStaffUsecase(repo domain.StaffRepository, jwtService auth.JWTService) domain.StaffUsecase {
	return &staffUsecase{repo: repo, jwtService: jwtService}
}

func (s *staffUsecase) CreateNewStaff(ctx context.Context, input *domain.CreateStaffInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถ hash password ได้"})
		return err
	}

	// หาโรงพยาบาล ถ้าไม่มีให้สร้างใหม่
	var hospital domain.Hospital
	if err := config.DB.Where("name = ?", input.HospitalName).FirstOrCreate(&hospital, domain.Hospital{Name: input.HospitalName}).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการเชื่อมต่อหรือบันทึกข้อมูลโรงพยาบาล"})
		return err
	}

	// สร้าง staff
	staff := domain.Staff{
		Username:   input.Username,
		Password:   string(hashedPassword),
		HospitalID: hospital.ID,
	}

	if err := s.repo.Create(ctx, &staff); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการเชื่อมต่อหรือบันทึกข้อมูล"})
		return err
	}

	// c.JSON(http.StatusOK, gin.H{"message": "สร้าง staff สำเร็จ"})
	return nil
}

func (s *staffUsecase) LoginStaff(ctx context.Context, input *domain.CreateStaffInput) string {
	staff, err := s.repo.FindByUsernameHospitalname(ctx,input)
	if err != nil {
		return ""
	}
	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(input.Password)); err != nil {
		return ""
	}

	return s.jwtService.GenerateToken(staff.Username)
}
