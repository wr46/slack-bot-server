package event

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/logger"
)

// HandleEvent the event handle and is responsible for event redirection
func HandleEvent(data interface{}, api *slack.Client) {
	logger.Log(logger.Debug, "Event will be handled!")
	switch ev := data.(type) {
	case *slack.HelloEvent:
		logger.Log(logger.Debug, "Hello event done!")

	case *slack.ConnectedEvent:
		logger.Log(logger.Debug, fmt.Sprintf("Infos: %v", ev.Info))
		logger.Log(logger.Debug, fmt.Sprintf("Connection counter: %v", ev.ConnectionCount))

	case *slack.MessageEvent:
		go HandleMessage(ev, api)

	case *slack.PresenceChangeEvent:
		logger.Log(logger.Debug, fmt.Sprintf("Presence Change: %v", ev))

	case *slack.LatencyReport:
		logger.Log(logger.Debug, fmt.Sprintf("Current latency: %v", ev.Value))

	case *slack.RTMError:
		logger.Log(logger.Warning, fmt.Sprintf("Error: %s", ev.Error()))

	case *slack.InvalidAuthEvent:
		logger.Log(logger.Warning, "Invalid credentials")
		return

	default:
		logger.Log(logger.Debug, "The Event is unknown!")
	}
}
