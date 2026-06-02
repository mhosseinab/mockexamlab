package repository

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TestStatRepository struct {
	db *gorm.DB
}

func ProvideTestStatRepository(DB *gorm.DB) TestStatRepository {
	return TestStatRepository{
		db: DB,
	}
}

func (t *TestStatRepository) Update(stat models.TestStat) (*models.TestStat, error) {
	var statToBeUpdated models.TestStat
	if err := t.db.Model(&statToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.TestStat{BaseModel: models.BaseModel{ID: stat.ID}}).
		Updates(models.TestStat{
			Listening: stat.Listening,
			Reading:   stat.Reading,
			Writing:   stat.Writing,
			Speaking:  stat.Speaking,
			Overall:   stat.Overall,
		}).
		Error; err != nil {
		return nil, err
	}
	return &statToBeUpdated, nil
}

func (t *TestStatRepository) FindById(testId uuid.UUID) (*models.TestStat, error) {
	var stat = models.TestStat{TestId: testId}
	if err := t.db.First(&stat).Error; err != nil {
		return nil, err
	}
	return &stat, nil
}
