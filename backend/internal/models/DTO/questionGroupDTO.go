package DTO

import (
	"MockExamLab/internal/models"
	"github.com/google/uuid"
	"time"
)

type QuestionGroupCreateRequest struct {
	SectionId              uuid.UUID                                    `json:"sectionId"`
	Title                  string                                       `json:"title"`
	Description            string                                       `json:"description"`
	QuestionType           int                                          `json:"componentType"`
	QStart                 int                                          `json:"q_start"`
	QEnd                   int                                          `json:"q_end"`
	MultipleChoice         []models.QuestionGroupMultipleChoice         `json:"multipleChoice"`
	IdentifyingInformation []models.QuestionGroupIdentifyingInformation `json:"identifyingInformation"`
	Map                    models.QuestionGroupMap                      `json:"map"`
	NoteCompletion         models.QuestionGroupNoteCompletion           `json:"noteCompletion"`
	SpeakingTopic          models.QuestionGroupSpeaking                 `json:"SpeakingTopic"`
	WritingTopic           models.QuestionGroupWriting                  `json:"writingTopic"`
}

type QuestionGroupCreateResponse struct {
	Id                     uuid.UUID                                     `json:"id"`
	Title                  string                                        `json:"title"`
	Description            string                                        `json:"description"`
	QuestionType           string                                        `json:"componentType"`
	MultipleChoice         *[]models.QuestionGroupMultipleChoice         `json:"multipleChoice"`
	IdentifyingInformation *[]models.QuestionGroupIdentifyingInformation `json:"identifyingInformation"`
	Map                    *models.QuestionGroupMap                      `json:"map"`
	NoteCompletion         *models.QuestionGroupNoteCompletion           `json:"noteCompletion"`
	SpeakingTopic          *models.QuestionGroupSpeaking                 `json:"SpeakingTopic"`
	WritingTopic           *models.QuestionGroupWriting                  `json:"writingTopic"`
	CreatedAt              time.Time                                     `json:"createdAt"`
}

type QuestionGroupUpdateRequest struct {
	Id                     uuid.UUID                                    `json:"id"`
	Title                  string                                       `json:"title"`
	Description            string                                       `json:"description"`
	QuestionType           int                                          `json:"componentType"`
	QStart                 int                                          `json:"q_start"`
	QEnd                   int                                          `json:"q_end"`
	MultipleChoice         []models.QuestionGroupMultipleChoice         `json:"multipleChoice"`
	IdentifyingInformation []models.QuestionGroupIdentifyingInformation `json:"identifyingInformation"`
	Map                    models.QuestionGroupMap                      `json:"map"`
	NoteCompletion         models.QuestionGroupNoteCompletion           `json:"noteCompletion"`
	SpeakingTopic          models.QuestionGroupSpeaking                 `json:"SpeakingTopic"`
	WritingTopic           models.QuestionGroupWriting                  `json:"writingTopic"`
}

type QuestionGroupUpdateResponse struct {
	Id                     uuid.UUID                                     `json:"id"`
	Title                  string                                        `json:"title"`
	Description            string                                        `json:"description"`
	QuestionType           string                                        `json:"componentType"`
	QStart                 int                                           `json:"q_start"`
	QEnd                   int                                           `json:"q_end"`
	MultipleChoice         *[]models.QuestionGroupMultipleChoice         `json:"multipleChoice"`
	IdentifyingInformation *[]models.QuestionGroupIdentifyingInformation `json:"identifyingInformation"`
	Map                    *models.QuestionGroupMap                      `json:"map"`
	NoteCompletion         *models.QuestionGroupNoteCompletion           `json:"noteCompletion"`
	SpeakingTopic          *models.QuestionGroupSpeaking                 `json:"SpeakingTopic"`
	WritingTopic           *models.QuestionGroupWriting                  `json:"writingTopic"`
	UpdateAt               time.Time                                     `json:"updateAt"`
}

type QuestionGroupResponse struct {
	Id                     uuid.UUID                                     `json:"id"`
	Title                  string                                        `json:"title"`
	Description            string                                        `json:"description"`
	QuestionType           string                                        `json:"componentType"`
	QStart                 int                                           `json:"q_start"`
	QEnd                   int                                           `json:"q_end"`
	Questions              []QuestionWithoutAnswerResponse               `json:"questions"`
	MultipleChoice         *[]models.QuestionGroupMultipleChoice         `json:"multipleChoice"`
	IdentifyingInformation *[]models.QuestionGroupIdentifyingInformation `json:"identifyingInformation"`
	Map                    *models.QuestionGroupMap                      `json:"map"`
	NoteCompletion         *models.QuestionGroupNoteCompletion           `json:"noteCompletion"`
	SpeakingTopic          *models.QuestionGroupSpeaking                 `json:"SpeakingTopic"`
	WritingTopic           *models.QuestionGroupWriting                  `json:"writingTopic"`
}
