package command

import (
	"fmt"
	"time"

	"github.com/nlopes/slack"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/logger"
	"github.com/wr46/slack-bot-server/utils"
)

// Errors
const noArgsMsg = "No arguments found! Check regex..."
const argsMissingMsg = "Arguments missing! Check regex..."
const valFTMissingMsg = "Values missing (From <> To)! Check regex..."
const valOnMissingMsg = "Values missing (On)! Check regex..."
const invalidValFromMsg = "Invalid value From date!"
const invalidValToMsg = "Invalid value To date!"
const invalidValOnMsg = "Invalid value On date!"
const invalidValDayMsg = "Invalid part of the day!"
const invalidValFromTodayMsg = "Value From date must be after today!"
const invalidValOnTodayMsg = "Value On date must be after today!"
const invalidValFromToMsg = "Value From date must be before To date!"

// Arguments constants
const argFrom = "from"
const argTo = "to"
const argOn = "on"
const argDay = "day"
const partOfDayMorning = "morning"
const partOfDayAfternoon = "afternoon"

// Command documentation constants
const syntax = "`vacation request { from dd/mm/yyyy to dd/mm/yyyy | on dd/mm/yyyy [{ " + partOfDayMorning + " | " + partOfDayAfternoon + " }] }`"
const regexpStr = "vacation\\s+request\\s+from\\s+(?P<" + argFrom + ">\\d{2}/\\d{2}/\\d{4})\\s+to\\s+(?P<" + argTo + ">\\d{2}/\\d{2}/\\d{4})\\s*$|on\\s+(?P<" + argOn + ">\\d{2}/\\d{2}/\\d{4})\\s*(?P<" + argDay + ">" + partOfDayMorning + "|" + partOfDayAfternoon + "|)\\s*$"

// Email template
const emailSubjectTemplate = "[Vacation] Bot vacation request from %s"
const emailHTMLBodyTemplate = `<p>Hello,</p><br>
								<p>This is a vacation request from %s at %s</p>
								%s<br>
								<p>Best Regards.</p>
								<p>Bot &#128540;</p>`
const emailHTMLBodyFromToPart = "<p>%s requested vacation from %s to %s.</p>"
const emailHTMLBodyOnPart = "<p>%s request vacation on %s</p>"
const emailHTMLBodyOnDayPart = "<p>%s request vacation on %s only in the %s</p>"
const emailSent = "Vacations requested!"
const emailNotSent = "Vacations requested was not sent! :vader:"

type vacationCmd struct {
	cmd command
}

var vacationDoc = commandDoc{
	name:            "_vacation_",
	syntax:          syntax,
	description:     "-- submit a vacation period, a day or half a day request",
	regexValidation: regexpStr,
	instance:        vacationCmd{},
}

var isValid = false

func (cmd vacationCmd) Run(user *slack.User) string {
	if !cmd.isValid() {
		return cmd.cmd.errorMsg
	}
	return sendMessage(user, cmd.cmd.args)
}

func (cmd vacationCmd) buildCommand(args map[string]string) Executable {
	logger.Log(logger.Debug, fmt.Sprintf("Command vacation arguments: %s", args))
	return vacationCmd{cmd: command{args: args, doc: vacationDoc, errorMsg: applyRules(args)}}
}

func (cmd vacationCmd) isValid() bool {
	return isValid
}

// Command rules
// - Arguments are required
// - Values must be valid dates
func applyRules(args map[string]string) string {
	// Validate missing arguments
	if args == nil {
		return errorMsg + noArgsMsg
	}

	valFrom, hasFrom := args[argFrom]
	valTo, hasTo := args[argTo]
	valOn, hasOn := args[argOn]
	valDay, hasDay := args[argDay]

	if !hasFrom || !hasTo || !hasOn || !hasDay {
		return errorMsg + argsMissingMsg
	}
	if (valFrom != "" && valTo == "") || (valFrom == "" && valTo != "") {
		return errorMsg + valFTMissingMsg
	}

	// Validate from and to dates
	if valFrom != "" && valTo != "" {
		if _, err := utils.ParseDate(valFrom); err != nil {
			return errorMsg + invalidValFromMsg
		}
		if _, err := utils.ParseDate(valTo); err != nil {
			return errorMsg + invalidValToMsg
		}
		if isFuture, _ := utils.IsFutureStringDate(valFrom); !isFuture {
			return errorMsg + invalidValFromTodayMsg
		}
		if result, _ := utils.CompareStringDates(valFrom, valTo); result >= 0 {
			return errorMsg + invalidValFromToMsg
		}

		isValid = true
		return ""
	}

	// Validate on dates
	if valOn == "" {
		return errorMsg + valOnMissingMsg
	}
	if _, err := utils.ParseDate(valOn); err != nil {
		return errorMsg + invalidValOnMsg
	}
	if isFuture, _ := utils.IsFutureStringDate(valOn); !isFuture {
		return errorMsg + invalidValOnTodayMsg
	}
	if valDay != "" && valDay != partOfDayMorning && valDay != partOfDayAfternoon {
		return errorMsg + invalidValDayMsg
	}

	isValid = true
	return ""
}

// This will send an email for the configurated email and cc for the caller
// The given user and values will be to compose the email message request
func sendMessage(user *slack.User, args map[string]string) string {
	username := user.RealName
	valFrom, _ := args[argFrom]
	valTo, _ := args[argTo]
	valOn, _ := args[argOn]
	valDay, _ := args[argDay]

	bodyPart := ""
	if valFrom != "" {
		bodyPart = fmt.Sprintf(emailHTMLBodyFromToPart, username, valFrom, valTo)
	} else if valOn != "" {
		bodyPart = fmt.Sprintf(emailHTMLBodyOnPart, username, valOn)
	} else if valOn != "" && valDay != "" {
		bodyPart = fmt.Sprintf(emailHTMLBodyOnDayPart, username, valOn, valDay)
	}

	t := time.Now()
	subject := fmt.Sprintf(emailSubjectTemplate, username)
	// recipients := []string{utils.ExtractEmails(message)}
	recipients := []string{configuration.Env.VacationRecipientEmail}
	htmlBody := fmt.Sprintf(emailHTMLBodyTemplate, username, t.Format("15:04:05 on 02/01/2006"), bodyPart)

	if utils.SendEmail(utils.BuildMessage(subject, user, htmlBody, recipients)) {
		return successMsg + emailSent
	}
	return errorMsg + emailNotSent
}
