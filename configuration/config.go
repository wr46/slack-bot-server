package configuration

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wr46/slack-bot-server/logger"
)

// Dotenv file variables labels
const (
	DebugLbl                  = "DEBUG"
	BotOauthTokenLbl          = "BOT_USER_OAUTH_ACCESS_TOKEN"
	VerificationTokenLbl      = "VERIFICATION_TOKEN"
	EmailSMTPServerLbl        = "EMAIL_SMTP_SERVER"
	EmailSMTPPortLbl          = "EMAIL_SMTP_PORT"
	EmailUserLbl              = "EMAIL_USER"
	EmailPasswordLbl          = "EMAIL_PASSWORD"
	VacationRecipientEmailLbl = "VACATION_RECIPIENT_EMAIL"
)

type SlackEnv struct {
	BotOauthToken     string
	VerificationToken string
	BotID             string
	BotChannelID      string
}

type EmailEnv struct {
	SMTPServer        string
	SMTPPort          int
	User              string
	Password          string
	VacationRecipient string
}

// Configuration container
type Configuration struct {
	Debug bool
	Slack SlackEnv
	Email EmailEnv
}

// Env environment data
var Env Configuration

// Setup configuration
func Setup() {
	// Initialize logger
	logger.Setup(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	setupApplication()
	setupSlack()
	setupEmail()

	logger.SetDebug(Env.Debug)
	logger.Log(logger.Info, "Environment is ready!")
}

func setupApplication() {
	fmt.Println("Hello " + os.Getenv("DEBUG"))

	isDebug, err := strconv.ParseBool(os.Getenv(DebugLbl))
	if err != nil {
		logger.Log(logger.Error, err.Error())
	}

	Env.Debug = isDebug

	logger.Log(logger.Info, "Debug mode: "+strconv.FormatBool(Env.Debug))
}

func setupSlack() {
	Env.Slack.BotOauthToken = os.Getenv(BotOauthTokenLbl)
	Env.Slack.VerificationToken = os.Getenv(VerificationTokenLbl)
}

func setupEmail() {
	Env.Email.SMTPServer = os.Getenv(EmailSMTPServerLbl)

	var smtpPort, errPrt = strconv.Atoi(os.Getenv(EmailSMTPPortLbl))
	if errPrt != nil {
		logger.Log(logger.Error, errPrt.Error())
	}

	Env.Email.SMTPPort = smtpPort
	Env.Email.User = os.Getenv(EmailUserLbl)
	Env.Email.Password = os.Getenv(EmailPasswordLbl)
	Env.Email.VacationRecipient = os.Getenv(VacationRecipientEmailLbl)
}
