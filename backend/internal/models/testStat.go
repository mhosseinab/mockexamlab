package models

import (
	"MockExamLab/internal/models/consts"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type TestStat struct {
	BaseModel
	TestId     uuid.UUID         `json:"userTestId"`
	TestModule consts.TestModule `json:"testModule"`
	Listening  pgtype.JSONB      `json:"listening" gorm:"type:jsonb;default:'[]';not null"`
	Reading    pgtype.JSONB      `json:"reading" gorm:"type:jsonb;default:'[]';not null"`
	Writing    pgtype.JSONB      `json:"writing" gorm:"type:jsonb;default:'[]';not null"`
	Speaking   pgtype.JSONB      `json:"speaking" gorm:"type:jsonb;default:'[]';not null"`
	Overall    pgtype.JSONB      `json:"overall" gorm:"type:jsonb;default:'[]';not null"`
}

type TestComponentStat struct {
	Score float32 `json:"score"`
	Count float32 `json:"count"`
}
type TestSkillsStat struct {
	Listening *[]TestComponentStat `json:"listening"`
	Reading   *[]TestComponentStat `json:"reading"`
	Writing   *[]TestComponentStat `json:"writing"`
	Speaking  *[]TestComponentStat `json:"speaking"`
	Overall   *[]TestComponentStat `json:"overall"`
}

func (u *TestStat) TestStatJsonToDTOModel() (*TestSkillsStat, error) {
	var jsonStringListening string
	var jsonStringReading string
	var jsonStringWriting string
	var jsonStringSpeaking string
	var jsonStringOverall string
	var scoreListening []TestComponentStat
	var scoreReading []TestComponentStat
	var scoreWriting []TestComponentStat
	var scoreSpeaking []TestComponentStat
	var scoreOverall []TestComponentStat

	if errJsonStringToModel := u.Listening.AssignTo(&jsonStringListening); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	if errJsonConverter := json.Unmarshal([]byte(jsonStringListening), &scoreListening); errJsonConverter != nil {
		return nil, errJsonConverter
	}

	if errJsonStringToModel := u.Reading.AssignTo(&jsonStringReading); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	if errJsonConverter := json.Unmarshal([]byte(jsonStringReading), &scoreReading); errJsonConverter != nil {
		return nil, errJsonConverter
	}

	if errJsonStringToModel := u.Listening.AssignTo(&jsonStringWriting); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	if errJsonConverter := json.Unmarshal([]byte(jsonStringWriting), &scoreWriting); errJsonConverter != nil {
		return nil, errJsonConverter
	}

	if errJsonStringToModel := u.Listening.AssignTo(&jsonStringSpeaking); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	if errJsonConverter := json.Unmarshal([]byte(jsonStringSpeaking), &scoreSpeaking); errJsonConverter != nil {
		return nil, errJsonConverter
	}

	if errJsonStringToModel := u.Listening.AssignTo(&jsonStringOverall); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	if errJsonConverter := json.Unmarshal([]byte(jsonStringOverall), &scoreOverall); errJsonConverter != nil {
		return nil, errJsonConverter
	}

	return &TestSkillsStat{
		Listening: &scoreListening,
		Reading:   &scoreReading,
		Writing:   &scoreWriting,
		Speaking:  &scoreSpeaking,
		Overall:   &scoreOverall,
	}, nil

}

func (u *TestStat) UserTestModelToJson(d TestSkillsStat) error {
	var jsonString string
	if jsonData, errJsonConverter := json.Marshal(&d.Listening); errJsonConverter != nil {
		return errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := u.Listening.Set(&jsonString); errJsonStringToModel != nil {
			return errJsonStringToModel
		}
	}

	if jsonData, errJsonConverter := json.Marshal(&d.Reading); errJsonConverter != nil {
		return errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := u.Reading.Set(&jsonString); errJsonStringToModel != nil {
			return errJsonStringToModel
		}
	}

	if jsonData, errJsonConverter := json.Marshal(&d.Writing); errJsonConverter != nil {
		return errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := u.Writing.Set(&jsonString); errJsonStringToModel != nil {
			return errJsonStringToModel
		}
	}

	if jsonData, errJsonConverter := json.Marshal(&d.Speaking); errJsonConverter != nil {
		return errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := u.Speaking.Set(&jsonString); errJsonStringToModel != nil {
			return errJsonStringToModel
		}
	}

	if jsonData, errJsonConverter := json.Marshal(&d.Overall); errJsonConverter != nil {
		return errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := u.Overall.Set(&jsonString); errJsonStringToModel != nil {
			return errJsonStringToModel
		}
	}
	return nil
}
