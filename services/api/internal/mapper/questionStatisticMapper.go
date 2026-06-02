package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
)

func questionStatisticModelToResponse(m *models.QuestionStatistic) DTO.QuestionStatisticResponse {
	return DTO.QuestionStatisticResponse{
		Id:             m.ID,
		TAnswer:        m.TotalAnswer,
		TCorrectAnswer: m.TotalCorrectAnswer,
		QuestionType:   consts.QuestionType.ToString(m.QuestionType),
	}
}
