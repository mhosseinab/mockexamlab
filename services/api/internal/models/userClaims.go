package models

import (
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId      uuid.UUID       `json:"userId"`
	FirebaseId  string          `json:"firebaseId"`
	DisplayName string          `json:"displayName"`
	Email       string          `json:"email"`
	Role        consts.UserRole `json:"role"`
}
