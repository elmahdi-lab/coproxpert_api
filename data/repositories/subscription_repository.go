package repositories

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	models2 "ithumans.com/coproxpert/data/models"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository() *SubscriptionRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &SubscriptionRepository{db: db}
}

func (sr *SubscriptionRepository) Create(Subscription *models2.Subscription) error {
	return sr.db.Create(Subscription).Error
}

func (sr *SubscriptionRepository) Update(Subscription *models2.Subscription) error {
	return sr.db.Save(Subscription).Error
}

func (sr *SubscriptionRepository) Delete(Subscription *models2.Subscription) error {
	return sr.db.Delete(Subscription).Error
}

func (sr *SubscriptionRepository) DeleteByID(id uuid.UUID) bool {
	err := sr.db.Delete(&models2.Subscription{}, id)
	if err != nil {
		return true
	}
	return false
}

func (sr *SubscriptionRepository) FindByUserID(userID uuid.UUID) (*models2.Subscription, error) {
	var subscription *models2.Subscription
	if err := sr.db.Where("user_id = ?", userID).First(&subscription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return subscription, nil
}

func (sr *SubscriptionRepository) FindByUser(user *models2.User) (*models2.Subscription, error) {
	return sr.FindByUserID(user.ID)
}

func (sr *SubscriptionRepository) FindByID(id uuid.UUID) (*models2.Subscription, error) {
	var subscription *models2.Subscription
	if err := sr.db.First(&subscription, id).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

func (sr *SubscriptionRepository) FindAll() ([]models2.Subscription, error) {
	var subscriptions []models2.Subscription
	if err := sr.db.Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}
