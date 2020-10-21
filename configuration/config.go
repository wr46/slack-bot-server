package configuration

import (
	"os"
	"strconv"

	"github.com/wr46/slack-bot-server/logger"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Dotenv file variables labels
const (
	DebugLbl                  = "DEBUG"
	I18nTagLbl                = "I18N_TAG"
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

// Configuration container.
type Configuration struct {
	Debug       bool
	I18nPrinter *message.Printer
	Slack       SlackEnv
	Email       EmailEnv
}

// Env environment data.
var Env Configuration

// Setup configuration.
func Setup() {
	// Initialize logger
	logger.Setup(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	setupLanguage()
	setupApplication()
	setupSlack()
	setupEmail()

	logger.SetDebug(Env.Debug)
	logger.Log(logger.Debug, "Environment is ready!")
}

func setupLanguage() {
	tag := language.Make(os.Getenv(I18nTagLbl))
	if tag == language.Und {
		tag = language.English

		logger.Log(logger.Warning, "Unknown language! Default English used!")
	}

	Env.I18nPrinter = message.NewPrinter(tag)
}

func setupApplication() {
	isDebug, err := strconv.ParseBool(os.Getenv(DebugLbl))
	if err != nil {
		logger.Log(logger.Error, err.Error())
	}

	Env.Debug = isDebug

	logger.Log(logger.Debug, "Debug mode: "+strconv.FormatBool(Env.Debug))
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
