package repository

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TestRepository struct {
	db *gorm.DB
}

func ProvideTestRepository(DB *gorm.DB) TestRepository {
	return TestRepository{
		db: DB,
	}
}
func (t *TestRepository) Create(test models.Test) (*models.Test, error) {
	if err := t.db.Create(&test).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (t *TestRepository) Update(test models.Test) (*models.Test, error) {
	var testToBeUpdated models.Test
	if err := t.db.Model(&testToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.Test{BaseModel: models.BaseModel{ID: test.ID}}).
		Updates(models.Test{
			Name:          test.Name,
			Description:   test.Description,
			TestDate:      test.TestDate,
			Module:        test.Module,
			ExtraProperty: test.ExtraProperty,
		}).
		Error; err != nil {
		return nil, err
	}
	return &testToBeUpdated, nil
}

func (t *TestRepository) DeleteById(id uuid.UUID) error {
	return t.db.Delete(&models.Test{}, id).Error
}

func (t *TestRepository) FindById(id uuid.UUID) (*models.Test, error) {
	var test = models.Test{BaseModel: models.BaseModel{ID: id}}
	if err := t.db.Preload("Sections").First(&test).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (t *TestRepository) FindByIdWithPreload(id uuid.UUID) (*models.Test, error) {
	var test = models.Test{BaseModel: models.BaseModel{ID: id}}
	if err := t.db.Preload("Sections.QuestionGroups").First(&test).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (t *TestRepository) FindAll(params *models.QueryParams) (*[]models.Test, int64, error) {
	var tests []models.Test
	var count int64 = 0
	if err := t.db.Model(&models.Test{}).
		Scopes(FilterByQuery(params, consts.ALL)).
		Find(&tests).
		Count(&count).Error; err != nil {
		return nil, count, err
	}
	return &tests, count, nil
}
