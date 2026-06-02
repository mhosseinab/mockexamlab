package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type QuestionStatistic struct {
	BaseModel
	QuestionID         uuid.UUID           `json:"questionId"`
	TotalAnswer        int                 `json:"totalAnswer"`
	TotalCorrectAnswer int                 `json:"totalCorrectAnswer"`
	QuestionType       consts.QuestionType `json:"questionType"`
	ExtraProperty      pgtype.JSONB        `json:"extraProperty" gorm:"type:jsonb;default:'[]';not null"` // ExtraProperty
}
