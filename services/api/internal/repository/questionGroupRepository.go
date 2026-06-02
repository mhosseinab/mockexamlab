package repository

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuestionGroupRepository struct {
	db *gorm.DB
}

func ProvideQuestionGroupRepository(db *gorm.DB) QuestionGroupRepository {
	return QuestionGroupRepository{db: db}
}

func (r *QuestionGroupRepository) Create(group models.QuestionGroup) (*models.QuestionGroup, error) {
	if err := r.db.Create(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *QuestionGroupRepository) Update(group models.QuestionGroup) (*models.QuestionGroup, error) {
	var groupToBeUpdated models.QuestionGroup
	if err := r.db.Model(&groupToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.QuestionGroup{BaseModel: models.BaseModel{ID: group.ID}}).
		Updates(models.QuestionGroup{
			Title:         group.Title,
			Desc:          group.Desc,
			QuestionType:  group.QuestionType,
			ExtraProperty: group.ExtraProperty,
		}).
		Error; err != nil {
		return nil, err
	}
	return &groupToBeUpdated, nil
}

func (r *QuestionGroupRepository) DeleteById(id uuid.UUID) error {
	return r.db.Delete(&models.QuestionGroup{}, id).Error
}

func (r *QuestionGroupRepository) FindById(id uuid.UUID) (*models.QuestionGroup, error) {
	var group = models.QuestionGroup{BaseModel: models.BaseModel{ID: id}}
	if err := r.db.First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *QuestionGroupRepository) FindAllBySectionId(sectionId uuid.UUID) (*[]models.QuestionGroup, error) {
	var groups []models.QuestionGroup
	if err := r.db.Where(models.QuestionGroup{SectionID: sectionId}).Find(&groups).Error; err != nil {
		return nil, err
	}
	return &groups, nil
}
