package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
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

func (sr *SubscriptionRepository) Create(Subscription *models.Subscription) error {
	return sr.db.Create(Subscription).Error
}

func (sr *SubscriptionRepository) Update(Subscription *models.Subscription) error {
	return sr.db.Save(Subscription).Error
}

func (sr *SubscriptionRepository) Delete(Subscription *models.Subscription) error {
	return sr.db.Delete(Subscription).Error
}

func (sr *SubscriptionRepository) DeleteByID(id uuid.UUID) bool {
	err := sr.db.Delete(&models.Subscription{}, id)
	if err != nil {
		return true
	}
	return false
}

func (sr *SubscriptionRepository) FindByUser(user *models.User) ([]models.Subscription, error) {
	var subscription []models.Subscription
	if err := sr.db.Where("user_id = ?", user.ID).Find(&subscription).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

func (sr *SubscriptionRepository) FindByID(id uuid.UUID) (*models.Subscription, error) {
	var Subscription models.Subscription
	if err := sr.db.First(&Subscription, id).Error; err != nil {
		return nil, err
	}
	return &Subscription, nil
}

func (sr *SubscriptionRepository) FindAll() ([]models.Subscription, error) {
	var Subscriptions []models.Subscription
	if err := sr.db.Find(&Subscriptions).Error; err != nil {
		return nil, err
	}
	return Subscriptions, nil
}
