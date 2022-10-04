package DTO

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"time"
)

type TestCreateRequest struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Module        int               `json:"module"`
	TestDate      int64             `json:"testDate"`
	TotalTestTime int               `json:"totalTestTime"`
	Detail        models.TestDetail `json:"detail"`
}

type TestCreateResponse struct {
	Id          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Module      string             `json:"module"`
	TestDate    time.Time          `json:"testDate"`
	Detail      *models.TestDetail `json:"detail"`
	CreatedAt   time.Time          `json:"createdAt"`
}

type TestUpdateRequest struct {
	Id            uuid.UUID         `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Module        int               `json:"module"`
	TestDate      int64             `json:"testDate"`
	Detail        models.TestDetail `json:"detail"`
	TotalTestTime int               `json:"totalTestTime"`
}

type TestUpdateResponse struct {
	Id          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Module      string             `json:"module"`
	TestDate    time.Time          `json:"testDate"`
	Detail      *models.TestDetail `json:"detail"`
	UpdateAt    time.Time          `json:"updateAt"`
}

type TestResponse struct {
	Id            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Module        string             `json:"module"`
	TestDate      time.Time          `json:"testDate"`
	TotalTestTime int                `json:"totalTestTime"`
	Detail        *models.TestDetail `json:"detail"`
	Sections      []SectionResponse  `json:"sections"`
}
