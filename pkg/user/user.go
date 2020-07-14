// Package user provides all function related to user
package user

import (
	"github.com/slack-go/slack"
	"log"
	"sync"
)

// GetUserList generate list of user of a channel
// Returns slice of string with user id
// Update limit if your slack channel has more than 100 people
func GetUserList(channelId string, client *slack.Client) []string {

	param := slack.GetUsersInConversationParameters{
		ChannelID: channelId,
		Cursor:    "",
		Limit:     0, // Need to update this if channel has more that 100 members
	}

	conversation, _, err := client.GetUsersInConversation(&param)
	if err != nil {
		log.Fatalf("Error while fetching user list from conversion %s", err)
	}
	return conversation
}

// InitUserMap initialize a map from given slice of string that contains user ids
// Key of the map is user id and initial value is false
// It doesn't insert an user id If that id exists inside exclude map
func InitUserMap(users []string, exclude map[string]string) map[string]bool {
	m := make(map[string]bool)
	for _, userId := range users {
		if _, ok := exclude[userId]; !ok {
			m[userId] = false
		}
	}
	return m
}

// MarkUserAsGreen update userMap
// Update user id value from false to true, if slack.Message slice contains message from that user
func MarkUserAsGreen(messages []slack.Message, userMap map[string]bool) {
	for _, msg := range messages {
		user := msg.User
		if _, ok := userMap[user]; ok {
			userMap[user] = true
		}
	}
}

// GetUserToRemind search on userMap and construct a slice of user id
// Then use getUserNames function and return usernames with count
// usernames contains name of users, who didn't provide their status today and need to get a reminder
func GetUserToRemind(userMap map[string]bool, client *slack.Client) ([]string, int) {
	count := 0
	var userIds []string

	for key := range userMap {
		userId := key
		if isTrue, ok := userMap[userId]; ok && !isTrue {
			userIds = append(userIds, key)
			count++
		}
	}
	usernames := getUserNames(count, userIds, client)
	return usernames, count
}

// GerUserDetails finds user info by userId
func GetUserDetails(userId string, client *slack.Client) *slack.User {
	info, err := client.GetUserInfo(userId)
	if err != nil {
		log.Fatalf("error while get user info userId[%s], %s", userId, err)
	}
	return info
}

// getUserNames fetch user name from userId list asynchronously using goroutines
func getUserNames(count int, userIds []string, slackClient *slack.Client) []string {
	userNames := make([]string, count)
	var wg sync.WaitGroup
	wg.Add(count)
	for idx, userId := range userIds {
		i, id := idx, userId
		go func() {
			details := GetUserDetails(id, slackClient)
			userNames[i] = details.Name
			wg.Done()
		}()
	}
	wg.Wait()
	return userNames
}
