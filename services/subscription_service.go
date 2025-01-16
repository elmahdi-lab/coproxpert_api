package services

import (
	"fmt"

	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateSubscription(user *models.User, subscriptionType models.SubscriptionType) (*models.Subscription, error) {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	subscription, err := subscriptionRepository.FindByUser(user)

	if err != nil {
		return nil, fmt.Errorf("error while fetching subscription: %w", err)
	}

	if subscription != nil {
		return nil, fmt.Errorf("user already has a subscription")
	}

	subscription = &models.Subscription{}
	subscription.CreateTrialSubscription(user, subscriptionType)

	err = subscriptionRepository.Create(subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func FindSubscriptionByUser(user *models.User) *models.Subscription {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	subscription, err := subscriptionRepository.FindByUser(user)

	if err != nil {
		return nil
	}

	return subscription
}
