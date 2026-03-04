package repository

import (
	"agnos/internal/domain"
	"context"
	// "fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresPatientRepository struct {
	db *gorm.DB
}

func NewPostgresPatientRepository(db *gorm.DB) domain.PatientRepository {
	return &postgresPatientRepository{db: db}
}

func (p *postgresPatientRepository) Create(patient *domain.Patient) error {
	return p.db.Create(patient).Error
}

func (p *postgresPatientRepository) FindPatientRepo(ctx context.Context, hospitalID uuid.UUID, patient *domain.PatientSearchInput) ([]domain.Patient, error) {
	query := p.db.Where("hospital_id = ?", hospitalID)
	if patient.FirstNameTH != nil {
		query = query.Where("first_name_th ilike ?", "%"+*patient.FirstNameTH+"%")
	}
	if patient.MiddleNameTH != nil {
		query = query.Where("middle_name_th ilike ?", "%"+*patient.MiddleNameTH+"%")
	}
	if patient.LastNameTH != nil {
		query = query.Where("last_name_th ilike ?", "%"+*patient.LastNameTH+"%")
	}
	if patient.FirstNameEN != nil {
		query = query.Where("first_name_en ilike ?", "%"+*patient.FirstNameEN+"%")
	}
	if patient.MiddleNameEN != nil {
		query = query.Where("middle_name_en ilike ?", "%"+*patient.MiddleNameEN+"%")
	}
	if patient.LastNameEN != nil {
		query = query.Where("last_name_en ilike ?", "%"+*patient.LastNameEN+"%")
	}
	if patient.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *patient.DateOfBirth)
	}
	if patient.PatientHN != nil {
		query = query.Where("patient_hn = ?", *patient.PatientHN)
	}
	if patient.NationalID != nil {
		query = query.Where("national_id = ?", *patient.NationalID)
	}
	if patient.PassportID != nil {
		query = query.Where("passport_id = ?", *patient.PassportID)
	}
	if patient.PhoneNumber != nil {
		query = query.Where("phone_number = ?", *patient.PhoneNumber)
	}
	if patient.Email != nil {
		query = query.Where("email = ?", *patient.Email)
	}
	if patient.Gender != nil {
		query = query.Where("gender = ?", *patient.Gender)
	}
	var result []domain.Patient
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *postgresPatientRepository) FindPatientByIDRepo(ctx context.Context, id string) (*domain.Patient, error) {
	var result domain.Patient
	if err := p.db.Where("national_id = ?", id).Or("passport_id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
