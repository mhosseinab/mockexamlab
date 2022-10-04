package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

func PaddleDtoToModel(userId uuid.UUID, dto *DTO.BasePaddleResponse, extra interface{}) (*models.SubscriptionEvent, error) {
	var extraProperties pgtype.JSONB
	var jsonData string
	if data, err := json.Marshal(extra); err != nil {
		return nil, err
	} else {
		jsonData = string(data)
	}
	if err := extraProperties.Set(jsonData); err != nil {
		return nil, err
	}
	return &models.SubscriptionEvent{
		UserID:             userId,
		Email:              dto.Email,
		AlertId:            dto.AlertId,
		AlertName:          dto.AlertName,
		Status:             dto.Status,
		CheckoutId:         dto.CheckoutId,
		CancelUrl:          dto.CancelUrl,
		Currency:           dto.Currency,
		CustomData:         dto.CustomData,
		EventTime:          dto.EventTime,
		PassThrough:        dto.Passthrough,
		PSignature:         dto.PSignature,
		PUserId:            dto.UserId,
		SubscriptionId:     dto.SubscriptionId,
		SubscriptionPlanId: dto.SubscriptionPlanId,
		ExtraProperty:      extraProperties,
	}, nil
}
