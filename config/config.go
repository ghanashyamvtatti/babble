package config

import (
	"ds-project/models"
	"time"
)

type ApplicationConfig struct {
	Users         map[string]models.User
	Tokens        map[string]string
	Posts         map[string][]*models.Post
	Subscriptions map[string][]string
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Users:         map[string]models.User{},
		Tokens:        map[string]string{},
		Posts:         map[string][]*models.Post{},
		Subscriptions: map[string][]string{},
	}

	// Add User 1
	appConfig.Users["ghanu"] = models.User{
		FullName:  "Ghanashyam",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	// Add User 2
	appConfig.Users["varun"] = models.User{
		FullName:  "Varun",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	// Add User 3
	appConfig.Users["pratik"] = models.User{
		FullName:  "Pratik",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	// Add universal token for users
	appConfig.Tokens["ghanu"] = "MASTER-TOKEN"
	appConfig.Tokens["varun"] = "MASTER-TOKEN"
	appConfig.Tokens["pratik"] = "MASTER-TOKEN"

	// Add default posts for each user
	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "Hello World! This is Ghanashyam.",
		Username:  "ghanu",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "WOLOLO!",
		Username:  "ghanu",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["ghanu"] = append(appConfig.Posts["ghanu"], &models.Post{
		Post:      "Knock Knock. Anybody there?",
		Username:  "ghanu",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})

	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "My name is Varun.",
		Username:  "varun",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "I hope this application works well",
		Username:  "varun",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["varun"] = append(appConfig.Posts["varun"], &models.Post{
		Post:      "Hey! I'm here!",
		Username:  "varun",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})

	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "Pratik is here!",
		Username:  "pratik",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "I wonder what time it is in Mars",
		Username:  "pratik",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})
	appConfig.Posts["pratik"] = append(appConfig.Posts["pratik"], &models.Post{
		Post:      "lorem ipsum",
		Username:  "pratik",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	})

	// Subscriptions
	appConfig.Subscriptions["ghanu"] = append(appConfig.Subscriptions["ghanu"], "ghanu")
	appConfig.Subscriptions["ghanu"] = append(appConfig.Subscriptions["ghanu"], "varun")

	appConfig.Subscriptions["varun"] = append(appConfig.Subscriptions["varun"], "varun")
	appConfig.Subscriptions["varun"] = append(appConfig.Subscriptions["varun"], "pratik")

	return appConfig
}
