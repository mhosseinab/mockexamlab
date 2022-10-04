package DTO

import (
	"github.com/google/uuid"
	"time"
)

type UserAnswerCreateRequest struct {
	UserTestId uuid.UUID `json:"userTestId"`
	QuestionId uuid.UUID `json:"questionId"`
	Answer     string    `json:"answer"`
}

type UserAnswerCreateResponse struct {
	Id         uuid.UUID `json:"id"`
	UserTestId uuid.UUID `json:"userTestId"`
	QuestionId uuid.UUID `json:"questionId"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"createdAt"`
}

type UserAnswerUpdateAnswerRequest struct {
	Id     uuid.UUID `json:"id"`
	Answer string    `json:"answer"`
}

type UserAnswerUpdateAnswerResponse struct {
	Id         uuid.UUID `json:"id"`
	UserTestId uuid.UUID `json:"userTestId"`
	QuestionId uuid.UUID `json:"questionId"`
	Answer     string    `json:"answer"`
	UpdateAt   time.Time `json:"updateAt"`
}

type UserAnswerUpdateMarkerRequest struct {
	Id            uuid.UUID `json:"id"`
	MarkerId      uuid.UUID `json:"markerId"`
	MarkerScore   float32   `json:"markerScore"`
	MarkerComment string    `json:"markerComment"`
}

type UserAnswerUpdateMarkerResponse struct {
	Id            uuid.UUID `json:"id"`
	UserTestId    uuid.UUID `json:"userTestId"`
	QuestionId    uuid.UUID `json:"questionId"`
	Answer        string    `json:"answer"`
	MarkerId      uuid.UUID `json:"markerId"`
	MarkerScore   float32   `json:"markerScore"`
	MarkerComment string    `json:"markerComment"`
	UpdateAt      time.Time `json:"updateAt"`
}

type UserAnswerResponse struct {
	Id            uuid.UUID `json:"id"`
	UserTestId    uuid.UUID `json:"userTestId"`
	QuestionId    uuid.UUID `json:"questionId"`
	Answer        string    `json:"answer"`
	MarkerId      uuid.UUID `json:"markerId"`
	MarkerScore   float32   `json:"markerScore"`
	MarkerComment string    `json:"markerComment"`
}

/*type UserAnswerWithScoreResponse struct {
	Id            uuid.UUID `json:"id"`
	UserTestId    uuid.UUID `json:"userTestId"`
	QuestionId    uuid.UUID `json:"questionId"`
	Answer        string    `json:"answer"`
	MarkerId      uuid.UUID `json:"markerId"`
	MarkerScore   int       `json:"markerScore"`
	MarkerComment string       `json:"markerComment"`
}*/
