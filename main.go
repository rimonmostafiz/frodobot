package main

import (
	"fmt"
	"github.com/rimonmsotafiz/frodobot/pkg/cfg"
	"github.com/rimonmsotafiz/frodobot/pkg/convo"
	"github.com/rimonmsotafiz/frodobot/pkg/user"
	"github.com/robfig/cron/v3"
	"github.com/slack-go/slack"
	"io"
	"log"
	"os"
)

// InfoLogger for writing info log
var (
	InfoLogger *log.Logger
)

func main() {
	initLogger()
	InfoLogger.Printf("FrodoBot Application Start...")
	c := cron.New()
	cfg.InitViper(".env")

	token := cfg.ReadFromEnv("SLACK_TOKEN")
	channelId := cfg.ReadFromEnv("CHANNEL_ID")
	//testChannelId := cfg.ReadFromEnv("TEST_CHANNEL_ID")
	excludeUserMap := cfg.ReadStringMapFromEnv("EXCLUDE_USER_VALUES")
	slackClient := slack.New(token)

	_, _ = c.AddFunc("45 10 * * 0-4", func() {
		InfoLogger.Printf("FrodoBot Cron Started")
		users := user.GetUserList(channelId, slackClient)
		userMap := user.InitUserMap(users, excludeUserMap)
		messages := convo.GetConversationsHistory(channelId, slackClient)
		user.MarkUserAsGreen(messages, userMap)
		usernames, count := user.GetUserToRemind(userMap, slackClient)

		fmt.Printf("Total %v User to remind\n", count)

		if count > 0 {
			msg := convo.ConstructReminderMsg(usernames)
			InfoLogger.Println(msg)
			convo.SendReminder(channelId, msg, slackClient)
		}
		InfoLogger.Printf("FrodoBot Cron Ended")
	})

	c.Start()
	<-make(chan struct{})
}

func initLogger() {
	file, err := os.OpenFile("frodobot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
