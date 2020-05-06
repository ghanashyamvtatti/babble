package config

type ApplicationConfig struct {
	Tokens map[string]string
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Tokens: map[string]string{},
	}
	// Add universal token for users
	appConfig.Tokens["ghanu"] = "MASTER-TOKEN"
	appConfig.Tokens["varun"] = "MASTER-TOKEN"
	appConfig.Tokens["pratik"] = "MASTER-TOKEN"

	return appConfig
}
