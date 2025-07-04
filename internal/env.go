package internal

import (
	"errors"
	"os"
)

type Env struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	OpenAIAPIKey string
	Chat         string
}

func CheckEnv() (*Env, error) {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")
	if !exists {
		return nil, errors.New("ACCESS_TOKEN environment variable not set")
	}
	refreshToken, exists := os.LookupEnv("REFRESH_TOKEN")
	if !exists {
		return nil, errors.New("REFRESH_TOKEN environment variable not set")
	}
	clientID, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		return nil, errors.New("CLIENT_ID environment variable not set")
	}
	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		return nil, errors.New("CLIENT_SECRET environment variable not set")
	}
	openAIAPIKey, exists := os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		return nil, errors.New("OPENAI_API_KEY environment variable not set")
	}
	chat, exists := os.LookupEnv("CHAT")
	if !exists {
		return nil, errors.New("CHAT environment variable not set")
	}
	env := &Env{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		OpenAIAPIKey: openAIAPIKey,
		Chat:         chat,
	}
	return env, nil
}
