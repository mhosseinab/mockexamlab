package service

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"MockExamLab/internal/repository"
	"time"
)

type PaddleService struct {
	r  repository.SubscriptionEventRepositoryRepository
	ru repository.UserRepository
}

func ProvidePaddleService(r repository.SubscriptionEventRepositoryRepository, ru repository.UserRepository) PaddleService {
	return PaddleService{
		r:  r,
		ru: ru,
	}
}

func (s *PaddleService) SubscriptionCreated(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionCreate) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	if errUpdateUserSub := s.updateStatus(*u, dto.NextBillDate, consts.Subscriber); errUpdateUserSub != nil {
		return errUpdateUserSub
	}
	return nil
}

func (s *PaddleService) SubscriptionUpdate(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionUpdate) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	status := consts.Expired
	if consts.PaddleStatus(dtoBase.Status) == consts.Active {
		status = consts.Expired
	}
	if consts.PaddleStatus(dtoBase.Status) == consts.PaddlePaused {
		status = consts.Paused
	}
	if errUpdateUserSub := s.updateStatus(*u, dto.NextBillDate, status); errUpdateUserSub != nil {
		return errUpdateUserSub
	}
	return nil
}

func (s *PaddleService) SubscriptionCancelled(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionCanceled) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	if errUpdateUserSub := s.updateStatus(*u, time.UnixMilli(u.SubscriptionExpirationMs).Format("2006-01-02"), consts.Canceled); errUpdateUserSub != nil {
		return errUpdateUserSub
	}
	return nil
}

func (s *PaddleService) SubscriptionPaymentSucceeded(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionPaymentSucceeded) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	if errUpdateUserSub := s.updateStatus(*u, dto.NextBillDate, consts.Subscriber); errUpdateUserSub != nil {
		return errUpdateUserSub
	}
	return nil
}

func (s *PaddleService) SubscriptionPaymentFailed(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionPaymentFailed) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	return nil
}

func (s *PaddleService) SubscriptionRefunded(dtoBase *DTO.BasePaddleResponse, dto *DTO.SubscriptionPaymentRefunded) error {
	u, errFetchUser := s.ru.GetUserByEmail(dtoBase.Email)
	if errFetchUser != nil {
		return errFetchUser
	}
	e, errMapper := mapper.PaddleDtoToModel(u.ID, dtoBase, dto)
	if errFetchUser != nil {
		return errMapper
	}
	_, errSaveEvent := s.r.Create(*e)
	if errSaveEvent != nil {
		return errSaveEvent
	}
	if errUpdateUserSub := s.updateStatus(*u, time.Now().Format("2006-01-02"), consts.Refund); errUpdateUserSub != nil {
		return errUpdateUserSub
	}
	return nil
}

func (s *PaddleService) updateStatus(u models.User, nextBillingDate string, status consts.SubscriptionState) error {
	date, errTimeConvert := time.Parse("2006-01-02", nextBillingDate)
	if errTimeConvert != nil {
		return errTimeConvert
	}
	u.SubscriptionExpirationMs = date.Unix()
	u.SubscriptionState = status
	if err := s.ru.UpdateSubscription(&u); err != nil {
		return err
	}
	return nil
}
