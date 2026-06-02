package repository

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserTestRepository struct {
	db *gorm.DB
}

func ProvideUserTestRepository(DB *gorm.DB) UserTestRepository {
	return UserTestRepository{
		db: DB,
	}
}
func (t *UserTestRepository) Create(test models.UserTest) (*models.UserTest, error) {
	if err := t.db.Create(&test).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

/*func (t *UserTestRepository) Update(test models.UserTest) (*models.UserTest, error) {
	var testToBeUpdated models.UserTest
	if err := t.db.Model(&testToBeUpdated).
		Clauses(clause.Returning{}).
		Where(models.UserTest{BaseModel: models.BaseModel{ID: test.ID}}).
		Updates(models.UserTest{
			Name:        test.Name,
			Description: test.Description,
			TestDate:    test.TestDate,
			Module:      test.Module}).
		Error; err != nil {
		return nil, err
	}
	return &testToBeUpdated, nil
}*/

func (t *UserTestRepository) UpdateScore(id uuid.UUID, score models.UserTestScore) (*models.UserTest, error) {
	var userTestUpdated models.UserTest
	if err := t.db.Model(&userTestUpdated).
		Clauses(clause.Returning{}).
		Where(models.UserTest{BaseModel: models.BaseModel{ID: id}}).
		Updates(models.UserTest{
			OverallScore:   score.OverallScore,
			ListeningScore: score.ListeningScore,
			ReadingScore:   score.ReadingScore,
			WritingScore:   score.WritingScore,
			SpeakingScore:  score.SpeakingScore,
		}).
		Error; err != nil {
		return nil, err
	}
	return &userTestUpdated, nil
}

func (t *UserTestRepository) OverallScoreStat(testId uuid.UUID) (*[]models.TestComponentStat, error) {
	var res []models.TestComponentStat
	if err := t.db.Model(&models.UserTest{}).
		Select("overall_score,count(*)").
		Where(&models.UserTest{TestID: testId}).
		Group("overall_score").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) ListeningScoreStat(testId uuid.UUID) (*[]models.TestComponentStat, error) {
	var res []models.TestComponentStat
	if err := t.db.Model(&models.UserTest{}).
		Select("listening_score,count(*)").
		Where(&models.UserTest{TestID: testId}).
		Group("listening_score").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) ReadingScoreStat(testId uuid.UUID) (*[]models.TestComponentStat, error) {
	var res []models.TestComponentStat
	if err := t.db.Model(&models.UserTest{}).
		Select("reading_score,count(*)").
		Where(&models.UserTest{TestID: testId}).
		Group("reading_score").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) WritingScoreStat(testId uuid.UUID) (*[]models.TestComponentStat, error) {
	var res []models.TestComponentStat
	if err := t.db.Model(&models.UserTest{}).
		Select("writing_score,count(*)").
		Where(&models.UserTest{TestID: testId}).
		Group("writing_score").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) SpeakingScoreStat(testId uuid.UUID) (*[]models.TestComponentStat, error) {
	var res []models.TestComponentStat
	if err := t.db.Model(&models.UserTest{}).
		Select("speaking_score,count(*)").
		Where(&models.UserTest{TestID: testId}).
		Group("speaking_score").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) AverageStatWithFilter(userId uuid.UUID, module consts.TestModule) (*DTO.UserSkillsAvgStat, error) {
	var res DTO.UserSkillsAvgStat
	if err := t.db.Model(&models.UserTest{}).
		Select("AVG(listening_score),AVG(reading_score),AVG(writing_score),AVG(speaking_score)").
		Where(&models.UserTest{UserID: userId}).
		Joins("JOIN tests ON user_tests.test_id=tests.id").
		Where("tests.module=?", module).
		Where("user_test_type=?", 0).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) AverageStat(userId uuid.UUID) (*DTO.UserSkillsAvgStat, error) {
	var res DTO.UserSkillsAvgStat
	if err := t.db.Model(&models.UserTest{}).
		Select("AVG(listening_score),AVG(reading_score),AVG(writing_score),AVG(speaking_score)").
		Where(&models.UserTest{UserID: userId}).
		Joins("JOIN tests ON user_tests.test_id=tests.id").
		Where("user_test_type=?", 0).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (t *UserTestRepository) DeleteById(id uuid.UUID) error {
	return t.db.Delete(&models.UserTest{}, id).Error
}

func (t *UserTestRepository) FindById(id uuid.UUID) (*models.UserTest, error) {
	var test = models.UserTest{BaseModel: models.BaseModel{ID: id}}
	if err := t.db.Preload("UserAnswers").First(&test).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (t *UserTestRepository) FindAllByUserId(userId uuid.UUID, params *models.QueryParams) (*[]models.UserTest, int64, error) {
	var tests []models.UserTest
	var count int64 = 0
	if err := t.db.Model(&models.UserTest{}).
		Scopes(FilterByQuery(params, consts.PAGINATE)).
		Where(models.UserTest{UserID: userId}).
		Find(&tests).
		Count(&count).Error; err != nil {
		return nil, count, err
	}
	return &tests, count, nil
}
