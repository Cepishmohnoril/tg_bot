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

func buttonHandler(c telebot.Context, nStep string, msg string) error {
	session := NewSession()
	session.setNextStep(nStep)
	setSession(c.Chat().ID, session)

	return c.Send(msg)
}

func handleText(c telebot.Context) error {
	var (
		sesId   = c.Chat().ID
		session = getSession(sesId)
	)

	switch session.nextStep {
	case "description":
		session.addData(c.Message().ID)
		session.setNextStep("date")
		setSession(sesId, session)
		return c.Send("Додайте дату коли було зроблено фото/відео.")
	case "date":
		session.addData(c.Message().ID)
		forwardSessionDataToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Данні відправлено.")
	case "complain":
		session.addData(c.Message().ID)
		sendTextToAdmin("Скарга!")
		forwardSessionDataToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Данні відправлено.")
	case "suggestion":
		session.addData(c.Message().ID)
		sendTextToAdmin("Побажання.")
		forwardSessionDataToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Данні відправлено.")
	}

	return nil
}

func handleVid(c telebot.Context) error {
	var (
		sesId   = c.Chat().ID
		session = getSession(sesId)
	)

	if session.nextStep == "video" {
		session.addData(c.Message().ID)
		session.setNextStep("description")
		setSession(sesId, session)
		return c.Send("Надайте короткий опис.")
	}

	return nil
}

func handleImg(c telebot.Context) error {
	var (
		sesId   = c.Chat().ID
		session = getSession(sesId)
	)

	if session.nextStep == "image" {
		session.addData(c.Message().ID)
		session.setNextStep("description")
		setSession(sesId, session)
		return c.Send("Надайте короткий опис.")
	}

	return nil
}

func sendTextToAdmin(text string) {
	recipient := getAdminRecipient()
	b.Send(recipient, text)

}

func forwardSessionDataToAdmin(data []int, chatId int64) {
	recipient := getAdminRecipient()

	for _, msgId := range data {
		admMsg := AdminMsg{
			msgId:  fmt.Sprint(msgId),
			chatId: chatId,
		}

		b.Copy(recipient, admMsg)
	}
}

func getAdminRecipient() telebot.Recipient {
	admChatId, _ := os.LookupEnv("OUTPUT_CHAT_ID")

	return AdminChat{
		chatId: admChatId,
	}
}
