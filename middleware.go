package main

import (
	"fmt"
	"os"

	"gopkg.in/telebot.v3"
)

func checkIsAdminChat(next telebot.HandlerFunc) telebot.HandlerFunc {

	return func(c telebot.Context) error {
		chatId := c.Chat().ID
		admChatId, _ := os.LookupEnv("OUTPUT_CHAT_ID")

		//Ignore output chat
		if admChatId == fmt.Sprint(chatId) {
			return c.Respond()
		}

		return next(c)
	}
}

func checkSession(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		chatId := c.Chat().ID

		if sessionExists(chatId) {
			return next(c)
		}

		return c.Send("Оберіть дію.")
	}
}
