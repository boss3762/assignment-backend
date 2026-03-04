package domain

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	HospitalID uuid.UUID `json:"hospital_id" binding:"required"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID" json:"-"`

	FirstNameTH  string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th" binding:"required"`
	FirstNameEN  string `json:"first_name_en" binding:"required"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en" binding:"required"`
	DateOfBirth  string `json:"date_of_birth"`
	PatientHN    string `gorm:"unique" json:"patient_hn" binding:"required"`
	NationalID   string `gorm:"unique" json:"national_id" binding:"required"`
	PassportID   string `gorm:"unique;default:null" json:"passport_id"`
	PhoneNumber  string `gorm:"unique" json:"phone_number"`
	Email        string `gorm:"unique" json:"email"`
	Gender       string `json:"gender"`
}

type PatientInput struct {
	FirstNameTH  string `json:"first_name_th" binding:"required"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th" binding:"required"`
	FirstNameEN  string `json:"first_name_en" binding:"required"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en" binding:"required"`
	DateOfBirth  string `json:"date_of_birth"`
	PatientHN    string `gorm:"unique" json:"patient_hn" binding:"required"`
	NationalID   string `gorm:"unique" json:"national_id" binding:"required"`
	PassportID   string `gorm:"unique" json:"passport_id"`
	PhoneNumber  string `gorm:"unique" json:"phone_number"`
	Email        string `gorm:"unique" json:"email"`
	Gender       string `json:"gender"`
}

type PatientSearchInput struct {
	FirstNameTH  *string `json:"first_name_th"`
	MiddleNameTH *string `json:"middle_name_th"`
	LastNameTH   *string `json:"last_name_th"`
	FirstNameEN  *string `json:"first_name_en"`
	MiddleNameEN *string `json:"middle_name_en"`
	LastNameEN   *string `json:"last_name_en"`
	DateOfBirth  *string `json:"date_of_birth"`
	PatientHN    *string `json:"patient_hn"`
	NationalID   *string `json:"national_id"`
	PassportID   *string `json:"passport_id"`
	PhoneNumber  *string `json:"phone_number"`
	Email        *string `json:"email"`
	Gender       *string `json:"gender"`
}

type PatientRepository interface {
	Create(patient *Patient) error
	FindPatientRepo(ctx context.Context, hospitalID uuid.UUID, patient *PatientSearchInput) ([]Patient, error)
	FindPatientByIDRepo(ctx context.Context, id string) (*Patient, error)
}

type PatientUsecase interface {
	CreateNewPatient(ctx context.Context, staffname string, patient *PatientInput) error
	FindPatient(ctx context.Context, hospitalName string, patient *PatientSearchInput) ([]Patient, error)
	FindPatientByID(ctx context.Context, id string) (*Patient, error)
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
