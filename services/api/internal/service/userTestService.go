package service

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"MockExamLab/internal/repository"
	"github.com/google/uuid"
	"math"
)

type UserTestService struct {
	r  repository.UserTestRepository
	rq repository.QuestionRepository
	rs repository.TestStatRepository
	ru repository.UserRepository
}

func ProvideUserTestService(r repository.UserTestRepository, rq repository.QuestionRepository, rs repository.TestStatRepository, ru repository.UserRepository) UserTestService {
	return UserTestService{r: r,
		rq: rq,
		rs: rs,
		ru: ru}
}

func (t *UserTestService) Create(test models.UserTest) (*models.UserTest, error) {
	/*	if stat, _, err := t.ru.CheckUserSubscription(test.UserID); err != nil {
			return nil, err
		} else {
			if stat != consts.Subscriber {
				return nil, &customErrors.SubscriptionError{Status: stat}
			}
		}*/
	return t.r.Create(test)
}

func (t *UserTestService) SubmitTest(testId uuid.UUID, userTestId uuid.UUID) (*DTO.UserTestResultResponse, error) {
	question, errFetchQuestion := t.rq.FindAllByTestId(testId, userTestId)
	if errFetchQuestion != nil {
		return nil, errFetchQuestion
	}
	questionDTO := mapper.QuestionModelWithUserAnswerToArrayResponse(question)
	testScore, errCalculate := t.calculateScore(consts.IeltsAcademic, question)
	if errCalculate != nil {
		return nil, errCalculate
	}
	updatedUserTest, errUpdate := t.UpdateScore(userTestId, testScore)
	if errUpdate != nil {
		return nil, errUpdate
	}

	stat, errStat := t.getStat(userTestId, updatedUserTest.UserID, testScore)
	if errStat != nil {
		return nil, errStat
	}

	return &DTO.UserTestResultResponse{
		Id:             updatedUserTest.ID,
		TestId:         updatedUserTest.TestID,
		UserId:         updatedUserTest.UserID,
		OverallScore:   testScore.OverallScore,
		ListeningScore: testScore.ListeningScore,
		ReadingScore:   testScore.ReadingScore,
		WritingScore:   testScore.WritingScore,
		SpeakingScore:  testScore.SpeakingScore,
		AnswerSheet:    questionDTO,
		Stat:           stat,
	}, nil
}

func (t *UserTestService) getStat(testId uuid.UUID, userId uuid.UUID, testScore models.UserTestScore) (*DTO.UserTestStat, error) {
	/*	overalStat, errFetchOverall := t.r.OverallScoreStat(testId)
		if errFetchOverall != nil {
			return nil, errFetchOverall
		}

		listeningStat, errFetchListening := t.r.ListeningScoreStat(testId)
		if errFetchListening != nil {
			return nil, errFetchListening
		}

		readingStat, errFetchReading := t.r.ReadingScoreStat(testId)
		if errFetchReading != nil {
			return nil, errFetchReading
		}

		writingStat, errFetchWriting := t.r.WritingScoreStat(testId)
		if errFetchWriting != nil {
			return nil, errFetchWriting
		}

		speakingStat, errFetchSpeaking := t.r.SpeakingScoreStat(testId)
		if errFetchSpeaking != nil {
			return nil, errFetchSpeaking
		}*/

	Stat, errFetchStat := t.rs.FindById(testId)
	if errFetchStat != nil {
		return nil, errFetchStat
	}
	testStat, err := Stat.TestStatJsonToDTOModel()
	if err != nil {
		return nil, err
	}
	if errUpdateStat := t.updateTestStat(Stat.ID, testStat, testScore); errUpdateStat != nil {
		return nil, errUpdateStat
	}

	skillStat, errFetchSkill := t.r.AverageStatWithFilter(userId, consts.IeltsAcademic)
	if errFetchSkill != nil {
		return nil, errFetchSkill
	}

	skillStatAllTests, errFetchSkillAllTest := t.r.AverageStat(userId)
	if errFetchSkillAllTest != nil {
		return nil, errFetchSkillAllTest
	}
	return &DTO.UserTestStat{
		Skill:             *testStat,
		Average:           *skillStat,
		AverageTotalSkill: *skillStatAllTests,
	}, nil
}

func (t *UserTestService) updateTestStat(id uuid.UUID, stat *models.TestSkillsStat, testScore models.UserTestScore) error {
	if stat.Overall != nil {
		for _, item := range *stat.Overall {
			if item.Score == testScore.OverallScore {
				item.Count++
			}
		}
	}
	if stat.Listening != nil {
		for _, item := range *stat.Listening {
			if item.Score == testScore.ListeningScore {
				item.Count++
			}
		}
	}
	if stat.Reading != nil {
		for _, item := range *stat.Reading {
			if item.Score == testScore.ReadingScore {
				item.Count++
			}
		}
	}
	if stat.Writing != nil {
		for _, item := range *stat.Writing {
			if item.Score == testScore.WritingScore {
				item.Count++
			}
		}
	}
	if stat.Speaking != nil {
		for _, item := range *stat.Speaking {
			if item.Score == testScore.SpeakingScore {
				item.Count++
			}
		}
	}
	if m, err := mapper.TestStatUpdateRequestToModel(stat, id); err != nil {
		return err
	} else {
		if _, errUpdate := t.rs.Update(*m); errUpdate != nil {
			return errUpdate
		}
	}
	return nil
}

