package main

import (
	"fmt"
	"os"

	"gopkg.in/telebot.v3"
)

type AdminChat struct {
	chatId string
}

type AdminMsg struct {
	msgId  string
	chatId int64
}

func (m AdminMsg) MessageSig() (messageID string, chatID int64) {
	return m.msgId, m.chatId
}

func (r AdminChat) Recipient() string {
	return r.chatId
}

func getVid(c telebot.Context) error {
	session := createSession(c.Chat().ID)

	session.setType("video")
	session.setNextStep("video")

	return c.Send("Додайте відео і надайте короткий опис.")
}

func sendImg(c telebot.Context) error {
	return nil
}

func sendComplain(c telebot.Context) error {
	return nil
}

func sendContentRequest(c telebot.Context) error {
	return nil
}

func handleText(c telebot.Context) error {
	session := getSession(c.Chat().ID)

	switch session.nextStep {
	case "description":
		session.addData(c.Message().ID)
		session.setNextStep("date")
		return c.Send("Додайте дату коли було зроблено фото/відео.")
	case "date":
		session.addData(c.Message().ID)
		sendToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Данні відправлено.")
	}

	return nil
}

func handleVid(c telebot.Context) error {
	session := getSession(c.Chat().ID)

	if session.nextStep == "video" {
		session.addData(c.Message().ID)
		session.setNextStep("description")
		return c.Send("Додайте відео і надайте короткий опис.")
	}

	return nil
}

func handleImg(c telebot.Context) error {
	return nil
}

func sendToAdmin(data []int, chatId int64) {
	admChatId, _ := os.LookupEnv("OUTPUT_CHAT_ID")

	recipient := AdminChat{
		chatId: admChatId,
	}

	for _, msgId := range data {
		admMsg := AdminMsg{
			msgId:  fmt.Sprint(msgId),
			chatId: chatId,
		}

		b.Copy(recipient, admMsg)
	}
}
