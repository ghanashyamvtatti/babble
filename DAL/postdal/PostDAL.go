package postdal

import (
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	"ds-project/common/utilities"
	"ds-project/config"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"sort"
	"sync"
)

type PostDB struct {
	Posts map[string][]*models.Post
}

var (
	mutex sync.Mutex
)

func AddPost(request config.DALRequest, username string, post string, result chan bool) {
	mutex.Lock()
	defer mutex.Unlock()

	newPost := &models.Post{
		Post:      post,
		Username:  username,
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}
	postDB := getPostCollection(request)
	postDB.Posts[username] = append(postDB.Posts[username], newPost)

	updatePosts(postDB, request, result)
}

func GetPosts(request config.DALRequest, username string, result chan *posts.GetPostsResponse) {
	mutex.Lock()
	defer mutex.Unlock()

	postDB := getPostCollection(request)

	result <- &posts.GetPostsResponse{Posts: postDB.Posts[username]}
}

func GetFeed(request config.DALRequest, subscriptions []string, result chan []*models.Post) {
	mutex.Lock()
	defer mutex.Unlock()

	postDB := getPostCollection(request)
	var responsePosts []*models.Post

	for _, subscription := range subscriptions {
		userPosts := postDB.Posts[subscription]
		responsePosts = append(responsePosts, userPosts...)
	}
	sort.Slice(responsePosts, func(i, j int) bool {
		iTime, _ := ptypes.Timestamp(responsePosts[i].CreatedAt)
		jTime, _ := ptypes.Timestamp(responsePosts[j].CreatedAt)
		return iTime.After(jTime)
	})

	result <- responsePosts
}

func getPostCollection(request config.DALRequest) *PostDB {
	bytes := utilities.GetKey(request.Ctx, request.Client, "posts")
	var postDB PostDB
	if err := json.Unmarshal(bytes, &postDB); err != nil {
		request.ErrorChan <- err
	}
	return &postDB
}

func updatePosts(postDB *PostDB, request config.DALRequest, result chan bool) {

	bytes, err := json.Marshal(postDB)
	if err != nil {
		request.ErrorChan <- err
		result <- false
	}

	utilities.PutKey(request.Ctx, request.Client, "posts", bytes)
	result <- true
}
