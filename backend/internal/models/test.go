package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"time"
)

type Test struct {
	BaseModel
	Name           string            `json:"name" filter:"param:name;searchable" gorm:"column:name"`
	Description    string            `json:"description"`
	TestDate       time.Time         `json:"testDate"` // Original Exam Date
	Module         consts.TestModule `json:"module" gorm:"column:module" filter:"filterable"`
	Sections       []Section         `json:"sections" gorm:"foreignKey:TestID"`
	UserTests      []UserTest        `json:"userTests" gorm:"foreignKey:TestID"`
	Stat           *TestStat         `json:"stat" gorm:"foreignKey:TestId"`
	TestTime       int               `json:"testTime"`                                              // test time in minute
	ExtraProperty  pgtype.JSONB      `json:"extraProperty" gorm:"type:jsonb;default:'{}';not null"` // ExtraProperty
	CreatorId      uuid.UUID         `json:"creatorId"`
	LastModifierId uuid.UUID         `json:"lastModifierId"`
}

type TestDetail struct {
	ComponentsDetail []ComponentDetail `json:"componentsDetail"`
}

type ComponentDetail struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        int    `json:"time"`
}
