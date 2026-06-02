package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"time"
)

func ArrayUserAnswerCreateRequestToModel(DTOs *[]DTO.UserAnswerCreateRequest) []models.UserAnswer {
	answerDTOs := make([]models.UserAnswer, len(*DTOs))
	for i, item := range *DTOs {
		answerDTOs[i] = UserAnswerCreateRequestToModel(&item)
	}
	return answerDTOs
}

func UserAnswerCreateRequestToModel(DTO *DTO.UserAnswerCreateRequest) models.UserAnswer {
	return models.UserAnswer{
		UserTestID: DTO.UserTestId,
		QuestionID: DTO.QuestionId,
		Answer:     DTO.Answer}
}

func ArrayUserAnswerModelToCreateResponse(models *[]models.UserAnswer) []DTO.UserAnswerCreateResponse {
	answerDTOs := make([]DTO.UserAnswerCreateResponse, len(*models))
	for i, item := range *models {
		answerDTOs[i] = UserAnswerModelToCreateResponse(&item)
	}
	return answerDTOs
}

func UserAnswerModelToCreateResponse(m *models.UserAnswer) DTO.UserAnswerCreateResponse {
	return DTO.UserAnswerCreateResponse{
		Id:         m.ID,
		UserTestId: m.UserTestID,
		QuestionId: m.QuestionID,
		Answer:     m.Answer,
		CreatedAt:  time.Unix(m.CreatedAt, 0)}
}

func UserAnswerModelToUpdateAnswerResponse(m *models.UserAnswer) DTO.UserAnswerUpdateAnswerResponse {
	return DTO.UserAnswerUpdateAnswerResponse{
		Id:         m.ID,
		UserTestId: m.UserTestID,
		QuestionId: m.QuestionID,
		Answer:     m.Answer,
		UpdateAt:   time.Unix(m.CreatedAt, 0)}
}

func UserAnswerModelToUpdateMarkerResponse(m *models.UserAnswer) DTO.UserAnswerUpdateMarkerResponse {
	return DTO.UserAnswerUpdateMarkerResponse{
		Id:            m.ID,
		UserTestId:    m.UserTestID,
		QuestionId:    m.QuestionID,
		Answer:        m.Answer,
		MarkerId:      m.MarkerID,
		MarkerScore:   m.MarkerScore,
		MarkerComment: m.MarkerComment,
		UpdateAt:      time.Unix(m.CreatedAt, 0)}
}

func UserAnswerModelToResponse(m *models.UserAnswer) *DTO.UserAnswerResponse {
	return &DTO.UserAnswerResponse{
		Id:            m.ID,
		UserTestId:    m.UserTestID,
		QuestionId:    m.QuestionID,
		Answer:        m.Answer,
		MarkerId:      m.MarkerID,
		MarkerScore:   m.MarkerScore,
		MarkerComment: m.MarkerComment,
	}
}

func UserAnswerModelsToArrayResponse(models *[]models.UserAnswer) *[]DTO.UserAnswerResponse {
	answerDTOs := make([]DTO.UserAnswerResponse, len(*models))
	for i, item := range *models {
		answerDTOs[i] = *UserAnswerModelToResponse(&item)
	}
	return &answerDTOs
}
