package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"time"
)

func QuestionCreateRequestToModel(DTO *DTO.QuestionCreateRequest, creatorId uuid.UUID) models.Question {
	return models.Question{
		GroupID:        DTO.GroupId,
		QNumber:        DTO.QNumber,
		CorrectAnswer:  DTO.CorrectAnswer,
		Component:      consts.TestComponent(DTO.TestComponent),
		CreatorId:      creatorId,
		LastModifierId: creatorId,
		Statistic: models.QuestionStatistic{
			BaseModel:          models.BaseModel{ID: uuid.New()},
			QuestionType:       consts.QuestionType(DTO.QuestionType),
			TotalAnswer:        0,
			TotalCorrectAnswer: 0,
		},
	}
}

func QuestionModelToCreateResponse(m *models.Question) DTO.QuestionCreateResponse {
	return DTO.QuestionCreateResponse{
		Id:            m.ID,
		StatId:        m.Statistic.ID,
		QNumber:       m.QNumber,
		TestComponent: consts.TestComponent.ToString(m.Component),
		CorrectAnswer: m.CorrectAnswer,
		Statistic:     questionStatisticModelToResponse(&m.Statistic),
		CreatedAt:     time.Unix(m.CreatedAt, 0),
	}
}

func QuestionUpdateRequestToModel(DTO *DTO.QuestionUpdateRequest, modifierId uuid.UUID) models.Question {
	return models.Question{
		BaseModel:      models.BaseModel{ID: DTO.Id},
		QNumber:        DTO.QNumber,
		CorrectAnswer:  DTO.CorrectAnswer,
		Component:      consts.TestComponent(DTO.TestComponent),
		LastModifierId: modifierId}
}

func QuestionModelToUpdateResponse(m *models.Question) DTO.QuestionUpdateResponse {
	return DTO.QuestionUpdateResponse{
		Id:            m.ID,
		QNumber:       m.QNumber,
		CorrectAnswer: m.CorrectAnswer,
		UpdateAt:      time.Unix(m.CreatedAt, 0)}
}

func QuestionModelToResponse(m *models.Question) DTO.QuestionResponse {

	return DTO.QuestionResponse{
		Id:            m.ID,
		QNumber:       m.QNumber,
		CorrectAnswer: m.CorrectAnswer,
		Statistic:     questionStatisticModelToResponse(&m.Statistic),
	}
}

func QuestionModelsToArrayResponse(models *[]models.Question) []DTO.QuestionResponse {
	DTOs := make([]DTO.QuestionResponse, len(*models))
	for i, item := range *models {
		DTOs[i] = QuestionModelToResponse(&item)
	}
	return DTOs
}

func QuestionModelWithoutAnswerToResponse(m *models.Question) DTO.QuestionWithoutAnswerResponse {

	return DTO.QuestionWithoutAnswerResponse{
		Id:      m.ID,
		QNumber: m.QNumber,
	}
}

func QuestionModelsWithoutAnswerToArrayResponse(models *[]models.Question) []DTO.QuestionWithoutAnswerResponse {
	DTOs := make([]DTO.QuestionWithoutAnswerResponse, len(*models))
	for i, item := range *models {
		DTOs[i] = QuestionModelWithoutAnswerToResponse(&item)
	}
	return DTOs
}

func QuestionModelWithUserAnswerToResponse(m *models.Question) DTO.QuestionWithUserAnswerResponse {
	var userAnswer *DTO.UserAnswerResponse
	if m.UserAnswers != nil && len(m.UserAnswers) > 0 {
		userAnswer = UserAnswerModelToResponse(&m.UserAnswers[0])
	}
	return DTO.QuestionWithUserAnswerResponse{
		Id:            m.ID,
		QNumber:       m.QNumber,
		CorrectAnswer: m.CorrectAnswer,
		Statistic:     questionStatisticModelToResponse(&m.Statistic),
		UserAnswer:    userAnswer,
	}
}

func QuestionModelWithUserAnswerToArrayResponse(models *[]models.Question) []DTO.QuestionWithUserAnswerResponse {
	DTOs := make([]DTO.QuestionWithUserAnswerResponse, len(*models))
	for i, item := range *models {
		DTOs[i] = QuestionModelWithUserAnswerToResponse(&item)
	}
	return DTOs
}
