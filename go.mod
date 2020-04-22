module ds-project

go 1.13

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0 // indirect

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.1
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.0-rc.4
	github.com/google/uuid v1.1.1
	github.com/mattn/goreman v0.3.5 // indirect
	github.com/stretchr/testify v1.4.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200420175359-c1e7f73a0232 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200406155108-e3b113bbe6a4 // indirect
	google.golang.org/genproto v0.0.0-20200406120821-33397c535dc2 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.20.1
)
