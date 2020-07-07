package main

import (
	"fmt"
	"github.com/rimonmsotafiz/flott-bot/conversation"
	"github.com/rimonmsotafiz/flott-bot/user"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
	"log"
)

var (
	Viper = viper.New()
)

func main() {
	Viper.SetConfigFile(".env")
	token := readFromEnv("SLACK_TOKEN")
	channelId := readFromEnv("CHANNEL_ID")
	excludeUserMap := Viper.GetStringMapString("EXCLUDE_USER_VALUES")

	slackClient := slack.New(token)

	users := user.GetUserList(channelId, slackClient)
	userMap := user.CreateUserMap(users, excludeUserMap)
	messages := conversation.GetConversationsHistory(channelId, slackClient)
	user.MarkUserAsGreen(messages, userMap)
	userToRemind, count := user.GetUserToRemind(userMap)

	fmt.Println(userToRemind, count)
}

func readFromEnv(key string) string {

	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := Viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
