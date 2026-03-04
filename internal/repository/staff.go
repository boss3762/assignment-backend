package repository

import (
	"agnos/config"
	"agnos/internal/domain"
	"context"

	"gorm.io/gorm"
)

type postgresStaffRepository struct {
	db *gorm.DB
}

func NewPostgresStaffRepository(db *gorm.DB) domain.StaffRepository {
	return &postgresStaffRepository{db: db}
}

func (p *postgresStaffRepository) Create(ctx context.Context, staff *domain.Staff) error {
	if err := config.DB.Create(staff).Error; err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Username นี้มีอยู่ในระบบแล้ว หรือเกิดข้อผิดพลาดอื่น"})
		return err
	}
	// c.JSON(http.StatusOK, gin.H{"message": "สร้าง staff สำเร็จ"})
	return nil
}

func (p *postgresStaffRepository) FindByUsername(ctx context.Context, username string) (*domain.Staff, error) {
	var staff domain.Staff
	if err := config.DB.Preload("Hospital").Where("username = ?", username).First(&staff).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (p *postgresStaffRepository) FindByUsernameHospitalname(ctx context.Context, input *domain.CreateStaffInput) (*domain.Staff, error) {
	var query domain.Staff
	err := config.DB.
		Preload("Hospital").
		Joins("JOIN hospitals ON hospitals.id = staffs.hospital_id").
		Where("staffs.username = ?", input.Username).
		Where("hospitals.name = ?", input.HospitalName).
		First(&query).Error
	if err != nil {
		return nil, err
	}
	return &query, nil
}
