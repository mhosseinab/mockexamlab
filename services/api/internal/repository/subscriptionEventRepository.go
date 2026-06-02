package repository

import (
	"MockExamLab/internal/models"
	"gorm.io/gorm"
)

type SubscriptionEventRepositoryRepository struct {
	db *gorm.DB
}

func ProvideSubscriptionEventRepositoryRepository(db *gorm.DB) SubscriptionEventRepositoryRepository {
	return SubscriptionEventRepositoryRepository{db: db}
}

func (r *SubscriptionEventRepositoryRepository) Create(event models.SubscriptionEvent) (*models.SubscriptionEvent, error) {
	if err := r.db.Create(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
