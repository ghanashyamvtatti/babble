package services

import (
	"ds-project/config"
	// "sync"
)
// var (
// 	mutex sync.Mutex
// )
func Subscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {
	
	mutex.Lock()
	appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber], publisher)
	mutex.Unlock()
}

func Unsubscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {
	
	mutex.Lock()
	for index, pub := range appConfig.Subscriptions[subscriber] {
		if publisher == pub {
			appConfig.Subscriptions[subscriber] = append(appConfig.Subscriptions[subscriber][:index], appConfig.Subscriptions[subscriber][index+1:]...)
		}
	}
	mutex.Unlock()
}

func GetSubscriptionsForUsername(appConfig *config.ApplicationConfig, username string) []string {
	// TODO: Pratik needs to implement this
	//return nil
	return appConfig.Subscriptions[username]
}