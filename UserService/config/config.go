package config

import (
	"ds-project/common/proto/models"
	"github.com/golang/protobuf/ptypes"
)

type ApplicationConfig struct {
	Users map[string]*models.User
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Users: map[string]*models.User{},
	}

	// Add User 1
	appConfig.Users["ghanu"] = &models.User{
		FullName:  "Ghanashyam",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	// Add User 2
	appConfig.Users["varun"] = &models.User{
		FullName:  "Varun",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	// Add User 3
	appConfig.Users["pratik"] = &models.User{
		FullName:  "Pratik",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return appConfig
}
