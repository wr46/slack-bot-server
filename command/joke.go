package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/logger"
)

// Errors
const badResponse = "Oops... Bad response!"

// Arguments constants.
const jokeResponse = "response"
const jokeCategories = "categories"

// Command documentation constants.
const jokeSyntax = "`joke { " + jokeCategories + " | <category> }`"
const jokeRegexpStr = "joke\\s+(?P<" + jokeResponse + ">[a-z]+)\\s*$"

// Jokes API.
const jokeAPIBaseURLStr = "https://api.jokes.one/jod"

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

	response := cmd.cmd.args[jokeResponse]
	baseURL, err := url.Parse(jokeAPIBaseURLStr)

	if err != nil {
		return errorMsg + "Malformed URL!"
	}

	if response == jokeCategories {
		msg := buildJokeHelpMsg(baseURL)

		return msg
	}

	return buildJokeMsg(baseURL, response)
}

func (cmd jokeCmd) buildCommand(args map[string]string) Executable {
	logger.Log(logger.Debug, fmt.Sprintf("Command joke arguments: %s", args))

	return jokeCmd{cmd: command{args: args, doc: jokeDoc, errorMsg: applyJokeRules(args)}}
}

func (cmd jokeCmd) isValid() bool {
	return isValid
}

// Gives a message with all commands and options by template
func buildJokeMsg(baseURL *url.URL, category string) string {
	params := url.Values{}
	params.Add("category", category)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return errorMsg + "Oops... there was a problem in the request. Is it a valid category?"
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return errorMsg + "Oops... there was a problem in parsing response"
	}
	defer resp.Body.Close()

	contents, ok := result["contents"].(map[string]interface{})

	if !ok {
		return errorMsg + badResponse
	}

	jokes, ok := contents["jokes"].([]interface{})

	if !ok {
		return errorMsg + badResponse
	}

	joke := jokes[0].(map[string]interface{})

	if !ok {
		return errorMsg + badResponse
	}

	text := joke["joke"].(map[string]interface{})

	return text["text"].(string) + "\n :rolling_on_the_floor_laughing:"
}

func buildJokeHelpMsg(baseURL *url.URL) string {
	resp, err := http.Get(baseURL.String() + "/" + jokeCategories)
	if err != nil {
		return errorMsg + "Oops... there was a problem in the request"
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return errorMsg + "Oops... there was a problem in parsing response"
	}
	defer resp.Body.Close()

	contents, ok := result["contents"].(map[string]interface{})

	if !ok {
		return errorMsg + badResponse + "Missing 'contents'!"
	}

	categories, ok := contents["categories"].([]interface{})

	if !ok {
		return errorMsg + badResponse + "Missing 'categories'!"
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

	response, hasStr := args[jokeResponse]

	if !hasStr || response == "" {
		return errorMsg + "There is some problem in regex!"
	}

	isValid = true

	return ""
}
