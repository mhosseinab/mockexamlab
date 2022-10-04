package repository

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/consts"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where(models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByFirebaseId(firebaseId string) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where(models.User{FirebaseId: firebaseId}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUserSubscription(id uuid.UUID) (consts.SubscriptionState, int64, error) {
	var u models.User
	if err := r.db.Model(&models.User{}).
		Where(models.User{BaseModel: models.BaseModel{ID: id}}).
		First(&u).
		Error; err != nil {
		return consts.NonSubscriber, 0, err
	}
	return u.SubscriptionState, u.SubscriptionExpirationMs, nil
}

func (r *UserRepository) CheckUserByEmail(email string) (bool, error) {
	var count int64 = 0

	if err := r.db.Model(&models.User{}).
		Where(models.User{Email: email}).
		Count(&count).
		Error; err != nil {
		return true, err
	}
	return count > 0, nil
}

func (r *UserRepository) UpdateLastLogin(user *models.User) error {
	currentTime := time.Now().Unix()
	//return r.db.Model(&user).Update("last_login", currentTime).Error
	return r.db.Model(&user).Updates(models.User{LastLogin: currentTime}).Error
}

func (r *UserRepository) UpdateSubscription(user *models.User) error {
	return r.db.Model(&user).Updates(user).Error
}

func (t *UserRepository) FindAllFilterPagination(params *models.QueryParams) (*[]models.User, int64, error) {
	var users []models.User
	var count int64 = 0
	if err := t.db.Model(&models.User{}).Scopes(FilterByQuery(params, consts.ALL)).Find(&users).Count(&count).Error; err != nil {
		return nil, count, err
	}
	return &users, count, nil
}
