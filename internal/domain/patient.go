package domain

import (
	"github.com/google/uuid"
	"context"
	"gorm.io/gorm"
)

type Patient struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	HospitalID uuid.UUID `json:"hospital_id" binding:"required"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID" json:"-"`

	FirstNameTH string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH  string `json:"last_name_th" binding:"required"`
	FirstNameEN string `json:"first_name_en" binding:"required"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN  string `json:"last_name_en" binding:"required"`
	DateOfBirth string `json:"date_of_birth"`
	PatientHN string `json:"patient_hn" binding:"required"`
	NationalID string `json:"national_id" binding:"required"`
	PassportID string `json:"passport_id"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}

type PatientInput struct {
	FirstNameTH string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH  string `json:"last_name_th" binding:"required"`
	FirstNameEN string `json:"first_name_en" binding:"required"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN  string `json:"last_name_en" binding:"required"`
	DateOfBirth string `json:"date_of_birth"`
	PatientHN string `json:"patient_hn" binding:"required"`
	NationalID string `json:"national_id" binding:"required"`
	PassportID string `json:"passport_id"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}

type PatientRepository interface {
	Create(patient *Patient) error
	FindPatientRepo(ctx context.Context, hospitalID uuid.UUID, patient *PatientInput) (*Patient, error)
}

type PatientUsecase interface {
	CreateNewPatient(ctx context.Context, staffname string, patient *PatientInput) error
	FindPatient(ctx context.Context, hospitalName string, patient *PatientInput) (*Patient, error)
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}


