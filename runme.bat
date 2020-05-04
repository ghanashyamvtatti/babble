start runetcd.bat
start go run AuthService/authserver.go
start go run PostService/postsserver.go
start go run UserService/userserver.go
start go run SubscriptionService/subscriptionsserver.go
start npm start --prefix UI/babble
go run APIGateway/web.go