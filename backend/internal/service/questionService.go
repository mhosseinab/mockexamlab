package service

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
)

type QuestionService struct {
	r repository.QuestionRepository
}

func ProvideQuestionService(r repository.QuestionRepository) QuestionService {
	return QuestionService{r: r}
}

func (t *QuestionService) Create(question models.Question) (*models.Question, error) {
	return t.r.Create(question)
}

func (t *QuestionService) Update(question models.Question) (*models.Question, error) {
	return t.r.Update(question)
}

func (t *QuestionService) DeleteById(id uuid.UUID) error {
	return t.r.DeleteById(id)
}

func (t *QuestionService) FindById(id uuid.UUID) (*models.Question, error) {
	return t.r.FindById(id)
}

func (t *QuestionService) FindAllByGroupId(groupId uuid.UUID) (*[]models.Question, error) {
	return t.r.FindAllByGroupId(groupId)
}

func (t *QuestionService) FindAllByTestId(testId uuid.UUID, userTestId uuid.UUID) (*[]models.Question, error) {
	return t.r.FindAllByTestId(testId, userTestId)
}
