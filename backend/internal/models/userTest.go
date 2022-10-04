package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"time"
)

type UserTest struct {
	BaseModel
	TestID         uuid.UUID           `json:"testId"`
	UserID         uuid.UUID           `json:"userId"`
	TestDate       time.Time           `json:"testDate"`
	Duration       int                 `json:"duration"`
	UserTestType   consts.UserTestType `json:"userTestType"`
	OverallScore   float32             `json:"overallScore"`
	ListeningScore float32             `json:"listeningScore"`
	ReadingScore   float32             `json:"readingScore"`
	WritingScore   float32             `json:"writingScore"`
	SpeakingScore  float32             `json:"speakingScore"`
	CreatorId      uuid.UUID           `json:"creatorId"`
	LastModifierId uuid.UUID           `json:"lastModifierId"`
	UserAnswers    []UserAnswer        `json:"userAnswers" gorm:"foreignKey:UserTestID;"`
	//Score          pgtype.JSONB        `json:"score" gorm:"type:jsonb;default:'[]';not null"`
	//Test           *Test               `json:"test"`
}
type UserTestScore struct {
	OverallScore   float32 `json:"overallScore"`
	ListeningScore float32 `json:"listeningScore"`
	ReadingScore   float32 `json:"readingScore"`
	WritingScore   float32 `json:"writingScore"`
	SpeakingScore  float32 `json:"speakingScore"`
}

/*func (u *UserTest) UserTestJsonToModel() (*UserTestScore, error) {
	var jsonString string
	if errJsonStringToModel := u.Score.AssignTo(&jsonString); errJsonStringToModel != nil {
		return nil, errJsonStringToModel
	}
	score := UserTestScore{}
	if errJsonConverter := json.Unmarshal([]byte(jsonString), &score); errJsonConverter != nil {
		return nil, errJsonConverter
	} else {
		return &score, nil
	}

}

func (u *UserTest) UserTestModelToJson(d interface{}) (*pgtype.JSONB, error) {
	var jsonString string
	var pgJson pgtype.JSONB
	if jsonData, errJsonConverter := json.Marshal(&d); errJsonConverter != nil {
		return nil, errJsonConverter
	} else {
		jsonString = string(jsonData)
		if errJsonStringToModel := pgJson.Set(&jsonString); errJsonStringToModel != nil {
			return nil, errJsonStringToModel
		} else {
			return &pgJson, nil
		}
	}
}*/
