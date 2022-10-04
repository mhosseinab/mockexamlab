package service

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models"
	"errors"
	"gorm.io/gorm"

	//"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/repository"
)

type UserService struct {
	r  repository.UserRepository
	fr repository.FirebaseRepository
}

func ProvideUserService(r repository.UserRepository, fr repository.FirebaseRepository) UserService {
	return UserService{
		r:  r,
		fr: fr,
	}
}

/*func (s *UserService) SignUpUser(user models.User) (*models.User, error) {
	if isExistUser, err := s.r.CheckUserByEmail(user.Email); err != nil {
		return nil, err
	} else {
		if isExistUser {
			return nil, errors.New("this email currently exist")
		}
	}
	if m, err := s.r.CreateUser(user); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (s *UserService) LoginUser(req DTO.UserLoginRequest) (*models.User, error) {
	m, err := s.r.GetUserByEmail(strings.TrimSpace(strings.ToLower(req.Email)))
	if err != nil {
		return nil, err
	}
	if passErr := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(req.Password)); passErr != nil {
		return nil, errors.New("error_msg\": \"invalid password data")
	}
	if updateLoginErr := s.r.UpdateLastLogin(m); updateLoginErr != nil {
		return nil, updateLoginErr
	}
	return m, nil
}*/

func (s *UserService) LoginOrSignUpUser(u models.UserClaims) (*models.User, error) {
	m, err := s.r.GetUserByFirebaseId(u.FirebaseId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if m == nil {
		m, err = s.r.CreateUser(mapper.UserLoginOrSignUpRequestToModel(&u))
		if err != nil {
			return nil, err
		}
	}
	if u.UserId != m.ID {
		claims := map[string]interface{}{
			"role":        u.Role,
			"app_user_id": m.ID,
		}
		err = s.fr.UpdateClaims(u.FirebaseId, claims)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (s *UserService) FindAllFilterPagination(params *models.QueryParams) (*[]models.User, int64, error) {
	users, count, err := s.r.FindAllFilterPagination(params)
	if err != nil {
		return nil, count, err
	}
	return users, count, nil
}
