package command

import "github.com/slack-go/slack"

// Executable interface will give command abstraction
// Run for command execution
// buildCommand for validation and instantiation
// isValid to get command validation.
type Executable interface {
	Run(user *slack.User) string
	buildCommand(map[string]string) Executable
	isValid() bool
}

// COMMAND DOCUMENTATION TEMPLATE
// name - The name of the command
// syntax - The command as it must be typed
// description - A small description of the command objective
// regexValidation - Regex used to capture the command string and arguments
// instance - The instance will call itself to use is interface method.
type commandDoc struct {
	name            string
	syntax          string
	description     string
	regexValidation string
	instance        Executable
}

// COMMAND DOCUMENTATION TEMPLATE
// args - This will store the command arguments
// doc - The related command documentation
// errorMsg - Error messages related with command process.
type command struct {
	args     map[string]string
	doc      commandDoc
	errorMsg string
}
