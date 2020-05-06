package postdal

import (
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
)

type PostDAL interface {
	AddPost(request common.DALRequest, username string, post string, result chan bool)
	DeletePost(request common.DALRequest, username string, delPost string, result chan bool)
	GetPosts(request common.DALRequest, username string, result chan *posts.GetPostsResponse)
	GetFeed(request common.DALRequest, subscriptions []string, result chan []*models.Post)
}
