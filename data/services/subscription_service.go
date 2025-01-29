package services

import (
	"fmt"

	models2 "ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/repositories"
)

func CreateSubscription(user *models2.User, subscriptionType models2.SubscriptionType) (*models2.Subscription, error) {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	subscription, err := subscriptionRepository.FindByUser(user)

	if err != nil {
		return nil, fmt.Errorf("error while fetching subscription: %w", err)
	}

	if subscription != nil {
		return nil, fmt.Errorf("user already has a subscription")
	}

	subscription = &models2.Subscription{}
	subscription.CreateTrialSubscription(user, subscriptionType)

	err = subscriptionRepository.Create(subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func FindSubscriptionByUser(user *models2.User) *models2.Subscription {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	subscription, err := subscriptionRepository.FindByUser(user)

	if err != nil {
		return nil
	}

	return subscription
}
