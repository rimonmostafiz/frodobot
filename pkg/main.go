package main

import (
	"fmt"
	"github.com/rimonmsotafiz/flott-bot/pkg/cfg"
	"github.com/rimonmsotafiz/flott-bot/pkg/convo"
	"github.com/rimonmsotafiz/flott-bot/pkg/user"
	"github.com/slack-go/slack"
)

func main() {
	cfg.InitViper(".env")

	token := cfg.ReadFromEnv("SLACK_TOKEN")
	channelId := cfg.ReadFromEnv("CHANNEL_ID")
	excludeUserMap := cfg.ReadStringMapFromEnv("EXCLUDE_USER_VALUES")

	slackClient := slack.New(token)

	users := user.GetUserList(channelId, slackClient)
	userMap := user.InitUserMap(users, excludeUserMap)
	messages := convo.GetConversationsHistory(channelId, slackClient)
	user.MarkUserAsGreen(messages, userMap)
	userToRemind, count := user.GetUserToRemind(userMap)

	fmt.Println(userToRemind, count)
}
