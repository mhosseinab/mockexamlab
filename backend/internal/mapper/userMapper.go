package mapper

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"time"
)

/*func UserSignUpRequestToModel(dto *DTO.UserSignUpRequest) models.User {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)
	return models.User{
		Email:     dto.Email,
		Name:      dto.Name,
		Family:    dto.Family,
		Password:  string(hashPassword),
		LastLogin: time.Now().Unix(),
	}
}*/

func UserLoginOrSignUpRequestToModel(dto *models.UserClaims) models.User {
	return models.User{
		FirebaseId:  dto.FirebaseId,
		Email:       dto.Email,
		DisplayName: dto.DisplayName,
		Role:        dto.Role,
		LastLogin:   time.Now().Unix(),
	}
}

func UserModelToSignupOrLoginResponse(m *models.User) DTO.UserSignUpOrLoginUserResponse {
	return DTO.UserSignUpOrLoginUserResponse{
		UserId:                   m.ID,
		Email:                    m.Email,
		Alias:                    m.DisplayName,
		SubscriptionState:        m.SubscriptionState.ToString(),
		SubscriptionExpirationMs: m.SubscriptionExpirationMs,
	}
}

func UserModelToArrayResponse(models *[]models.User, count int64) DTO.PaginateResponse {
	DTOs := make([]DTO.UserSignUpOrLoginUserResponse, len(*models))
	for i, item := range *models {
		DTOs[i] = UserModelToSignupOrLoginResponse(&item)
	}
	return DTO.PaginateResponse{
		Data:  DTOs,
		Count: count,
	}
}
