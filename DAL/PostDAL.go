package DAL

import (
	"ds-project/common/proto/models"
	"ds-project/config"
	"github.com/golang/protobuf/ptypes"
)

func AddPost(appConfig *config.ApplicationConfig, username string, post string) bool {
	newPost := &models.Post{
		Post:      post,
		Username:  username,
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}
	appConfig.Posts[username] = append(appConfig.Posts[username], newPost)

	return true
}

func GetPosts(appConfig *config.ApplicationConfig, username string) []*models.Post {
	return appConfig.Posts[username]
}
