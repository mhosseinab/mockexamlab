package DTO

import (
	"github.com/google/uuid"
	"time"
)

type QuestionCreateRequest struct {
	GroupId       uuid.UUID `json:"groupId"`
	QNumber       int       `json:"q_number"`
	TestComponent int       `json:"testComponent"`
	QuestionType  int       `json:"questionType"`
	CorrectAnswer string    `json:"correctAnswer"`
}

type QuestionCreateResponse struct {
	Id             uuid.UUID                 `json:"id"`
	StatId         uuid.UUID                 `json:"statId"`
	QNumber        int                       `json:"q_number"`
	TAnswer        int                       `json:"t_answer"`
	TCorrectAnswer int                       `json:"t_correct_answer"`
	TestComponent  string                    `json:"testComponent"`
	CorrectAnswer  string                    `json:"correctAnswer"`
	Statistic      QuestionStatisticResponse `json:"statistic"`
	CreatedAt      time.Time                 `json:"createdAt"`
}

type QuestionUpdateRequest struct {
	Id            uuid.UUID `json:"id"`
	QNumber       int       `json:"q_number"`
	CorrectAnswer string    `json:"correctAnswer"`
	TestComponent int       `json:"testComponent"`
}

type QuestionUpdateResponse struct {
	Id            uuid.UUID `json:"id"`
	QNumber       int       `json:"q_number"`
	CorrectAnswer string    `json:"correctAnswer"`
	UpdateAt      time.Time `json:"updateAt"`
}

type QuestionResponse struct {
	Id            uuid.UUID                 `json:"id"`
	QNumber       int                       `json:"q_number"`
	CorrectAnswer string                    `json:"correctAnswer"`
	Statistic     QuestionStatisticResponse `json:"statistic"`
}

type QuestionWithoutAnswerResponse struct {
	Id      uuid.UUID `json:"id"`
	QNumber int       `json:"q_number"`
}

type QuestionWithUserAnswerResponse struct {
	Id            uuid.UUID                 `json:"id"`
	QNumber       int                       `json:"q_number"`
	CorrectAnswer string                    `json:"correctAnswer"`
	Statistic     QuestionStatisticResponse `json:"statistic"`
	UserAnswer    *UserAnswerResponse       `json:"userAnswer"`
}
