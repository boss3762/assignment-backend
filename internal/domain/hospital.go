package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hospital struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"unique" json:"name" binding:"required"`
}

type HospitalRepository interface {
	Create(hospital *Hospital) error
	FindByName(name string) (*Hospital, error)
}

type HospitalUsecase interface {
	FindUuid(c *gin.Context) (uuid.UUID, error)
}

func (h *Hospital) BeforeCreate(tx *gorm.DB) error {
    h.ID = uuid.New()
    return nil
}
