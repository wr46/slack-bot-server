package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/logger"
)

var (
	api *slack.Client
	rtm *slack.RTM
)

// Setup connect to Slack API and instantiate Real Time Messaging
func Setup() {
	logger.Log(logger.Debug, configuration.Env.SlackToken)
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
		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "TOKEN"}))
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

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
