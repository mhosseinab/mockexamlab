// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"MockExamLab/internal/repository"
	"MockExamLab/internal/service"
	"firebase.google.com/go/v4/auth"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initUserAPI(db *gorm.DB, auth2 *auth.Client) UserAPI {
	userRepository := repository.ProvideUserRepository(db)
	firebaseRepository := repository.ProvideFirebaseRepository(auth2)
	userService := service.ProvideUserService(userRepository, firebaseRepository)
	userAPI := ProvideUserAPI(userService)
	return userAPI
}

func initTestAPI(db *gorm.DB) TestAPI {
	testRepository := repository.ProvideTestRepository(db)
	testService := service.ProvideTestService(testRepository)
	testAPI := ProvideTestAPI(testService)
	return testAPI
}

func initSectionAPI(db *gorm.DB) SectionAPI {
	sectionRepository := repository.ProvideSectionRepository(db)
	sectionService := service.ProvideSectionService(sectionRepository)
	sectionAPI := ProvideSectionAPI(sectionService)
	return sectionAPI
}

func initQuestionGroupAPI(db *gorm.DB) QuestionGroupAPI {
	questionGroupRepository := repository.ProvideQuestionGroupRepository(db)
	questionGroupService := service.ProvideQuestionGroupService(questionGroupRepository)
	questionGroupAPI := ProvideQuestionGroupAPI(questionGroupService)
	return questionGroupAPI
}

func initQuestionAPI(db *gorm.DB) QuestionAPI {
	questionRepository := repository.ProvideQuestionRepository(db)
	questionService := service.ProvideQuestionService(questionRepository)
	questionAPI := ProvideQuestionAPI(questionService)
	return questionAPI
}

func initUserTestAPI(db *gorm.DB) UserTestAPI {
	userTestRepository := repository.ProvideUserTestRepository(db)
	questionRepository := repository.ProvideQuestionRepository(db)
	testStatRepository := repository.ProvideTestStatRepository(db)
	userRepository := repository.ProvideUserRepository(db)
	userTestService := service.ProvideUserTestService(userTestRepository, questionRepository, testStatRepository, userRepository)
	userTestAPI := ProvideUserTestAPI(userTestService)
	return userTestAPI
}

func initUserAnswerAPI(db *gorm.DB) UserAnswerAPI {
	userAnswerRepository := repository.ProvideUserAnswerRepository(db)
	userAnswerService := service.ProvideUserAnswerService(userAnswerRepository)
	userAnswerAPI := ProvideUserAnswerAPI(userAnswerService)
	return userAnswerAPI
}

func initProvidePaddleAPI(db *gorm.DB) PaddleAPI {
	subscriptionEventRepositoryRepository := repository.ProvideSubscriptionEventRepositoryRepository(db)
	userRepository := repository.ProvideUserRepository(db)
	paddleService := service.ProvidePaddleService(subscriptionEventRepositoryRepository, userRepository)
	paddleAPI := ProvidePaddleAPI(paddleService)
	return paddleAPI
}
