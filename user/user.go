package user

import (
	"github.com/slack-go/slack"
	"log"
)

func GetUserList(channelId string, client *slack.Client) []string {
	usersInConversionsParam := slack.GetUsersInConversationParameters{
		ChannelID: channelId,
		Cursor:    "",
		Limit:     0,
	}

	conversation, _, err := client.GetUsersInConversation(&usersInConversionsParam)
	if err != nil {
		log.Fatalf("Error while fetching user list from conversion %s", err)
	}
	return conversation
}

func CreateUserMap(users []string, exclude map[string]string) map[string]bool {
	m := make(map[string]bool)
	for _, userId := range users {
		if _, ok := exclude[userId]; !ok {
			m[userId] = false
		}
	}
	return m
}

func MarkUserAsGreen(messages []slack.Message, userMap map[string]bool) {
	for _, msg := range messages {
		user := msg.User
		if _, ok := userMap[user]; ok {
			userMap[user] = true
		}
	}
}

func GetUserToRemind(userMap map[string]bool) ([]string, int) {
	count := 0
	userToRemind := make([]string, 100, 200)

	for key, _ := range userMap {
		userId := key
		if isTrue, ok := userMap[userId]; ok && !isTrue {
			userToRemind[count] = key
			count++
		}
	}
	return userToRemind, count
}
