package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"time"
)

func TestCreateRequestToModel(DTO *DTO.TestCreateRequest, creatorId uuid.UUID) (*models.Test, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	if data, err := json.Marshal(DTO.Detail); err != nil {
		return nil, err
	} else {
		jsonData = string(data)
	}
	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}
	return &models.Test{
		Name:           DTO.Name,
		Description:    DTO.Description,
		Module:         consts.TestModule(DTO.Module),
		CreatorId:      creatorId,
		LastModifierId: creatorId,
		TestTime:       DTO.TotalTestTime,
		ExtraProperty:  extraProperties,
		Stat: &models.TestStat{
			BaseModel:  models.BaseModel{ID: uuid.New()},
			TestModule: consts.TestModule(DTO.Module),
		},
		TestDate: time.Unix(DTO.TestDate, 0)}, nil
}

func TestModelToCreateResponse(m *models.Test) (*DTO.TestCreateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.TestCreateResponse{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Module:      consts.TestModule.ToString(m.Module),
		TestDate:    m.TestDate,
		CreatedAt:   time.Unix(m.CreatedAt, 0)}

	if err := json.Unmarshal([]byte(jsonString), &result.Detail); err != nil {
		return nil, err
	}
	return &result, nil

}

func TestUpdateRequestToModel(DTO *DTO.TestUpdateRequest, modifierId uuid.UUID) (*models.Test, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	if data, err := json.Marshal(DTO.Detail); err != nil {
		return nil, err
	} else {
		jsonData = string(data)
	}
	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}
	return &models.Test{
		BaseModel:      models.BaseModel{ID: DTO.Id},
		Name:           DTO.Name,
		Description:    DTO.Description,
		Module:         consts.TestModule(DTO.Module),
		LastModifierId: modifierId,
		TestTime:       DTO.TotalTestTime,
		ExtraProperty:  extraProperties,
		TestDate:       time.Unix(DTO.TestDate, 0)}, nil
}

func TestModelToUpdateResponse(m *models.Test) (*DTO.TestUpdateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.TestUpdateResponse{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Module:      consts.TestModule.ToString(m.Module),
		TestDate:    m.TestDate,
		UpdateAt:    time.Unix(m.CreatedAt, 0)}
	if err := json.Unmarshal([]byte(jsonString), &result.Detail); err != nil {
		return nil, err
	}
	return &result, nil
}

func TestModelToResponse(m *models.Test) (*DTO.TestResponse, error) {
	var sectionResponse []DTO.SectionResponse
	if m.Sections != nil && len(m.Sections) > 0 {
		data, _ := SectionModelsToArrayResponse(&m.Sections)
		sectionResponse = *data
	}
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.TestResponse{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Module:      consts.TestModule.ToString(m.Module),
		TestDate:    m.TestDate,
		Sections:    sectionResponse,
	}
	if err := json.Unmarshal([]byte(jsonString), &result.Detail); err != nil {
		return nil, err
	}
	return &result, nil
}

func TestModelsToArrayResponse(models *[]models.Test, count int64) (*DTO.PaginateResponse, error) {
	testDTOs := make([]DTO.TestResponse, len(*models))
	for i, item := range *models {
		if data, err := TestModelToResponse(&item); err != nil {
			return nil, err
		} else {
			testDTOs[i] = *data
		}
	}
	return &DTO.PaginateResponse{
		Data:  testDTOs,
		Count: count,
	}, nil
}
