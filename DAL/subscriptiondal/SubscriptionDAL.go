package subscriptiondal

import (
	"context"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/utilities"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

type SubscriptionStorage struct {
	subscriptions map[string][]string
}

var (
	mutex sync.Mutex
)

func GetSubscriptions(ctx context.Context, client *clientv3.Client, username string, result chan *subscriptions.GetSubscriptionsResponse, errorChan chan error) {
	mutex.Lock()
	defer mutex.Unlock()

	subscriptionStorage := getSubscriptionCollection(ctx, client, errorChan)

	result <- &subscriptions.GetSubscriptionsResponse{Subscriptions: subscriptionStorage.subscriptions[username]}
}

func Subscribe(ctx context.Context, client *clientv3.Client, subscriber string, publisher string, result chan bool, errorChan chan error) {
	mutex.Lock()
	defer mutex.Unlock()

	subscriptionStorage := getSubscriptionCollection(ctx, client, errorChan)
	subscriptionStorage.subscriptions[publisher] = append(subscriptionStorage.subscriptions[publisher], subscriber)
	updateSubscriptions(subscriptionStorage, ctx, client, result, errorChan)
}

func Unsubscribe(ctx context.Context, client *clientv3.Client, subscriber string, publisher string, result chan bool, errorChan chan error) {
	mutex.Lock()
	defer mutex.Unlock()
	subscriptionStorage := getSubscriptionCollection(ctx, client, errorChan)
	for index, pub := range subscriptionStorage.subscriptions[subscriber] {
		if publisher == pub {
			subscriptionStorage.subscriptions[subscriber] = append(subscriptionStorage.subscriptions[subscriber][:index], subscriptionStorage.subscriptions[subscriber][index+1:]...)
		}
	}
	updateSubscriptions(subscriptionStorage, ctx, client, result, errorChan)

}

func getSubscriptionCollection(ctx context.Context, client *clientv3.Client, errorChan chan error) *SubscriptionStorage {
	bytes := utilities.GetKey(ctx, client, "subscriptions")
	var subscriptionStorage SubscriptionStorage
	if err := json.Unmarshal(bytes, &subscriptionStorage); err != nil {
		errorChan <- err
	}
	return &subscriptionStorage
}

func updateSubscriptions(storage *SubscriptionStorage, ctx context.Context, client *clientv3.Client, result chan bool, errorChan chan error) {

	subscriptionsBytes, err := json.Marshal(storage)
	if err != nil {
		errorChan <- err
		result <- false
	}

	utilities.PutKey(ctx, client, "subscriptions", subscriptionsBytes)
	result <- true
}
