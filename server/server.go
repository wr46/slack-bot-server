package server

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/wr46/slack-bot-server/configuration"
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
	/*
		go rtm.ManageConnection()

		for evt := range rtm.IncomingEvents {
			event.HandleEvent(evt.Data, api)
		}
		logger.Log(logger.Info, "Server terminated!")
	*/

	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		logger.Log(logger.Info, body)
		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: configuration.Env.Slack.VerificationToken}))
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text")
			w.Write([]byte(r.Challenge))
		}
		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
			}
		}
	})

	logger.Log(logger.Info, "Server listening!")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		logger.Log(logger.Error, err.Error())
	}
}
