package subscriptiondal

import (
	"ds-project/common"
	"ds-project/common/proto/subscriptions"
)

type SubscriptionDAL interface {
	GetSubscriptions(request common.DALRequest, username string, result chan *subscriptions.GetSubscriptionsResponse)
	Subscribe(request common.DALRequest, subscriber string, publisher string, result chan bool)
	Unsubscribe(request common.DALRequest, subscriber string, publisher string, result chan bool)
}
