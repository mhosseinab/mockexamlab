package cmd

import (
	_ "MockExamLab/docs"
	"MockExamLab/internal/api"
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/consts"
	"MockExamLab/internal/tools"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/errorutils"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Runner() {
	dsn := tools.GetEnv("DB_DSN", "host=localhost user=postgres password=Ali1230272126 dbname=MockExam port=5432 sslmode=disable TimeZone=Asia/Tehran")
	//dsn := tools.GetEnv("DB_DSN", "host=app-db user=mockexam password=OzOdGmEYJ dbname=MockExam port=5432 sslmode=disable TimeZone=Asia/Tehran")
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	myDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println(err)
	}

	if err2 := myDb.AutoMigrate(
		&models.User{},
		&models.Test{},
		&models.TestStat{},
		&models.Section{},
		&models.QuestionGroup{},
		&models.Question{},
		&models.QuestionStatistic{},
		&models.User{},
		&models.UserTest{},
		&models.SubscriptionEvent{},
		&models.UserAnswer{}); err2 != nil {
		log.Println(err2.Error())
	}

	if auth, authErr := accessServicesSingleApp(); authErr != nil {
		log.Fatal(authErr)
	} else {
		if errCreateAdmin := createAdminUser(auth); errCreateAdmin != nil {
			log.Fatal(errCreateAdmin)
		} else {
			router := api.SetupRouter(myDb, auth)
			log.Fatal(router.Run())
		}
	}
}

func accessServicesSingleApp() (*auth.Client, error) {
	// [START access_services_single_app_golang]
	// Initialize default app
	opt := option.WithCredentialsFile("mock-exam-20921-firebase-adminsdk-jgdbs-37c535e3bf.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	// [END access_services_single_app_golang]

	return client, err
}

func createAdminUser(client *auth.Client) error {
	params := (&auth.UserToCreate{}).
		Email("admin@go7.ir").
		EmailVerified(true).
		Password("Ali123456").
		DisplayName("admin").
		Disabled(false)
	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		if errorutils.IsAlreadyExists(err) {
			return nil
		}
		return err
	}
	claims := map[string]interface{}{
		"role": consts.Admin,
	}
	if errClaim := client.SetCustomUserClaims(context.Background(), u.UID, claims); errClaim != nil {
		return errClaim
	}
	return nil
}
