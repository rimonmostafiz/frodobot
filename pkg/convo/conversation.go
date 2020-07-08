// Package convo contains all the function related to conversion
package convo

import (
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"time"
)

// Values for oldest conversion time calculation
var (
	year, month, day     = time.Now().Date()
	hour, min, sec, nsec = 06, 00, 00, 000
	today                = time.Date(year, month, day, hour, min, sec, nsec, time.Now().Location())
	todayUnix            = strconv.FormatInt(today.Unix(), 10)
)

// GetConversationsHistory returns slice of slack.Message for a given channel
// As a param we send channelId and Oldest time of the message to consider
func GetConversationsHistory(channelId string, client *slack.Client) []slack.Message {
	historyParam := slack.GetConversationHistoryParameters{
		ChannelID: channelId,
		Cursor:    "",
		Inclusive: false,
		Latest:    "",
		Limit:     0,
		Oldest:    todayUnix,
	}

	history, err := client.GetConversationHistory(&historyParam)
	if err != nil {
		log.Fatalf("Error while fetching channel history %s", err)
	}
	return history.Messages
}
