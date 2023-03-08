package main

import (
	"github.com/wr46/slack-bot-server/cmd/servers"
	"github.com/wr46/slack-bot-server/configuration"
)

func main() {

	// Set environment variables and init logger
	configuration.Setup()

	// Connect to Slack by socket mode
	api := servers.SetupSlackApi()

	// Listen Slack Api events
	servers.RunSlack(api)
}
