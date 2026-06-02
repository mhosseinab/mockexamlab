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

func QuestionGroupCreateRequestToModel(DTO *DTO.QuestionGroupCreateRequest, creatorId uuid.UUID) (*models.QuestionGroup, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	switch consts.QuestionType(DTO.QuestionType) {
	case consts.MultipleChoice:
		{
			if data, err := json.Marshal(DTO.MultipleChoice); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.IdentifyingInformation:
		{
			if data, err := json.Marshal(DTO.IdentifyingInformation); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Map:
		{
			if data, err := json.Marshal(DTO.Map); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.NoteCompletion:
		{
			if data, err := json.Marshal(DTO.NoteCompletion); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.SpeakingTopic:
		{
			if data, err := json.Marshal(DTO.SpeakingTopic); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.WritingTopic:
		{
			if data, err := json.Marshal(DTO.WritingTopic); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	}
	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}

	return &models.QuestionGroup{
		SectionID:      DTO.SectionId,
		Title:          DTO.Title,
		Desc:           DTO.Description,
		QStart:         DTO.QStart,
		QEnd:           DTO.QEnd,
		CreatorId:      creatorId,
		LastModifierId: creatorId,
		QuestionType:   consts.QuestionType(DTO.QuestionType),
		ExtraProperty:  extraProperties,
	}, nil
}

func QuestionGroupModelToCreateResponse(m *models.QuestionGroup) (*DTO.QuestionGroupCreateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.QuestionGroupCreateResponse{
		Id:           m.ID,
		Title:        m.Title,
		Description:  m.Desc,
		QuestionType: consts.QuestionType.ToString(m.QuestionType),
		CreatedAt:    time.Unix(m.CreatedAt, 0),
	}
	switch m.QuestionType {
	case consts.MultipleChoice:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.MultipleChoice); err != nil {
				return nil, err
			}
		}
	case consts.IdentifyingInformation:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.IdentifyingInformation); err != nil {
				return nil, err
			}
		}
	case consts.Map:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.Map); err != nil {
				return nil, err
			}
		}
	case consts.NoteCompletion:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.NoteCompletion); err != nil {
				return nil, err
			}
		}
	case consts.SpeakingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SpeakingTopic); err != nil {
				return nil, err
			}
		}
	case consts.WritingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.WritingTopic); err != nil {
				return nil, err
			}
		}
	}
	return &result, nil
}

func QuestionGroupUpdateRequestToModel(DTO *DTO.QuestionGroupUpdateRequest, modifierId uuid.UUID) (*models.QuestionGroup, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	switch consts.QuestionType(DTO.QuestionType) {
	case consts.MultipleChoice:
		{
			if data, err := json.Marshal(DTO.MultipleChoice); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.IdentifyingInformation:
		{
			if data, err := json.Marshal(DTO.IdentifyingInformation); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.Map:
		{
			if data, err := json.Marshal(DTO.Map); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.NoteCompletion:
		{
			if data, err := json.Marshal(DTO.NoteCompletion); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.SpeakingTopic:
		{
			if data, err := json.Marshal(DTO.SpeakingTopic); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	case consts.WritingTopic:
		{
			if data, err := json.Marshal(DTO.WritingTopic); err != nil {
				return nil, err
			} else {
				jsonData = string(data)
			}
		}
	}
	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}

	return &models.QuestionGroup{
		BaseModel:      models.BaseModel{ID: DTO.Id},
		Title:          DTO.Title,
		Desc:           DTO.Description,
		QStart:         DTO.QStart,
		QEnd:           DTO.QEnd,
		LastModifierId: modifierId,
		QuestionType:   consts.QuestionType(DTO.QuestionType),
		ExtraProperty:  extraProperties,
	}, nil
}

func QuestionGroupModelToUpdateResponse(m *models.QuestionGroup) (*DTO.QuestionGroupUpdateResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}
	result := DTO.QuestionGroupUpdateResponse{
		Id:           m.ID,
		Title:        m.Title,
		Description:  m.Desc,
		QStart:       m.QStart,
		QEnd:         m.QEnd,
		QuestionType: consts.QuestionType.ToString(m.QuestionType),
		UpdateAt:     time.Unix(m.CreatedAt, 0),
	}
	switch m.QuestionType {
	case consts.MultipleChoice:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.MultipleChoice); err != nil {
				return nil, err
			}
		}
	case consts.IdentifyingInformation:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.IdentifyingInformation); err != nil {
				return nil, err
			}
		}
	case consts.Map:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.Map); err != nil {
				return nil, err
			}
		}
	case consts.NoteCompletion:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.NoteCompletion); err != nil {
				return nil, err
			}
		}
	case consts.SpeakingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SpeakingTopic); err != nil {
				return nil, err
			}
		}
	case consts.WritingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.WritingTopic); err != nil {
				return nil, err
			}
		}
	}
	return &result, nil
}

func QuestionGroupModelToResponse(m *models.QuestionGroup) (*DTO.QuestionGroupResponse, error) {
	var jsonString string
	if err := m.ExtraProperty.AssignTo(&jsonString); err != nil {
		return nil, err
	}

	result := DTO.QuestionGroupResponse{
		Id:           m.ID,
		Title:        m.Title,
		Description:  m.Desc,
		QStart:       m.QStart,
		QEnd:         m.QEnd,
		QuestionType: consts.QuestionType.ToString(m.QuestionType),
	}
	switch m.QuestionType {
	case consts.MultipleChoice:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.MultipleChoice); err != nil {
				return nil, err
			}
		}
	case consts.IdentifyingInformation:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.IdentifyingInformation); err != nil {
				return nil, err
			}
		}
	case consts.Map:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.Map); err != nil {
				return nil, err
			}
		}
	case consts.NoteCompletion:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.NoteCompletion); err != nil {
				return nil, err
			}
		}
	case consts.SpeakingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.SpeakingTopic); err != nil {
				return nil, err
			}
		}
	case consts.WritingTopic:
		{
			if err := json.Unmarshal([]byte(jsonString), &result.WritingTopic); err != nil {
				return nil, err
			}
		}
	}

	if m.Questions != nil && len(m.Questions) > 0 {
		data := QuestionModelsWithoutAnswerToArrayResponse(&m.Questions)
		result.Questions = data
	}
	return &result, nil
}

func QuestionGroupModelsToArrayResponse(models *[]models.QuestionGroup) (*[]DTO.QuestionGroupResponse, error) {
	DTOs := make([]DTO.QuestionGroupResponse, len(*models))
	for i, item := range *models {
		if data, err := QuestionGroupModelToResponse(&item); err != nil {
			return nil, err
		} else {
			DTOs[i] = *data
		}
	}
	return &DTOs, nil
}
