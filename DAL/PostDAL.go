package DAL

import (
	"ds-project/config"
	"ds-project/models"
	"time"
)

func AddPost(appConfig *config.ApplicationConfig, username string, post string) bool {
	newPost := &models.Post{
		Post:      post,
		Username:  username,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	appConfig.Posts[username] = append(appConfig.Posts[username], newPost)

	return true
}

func GetPosts(appConfig *config.ApplicationConfig, username string) []*models.Post {
	return appConfig.Posts[username]
}
