package event

import (
	"fmt"
	"geekbot/bot"
	"geekbot/configuration"
	"geekbot/logger"
	"strings"

	"github.com/nlopes/slack"
)

func HandleMessage(event *slack.MessageEvent, api *slack.Client) {
	if isToRejectMessage(event) {
		return
	}
	// Bot will handle the user message
	message := bot.AnswerMessage(event, api)
	sendMessage(api, event.Channel, message)
}

/**
 * Evaluate if is related to a Bot or an User message
 */
func isABotMessage(event *slack.MessageEvent) bool {
	if len(event.BotID) > 0 {
		logger.Log(logger.Debug, fmt.Sprintf("BotId = '%s'", event.BotID))
		logger.Log(logger.Debug, fmt.Sprintf("BotChannelId = '%s'", event.Channel))
		return true
	}
	return false
}

/**
 * Evaluate if message is to the Bot
 */
func isMessageToBot(event *slack.MessageEvent) bool {
	// Message sent into Bot chat window or with Bot tagged
	if configuration.ENV.BotChannelId == event.Channel ||
		strings.Contains(event.Text, "<@"+configuration.ENV.BotId+">") {
		return true
	}

	return false
}

/**
 * Evaluate if message is to be handled
 */
func isToRejectMessage(event *slack.MessageEvent) bool {
	// Only user messages are handled
	if isABotMessage(event) || !isMessageToBot(event) {
		return true
	}
	return false
}

/**
 * Send a message to the given user
 */
func sendMessage(api *slack.Client, recipientId string, message string) {
	channelID, timestamp, err := api.PostMessage(
		recipientId,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	logger.Log(logger.Info, fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))
}
