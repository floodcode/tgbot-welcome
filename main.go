package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/floodcode/tgbot"
)

var bot tgbot.TelegramBot

type botConfig struct {
	Token string `json:"token"`
}

func main() {
	configData, err := ioutil.ReadFile("config.json")
	checkError(err)

	var config botConfig
	err = json.Unmarshal(configData, &config)
	checkError(err)

	bot, err = tgbot.New(config.Token)
	checkError(err)

	bot.Poll(tgbot.PollConfig{
		Delay:    100,
		Callback: updatesCallback,
	})
}

func updatesCallback(updates []tgbot.Update) {
	for _, update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.NewChatMembers != nil {
			processNewChatMembers(update.Message)
		}
	}
}

func processNewChatMembers(message *tgbot.Message) {
	if message.Chat.ID > 0 {
		return
	}

	for _, user := range message.NewChatMembers {
		bot.SendMessage(tgbot.SendMessageConfig{
			ChatID: tgbot.ChatID(message.Chat.ID),
			Text:   fmt.Sprintf("Welcome, %s!", user.FirstName),
		})
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
