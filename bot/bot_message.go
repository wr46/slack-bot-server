package bot

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/wr46/slack-bot-server/command"
	"github.com/wr46/slack-bot-server/logger"
)

// AnswerMessage bot will answer to given message
func AnswerMessage(event *slack.MessageEvent, api *slack.Client) string {

	user, err := api.GetUserInfo(event.User)

	if err != nil {
		logger.Log(logger.Error, fmt.Sprintf("Failed to get user info: %s", err))
	}

	var username = user.RealName
	logger.Log(logger.Debug, fmt.Sprintf("Event from User = %s and Email = %s", username, user.Profile.Email))

	// Execute command if found!
	var cmd command.Executable = command.GetCommand(event.Text)
	if cmd == nil {
		return command.BuildUnknownCmdMsg(username)
	}

	return cmd.Run(user)
}
