package mapper

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
)

func TestStatUpdateRequestToModel(DTO *models.TestSkillsStat, id uuid.UUID) (*models.TestStat, error) {
	stat := &models.TestStat{
		BaseModel: models.BaseModel{ID: id},
	}
	if err := stat.UserTestModelToJson(*DTO); err != nil {
		return nil, err
	}
	return stat, nil
}
