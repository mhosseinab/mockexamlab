package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type QuestionGroup struct {
	BaseModel
	SectionID      uuid.UUID           `json:"sectionId"`
	Title          string              `json:"title"`
	Desc           string              `json:"desc"`
	QStart         int                 `json:"q_start"`
	QEnd           int                 `json:"q_end"`
	QuestionType   consts.QuestionType `json:"questionType"`                                          // Question Type
	ExtraProperty  pgtype.JSONB        `json:"extraProperty" gorm:"type:jsonb;default:'[]';not null"` // ExtraProperty
	Questions      []Question          `json:"questions" gorm:"foreignKey:GroupID;"`
	CreatorId      uuid.UUID           `json:"creatorId"`
	LastModifierId uuid.UUID           `json:"lastModifierId"`
}

type QuestionGroupMultipleChoice struct {
	QNumber int               `json:"q_number"`
	Title   string            `json:"title"`
	Answers map[string]string `json:"answers"`
}

type QuestionGroupMap struct {
	MapUrl string                 `json:"mapUrl"`
	Items  []QuestionGroupMapItem `json:"items"`
}

type QuestionGroupMapItem struct {
	QNumber int    `json:"q_number"`
	Title   string `json:"title"`
}

type QuestionGroupNoteCompletion struct {
	Note string `json:"note"`
}

type QuestionGroupIdentifyingInformation struct {
	QNumber int    `json:"q_number"`
	Title   string `json:"title"`
}
type QuestionGroupSpeaking struct {
	QNumber int    `json:"q_number"`
	Topic   string `json:"topic"`
}
type QuestionGroupWriting struct {
	QNumber int    `json:"q_number"`
	Topic   string `json:"topic"`
}
