package command

import (
	"fmt"
	"regexp"
)

// GetCommand this is the main function to return the command by given string
func GetCommand(text string) Executable {
	for _, cmd := range commandsDoc {
		regx := regexp.MustCompile(cmd.regexValidation)
		match := regx.FindStringSubmatch(text)
		if len(match) > 0 {
			args := make(map[string]string)
			for i, label := range regx.SubexpNames() {
				if i > 0 && i <= len(match) {
					args[label] = match[i]
				}
			}
			// Feel free to improve this solution for abstraction
			return cmd.instance.buildCommand(args)
		}
	}
	return nil
}

// BuildUnknownCmdMsg gives a message for unknown command
func BuildUnknownCmdMsg(username string) string {
	return fmt.Sprintf(msgUnknownCmd, username)
}
