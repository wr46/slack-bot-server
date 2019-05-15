package configuration

import (
	"geekbot/logger"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	SlackToken   string
	BotId        string
	BotChannelId string
	Debug        bool
}

var ENV Configuration

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
