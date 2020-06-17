# slack-bot-server

A simple slack bot server template in Go for messaging

## Introduction

This project is a simple [Go](https://golang.org/) server application for a Slack Bot. \
The main goal here was to implement a Bot that answer to users in same workspace, by direct chat (user<>bot) \
or mention the bot in workspace channels.

## Usage

In order for the server to work, you need to configure the config.json file. (./configuration/config.development.json)

```javascript
{
    "SlackToken" : "YOUR SLACK TOKEN HERE",
    "BotID" : "YOUR SLACK BOT ID HERE",
    "EmailSMTPServer": "SMTP SERVER",
    "EmailSMTPPort": "SMTP PORT",
    "EmailUser": "YOUR SENDER EMAIL",
    "EmailPassword": "YOUR SENDER EMAIL PASSWORD",
    "VacationRecipientEmail": "YOUR RECIPIENT EMAIL FOR VACATION REQUEST",
    "Debug" : false
}
```

- **SlackToken** can be found after bot integration in your Slack workspace website
- **BotID** can be found simply by execute this server in debug mode and chat directly with the bot

**Note:** \
The email configuration was intended to be used as a vacation request process. \
Change it to suit your needs.

## Docker

**Docker build and run**
Run the following command:

```bash
docker-compose up
docker build -t slack-bot-server . 
```

```bash
docker run --name slack-bot-server -d slack-bot-server
```

http://localhost:4040/api/tunnels

**Usefull commands:**

- List docker containers

```bash
docker ps
```

- Stop one or more running docker containers

```bash
docker stop slack-bot-server
```

- Remove one or more docker containers

```bash
docker rm slack-bot-server
```

## Useful documentation

- [dep](https://github.com/golang/dep) (dependency manager)
- [Slack API in Go](https://github.com/nlopes/slack)
- [Gonfig](https://github.com/tkanos/gonfig) (JSON configs and enviornment variables)

## License

slack-bot-server is licensed by an MIT license as can be found in the [LICENSE](https://github.com/wr46/slack-bot-server/blob/master/LICENSE) file.