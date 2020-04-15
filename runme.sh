#!/bin/sh
nohup go run servers/dataserver/dataserver.go &
nohup go run servers/authserver/authserver.go &
nohup go run servers/postsserver/postsserver.go &
nohup go run servers/userserver/userserver.go &
nohup go run servers/subscriptionsserver/subscriptionsserver.go &
nohup npm start --prefix UI/babble &
go run web.go