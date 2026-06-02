package repository

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuestionRepository struct {
	db *gorm.DB
}

func ProvideQuestionRepository(db *gorm.DB) QuestionRepository {
	return QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(question models.Question) (*models.Question, error) {
	if err := r.db.Create(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) Update(question models.Question) (*models.Question, error) {
	var toBeUpdated models.Question
	if err := r.db.Model(&toBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.Question{BaseModel: models.BaseModel{ID: question.ID}}).
		Updates(models.Question{
			CorrectAnswer:  question.CorrectAnswer,
			QNumber:        question.QNumber,
			Component:      question.Component,
			LastModifierId: question.LastModifierId,
		}).
		Error; err != nil {
		return nil, err
	}
	return &toBeUpdated, nil
}

func (r *QuestionRepository) DeleteById(id uuid.UUID) error {
	return r.db.Delete(&models.Question{}, id).Error
}

func (r *QuestionRepository) FindById(id uuid.UUID) (*models.Question, error) {
	var question = models.Question{BaseModel: models.BaseModel{ID: id}}
	if err := r.db.First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) FindAllByGroupId(groupId uuid.UUID) (*[]models.Question, error) {
	var questions []models.Question
	if err := r.db.Preload("Statistic").Where(models.Question{GroupID: groupId}).Find(&questions).Error; err != nil {
		return nil, err
	}
	return &questions, nil
}

func (r *QuestionRepository) FindAllByTestId(testId uuid.UUID, userTestId uuid.UUID) (*[]models.Question, error) {
	var questions []models.Question
	if err := r.db.Preload("Statistic").
		Preload("UserAnswers", "user_test_id = ? ", userTestId).
		Joins("JOIN question_groups ON questions.group_id=question_groups.id").
		Joins("JOIN sections ON question_groups.section_id=sections.id").
		Joins("JOIN tests ON sections.test_id=tests.id").
		Where("tests.id = ?", testId).
		Order("q_number ASC").
		Find(&questions).Error; err != nil {
		return nil, err
	}
	return &questions, nil
}
