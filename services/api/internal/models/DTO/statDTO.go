package DTO

import "MockExamLab/internal/models"

type UserSkillsAvgStat struct {
	ListeningOverall float32 `json:"listeningOverall"`
	ReadingOverall   float32 `json:"readingOverall"`
	WritingOverall   float32 `json:"writingOverall"`
	SpeakingOverall  float32 `json:"speakingOverall"`
}

type UserTestStat struct {
	Skill             models.TestSkillsStat `json:"skill"`
	Average           UserSkillsAvgStat     `json:"average"`
	AverageTotalSkill UserSkillsAvgStat     `json:"averageTotalSkill"`
}
