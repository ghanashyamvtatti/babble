package config

import (
	"ds-project/common/proto/auth"
	"ds-project/common/proto/posts"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/proto/users"
)

type ServiceClients struct {
	UserClient         users.UserServiceClient
	PostClient         posts.PostsServiceClient
	AuthClient         auth.AuthServiceClient
	SubscriptionClient subscriptions.SubscriptionServiceClient
}
