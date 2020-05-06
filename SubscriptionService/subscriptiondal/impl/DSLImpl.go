package impl

import (
	"ds-project/SubscriptionService/config"
	"ds-project/common"
	"ds-project/common/proto/subscriptions"
	"sync"
)

type DSLSubscriptionDAL struct {
	Mutex     sync.Mutex
	AppConfig *config.ApplicationConfig
}

func (dal *DSLSubscriptionDAL) GetSubscriptions(request common.DALRequest, username string, result chan *subscriptions.GetSubscriptionsResponse) {
	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	result <- &subscriptions.GetSubscriptionsResponse{Subscriptions: dal.AppConfig.Subscriptions[username]}
}

func (dal *DSLSubscriptionDAL) Subscribe(request common.DALRequest, subscriber string, publisher string, result chan bool) {
	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	// Ensures that there are no duplicate entries
	for _, pub := range dal.AppConfig.Subscriptions[subscriber] {
		if publisher == pub {
			result <- true
			return
		}
	}
	dal.AppConfig.Subscriptions[subscriber] = append(dal.AppConfig.Subscriptions[subscriber], publisher)

	result <- true
}

func (dal *DSLSubscriptionDAL) Unsubscribe(request common.DALRequest, subscriber string, publisher string, result chan bool) {
	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()
	for index, pub := range dal.AppConfig.Subscriptions[subscriber] {
		if publisher == pub {
			dal.AppConfig.Subscriptions[subscriber] = append(dal.AppConfig.Subscriptions[subscriber][:index], dal.AppConfig.Subscriptions[subscriber][index+1:]...)
		}
	}
	result <- true
}
