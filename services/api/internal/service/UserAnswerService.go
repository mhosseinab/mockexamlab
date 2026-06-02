package service

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
)

type UserAnswerService struct {
	repository repository.UserAnswerRepository
}

func ProvideUserAnswerService(service repository.UserAnswerRepository) UserAnswerService {
	return UserAnswerService{repository: service}
}

func (t *UserAnswerService) Create(m models.UserAnswer) (*models.UserAnswer, error) {
	return t.repository.Create(m)
}

func (t *UserAnswerService) BatchCreate(m []models.UserAnswer) (*[]models.UserAnswer, error) {
	return t.repository.BatchCreate(m)
}

func (t *UserAnswerService) UpdateAnswer(id uuid.UUID, answer string) (*models.UserAnswer, error) {
	return t.repository.UpdateAnswer(id, answer)
}

func (t *UserAnswerService) UpdateMarker(id uuid.UUID, markerScore float32, markerComment string) (*models.UserAnswer, error) {
	return t.repository.UpdateMarker(id, markerScore, markerComment)
}

func (t *UserAnswerService) DeleteById(id uuid.UUID) error {
	return t.repository.DeleteById(id)
}

func (t *UserAnswerService) FindById(id uuid.UUID) (*models.UserAnswer, error) {
	return t.repository.FindById(id)
}

func (t *UserAnswerService) FindAllByUserTestId(userTestId uuid.UUID) (*[]models.UserAnswer, error) {
	return t.repository.FindAllByUserTestId(userTestId)
}
