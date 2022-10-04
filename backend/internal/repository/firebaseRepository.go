package repository

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/consts"
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
)

type FirebaseRepository struct {
	auth *auth.Client
}

func ProvideFirebaseRepository(auth *auth.Client) FirebaseRepository {
	return FirebaseRepository{auth: auth}
}

func (r *FirebaseRepository) UpdateClaims(fbId string, claims map[string]interface{}) error {
	if err := r.auth.SetCustomUserClaims(context.Background(), fbId, claims); err != nil {
		return err
	}
	return nil
}

func (r *FirebaseRepository) VerifyIdToken(idToken string) (*models.UserClaims, error) {
	var userId uuid.UUID
	role := consts.Basic
	email := ""
	displayName := ""
	token, err := r.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	if val, ok := token.Claims["app_user_id"]; ok {
		var errParsUserId error
		userId, errParsUserId = uuid.Parse(val.(string))
		if errParsUserId != nil {
			userId = uuid.Nil
		}
	}

	if val, ok := token.Claims["role"]; ok {
		role = consts.UserRole(int(val.(float64)))
	}

	if val, ok := token.Claims["name"]; ok {
		displayName = val.(string)
	}
	if val, ok := token.Claims["email"]; ok {
		email = val.(string)
	}

	return &models.UserClaims{
		UserId:      userId,
		FirebaseId:  token.UID,
		Role:        role,
		DisplayName: displayName,
		Email:       email,
	}, nil
}

func (r *FirebaseRepository) GetUser(fbId string) (*auth.UserRecord, error) {
	if user, err := r.auth.GetUser(context.Background(), fbId); err != nil {
		return nil, err
	} else {
		return user, nil
	}

}
