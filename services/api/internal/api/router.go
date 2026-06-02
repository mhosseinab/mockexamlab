package api

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/repository"
	"errors"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

func SetupRouter(db *gorm.DB, auth *auth.Client) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"https://www.mockexamlab.com",
		"https://app.mockexamlab.com",
		"https://dev.mockexamlab.com",
	}
	config.AllowMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}
	config.AllowHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	}
	config.AddExposeHeaders("X-Total-Count")
	r.Use(cors.New(config))

	userAPI := initUserAPI(db, auth)
	testAPI := initTestAPI(db)
	sectionAPI := initSectionAPI(db)
	questionGroupAPI := initQuestionGroupAPI(db)
	questionAPI := initQuestionAPI(db)
	userTestAPI := initUserTestAPI(db)
	userAnswerAPI := initUserAnswerAPI(db)
	paddleAPI := initProvidePaddleAPI(db)

	v1 := r.Group("/api/v1")
	{
		v1.Use(authMiddleware(db, auth))

		v1.GET("/test/:id", testAPI.FindById)
		v1.GET("/test/full/:id", testAPI.FindByIdWithPreload)
		v1.GET("/test/all", testAPI.FindAll)
		v1.POST("/test", testAPI.Create)
		v1.PUT("/test", testAPI.Update)
		v1.DELETE("/test/:id", testAPI.DeleteById)

		//Section
		v1.GET("/section/:id", sectionAPI.FindById)
		v1.GET("/section/all", sectionAPI.FindAll)
		v1.POST("/section", sectionAPI.Create)
		v1.PUT("/section", sectionAPI.Update)
		v1.DELETE("/section/:id", sectionAPI.DeleteById)

		//Question Group
		v1.GET("/question-group/:id", questionGroupAPI.FindById)
		v1.GET("/question-group/all", questionGroupAPI.FindAllBySectionId)
		v1.POST("/question-group", questionGroupAPI.Create)
		v1.PUT("/question-group", questionGroupAPI.Update)
		v1.DELETE("/question-group/:id", questionGroupAPI.DeleteById)

		//Question
		v1.GET("/question/:id", questionAPI.FindById)
		v1.GET("/question/all", questionAPI.FindAllByGroupId)
		v1.GET("/question/test/all", questionAPI.FindAllByTestId)
		v1.POST("/question", questionAPI.Create)
		v1.PUT("/question", questionAPI.Update)
		v1.DELETE("/question/:id", sectionAPI.DeleteById)

		//User Test
		v1.GET("/user-test/:id", userTestAPI.FindById)
		v1.GET("/user-test/all", userTestAPI.FindAllByUserId)
		v1.POST("/user-test", userTestAPI.Create)
		v1.PUT("/user-test/submit", userTestAPI.SubmitTest)
		v1.DELETE("/user-test/:id", userTestAPI.DeleteById)

		//User Answer
		v1.GET("/user-answer/:id", userAnswerAPI.FindById)
		v1.GET("/user-answer/all", userAnswerAPI.FindAllByUserTestId)
		v1.POST("/user-answer", userAnswerAPI.Create)
		v1.POST("/user-answer/speaking", userAnswerAPI.CreateSpeaking)
		v1.POST("/user-answer/batch", userAnswerAPI.BatchCreate)
		v1.PUT("/user-answer/answer", userAnswerAPI.UpdateAnswer)
		v1.PUT("/user-answer/answer/speaking", userAnswerAPI.UpdateSpeakingAnswer)
		v1.PUT("/user-answer/marker", userAnswerAPI.UpdateMarker)
		v1.DELETE("/user-answer/:id", userAnswerAPI.DeleteById)

		v1.POST("/user/auth", userAPI.UserLoginOrSignUp)
		v1.GET("/user", userAPI.FindWithFilterAndPagination)

	}
	r.POST("/webhook", paddleAPI.SubscriptionWebhook)

	//r.POST("/api/v1/user/login", userAPI.UserLogin)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func authMiddleware(db *gorm.DB, auth *auth.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := getClaimsFromJWTHeader(context, db, auth)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "authentication error")
			return
		}
		context.Set("user", user) // add user info to session
		context.Next()
	}
}

func getClaimsFromJWTHeader(c *gin.Context, db *gorm.DB, auth *auth.Client) (*models.UserClaims, error) {
	t := c.GetHeader("Authorization")
	if t == "" {
		t = c.Query("token")
	}
	fr := repository.ProvideFirebaseRepository(auth)
	userClaim, errVerifyToken := fr.VerifyIdToken(t)
	if errVerifyToken != nil {
		return nil, errVerifyToken
	}
	if userClaim.UserId == uuid.Nil {
		ur := repository.ProvideUserRepository(db)
		if user, err := ur.GetUserByFirebaseId(userClaim.FirebaseId); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		} else {
			userClaim.UserId = user.ID
		}
	}
	return userClaim, nil
}
