package impl

import (
	"ds-project/PostService/config"
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	"github.com/golang/protobuf/ptypes"
	"sort"
	"sync"
)

type DSLPostDAL struct {
	Mutex     sync.Mutex
	AppConfig *config.ApplicationConfig
}

func (postDAL *DSLPostDAL) AddPost(request common.DALRequest, username string, post string, result chan bool) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	newPost := &models.Post{
		Post:      post,
		Username:  username,
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}
	postDAL.AppConfig.Posts[username] = append(postDAL.AppConfig.Posts[username], newPost)
	result <- true
}

func (postDAL *DSLPostDAL) DeletePost(request common.DALRequest, username string, delPost string, result chan bool) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	for index, post := range postDAL.AppConfig.Posts[username] {
		if post.Post == delPost {
			postDAL.AppConfig.Posts[username] = append(postDAL.AppConfig.Posts[username][:index], postDAL.AppConfig.Posts[username][index+1:]...)
		}
	}
	result <- true
}

func (postDAL *DSLPostDAL) GetPosts(request common.DALRequest, username string, result chan *posts.GetPostsResponse) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	result <- &posts.GetPostsResponse{Posts: postDAL.AppConfig.Posts[username]}
}

func (postDAL *DSLPostDAL) GetFeed(request common.DALRequest, subscriptions []string, result chan []*models.Post) {
	postDAL.Mutex.Lock()
	defer postDAL.Mutex.Unlock()

	var responsePosts []*models.Post

	for _, subscription := range subscriptions {
		userPosts := postDAL.AppConfig.Posts[subscription]
		responsePosts = append(responsePosts, userPosts...)
	}
	sort.Slice(responsePosts, func(i, j int) bool {
		iTime, _ := ptypes.Timestamp(responsePosts[i].CreatedAt)
		jTime, _ := ptypes.Timestamp(responsePosts[j].CreatedAt)
		return iTime.After(jTime)
	})

	result <- responsePosts
}
