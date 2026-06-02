package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type SubscriptionEvent struct {
	BaseModel
	UserID             uuid.UUID    `json:"userId"`
	AlertId            string       `json:"alert_id"`
	AlertName          string       `json:"alert_name"`
	CancelUrl          string       `json:"cancel_url"`
	CheckoutId         string       `json:"checkout_id"`
	Currency           string       `json:"currency"`
	CustomData         string       `json:"custom_data"`
	Email              string       `json:"email"`
	EventTime          string       `json:"event_time"`
	Status             string       `json:"status"`
	PSignature         string       `json:"p_signature"`
	PassThrough        string       `json:"pass_through"`
	SubscriptionId     string       `json:"subscription_id" gorm:"index"`
	SubscriptionPlanId string       `json:"subscription_plan_id"`
	PUserId            string       `json:"paddle_user_id"`
	ExtraProperty      pgtype.JSONB `json:"extraProperty" gorm:"type:jsonb;default:'{}';not null"` // ExtraProperty

}
