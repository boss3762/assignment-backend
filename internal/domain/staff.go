package domain

import (
	"context"

	// "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Staff struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"unique" json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	HospitalID uuid.UUID `json:"hospital_name" binding:"required"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID" json:"-"`
}

type StaffRepository interface {
	// FindByID(c *gin.Context, id uuid.UUID) (*Staff, error)
	Create(ctx context.Context, p *Staff) error
	// Update(c *gin.Context, p *Staff) error
	// Delete(c *gin.Context, id uuid.UUID) error
	FindByUsername(ctx context.Context, username string) (*Staff, error)
}

type StaffUsecase interface {
	CreateNewStaff(ctx context.Context, input *CreateStaffInput) error
	LoginStaff(ctx context.Context, input *CreateStaffInput) string
}

type CreateStaffInput struct {
	Username     string `gorm:"unique" json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalName string `json:"hospital_id" binding:"required"`
}

func (s *Staff) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
