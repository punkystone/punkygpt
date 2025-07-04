package main

import (
	"log"
	"punkygpt/internal"
	"regexp"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	client "github.com/punkystone/go-twitch-irc"
)

var chatRegex = regexp.MustCompile(`((?i)@punkygpt),?\s(.+)`)

func main() {
	env, err := internal.CheckEnv()
	if err != nil {
		panic(err)
	}
	client := client.NewClient(env.ClientID, env.ClientSecret, "punkystone", env.AccessToken, env.RefreshToken, nil)

	go func() {
		for err := range client.ErrorChannel {
			log.Printf("Error: %v\n", err)
		}
	}()

	client.IRCClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		trimmedMessage := strings.TrimSpace(message.Message)
		matches := chatRegex.FindAllStringSubmatch(trimmedMessage, -1)
		if len(matches) == 0 || len(matches[0]) != 3 {
			return
		}
		actualMessage := matches[0][2]
		response, err := internal.GetResponse(message.User.DisplayName, actualMessage, env.OpenAIAPIKey)
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}
		client.IRCClient.Reply(env.Chat, message.ID, response)
	})

	client.IRCClient.OnConnect(func() {
		log.Println("Connected to Twitch IRC")
	})

	client.IRCClient.Join(env.Chat)

	client.Connect()
}
