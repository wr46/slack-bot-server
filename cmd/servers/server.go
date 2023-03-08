package servers

import (
	"net/http"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/event"
	"github.com/wr46/slack-bot-server/logger"
)

var api *slack.Client

// Setup Slack API
func Setup() {
	slackToken := configuration.Env.Slack.BotOauthToken
	logger.Log(logger.Debug, slackToken)
	api = slack.New(slackToken)

	logger.Log(logger.Debug, "Slack API setup ready!")
}

// Run Listen Slack Api events and handle the events
func Run() {
	//event.SetContext(api)
	http.HandleFunc("/events-endpoint", event.HandleEvent)
	logger.Log(logger.Info, "Server listening!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log(logger.Error, err.Error())
	}
}
