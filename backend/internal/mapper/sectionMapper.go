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

func SectionCreateRequestToModel(DTO *DTO.SectionCreateRequest, creatorId uuid.UUID) (*models.Section, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	switch consts.TestComponent(DTO.ComponentType) {
	case consts.Speaking:
		{
			if data, err := json.Marshal(DTO.SectionSpeaking); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Listening:
		{
			if data, err := json.Marshal(DTO.SectionListening); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Reading:
		{
			if data, err := json.Marshal(DTO.SectionReading); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Writing:
		{
			if data, err := json.Marshal(DTO.SectionWriting); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	}

	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}
	return &models.Section{
		TestID:         DTO.TestId,
		Title:          DTO.Title,
		Desc:           DTO.Description,
		ComponentType:  consts.TestComponent(DTO.ComponentType),
		Time:           DTO.TotalSectionTime,
		QStart:         DTO.QStart,
		QEnd:           DTO.QEnd,
		CreatorId:      creatorId,
		LastModifierId: creatorId,
		ExtraProperty:  extraProperties,
	}, nil
}

func SectionModelToCreateResponse(m *models.Section) (*DTO.SectionCreateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.SectionCreateResponse{
		Id:            m.ID,
		Title:         m.Title,
		Description:   m.Desc,
		ComponentType: consts.TestComponent.ToString(m.ComponentType),
		CreatedAt:     time.Unix(m.CreatedAt, 0)}
	switch m.ComponentType {
	case consts.Speaking:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionSpeaking); err != nil {
				return nil, err
			}
		}
	case consts.Listening:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionListening); err != nil {
				return nil, err
			}
		}
	case consts.Reading:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionReading); err != nil {
				return nil, err
			}
		}
	case consts.Writing:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionWriting); err != nil {
				return nil, err
			}
		}
	}
	return &result, nil
}

func SectionUpdateRequestToModel(DTO *DTO.SectionUpdateRequest, modifierId uuid.UUID) (*models.Section, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	switch consts.TestComponent(DTO.ComponentType) {
	case consts.Speaking:
		{
			if data, err := json.Marshal(DTO.SectionSpeaking); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Listening:
		{
			if data, err := json.Marshal(DTO.SectionListening); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Reading:
		{
			if data, err := json.Marshal(DTO.SectionReading); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Writing:
		{
			if data, err := json.Marshal(DTO.SectionWriting); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	}

	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}
	return &models.Section{
		BaseModel:      models.BaseModel{ID: DTO.Id},
		Title:          DTO.Title,
		Desc:           DTO.Description,
		Time:           DTO.TotalSectionTime,
		QStart:         DTO.QStart,
		QEnd:           DTO.QEnd,
		ComponentType:  consts.TestComponent(DTO.ComponentType),
		LastModifierId: modifierId,
		ExtraProperty:  extraProperties,
	}, nil
}

func SectionModelToUpdateResponse(m *models.Section) (*DTO.SectionUpdateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.SectionUpdateResponse{
		Id:            m.ID,
		Title:         m.Title,
		Description:   m.Desc,
		ComponentType: consts.TestComponent.ToString(m.ComponentType),
		UpdateAt:      time.Unix(m.UpdatedAt, 0)}
	switch m.ComponentType {
	case consts.Speaking:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionSpeaking); err != nil {
				return nil, err
			}
		}
	case consts.Listening:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionListening); err != nil {
				return nil, err
			}
		}
	case consts.Reading:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionReading); err != nil {
				return nil, err
			}
		}
	case consts.Writing:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionWriting); err != nil {
				return nil, err
			}
		}
	}
	return &result, nil
}

func SectionModelToResponse(m *models.Section) (*DTO.SectionResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	var groupsResponse []DTO.QuestionGroupResponse
	if m.QuestionGroups != nil && len(m.QuestionGroups) > 0 {
		data, _ := QuestionGroupModelsToArrayResponse(&m.QuestionGroups)
		groupsResponse = *data
	}
	result := DTO.SectionResponse{
		Id:             m.ID,
		Title:          m.Title,
		Description:    m.Desc,
		QStart:         m.QStart,
		QEnd:           m.QEnd,
		QuestionGroups: groupsResponse,
		ComponentType:  consts.TestComponent.ToString(m.ComponentType),
	}
	switch m.ComponentType {
	case consts.Speaking:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionSpeaking); err != nil {
				return nil, err
			}
		}
	case consts.Listening:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionListening); err != nil {
				return nil, err
			}
		}
	case consts.Reading:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionReading); err != nil {
				return nil, err
			}
		}
	case consts.Writing:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SectionWriting); err != nil {
				return nil, err
			}
		}
	}
	return &result, nil
}

func SectionModelsToArrayResponse(models *[]models.Section) (*[]DTO.SectionResponse, error) {
	sectionDTOs := make([]DTO.SectionResponse, len(*models))
	for i, item := range *models {
		if data, err := SectionModelToResponse(&item); err != nil {
			return nil, err
		} else {
			sectionDTOs[i] = *data
		}
	}
	return &sectionDTOs, nil
}
