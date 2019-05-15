package bot

import (
	"github.com/nlopes/slack"
)

/**
 * Bot will answer to given message
 */
func AnswerMessage(event *slack.MessageEvent, api *slack.Client) string {

	// Do something

	return "Hello World! [Answer to: '" + event.Text + "']"
}
