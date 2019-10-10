package server

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/event"
	"github.com/wr46/slack-bot-server/logger"
)

var (
	api *slack.Client
	rtm *slack.RTM
)

// Setup connect to Slack API and instantiate Real Time Messaging
func Setup() {
	api = slack.New(
		configuration.Env.SlackToken,
		slack.OptionDebug(logger.IsDebug),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)
	logger.Log(logger.Debug, "API instance ready!")

	rtm = api.NewRTM()
	logger.Log(logger.Debug, "Real Time Messaging instance ready!")
}

// Run Listen Slack Api events and handle the events
func Run() {
	go rtm.ManageConnection()

	for evt := range rtm.IncomingEvents {
		event.HandleEvent(evt.Data, api)
	}
	logger.Log(logger.Info, "Server terminated!")
}
