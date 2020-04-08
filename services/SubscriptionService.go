package services

import (
	"ds-project/DAL"
	"ds-project/config"
	// "sync"
)

// var (
// 	mutex sync.Mutex
// )
func Subscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {

	mutex.Lock()
	DAL.Subscribe(appConfig, subscriber, publisher)
	mutex.Unlock()
}

func Unsubscribe(appConfig *config.ApplicationConfig, subscriber string, publisher string) {

	mutex.Lock()
	DAL.Unsubscribe(appConfig, subscriber, publisher)
	mutex.Unlock()
}

func GetSubscriptionsForUsername(appConfig *config.ApplicationConfig, username string) []string {
	return DAL.GetSubscriptions(appConfig, username)
}
