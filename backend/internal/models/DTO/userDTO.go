package DTO

import "github.com/google/uuid"

type UserSignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Family   string `json:"family" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpOrLoginUserResponse struct {
	UserId                   uuid.UUID `json:"userId"`
	Email                    string    `json:"email"`
	Alias                    string    `json:"alias"`
	SubscriptionState        string    `json:"subscriptionState"`
	SubscriptionExpirationMs int64     `json:"subscriptionExpirationMs"`
	//Token  string    `json:"token"`
}
