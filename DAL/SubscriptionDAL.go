package DAL

import "ds-project/config"

func GetSubscriptions(appConfig *config.ApplicationConfig, username string) []string {
	return appConfig.Subscriptions[username]
}

func Subscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) bool {
	appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber], publisher)
	return true
}

func Unsubscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) bool {
	for index, pub := range appConfig.Subscriptions[subscriber] {
		if publisher == pub {
			appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber][:index], appConfig.Subscriptions[subscriber][index+1:]...)
		}
	}
	return true
}