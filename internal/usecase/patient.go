package usecase

import (
	"agnos/internal/domain"
	"context"
	"github.com/google/uuid"
	// "fmt"
)

type patientUsecase struct {
	patientRepo domain.PatientRepository
	staffRepo domain.StaffRepository
}

func NewPatientUsecase(patientRepo domain.PatientRepository, staffRepo domain.StaffRepository) domain.PatientUsecase {
	return &patientUsecase{patientRepo: patientRepo, staffRepo: staffRepo}
}

func convertToPatient(patient *domain.PatientInput, hospitalID uuid.UUID) *domain.Patient {
	return &domain.Patient{
		HospitalID: hospitalID,
		FirstNameTH: patient.FirstNameTH,
		MiddleNameTH: patient.MiddleNameTH,
		LastNameTH: patient.LastNameTH,
		FirstNameEN: patient.FirstNameEN,
		MiddleNameEN: patient.MiddleNameEN,
		LastNameEN: patient.LastNameEN,
		DateOfBirth: patient.DateOfBirth,
		PatientHN: patient.PatientHN,
		NationalID: patient.NationalID,
		PassportID: patient.PassportID,
		PhoneNumber: patient.PhoneNumber,
		Email: patient.Email,
		Gender: patient.Gender,
	}
}

func (p *patientUsecase) CreateNewPatient(ctx context.Context, staffname string, patient *domain.PatientInput) error {
	staff, err := p.staffRepo.FindByUsername(ctx, staffname)
	if err != nil {
		return err
	}
	return p.patientRepo.Create(convertToPatient(patient, staff.HospitalID))
}

func (p *patientUsecase) FindPatient(ctx context.Context, staffname string, patient *domain.PatientInput) (*domain.Patient, error) {
	staff, err := p.staffRepo.FindByUsername(ctx, staffname)
	if err != nil {
		return nil, err
	}
	return p.patientRepo.FindPatientRepo(ctx,staff.HospitalID, patient)
}

