package command

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/logger"
)

// Errors
const invalidValuesMsg = "Invalid values! Check help..."
const failedToRunScriptMsg = "Script failed to execute! Check logs..."

// Arguments constants
const scriptTypes = "aa | bb | cc"
const scriptArgType = "type"
const scriptArgUser = "user"

// Command documentation constants
const scriptSyntax = "`script { " + scriptTypes + " } <user>`"
const scriptRegexpStr = "script\\s+(?P<" + scriptArgType + ">[a-z]{2,})\\s+(?P<" + scriptArgUser + ">[a-z]+)\\s*$"

const scriptPath = "/app/scripts/script.sh"

type scriptCmd struct {
	cmd command
}

var scriptDoc = commandDoc{
	name:            "_script_",
	syntax:          scriptSyntax,
	description:     "-- execute a script by a given script type and user",
	regexValidation: scriptRegexpStr,
	instance:        scriptCmd{},
}

var isScriptValid = false

func (cmd scriptCmd) Run(user *slack.User) string {
	if !cmd.isValid() {
		return cmd.cmd.errorMsg
	}

	result := executeScript(cmd.cmd.args[scriptArgType], cmd.cmd.args[scriptArgUser])

	return result
}

func (cmd scriptCmd) buildCommand(args map[string]string) Executable {
	logger.Log(logger.Debug, fmt.Sprintf("Command script arguments: %s", args))
	return scriptCmd{cmd: command{args: args, doc: scriptDoc, errorMsg: applyScriptRules(args)}}
}

func (cmd scriptCmd) isValid() bool {
	return isScriptValid
}

// Command rules
// - Script type must be known
func applyScriptRules(args map[string]string) string {
	// Validate missing arguments
	if args == nil {
		return errorMsg + noArgsMsg
	}

	if !strings.Contains(scriptTypes, args[scriptArgType]) {
		return errorMsg + invalidValuesMsg
	}

	isScriptValid = true

	return ""
}

func executeScript(sType string, user string) string {
	logger.Log(logger.Debug, "/app/scripts/script.sh")

	cmd := exec.Command("/bin/sh", scriptPath, sType, user)
	output, err := cmd.Output()

	if err != nil {
		logger.Log(logger.Error, err.Error())
		return errorMsg + failedToRunScriptMsg
	}

	logger.Log(logger.Debug, string(output))

	return string(output)
}
