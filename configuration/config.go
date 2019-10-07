package configuration

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
	"github.com/wr46/slack-bot-server/logger"
)

// Configuration container
type Configuration struct {
	SlackToken   string
	BotID        string
	BotChannelID string
	Debug        bool
}

// ENV environment data
var ENV Configuration

// Setup configuration
func Setup() {

	// Initialize logger
	logger.Setup(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)

	if err != nil {
		logger.Log(logger.Error, err.Error())
	}
	logger.SetDebug(configuration.Debug)
	ENV = configuration
	logger.Log(logger.Info, "Environment is ready!")
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"./", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
