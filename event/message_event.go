package event

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/bot"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/logger"
)

const privateChanelPrefix = 'D'

// HandleMessage message processing
func HandleMessage(event *slack.MessageEvent, api *slack.Client) {
	if isToRejectMessage(event) {
		return
	}
	// Bot will handle the user message
	message := bot.AnswerMessage(event, api)
	sendMessage(api, event.Channel, message)
}

// Evaluate if is related to a Bot or an User message
func isABotMessage(event *slack.MessageEvent) bool {
	if len(event.BotID) > 0 {
		logger.Log(logger.Debug, fmt.Sprintf("BotID = '%s'", event.BotID))
		return true
	}
	return false
}

// Evaluate if message is to the Bot
func isMessageToBot(event *slack.MessageEvent) bool {
	// Message sent into Bot chat window or with Bot tagged
	if isPrivateChat(event.Channel) ||
		strings.Contains(event.Text, "<@"+configuration.Env.BotID+">") {
		return true
	}

	return false
}

// Evaluate if channel is private or public
// Private chat in Slack has a prefix letter
func isPrivateChat(channelID string) bool {
	if channelID != "" && channelID[0] == privateChanelPrefix {
		return true
	}
	return false
}

// Evaluate if message is to be handled
func isToRejectMessage(event *slack.MessageEvent) bool {
	// Only user messages are handled
	if isABotMessage(event) || !isMessageToBot(event) {
		return true
	}
	return false
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
