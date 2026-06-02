package DTO

import (
	"github.com/google/uuid"
	"time"
)

type UserTestCreateRequest struct {
	TestId   uuid.UUID `json:"testId"`
	UserId   uuid.UUID `json:"userId"`
	TestType int       `json:"testType"`
	TestDate int64     `json:"testDate"`
}

type UserTestCreateResponse struct {
	Id           uuid.UUID `json:"id"`
	TestId       uuid.UUID `json:"testId"`
	UserId       uuid.UUID `json:"userId"`
	TestType     string    `json:"testType"`
	TestDuration int       `json:"testDuration"`
	TestDate     time.Time `json:"testDate"`
	CreatedAt    time.Time `json:"createdAt"`
}

type UserTestUpdateRequest struct {
	Id           uuid.UUID `json:"id"`
	TestDuration int       `json:"testDuration"`
}

type UserTestUpdateResponse struct {
	Id           uuid.UUID `json:"id"`
	TestId       uuid.UUID `json:"testId"`
	UserId       uuid.UUID `json:"userId"`
	TestType     string    `json:"testType"`
	TestDuration int       `json:"testDuration"`
	TestDate     time.Time `json:"testDate"`
	UpdateAt     time.Time `json:"updateAt"`
}

type UserTestResponse struct {
	Id           uuid.UUID            `json:"id"`
	TestId       uuid.UUID            `json:"testId"`
	UserId       uuid.UUID            `json:"userId"`
	TestType     string               `json:"testType"`
	TestDuration int                  `json:"testDuration"`
	TestDate     time.Time            `json:"testDate"`
	UserAnswers  []UserAnswerResponse `json:"userAnswers"`
}

type UserTestResultResponse struct {
	Id             uuid.UUID                        `json:"id"`
	TestId         uuid.UUID                        `json:"testId"`
	UserId         uuid.UUID                        `json:"userId"`
	TestType       string                           `json:"testType"`
	TestDuration   int                              `json:"testDuration"`
	OverallScore   float32                          `json:"overallScore"`
	ListeningScore float32                          `json:"listeningScore"`
	ReadingScore   float32                          `json:"readingScore"`
	WritingScore   float32                          `json:"writingScore"`
	SpeakingScore  float32                          `json:"speakingScore"`
	TestDate       time.Time                        `json:"testDate"`
	AnswerSheet    []QuestionWithUserAnswerResponse `json:"answerSheet"`
	Stat           *UserTestStat                    `json:"stat"`
}
