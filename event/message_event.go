package event

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/bot"
	"github.com/wr46/slack-bot-server/logger"
)

type MessageEvent struct {
	user    string
	text    string
	channel string
}

// HandleMessage message processing
func HandleMessage(event MessageEvent, api *slack.Client) {
	// Bot will handle the user message
	message := bot.AnswerMessage(event.user, event.text, api)
	sendMessage(api, event.channel, message)
}

// Send a message to the given user
func sendMessage(api *slack.Client, recipientID string, message string) {
	channelID, timestamp, err := api.PostMessage(
		recipientID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	logger.Log(logger.Info, fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))
}
