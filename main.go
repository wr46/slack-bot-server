package main

import (
	"slack-bot-server/configuration"
	"slack-bot-server/server"
)

func main() {

	// Set environment variables and init logger
	configuration.Setup()

	// Connect to Slack API and instantiate RTM
	server.Setup()

	// Listen Slack Api events
	server.Run()
}
