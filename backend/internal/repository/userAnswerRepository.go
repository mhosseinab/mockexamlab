package repository

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAnswerRepository struct {
	db *gorm.DB
}

func ProvideUserAnswerRepository(DB *gorm.DB) UserAnswerRepository {
	return UserAnswerRepository{
		db: DB,
	}
}

func (t *UserAnswerRepository) BatchCreate(userAnswers []models.UserAnswer) (*[]models.UserAnswer, error) {
	if err := t.db.Model(&models.UserAnswer{}).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "user_test_id"}, {Name: "question_id"}},
			DoNothing: true}).
		Create(&userAnswers).Error; err != nil {
		return nil, err
	}
	return &userAnswers, nil
}

func (t *UserAnswerRepository) Create(userAnswer models.UserAnswer) (*models.UserAnswer, error) {
	if err := t.db.Model(&models.UserAnswer{}).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "user_test_id"}, {Name: "question_id"}},
			DoNothing: true}).
		Create(&userAnswer).Error; err != nil {
		return nil, err
	}
	return &userAnswer, nil
}

func (t *UserAnswerRepository) UpdateAnswer(id uuid.UUID, answer string) (*models.UserAnswer, error) {
	ToBeUpdated := models.UserAnswer{}
	if err := t.db.Model(&ToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.UserAnswer{BaseModel: models.BaseModel{ID: id}}).
		Updates(models.UserAnswer{
			Answer: answer,
		}).
		Error; err != nil {
		return nil, err
	}
	return &ToBeUpdated, nil
}

func (t *UserAnswerRepository) UpdateMarker(id uuid.UUID, markerScore float32, markerComment string) (*models.UserAnswer, error) {
	var ToBeUpdated models.UserAnswer
	if err := t.db.Model(&ToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.UserAnswer{BaseModel: models.BaseModel{ID: id}}).
		Updates(models.UserAnswer{
			MarkerScore:   markerScore,
			MarkerComment: markerComment,
		}).
		Error; err != nil {
		return nil, err
	}
	return &ToBeUpdated, nil
}

func (t *UserAnswerRepository) DeleteById(id uuid.UUID) error {
	return t.db.Delete(&models.UserAnswer{}, id).Error
}

func (t *UserAnswerRepository) FindById(id uuid.UUID) (*models.UserAnswer, error) {
	var userAnswer = models.UserAnswer{BaseModel: models.BaseModel{ID: id}}
	if err := t.db.First(&userAnswer).Error; err != nil {
		return nil, err
	}
	return &userAnswer, nil
}

func (t *UserAnswerRepository) FindAllByUserTestId(userTestId uuid.UUID) (*[]models.UserAnswer, error) {
	var userAnswers []models.UserAnswer
	if err := t.db.Where(models.UserAnswer{UserTestID: userTestId}).Find(&userAnswers).Error; err != nil {
		return nil, err
	}
	return &userAnswers, nil
}
