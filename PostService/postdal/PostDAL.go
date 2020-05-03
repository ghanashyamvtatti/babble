package postdal

import (
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	"ds-project/common/utilities"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"sort"
	"sync"
)

type PostDB struct {
	Posts map[string][]*models.Post
}

type PostDAL struct {
	Mutex sync.Mutex
}

func (postDAL *PostDAL) AddPost(request common.DALRequest, username string, post string, result chan bool) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

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

func (postDAL *PostDAL) DeletePost(request common.DALRequest, username string, delPost string, result chan bool) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	postDB := getPostCollection(request)
	for index, post := range postDB.Posts[username] {
		if post.Post == delPost {
			postDB.Posts[username] = append(postDB.Posts[username][:index], postDB.Posts[username][index+1:]...)
		}
	}
	updatePosts(postDB, request, result)
}

func (postDAL *PostDAL) GetPosts(request common.DALRequest, username string, result chan *posts.GetPostsResponse) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	postDB := getPostCollection(request)

	result <- &posts.GetPostsResponse{Posts: postDB.Posts[username]}
}

func (postDAL *PostDAL) GetFeed(request common.DALRequest, subscriptions []string, result chan []*models.Post) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

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

func getPostCollection(request common.DALRequest) *PostDB {
	bytes := utilities.GetKey(request.Ctx, request.Client, "posts")
	var postDB PostDB
	if err := json.Unmarshal(bytes, &postDB); err != nil {
		request.ErrorChan <- err
	}
	return &postDB
}

func updatePosts(postDB *PostDB, request common.DALRequest, result chan bool) {

	bytes, err := json.Marshal(postDB)
	if err != nil {
		request.ErrorChan <- err
		result <- false
	}

	utilities.PutKey(request.Ctx, request.Client, "posts", bytes)
	result <- true
}
