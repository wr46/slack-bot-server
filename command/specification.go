package command

// Add here all active commands
// This commands will be the entry point for command discovery
var commandsDoc = [...]commandDoc{
	helpDoc,
	vacationDoc,
	jokeDoc,
}

// Add here all common hardcoded strings
// The command strings related to a single command must be added in their file scope
const msgUnknownCmd = "Hey %s, I'm unable to comply!! \n Do *help* for options... :duck:"
const msgCommandFormat = "\n %s %s \n*Syntax:* \n %s \n"
const msgSeparatorFormat = "\n"
const errorMsg = "[ERROR]: Something went wrong! :bomb:\n"
const successMsg = "Success! :sunglasses:\n"
