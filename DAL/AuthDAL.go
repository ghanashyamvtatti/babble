package DAL

import "ds-project/config"

func SetAccessToken(appConfig *config.ApplicationConfig, username string, token string) bool {
	appConfig.Tokens[username] = token
	return true
}

func GetAccessToken(appConfig *config.ApplicationConfig, username string) (string, bool) {
	token, ok := appConfig.Tokens[username]
	return token, ok
}

func DeleteAccessToken(appConfig *config.ApplicationConfig, username string) bool {
	if _, ok := appConfig.Tokens[username]; ok {
		delete(appConfig.Tokens, username)
	}
	return true
}
