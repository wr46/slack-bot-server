package server

import (
	"log"
	"os"
	"slack-bot-server/configuration"
	"slack-bot-server/event"
	"slack-bot-server/logger"

	"github.com/nlopes/slack"
)

var (
	api *slack.Client
	rtm *slack.RTM
)

/**
 * Connect to Slack API and instantiate Real Time Messaging
 */
func Setup() {
	api = slack.New(
		configuration.ENV.SlackToken,
		slack.OptionDebug(logger.IsDebug),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)
	logger.Log(logger.Debug, "API instance ready!")

	rtm = api.NewRTM()
	logger.Log(logger.Debug, "Real Time Messaging instance ready!")
}

/**
 * Listen Slack Api events and handle the events
 */
func Run() {
	go rtm.ManageConnection()

	for evt := range rtm.IncomingEvents {
		event.HandleEvent(evt.Data, api)
	}
	logger.Log(logger.Info, "Server terminated!")
}
