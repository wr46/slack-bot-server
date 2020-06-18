package event

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack/slackevents"
	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/logger"
)

var api *slack.Client

func SetContext(apiContext *slack.Client) {
	api = apiContext
}

// HandleEvent the event handle and is responsible for event redirection
func HandleEvent(writer http.ResponseWriter, request *http.Request) {
	body := getBody(request)
	eventsAPIEvent, err := parseEvent(body)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		urlVerification(writer, body)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			logger.Log(logger.Debug, "Mention event will be handled!")

			go HandleMessage(MessageEvent{ev.User, ev.Text, ev.Channel}, api)
		case *slackevents.MessageEvent:
			logger.Log(logger.Debug, "Message event will be handled!")
			// Ignore if message event from Bot
			if len(ev.BotID) == 0 {
				go HandleMessage(MessageEvent{ev.User, ev.Text, ev.Channel}, api)
			}
		default:
			logger.Log(logger.Debug, "The Event is unknown!")
		}
	}
}

func getBody(request *http.Request) string {
	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(request.Body); err != nil {
		logger.Log(logger.Error, err.Error())
	}

	return buffer.String()
}

func parseEvent(body string) (slackevents.EventsAPIEvent, error) {
	comparator := &slackevents.TokenComparator{VerificationToken: configuration.Env.Slack.VerificationToken}
	return slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(comparator))
}

func urlVerification(writer http.ResponseWriter, body string) {
	var r *slackevents.ChallengeResponse
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "text")

	if _, err := writer.Write([]byte(r.Challenge)); err != nil {
		logger.Log(logger.Error, err.Error())
	}
}
