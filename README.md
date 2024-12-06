# Telegram - A simple Telegram bot library for Go

Read this article to learn how to get your Telegram bot token and chat id:
[Setup Telegram bot to get alert notifications](https://medium.com/linux-shots/setup-telegram-bot-to-get-alert-notifications-90be7da4444)


Sending a message to your Telegram chat is as simple as

```go
package main

import (
	"fmt"

	"github.com/mxmlkzdh/telegram"
)

const (
	token  = "<YOUR-TELEGRAM-BOT-TOKEN>"
	chatID = "<YOUR-TELEGRAM-CHAT-ID>"
)

func main() {
	telegram := telegram.New(token, chatID, false)
	if err := telegram.Send("Hello, World!"); err != nil {
		fmt.Println(err)
	}
}

```
