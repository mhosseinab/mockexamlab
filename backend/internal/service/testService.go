package service

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
)

type TestService struct {
	repository repository.TestRepository
}

func ProvideTestService(service repository.TestRepository) TestService {
	return TestService{repository: service}
}

func (t *TestService) Create(test models.Test) (*models.Test, error) {
	return t.repository.Create(test)
}

func (t *TestService) Update(test models.Test) (*models.Test, error) {
	return t.repository.Update(test)
}

func (t *TestService) DeleteById(id uuid.UUID) error {
	return t.repository.DeleteById(id)
}

func (t *TestService) FindById(id uuid.UUID) (*models.Test, error) {
	return t.repository.FindById(id)
}

func (t *TestService) FindByIdWithPreloads(id uuid.UUID) (*models.Test, error) {
	return t.repository.FindByIdWithPreload(id)
}

func (t *TestService) FindAll(params *models.QueryParams) (*[]models.Test, int64, error) {
	return t.repository.FindAll(params)
}
