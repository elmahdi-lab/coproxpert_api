package services

import (
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateSubscription(s *models.Subscription) error {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	err := subscriptionRepository.Create(s)
	if err != nil {
		return err
	}
	return nil
}
