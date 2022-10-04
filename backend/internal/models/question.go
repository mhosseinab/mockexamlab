package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
)

type Question struct {
	BaseModel
	GroupID        uuid.UUID            `json:"groupId"`
	CorrectAnswer  string               `json:"correctAnswer"`
	Component      consts.TestComponent `json:"component"`
	QNumber        int                  `json:"q_number"`
	CreatorId      uuid.UUID            `json:"creatorId"`
	LastModifierId uuid.UUID            `json:"lastModifierId"`
	Statistic      QuestionStatistic    `json:"statistic" gorm:"foreignKey:QuestionID;"`
	UserAnswers    []UserAnswer         `json:"userAnswers" gorm:"foreignKey:QuestionID;"`
}

/*func (u *Question) AfterCreate(tx *gorm.DB) (err error) {

	return
}*/
