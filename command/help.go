package command

import (
	"fmt"

	"github.com/slack-go/slack"
)

type helpCmd struct {
	cmd command
}

// COMMAND TEMPLATE
// Name - The name of the command
// Syntax - The command as it must be typed
// Description - A small description of the command objective
// RegexValidation - Regex used to capture the command string and arguments
// Instance - The instance will call itself to use is interface method.
var helpDoc = commandDoc{
	name:            "_help_",
	syntax:          "`help`",
	description:     "-- list all accepted commands",
	regexValidation: "help\\s*$",
	instance:        helpCmd{},
}

// Template for commands presentation.
const msgHeaderFormat = "*Commands list:* \n"

func (cmd helpCmd) Run(user *slack.User) string {
	return buildHelpMsg()
}

func (cmd helpCmd) buildCommand(arguments map[string]string) Executable {
	return helpCmd{cmd: command{args: arguments, doc: helpDoc, errorMsg: ""}}
}

func (cmd helpCmd) isValid() bool {
	return true
}

// Gives a message with all commands and options by template.
func buildHelpMsg() string {
	message := msgHeaderFormat + msgSeparatorFormat
	for _, command := range commandsDoc {
		message += fmt.Sprintf(msgCommandFormat, command.name, command.description, command.syntax)
		message += msgSeparatorFormat
	}

	return message
}
