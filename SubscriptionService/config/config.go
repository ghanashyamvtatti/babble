package config

type ApplicationConfig struct {
	Subscriptions map[string][]string
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Subscriptions: map[string][]string{},
	}
	// Subscriptions
	appConfig.Subscriptions["ghanu"] = append(appConfig.Subscriptions["ghanu"], "ghanu")
	appConfig.Subscriptions["ghanu"] = append(appConfig.Subscriptions["ghanu"], "varun")

	appConfig.Subscriptions["varun"] = append(appConfig.Subscriptions["varun"], "varun")
	appConfig.Subscriptions["varun"] = append(appConfig.Subscriptions["varun"], "pratik")

	return appConfig
}
