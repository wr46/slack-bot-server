package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/logger"
)

// Arguments constants
const jokeResponse = "response"
const jokeCategories = "categories"

// Command documentation constants
const jokeSyntax = "`joke { " + jokeCategories + " | <category> }`"
const jokeRegexpStr = "joke\\s+(?P<" + jokeResponse + ">[a-z]+)\\s*$"

type jokeCmd struct {
	cmd command
}

// COMMAND TEMPLATE
// Name - The name of the command
// Syntax - The command as it must be typed
// Description - A small description of the command objective
// RegexValidation - Regex used to capture the command string and arguments
// Instance - The instance will call itself to use is interface method
var jokeDoc = commandDoc{
	name:            "_joke_",
	syntax:          jokeSyntax,
	description:     "-- give a random joke by category",
	regexValidation: jokeRegexpStr,
	instance:        jokeCmd{},
}

// Template for commands presentation
const msgHeadFormat = "*Commands list:* \n"

func (cmd jokeCmd) Run(user *slack.User) string {
	if !cmd.isValid() {
		return cmd.cmd.errorMsg
	}

	response := cmd.cmd.args["response"]

	if response == jokeCategories {
		msg := buildJokeHelpMsg()
		return msg
	}

	return buildJokeMsg(response)
}

func (cmd jokeCmd) buildCommand(args map[string]string) Executable {
	logger.Log(logger.Debug, fmt.Sprintf("Command joke arguments: %s", args))
	return jokeCmd{cmd: command{args: args, doc: jokeDoc, errorMsg: applyJokeRules(args)}}
}

func (cmd jokeCmd) isValid() bool {
	return isValid
}

// Gives a message with all commands and options by template
// TODO: improve request parsing and validation
func buildJokeMsg(category string) string {

	resp, err := http.Get("https://api.jokes.one/jod?category=" + category)
	if err != nil {
		return errorMsg + "Oops... there was a problem in the request. Is it a valid category?"
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	contents, ok := result["contents"].(map[string]interface{})

	if !ok {
		return errorMsg + "Oops... Bad response!"
	}

	jokes, ok := contents["jokes"].([]interface{})

	if !ok {
		return errorMsg + "Oops... Bad response!"
	}

	joke := jokes[0].(map[string]interface{})

	if !ok {
		return errorMsg + "Oops... Bad response!"
	}

	text := joke["joke"].(map[string]interface{})

	return text["text"].(string) + "\n :rolling_on_the_floor_laughing:"
}

// TODO: improve request parsing and validation
func buildJokeHelpMsg() string {
	resp, err := http.Get("https://api.jokes.one/jod/categories")
	if err != nil {
		return errorMsg + "Oops... there was a problem in the request"
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	contents, ok := result["contents"].(map[string]interface{})

	if !ok {
		return errorMsg + "Oops... Bad response, missing 'contents'!"
	}

	categories, ok := contents["categories"].([]interface{})

	if !ok {
		return errorMsg + "Oops... Bad response, missing 'categories'!"
	}

	var text string = "*Categories:* \n"

	for _, cat := range categories {
		m := cat.(map[string]interface{})
		name := m["name"].(string)
		description := m["description"].(string)
		text = text + "`" + name + "` - " + description + "\n"
	}

	return text
}

// Command rules
// - Arguments are required
// - Value must be a command or a category
func applyJokeRules(args map[string]string) string {
	// Validate missing arguments
	if args == nil {
		return errorMsg + noArgsMsg
	}

	response, hasStr := args["response"]

	if !hasStr || response == "" {
		return errorMsg + "There is some problem in regex!"
	}

	isValid = true

	return ""
}
