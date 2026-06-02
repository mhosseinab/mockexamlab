package DTO

import (
	"github.com/google/uuid"
)

type QuestionStatisticResponse struct {
	Id             uuid.UUID `json:"id"`
	TAnswer        int       `json:"t_answer"`
	TCorrectAnswer int       `json:"t_correctAnswer"`
	QuestionType   string    `json:"questionType"`
}
