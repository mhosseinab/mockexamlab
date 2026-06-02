package api

/*import (
	"MockExamLab/internal/repository"
	"MockExamLab/internal/service"
	"firebase.google.com/go/v4/auth"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserAPI(db *gorm.DB, auth *auth.Client) UserAPI {
	wire.Build(repository.ProvideUserRepository, repository.ProvideFirebaseRepository, service.ProvideUserService, ProvideUserAPI)
	return UserAPI{}
}

func initTestAPI(db *gorm.DB) TestAPI {
	wire.Build(repository.ProvideTestRepository, service.ProvideTestService, ProvideTestAPI)
	return TestAPI{}
}

func initSectionAPI(db *gorm.DB) SectionAPI {
	wire.Build(repository.ProvideSectionRepository, service.ProvideSectionService, ProvideSectionAPI)
	return SectionAPI{}
}

func initQuestionGroupAPI(db *gorm.DB) QuestionGroupAPI {
	wire.Build(repository.ProvideQuestionGroupRepository, service.ProvideQuestionGroupService, ProvideQuestionGroupAPI)
	return QuestionGroupAPI{}
}

func initQuestionAPI(db *gorm.DB) QuestionAPI {
	wire.Build(repository.ProvideQuestionRepository, service.ProvideQuestionService, ProvideQuestionAPI)
	return QuestionAPI{}
}

func initUserTestAPI(db *gorm.DB) UserTestAPI {
	wire.Build(repository.ProvideUserTestRepository, repository.ProvideTestStatRepository, repository.ProvideQuestionRepository,
		repository.ProvideUserRepository, service.ProvideUserTestService, ProvideUserTestAPI)
	return UserTestAPI{}
}

func initUserAnswerAPI(db *gorm.DB) UserAnswerAPI {
	wire.Build(repository.ProvideUserAnswerRepository, service.ProvideUserAnswerService, ProvideUserAnswerAPI)
	return UserAnswerAPI{}
}

func initProvidePaddleAPI(db *gorm.DB) PaddleAPI {
	wire.Build(repository.ProvideSubscriptionEventRepositoryRepository, repository.ProvideUserRepository, service.ProvidePaddleService, ProvidePaddleAPI)
	return PaddleAPI{}
}
*/
