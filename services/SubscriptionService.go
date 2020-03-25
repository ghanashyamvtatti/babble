package services

import (
	"pubsubhub/config"
	"pubsubhub/models"
	"time"
)

func Subscribe(appConfig *config.ApplicationConfig, subscriberId int, publisherId int) {
	newSubscription := &models.Subscription{
		Id:         len(appConfig.Subscriptions),
		Subscriber: GetUserById(appConfig, subscriberId),
		Publisher:  GetUserById(appConfig, publisherId),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	appConfig.Subscriptions = append(appConfig.Subscriptions, newSubscription)
}

func Unsubscribe(appConfig *config.ApplicationConfig, subscriberId int, publisherId int) {
	for i, subscription := range appConfig.Subscriptions {
		if subscription.Subscriber.Id == subscriberId && subscription.Publisher.Id == publisherId {
			appConfig.Subscriptions = append(appConfig.Subscriptions[:i], appConfig.Subscriptions[i+1:]...)
			return
		}
	}
}

func GetSubscriptionsForUserId(appConfig *config.ApplicationConfig, userId int) []*models.Subscription {
	var subscriptions []*models.Subscription
	for _, subscription := range appConfig.Subscriptions {
		if subscription.Subscriber.Id == userId {
			subscriptions = append(subscriptions, subscription)
		}
	}
	return subscriptions
}
