# slack-bot-server
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

A simple slack bot server template in Go for messaging

## Introduction

This project is a simple [Go](https://golang.org/) server application for a Slack Bot. \
The main goal here was to implement a Bot that answer to users in same workspace, by direct chat (user<>bot) \
or mention the bot in workspace channels.

## Implemented features

- **Show all available commands**
- **Do a vacation request**
- **Tell a joke**
- **Run a local script**

## Usage

In order for the server to work, you need to setup a '.env' file, you can use 'example.env' as a template.

```javascript
# NGROK
NGROK_AUTH=YOUR NGROK TOKEN
NGROK_PORT=botapp:3000
NGROK_REGION=eu

# Application
DEBUG=true

# Slack
BOT_USER_OAUTH_ACCESS_TOKEN=YOUR BOT USER OAUTH ACCESS TOKEN
VERIFICATION_TOKEN=YOUR APP VERIFICATION TOKEN

# Email (Gmail default)
EMAIL_SMTP_SERVER=smtp.gmail.com
EMAIL_SMTP_PORT=587
EMAIL_USER=email@gmail.com
EMAIL_PASSWORD=YOUR EMAIL PASSWORD
VACATION_RECIPIENT_EMAIL=EMAIL FOR VACATION REQUEST
```

- **Ngrok Token** you must create a Ngrok account so you can have a private url and add it to Slack "Event Subscriptions" request URL. \

```bash
    ex: "https://NGROK_PUBLIC_URL_HERE/events-endpoint" (check Ngrok section below)
```

- **Bot User OAuth Token** can be found after you create an app on Slack workspace
- **Verification Token** can be found on "Basic Information" -> "App Credentials" of your Slack App

**Note:** \
The email configuration was intended to be used as a vacation request process. Change it to suit your needs.

## Docker

**Docker build and run** 

```bash
docker-compose up
```

## Ngrok

Get the Ngrok public URL using the following URL

```bash
http://localhost:4040/api/tunnels
```

## Useful documentation

- [dep](https://github.com/golang/dep) (dependency manager)
- [Slack API in Go](https://github.com/slack-go/slack)

## ❗ Code of Conduct

We expect everyone to abide by our [**Code of Conduct**](./CODE_OF_CONDUCT.md). Please read it. 🤝

## 🙌🏻 How to Contribute

Check out our [**Contributing Guide**](./CONTRIBUTING.md) for information on contributing.

## License

slack-bot-server is licensed by an MIT license as can be found in the [LICENSE](./LICENSE) file.

## 💜 Thanks

Thanks to everyone contributing in any manner for this repo and we extend this thanks to everyone working on Open Source.

**Sharing is Caring**

## Contributors ✨
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://pt.linkedin.com/in/wilsonmarodrigues"><img src="https://avatars0.githubusercontent.com/u/5550776?v=4" width="100px;" alt=""/><br /><sub><b>Wilson Rodrigues</b></sub></a><br /><a href="https://github.com/wr46/slack-bot-server/commits?author=wr46" title="Code">💻</a> <a href="#maintenance-wr46" title="Maintenance">🚧</a></td>
    <td align="center"><a href="http://people.virginia.edu/~ys4ea"><img src="https://avatars0.githubusercontent.com/u/6952923?v=4" width="100px;" alt=""/><br /><sub><b>Yosyp Schwab</b></sub></a><br /><a href="https://github.com/wr46/slack-bot-server/commits?author=yosyp" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->
