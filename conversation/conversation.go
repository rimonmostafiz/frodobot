package conversation

import (
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"time"
)

var (
	year, month, day     = time.Now().Date()
	hour, min, sec, nsec = 06, 00, 00, 000
	today                = time.Date(year, month, day-1, hour, min, sec, nsec, time.Now().Location())
	todayUnix            = strconv.FormatInt(today.Unix(), 10)
)

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
