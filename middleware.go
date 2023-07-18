package main

import (
	"gopkg.in/telebot.v3"
)

func checkSession(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {

		chatId := c.Chat().ID

		if sessionExists(chatId) {
			return next(c)
		}

		return c.Send("Оберіть дію.")
	}
}
