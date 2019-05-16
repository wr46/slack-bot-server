# slack-bot-server

A simple slack bot server template in Go for messaging

## Introduction

This project is a simple [Go](https://golang.org/) server application for a Slack Bot. \
The main goal here was to implement a Bot that answer to users in same workspace, by direct chat (user<>bot) \
or mention the bot in workspace channels.

## Usage

In order for the server to work, you need to configure the config.json file. (./geekbot/configuration/config.development.json)

```javascript
{
    "SlackToken" : "YOUR SLACK TOKEN HERE",
    "BotId" : "YOUR SLACK BOT ID HERE",
    "Debug" : false
}
```

- **SlackToken** can be found after bot integration in your Slack workspace
- **BotId** can be found simply by execute this server in debug mode and chat directly with the bot

**For the next step it's required** [golang/dep](https://github.com/golang/dep) \
Run the following command to setup dependencies:

```bash
dep ensure
```

## Useful documentation

- [dep](https://github.com/golang/dep) (dependency manager)
- [Slack API in Go](https://github.com/nlopes/slack)
- [Gonfig](https://github.com/tkanos/gonfig) (JSON configs and enviornment variables)

## License

slack-bot-server is licensed by an MIT license as can be found in the [LICENSE](https://github.com/wr46/slack-bot-server/blob/master/LICENSE) file.