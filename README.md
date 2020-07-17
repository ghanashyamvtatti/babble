# Babble

## Description
A reliable, fault-tolerant, pub-sub style social media powered by Go, etcd and ReactJS

## Project Structure

```
.
├── common                  # Contains utilities and protobufs common to the entire project
  ├── proto                 # Contains all the protocol buffers
  ├── utilities             # Contains methods and structures common to the entire project
  ├── requests.go           # A handy struct to hold commonly used params in DAL functions
├── config                  # Initial entries for the etcd data store
├── test                    # Test cases for the project
  ├── concurrency
  ├── context
  └── services 
|── APIGateway              # Main backend - communicates with other microservies and provides REST APIs
  ├── controllers             # Controllers for the various endpoints defined in web.go
  ├── dtos                    # Data Transfer Objects (shapes the data in the format expected by the UI)
  ├── common
    ├── clients.go            # A wrapper struct for holding all the gRPC client objects
  └── web.go                  # Contains the gin routes 
├── UI                      # ReactJS based UI
├── AuthService             # gRPC microservice for authorization
  ├── authdal               # etcd and in-memory implementations
  └── config                # Stores config for in-memory storage
├── PostService             # gRPC microservice for managing posts
  ├── postdal               # etcd and in-memory implementations
  └── config                # Stores config for in-memory storage
├── SubscriptionService     # gRPC microservice for managing subscriptions
  ├── subscriptiondal       # etcd and in-memory implementations
  └── config                # Stores config for in-memory storage
├── UserService             # gRPC microservice for user management
  ├── userdal               # etcd and in-memory implementations
  └── config                # Stores config for in-memory storage
├── go.mod
├── .gitignore
└── README.md
```

## Ports

| Server               | Port |
|----------------------|------|
| Auth Server          | 3004 |
| Posts Server         | 3003 |
| Users Server         | 3002 |
| Subscriptions Server | 3005 |
| UI                   | 3000 |
| API Gateway          | 8080 |
 
## Setup
(ensure you have yarn or npm installed)

1. Go to https://github.com/etcd-io/etcd/releases
    * (Linux and MacOS) Follow the instructions as mentioned in the website
    * (Windows) Download the windows zip file, extract it locally and add the folder to PATH
2. `cd` into your GOPATH
3. `git clone https://github.com/Distributed-Systems-CSGY9223/vo383-ppr239-gvt217-final-project ds-project`
4. `cd ds-project/UI/babble`
5. `npm install` or `yarn install`
6. `cd ../..`
7. `go build ds-project`
8. Running everything:
    * If you're using Windows, run `runme.bat`
    * If you're using MacOS/Linux, run `runme.sh` 
9. Open `localhost:8080` in the browser
