package service

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
)

type SectionService struct {
	r repository.SectionRepository
}

func ProvideSectionService(r repository.SectionRepository) SectionService {
	return SectionService{
		r: r,
	}
}

func (t *SectionService) Create(section models.Section) (*models.Section, error) {
	return t.r.Create(section)
}

func (t *SectionService) Update(section models.Section) (*models.Section, error) {
	return t.r.Update(section)
}

func (t *SectionService) DeleteById(id uuid.UUID) error {
	return t.r.DeleteById(id)
}

func (t *SectionService) FindById(id uuid.UUID) (*models.Section, error) {
	return t.r.FindById(id)
}

func (t *SectionService) FindAllByTestId(testId uuid.UUID) (*[]models.Section, error) {
	return t.r.FindAllBySectionId(testId)
}
