package models

import (
	"MockExamLab/internal/models/consts"
)

type User struct {
	BaseModel
	FirebaseId               string                   `json:"firebaseId" gorm:"index"`
	DisplayName              string                   `json:"displayName"`
	Email                    string                   `json:"email" filter:"param:email;searchable" gorm:"column:email"`
	LastLogin                int64                    `json:"lastLogin"`
	Role                     consts.UserRole          `json:"role"`
	InvitationCode           string                   `json:"invitationCode"`
	InvitedUserCode          string                   `json:"invitedUserCode"`
	SubscriptionExpirationMs int64                    `json:"subscriptionExpirationMs" gorm:"default:0"`
	SubscriptionState        consts.SubscriptionState `json:"subscriptionState" gorm:"default:2"`
	UserTests                []UserTest               `json:"userTests" gorm:"foreignKey:UserID;"`
	UserSubscriptionEvent    []SubscriptionEvent      `json:"userSubscriptionEvent" gorm:"foreignKey:UserID;"`
}
