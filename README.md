# FrodoBot

[![Go Report Card](https://goreportcard.com/badge/github.com/rimonmostafiz/frodobot)](https://goreportcard.com/report/github.com/rimonmostafiz/frodobot)

A harmless slack bot that reminds you to post your daily status 🤖

## Give a Star! ⭐
If you like this project please consider giving it a ⭐ star ⭐. Thanks!

## Prerequisite
- Create a slack app
![create_app](https://a.slack-edge.com/80588/img/api/articles/bolt/config_create_app.png)
- Install the app to your workspace once to obtain your OAuth token. You need to store the bot token(`xoxb-...`) in your `.env` file
![install_app](https://a.slack-edge.com/80588/img/api/articles/bolt/config_install.png)
- Specify the OAuth scopes.
![auth_permission](https://a.slack-edge.com/d0cef/img/api/articles/bolt/config_scopes.png)
- List of OAuth Scopes
    - *chat.write* - Enable the bot to send messages
    - *groups.history* - View messages and other content in private channels 
    - *groups.read* - View basic information about private channels
    - *users.info* - Gets information about a user.

## Dependencies 
- go 1.14
- Docker
- `.env`

## Configurations
```
SLACK_TOKEN=YOUR_SLACK_TOEKN
CHANNEL_ID=G******X
EXCLUDE_USER_VALUES="{\"U******1\":\"x\",\"U******2\":\"x\"}"
```
`.env` file should contain this 3 key

- SLACK_TOKEN - your OAuth token
- CHANNEL_ID - channel where you want to post the reminder message
- EXCLUDE_USER_VALUES - a map with `userId` as key and `x` as value, while sending message bot will not mention those users.

## Instructions
- Clone this repository
- Add a `.env` file

### Build
- `go mod download`
- `go build`

### Run
- `./frodobot`

### Deployment
- `docker build -t frodobot:1.0.0 .`
- `docker container run -d --name frodobot frodobot:1.0.0` 
