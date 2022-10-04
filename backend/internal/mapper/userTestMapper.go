package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"time"
)

func UserTestCreateRequestToModel(DTO *DTO.UserTestCreateRequest, creatorId uuid.UUID) models.UserTest {
	return models.UserTest{
		TestID:         DTO.TestId,
		UserID:         DTO.UserId,
		UserTestType:   consts.UserTestType(DTO.TestType),
		TestDate:       time.Unix(DTO.TestDate, 0),
		CreatorId:      creatorId,
		LastModifierId: creatorId}
}

func UserTestModelToCreateResponse(m *models.UserTest) DTO.UserTestCreateResponse {
	return DTO.UserTestCreateResponse{
		Id:        m.ID,
		TestId:    m.TestID,
		UserId:    m.UserID,
		TestType:  consts.UserTestType.ToString(m.UserTestType),
		TestDate:  m.TestDate,
		CreatedAt: time.Unix(m.CreatedAt, 0)}
}

func UserTestUpdateRequestToModel(DTO *DTO.UserTestUpdateRequest, modifierId uuid.UUID) models.UserTest {
	return models.UserTest{
		BaseModel:      models.BaseModel{ID: DTO.Id},
		Duration:       DTO.TestDuration,
		LastModifierId: modifierId}
}

func UserTestModelToUpdateResponse(m *models.UserTest) DTO.UserTestUpdateResponse {
	return DTO.UserTestUpdateResponse{
		Id:       m.ID,
		TestId:   m.TestID,
		UserId:   m.UserID,
		TestType: consts.UserTestType.ToString(m.UserTestType),
		TestDate: m.TestDate,
		UpdateAt: time.Unix(m.CreatedAt, 0)}
}

func UserTestModelToResponse(m *models.UserTest) DTO.UserTestResponse {
	var answersResponse []DTO.UserAnswerResponse
	if m.UserAnswers != nil && len(m.UserAnswers) > 0 {
		answersResponse = *UserAnswerModelsToArrayResponse(&m.UserAnswers)
	}
	return DTO.UserTestResponse{
		Id:          m.ID,
		TestId:      m.TestID,
		UserId:      m.UserID,
		TestType:    consts.UserTestType.ToString(m.UserTestType),
		TestDate:    m.TestDate,
		UserAnswers: answersResponse,
	}
}

func UserTestModelsToArrayResponse(models *[]models.UserTest, count int64) DTO.PaginateResponse {
	testDTOs := make([]DTO.UserTestResponse, len(*models))
	for i, item := range *models {
		testDTOs[i] = UserTestModelToResponse(&item)
	}
	return DTO.PaginateResponse{
		Data:  testDTOs,
		Count: count,
	}
}
