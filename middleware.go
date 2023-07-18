package main

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func checkSession(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {

		chatId := c.Chat().ID

		if sessionExists(chatId) {
			fmt.Printf("%v+ \n", getSession(chatId))
			return next(c)
		}

		return c.Respond()
	}
}
