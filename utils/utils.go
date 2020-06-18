package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/slack-go/slack"
	"github.com/wr46/slack-bot-server/configuration"
	"github.com/wr46/slack-bot-server/logger"
	"gopkg.in/gomail.v2"
)

// HasString search substring in a given string
func HasString(haystack string, needle string) bool {
	return strings.Contains(haystack, needle)
}

// ParseDate parse string date by "02/01/2006" date format
func ParseDate(date string) (time.Time, error) {
	layoutISO := "02/01/2006"
	return time.Parse(layoutISO, date)
}

// IsFutureStringDate check if giver string date is after today
// return true if date > today and false otherwise
// return error if string date is invalid
func IsFutureStringDate(date string) (bool, error) {
	t, err := ParseDate(date)
	if err != nil {
		return false, err
	}

	if t.After(time.Now()) {
		return true, nil
	}

	return false, nil
}

// CompareStringDates compare both given string dates
// return -1 if date1 before date2
// return 0 if date1 equal to date2
// return 1 if date1 after date2
// return error if dates are invalid
func CompareStringDates(date1 string, date2 string) (int, error) {
	t1, err := ParseDate(date1)
	if err != nil {
		return 0, err
	}

	t2, err := ParseDate(date2)
	if err != nil {
		return 0, err
	}

	if t1.Equal(t2) {
		return 0, nil
	}

	if t1.Before(t2) {
		return -1, nil
	}

	return 1, nil
}

// BuildMessage fill email headers with given data
func BuildMessage(subject string, user *slack.User, htmlBody string, recipients []string) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", configuration.Env.Email.User)
	message.SetHeader("To", recipients...)
	message.SetAddressHeader("Cc", user.Profile.Email, user.RealName)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", htmlBody)
	// message.Attach("/home/*.jpg")
	return message
}

// SendEmail send email with given message by configuration properties
func SendEmail(message *gomail.Message) bool {
	dialer := gomail.NewPlainDialer(
		configuration.Env.Email.SMTPServer,
		configuration.Env.Email.SMTPPort,
		configuration.Env.Email.User,
		configuration.Env.Email.Password)

	if err := dialer.DialAndSend(message); err != nil {
		logger.Log(logger.Warning, fmt.Sprintf("Send email has failed! %s", err))
		return false
	}

	return true
}

// ExtractEmails get the first email sent by a user Slack message string
func ExtractEmails(text string) string {
	re := regexp.MustCompile(`mailto:(.*)\|`)
	email := re.FindStringSubmatch(text)[1]
	logger.Log(logger.Debug, fmt.Sprintf("Extracted email: %s", email))

	return email
}
