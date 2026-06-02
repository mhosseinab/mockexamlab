package repository

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SectionRepository struct {
	db *gorm.DB
}

func ProvideSectionRepository(db *gorm.DB) SectionRepository {
	return SectionRepository{db: db}
}

func (r *SectionRepository) Create(section models.Section) (*models.Section, error) {
	if err := r.db.Create(&section).Error; err != nil {
		return nil, err
	}
	return &section, nil
}

func (r *SectionRepository) Update(section models.Section) (*models.Section, error) {
	var sectionToBeUpdated models.Section
	if err := r.db.Model(&sectionToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.Section{BaseModel: models.BaseModel{ID: section.ID}}).
		Updates(models.Section{
			Title:          section.Title,
			Desc:           section.Desc,
			ComponentType:  section.ComponentType,
			LastModifierId: section.LastModifierId,
			ExtraProperty:  section.ExtraProperty,
		}).
		Error; err != nil {
		return nil, err
	}
	return &sectionToBeUpdated, nil
}

func (r *SectionRepository) DeleteById(id uuid.UUID) error {
	return r.db.Delete(&models.Section{}, id).Error
}

func (r *SectionRepository) FindById(id uuid.UUID) (*models.Section, error) {
	var section = models.Section{BaseModel: models.BaseModel{ID: id}}
	if err := r.db.First(&section).Error; err != nil {
		return nil, err
	}
	return &section, nil
}

func (r *SectionRepository) FindAllBySectionId(testId uuid.UUID) (*[]models.Section, error) {
	var sections []models.Section
	if err := r.db.Model(&models.Section{}).Where(models.Section{TestID: testId}).Find(&sections).Error; err != nil {
		return nil, err
	}
	return &sections, nil
}
