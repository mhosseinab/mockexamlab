package DTO

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"time"
)

type SectionCreateRequest struct {
	TestId           uuid.UUID               `json:"testId"`
	Title            string                  `json:"title"`
	Description      string                  `json:"description"`
	ComponentType    int                     `json:"componentType"`
	TotalSectionTime int                     `json:"totalSectionTime"`
	QStart           int                     `json:"q_start"`
	QEnd             int                     `json:"q_end"`
	SectionListening models.SectionListening `json:"sectionListening"`
	SectionReading   models.SectionReading   `json:"sectionReading"`
	SectionWriting   models.SectionWriting   `json:"sectionWriting"`
	SectionSpeaking  models.SectionSpeaking  `json:"sectionSpeaking"`
}

type SectionCreateResponse struct {
	Id               uuid.UUID                `json:"id"`
	Title            string                   `json:"title"`
	Description      string                   `json:"description"`
	ComponentType    string                   `json:"componentType"`
	SectionListening *models.SectionListening `json:"sectionListening"`
	SectionReading   *models.SectionReading   `json:"sectionReading"`
	SectionWriting   *models.SectionWriting   `json:"sectionWriting"`
	SectionSpeaking  *models.SectionSpeaking  `json:"sectionSpeaking"`
	CreatedAt        time.Time                `json:"createdAt"`
}

type SectionUpdateRequest struct {
	Id               uuid.UUID               `json:"id"`
	Title            string                  `json:"title"`
	Description      string                  `json:"description"`
	ComponentType    int                     `json:"componentType"`
	TotalSectionTime int                     `json:"totalSectionTime"`
	QStart           int                     `json:"q_start"`
	QEnd             int                     `json:"q_end"`
	SectionListening models.SectionListening `json:"sectionListening"`
	SectionReading   models.SectionReading   `json:"sectionReading"`
	SectionWriting   models.SectionWriting   `json:"sectionWriting"`
	SectionSpeaking  models.SectionSpeaking  `json:"sectionSpeaking"`
}

type SectionUpdateResponse struct {
	Id               uuid.UUID                `json:"id"`
	Title            string                   `json:"title"`
	Description      string                   `json:"description"`
	ComponentType    string                   `json:"componentType"`
	SectionListening *models.SectionListening `json:"sectionListening"`
	SectionReading   *models.SectionReading   `json:"sectionReading"`
	SectionWriting   *models.SectionWriting   `json:"sectionWriting"`
	SectionSpeaking  *models.SectionSpeaking  `json:"sectionSpeaking"`
	UpdateAt         time.Time                `json:"updateAt"`
}

type SectionResponse struct {
	Id               uuid.UUID                `json:"id"`
	Title            string                   `json:"title"`
	Description      string                   `json:"description"`
	ComponentType    string                   `json:"componentType"`
	TotalSectionTime int                      `json:"totalSectionTime"`
	QStart           int                      `json:"q_start"`
	QEnd             int                      `json:"q_end"`
	QuestionGroups   []QuestionGroupResponse  `json:"questionGroups"`
	SectionListening *models.SectionListening `json:"sectionListening"`
	SectionReading   *models.SectionReading   `json:"sectionReading"`
	SectionWriting   *models.SectionWriting   `json:"sectionWriting"`
	SectionSpeaking  *models.SectionSpeaking  `json:"sectionSpeaking"`
}
