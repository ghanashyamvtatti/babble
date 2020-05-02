package subscriptiondal

import (
	"ds-project/common"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/utilities"
	"encoding/json"
	"sync"
)

type SubscriptionDB struct {
	Subscriptions map[string][]string
}

var (
	mutex sync.Mutex
)

func GetSubscriptions(request common.DALRequest, username string, result chan *subscriptions.GetSubscriptionsResponse) {
	mutex.Lock()
	defer mutex.Unlock()

	subscriptionDB := getSubscriptionCollection(request)

	result <- &subscriptions.GetSubscriptionsResponse{Subscriptions: subscriptionDB.Subscriptions[username]}
}

func Subscribe(request common.DALRequest, subscriber string, publisher string, result chan bool) {
	mutex.Lock()
	defer mutex.Unlock()

	subscriptionDB := getSubscriptionCollection(request)
	subscriptionDB.Subscriptions[subscriber] = append(subscriptionDB.Subscriptions[subscriber], publisher)
	updateSubscriptions(subscriptionDB, request, result)
}

func Unsubscribe(request common.DALRequest, subscriber string, publisher string, result chan bool) {
	mutex.Lock()
	defer mutex.Unlock()
	subscriptionDB := getSubscriptionCollection(request)
	for index, pub := range subscriptionDB.Subscriptions[subscriber] {
		if publisher == pub {
			subscriptionDB.Subscriptions[subscriber] = append(subscriptionDB.Subscriptions[subscriber][:index], subscriptionDB.Subscriptions[subscriber][index+1:]...)
		}
	}
	updateSubscriptions(subscriptionDB, request, result)
}

func getSubscriptionCollection(request common.DALRequest) *SubscriptionDB {
	bytes := utilities.GetKey(request.Ctx, request.Client, "subscriptions")
	var subscriptionDB SubscriptionDB
	if err := json.Unmarshal(bytes, &subscriptionDB); err != nil {
		request.ErrorChan <- err
	}
	return &subscriptionDB
}

func updateSubscriptions(storage *SubscriptionDB, request common.DALRequest, result chan bool) {

	subscriptionsBytes, err := json.Marshal(storage)
	if err != nil {
		request.ErrorChan <- err
		result <- false
	}

	utilities.PutKey(request.Ctx, request.Client, "subscriptions", subscriptionsBytes)
	result <- true
}
