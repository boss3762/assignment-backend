package repository

import (
	"gorm.io/gorm"
	"agnos/internal/domain"
	"context"
	"github.com/google/uuid"
	"fmt"
)

type postgresPatientRepository struct {
	db *gorm.DB
}

func NewPostgresPatientRepository(db *gorm.DB) domain.PatientRepository {
	return &postgresPatientRepository{db: db}
}

func (p *postgresPatientRepository) Create(patient *domain.Patient) error {
	fmt.Printf("%+v\n", patient)
	// fmt.Println(p.db.Create(patient).Error)
	return p.db.Create(patient).Error
}

func (p *postgresPatientRepository) FindPatientRepo(ctx context.Context, hospitalID uuid.UUID, patient *domain.PatientInput) (*domain.Patient, error) {
	fmt.Println(patient)
	query := p.db.Where("hospital_id = ?", hospitalID).Where(patient)
	if err := query.Error; err != nil {
		return nil, err
	}
	fmt.Println(query)
	return &domain.Patient{}, nil
}