func (t *UserTestService) DeleteById(id uuid.UUID) error {
	return t.r.DeleteById(id)
}

func (t *UserTestService) UpdateScore(id uuid.UUID, score models.UserTestScore) (*models.UserTest, error) {

	return t.r.UpdateScore(id, score)
}

func (t *UserTestService) FindById(id uuid.UUID) (*models.UserTest, error) {
	return t.r.FindById(id)
}

func (t *UserTestService) FindAllByUserId(userId uuid.UUID, params *models.QueryParams) (*[]models.UserTest, int64, error) {
	return t.r.FindAllByUserId(userId, params)
}

func (t *UserTestService) calculateScore(testType consts.TestModule, questions *[]models.Question) (models.UserTestScore, error) {
	var listening []models.Question
	var reading []models.Question
	var writing []models.Question
	var speaking []models.Question
	listeningCorrectAnswer := 0
	readingCorrectAnswer := 0
	var listeningScore float32 = 0
	var readingScore float32 = 0
	var writingScore float32 = 0
	var speakingScore float32 = 0
	var overallScore float32 = 0

	for _, item := range *questions {
		switch item.Component {
		case consts.Speaking:
			speaking = append(speaking, item)
		case consts.Listening:
			listening = append(listening, item)
		case consts.Reading:
			reading = append(reading, item)
		case consts.Writing:
			writing = append(writing, item)
		}
	}
	for _, item := range listening {
		if item.UserAnswers != nil && len(item.UserAnswers) > 0 {
			listeningCorrectAnswer += int(item.UserAnswers[0].MarkerScore)
		}
	}
	for _, item := range reading {
		if item.UserAnswers != nil && len(item.UserAnswers) > 0 {
			readingCorrectAnswer += int(item.UserAnswers[0].MarkerScore)
		}
	}
	for _, item := range reading {
		if item.UserAnswers != nil && len(item.UserAnswers) > 0 {
			userAnswer := item.UserAnswers[0]
			if userAnswer.MarkerID != uuid.Nil {
				writingScore += userAnswer.MarkerScore
			} else {
				writingScore = -1
			}
		}
	}
	for _, item := range speaking {
		if item.UserAnswers != nil && len(item.UserAnswers) > 0 {
			userAnswer := item.UserAnswers[0]
			if userAnswer.MarkerID != uuid.Nil {
				speakingScore += userAnswer.MarkerScore
			} else {
				speakingScore = -1
			}
		}
	}
	switch listeningCorrectAnswer {
	case 11, 12:
		listeningScore = 4
	case 13, 14, 15:
		listeningScore = 4.5
	case 16, 17:
		listeningScore = 5
	case 18, 19, 20, 21, 22:
		listeningScore = 5.5
	case 23, 24, 25:
		listeningScore = 6
	case 26, 27, 28, 29:
		listeningScore = 6.5
	case 30, 31:
		listeningScore = 7
	case 32, 33, 34:
		listeningScore = 7.5
	case 35, 36:
		listeningScore = 8
	case 37, 38:
		listeningScore = 8.5
	case 39, 40:
		listeningScore = 9
	default:
		listeningScore = 0
	}
	if testType == consts.IeltsAcademic {
		switch readingCorrectAnswer {
		case 4, 5:
			readingScore = 2.5
		case 6, 7:
			readingScore = 3
		case 8, 9:
			readingScore = 3.5
		case 10, 11, 12:
			readingScore = 4
		case 13, 14:
			readingScore = 4.5
		case 15, 16, 17, 18:
			readingScore = 5
		case 19, 20, 21, 22:
			readingScore = 5.5
		case 23, 24, 25, 26:
			readingScore = 6
		case 27, 28, 29:
			readingScore = 6.5
		case 30, 31, 32:
			readingScore = 7
		case 33, 34:
			readingScore = 7.5
		case 35, 36:
			readingScore = 8
		case 37, 38:
			readingScore = 8.5
		case 39, 40:
			readingScore = 9
		}
	} else {
		switch readingCorrectAnswer {
		case 6, 8:
			readingScore = 2.5
		case 9, 10, 11:
			readingScore = 3
		case 12, 13, 14:
			readingScore = 3.5
		case 15, 16, 17, 18:
			readingScore = 4
		case 19, 20, 21, 22:
			readingScore = 4.5
		case 23, 24, 25, 26:
			readingScore = 5
		case 27, 28, 29:
			readingScore = 5.5
		case 30, 31:
			readingScore = 6
		case 32, 33:
			readingScore = 6.5
		case 34, 35:
			readingScore = 7
		case 36:
			readingScore = 7.5
		case 37, 38:
			readingScore = 8
		case 39:
			readingScore = 8.5
		case 40:
			readingScore = 9
		}
	}
	var speak float32 = 0
	var write float32 = 0
	if speakingScore < 0 {
		speak = 0
	}
	if writingScore < 0 {
		write = 0
	}

	overallScore = float32(roundFloat(float64((listeningScore+readingScore+write+speak)/4), 1))
	return models.UserTestScore{
		OverallScore:   overallScore,
		ListeningScore: listeningScore,
		ReadingScore:   readingScore,
		WritingScore:   writingScore,
		SpeakingScore:  speakingScore,
	}, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
