// Package convo contains all the function related to conversion
package convo

import (
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"time"
)

// getConversationWindow will calculate and returns the time window(startTime and endTime) for conversation history
func getConversationWindow() (string, string) {
	year, month, day := time.Now().Date()
	startHour, startMin, startSec := 06, 00, 00
	startingTime := time.Date(year, month, day, startHour, startMin, startSec, 000, time.Now().Location())
	startingTimeUnix := strconv.FormatInt(startingTime.Unix(), 10)
	endHour, endMin, endSec := 10, 45, 00
	endTime := time.Date(year, month, day, endHour, endMin, endSec, 000, time.Now().Location())
	endTimeUnix := strconv.FormatInt(endTime.Unix(), 10)
	return startingTimeUnix, endTimeUnix
}

// GetConversationsHistory returns slice of slack.Message for a given channel from oldest time to latest time
// As a param we send channelId and Oldest time of the message to consider
func GetConversationsHistory(channelId string, client *slack.Client) []slack.Message {
	startTime, endTime := getConversationWindow()
	historyParam := slack.GetConversationHistoryParameters{
		ChannelID: channelId,
		Cursor:    "",
		Inclusive: false,
		Latest:    endTime,
		Limit:     0,
		Oldest:    startTime,
	}

	history, err := client.GetConversationHistory(&historyParam)
	if err != nil {
		log.Fatalf("Error while fetching channel history %s", err)
	}
	return history.Messages
}

// ConstructReminderMsg prepare final message to send for reminder
func ConstructReminderMsg(usernames []string) string {
	remindMessageBody := ""
	for _, usr := range usernames {
		remindMessageBody += "<@" + usr + ">, "
	}
	return "Hello " + remindMessageBody + "just to remind you, you need to post your updates before 10:45 AM."
}

// SendReminder send message to a channel
func SendReminder(channelId string, msg string, client *slack.Client) {
	msgOption := slack.MsgOptionText(msg, false)
	options := []slack.MsgOption{msgOption}
	_, _, err := client.PostMessage(channelId, options...)
	if err != nil {
		log.Fatalf("Error while sending msg to channel[%s], %s", channelId, err)
	}
}
