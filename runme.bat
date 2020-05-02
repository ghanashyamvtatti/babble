start runetcd.bat
start go run servers/authserver/authserver.go
start go run servers/postsserver/postsserver.go
start go run servers/userserver/userserver.go
start go run servers/subscriptionsserver/subscriptionsserver.go
start npm start --prefix UI/babble
go run web.go