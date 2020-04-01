package services

import (
	"ds-project/config"
)

func Subscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {
	// TODO: Pratik needs to implement this
	appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber], publisher)
}

func Unsubscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {
	// TODO: Pratik needs to implement this
	for index, pub := range appConfig.Subscriptions[subscriber] {
		if publisher == pub {
			appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber][:index], appConfig.Subscriptions[subscriber][index+1:]...)
		}
	}
}

func GetSubscriptionsForUsername(appConfig *config.ApplicationConfig, username string) []string {
	// TODO: Pratik needs to implement this
	//return nil
	return appConfig.Subscriptions[username]
}
