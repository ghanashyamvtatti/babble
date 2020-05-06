package config

import (
	"ds-project/common/proto/models"
	"github.com/golang/protobuf/ptypes"
)

type ApplicationConfig struct {
	Posts map[string][]*models.Post
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Posts: map[string][]*models.Post{},
	}

	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "Hello World! This is Ghanashyam.",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "WOLOLO!",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "Knock Knock. Anybody there?",
		Username:  "ghanu",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "My name is Varun.",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "I hope this application works well",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "Hey! I'm here!",
		Username:  "varun",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "Pratik is here!",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "I wonder what time it is in Mars",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "lorem ipsum",
		Username:  "pratik",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	})

	return appConfig
}
