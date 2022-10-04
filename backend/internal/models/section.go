package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Section struct {
	BaseModel
	Title          string               `json:"title"`
	Desc           string               `json:"desc"`
	TestID         uuid.UUID            `json:"testId"`
	ComponentType  consts.TestComponent `json:"componentType"`                                         //Component Type
	ExtraProperty  pgtype.JSONB         `json:"extraProperty" gorm:"type:jsonb;default:'[]';not null"` // ExtraProperty
	CreatorId      uuid.UUID            `json:"creatorId"`
	LastModifierId uuid.UUID            `json:"lastModifierId"`
	Time           int                  `json:"time"`
	QStart         int                  `json:"q_start"`
	QEnd           int                  `json:"q_end"`
	QuestionGroups []QuestionGroup      `json:"questionGroups" gorm:"foreignKey:SectionID"`
}

type SectionListening struct {
	MediaURL       string `json:"mediaURL"`
	MediaLengthMin int    `json:"mediaLengthMin"`
}

type SectionReading struct {
	MainText string `json:"mainText"`
}

type SectionWriting struct {
	MinimumWords int `json:"minimumWords"`
}

type SectionSpeaking struct {
	AudioUrl string `json:"audioUrl"`
}
