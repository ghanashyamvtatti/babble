package main

import (
	"context"
	"ds-project/common/proto/models"
	"ds-project/common/utilities"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/protobuf/ptypes"
	"time"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

type UsersDB struct {
	Users map[string]*models.User
}

type TokenDB struct {
	Tokens map[string]string
}

type SubscriptionDB struct {
	Subscriptions map[string][]string
}

type PostDB struct {
	Posts map[string][]*models.Post
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()

	cli.Delete(ctx, "users", clientv3.WithPrefix())
	cli.Delete(ctx, "tokens", clientv3.WithPrefix())

	users := &UsersDB{
		Users: map[string]*models.User{},
	}

	tokens := &TokenDB{
		Tokens: map[string]string{},
	}

	subscriptions := &SubscriptionDB{Subscriptions: map[string][]string{}}

	posts := &PostDB{Posts: map[string][]*models.Post{}}

	users.Users["ghanu"] = &models.User{
		FullName:  "Ghanashyam",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	users.Users["varun"] = &models.User{
		FullName:  "Varun",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	// Add User 3
	users.Users["pratik"] = &models.User{
		FullName:  "Pratik",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	marshalledUser, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	utilities.PutKey(ctx, cli, "users", marshalledUser)

	bt := utilities.GetKey(ctx, cli, "users")
	var c UsersDB
	er := json.Unmarshal(bt, &c)
	if er != nil {
		panic(er)
	}

	fmt.Println(c)

	tokens.Tokens["ghanu"] = "MASTER-TOKEN"
	tokens.Tokens["varun"] = "MASTER-TOKEN"
	tokens.Tokens["pratik"] = "MASTER-TOKEN"

	marshalledToken, err := json.Marshal(tokens)
	if err != nil {
		panic(err)
	}
	utilities.PutKey(ctx, cli, "tokens", marshalledToken)

	bt1 := utilities.GetKey(ctx, cli, "tokens")
	var c1 TokenDB
	er1 := json.Unmarshal(bt1, &c1)
	if er1 != nil {
		panic(er1)
	}

	fmt.Println(c1)

	// SubscriptionService
	subscriptions.Subscriptions["ghanu"] = append(subscriptions.Subscriptions["ghanu"], "ghanu")
	subscriptions.Subscriptions["ghanu"] = append(subscriptions.Subscriptions["ghanu"], "varun")

	subscriptions.Subscriptions["varun"] = append(subscriptions.Subscriptions["varun"], "varun")
	subscriptions.Subscriptions["varun"] = append(subscriptions.Subscriptions["varun"], "pratik")

	marshalledSubs, err := json.Marshal(subscriptions)
	if err != nil {
		panic(err)
	}

	utilities.PutKey(ctx, cli, "subscriptions", marshalledSubs)

	bt2 := utilities.GetKey(ctx, cli, "subscriptions")
	var c2 SubscriptionDB
	er2 := json.Unmarshal(bt2, &c2)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println(c2)

	posts.Posts["ghanu"] = append(posts.Posts["ghanu"], &models.Post{
		Post:      "Hello World! This is Ghanashyam.",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["ghanu"] = append(posts.Posts["ghanu"], &models.Post{
		Post:      "WOLOLO!",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["ghanu"] = append(posts.Posts["ghanu"], &models.Post{
		Post:      "Knock Knock. Anybody there?",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	posts.Posts["varun"] = append(posts.Posts["varun"], &models.Post{
		Post:      "My name is Varun.",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["varun"] = append(posts.Posts["varun"], &models.Post{
		Post:      "I hope this application works well",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["varun"] = append(posts.Posts["varun"], &models.Post{
		Post:      "Hey! I'm here!",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	posts.Posts["pratik"] = append(posts.Posts["pratik"], &models.Post{
		Post:      "Pratik is here!",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["pratik"] = append(posts.Posts["pratik"], &models.Post{
		Post:      "I wonder what time it is in Mars",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	posts.Posts["pratik"] = append(posts.Posts["pratik"], &models.Post{
		Post:      "lorem ipsum",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	marshalledPosts, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}

	utilities.PutKey(ctx, cli, "posts", marshalledPosts)

	bt3 := utilities.GetKey(ctx, cli, "posts")
	var c3 PostDB
	err = json.Unmarshal(bt3, &c3)
	if err != nil {
		panic(err)
	}
	fmt.Println(c3)
}
