package service

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
)

type QuestionGroupService struct {
	r repository.QuestionGroupRepository
}

func ProvideQuestionGroupService(r repository.QuestionGroupRepository) QuestionGroupService {
	return QuestionGroupService{
		r: r,
	}
}

func (t *QuestionGroupService) Create(section models.QuestionGroup) (*models.QuestionGroup, error) {
	return t.r.Create(section)
}

func (t *QuestionGroupService) Update(section models.QuestionGroup) (*models.QuestionGroup, error) {
	return t.r.Update(section)
}

func (t *QuestionGroupService) DeleteById(id uuid.UUID) error {
	return t.r.DeleteById(id)
}

func (t *QuestionGroupService) FindById(id uuid.UUID) (*models.QuestionGroup, error) {
	return t.r.FindById(id)
}

func (t *QuestionGroupService) FindAllBySectionId(sectionId uuid.UUID) (*[]models.QuestionGroup, error) {
	return t.r.FindAllBySectionId(sectionId)
}
