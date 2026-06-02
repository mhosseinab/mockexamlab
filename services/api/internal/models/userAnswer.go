package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type UserAnswer struct {
	BaseModel
	UserTestID    uuid.UUID `json:"userTestId"`
	QuestionID    uuid.UUID `json:"questionId"`
	MarkerID      uuid.UUID `json:"markerId"`
	Answer        string    `json:"answer"`
	MarkerScore   float32   `json:"markerScore"`
	MarkerComment string    `json:"markerComment"`
}

func (u *UserAnswer) BeforeCreate(tx *gorm.DB) (err error) {
	question := Question{}
	if errr := tx.Model(&Question{}).Preload("Statistic").Where(Question{BaseModel: BaseModel{ID: u.QuestionID}}).First(&question).Error; errr != nil {
		return errr
	}
	if question.Component == consts.Listening || question.Component == consts.Reading {
		if strings.Trim(strings.ToLower(question.CorrectAnswer), " ") == strings.Trim(strings.ToLower(u.Answer), " ") {
			u.MarkerScore++
			question.Statistic.TotalCorrectAnswer++
		}
		question.Statistic.TotalAnswer++
		if errUpdate := tx.Model(&question.Statistic).Updates(QuestionStatistic{
			TotalCorrectAnswer: question.Statistic.TotalCorrectAnswer,
			TotalAnswer:        question.Statistic.TotalAnswer,
		}).Error; errUpdate != nil {
			return errUpdate
		}
	}
	return
}

func (u *UserAnswer) AfterUpdate(tx *gorm.DB) (err error) {
	question := Question{}
	if errr := tx.Model(&Question{}).Preload("Statistic").Where(Question{BaseModel: BaseModel{ID: u.QuestionID}}).First(&question).Error; errr != nil {
		return errr
	}
	if question.Component == consts.Listening || question.Component == consts.Reading {
		if strings.Trim(strings.ToLower(question.CorrectAnswer), " ") == strings.Trim(strings.ToLower(u.Answer), " ") {
			u.MarkerScore++
			question.Statistic.TotalCorrectAnswer++
		} else {
			if u.MarkerScore > 0 {
				u.MarkerScore--
			}
			if question.Statistic.TotalCorrectAnswer > 0 {
				question.Statistic.TotalCorrectAnswer--
			}
		}

		if errUpdateScore := tx.Model(&u).UpdateColumn("marker_score", u.MarkerScore).Error; errUpdateScore != nil {
			return errUpdateScore
		}
		if errUpdateStat := tx.Model(&question.Statistic).Updates(QuestionStatistic{
			TotalCorrectAnswer: question.Statistic.TotalCorrectAnswer,
		}).Error; errUpdateStat != nil {
			return errUpdateStat
		}
	}
	return
}

/*func (u *UserAnswer) AfterUpdate(tx *gorm.DB) (err error) {
	if mErr := tx.Model(&UserTest{}).
		Where(UserTest{BaseModel: BaseModel{ID: u.UserTestID}}).
		Update("marker_score", gorm.Expr("marker_score + ?", u.MarkerScore)).
		Error; mErr != nil {
		return mErr
	}
	return
}
*/
