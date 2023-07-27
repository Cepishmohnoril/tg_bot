package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

type AdminChat struct {
	chatId string
}

type AdminMsg struct {
	msgId  string
	chatId int64
}

type AdminLib struct {
}

func (m AdminMsg) MessageSig() (messageID string, chatID int64) {
	return m.msgId, m.chatId
}

func (r AdminChat) Recipient() string {
	return r.chatId
}

func buttonHandler(c telebot.Context, nStep string, msg string) error {
	session := NewSession(c.Chat().ID)
	session.setNextStep(nStep)

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
		forwardSessionDataToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Дякую! Я все отримав. Щиро твій, бот Нацгвардії.")
	case "feedback":
		session.addData(c.Message().ID)
		sendTextToAdmin("Відгук!")
		forwardSessionDataToAdmin(session.data, c.Chat().ID)
		terminateSession(c.Chat().ID)
		return c.Send("Дякую! Я все отримав. Щиро твій, бот Нацгвардії.")
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
		c.Send("Додай дату, коли це було і короткий опис подій.")
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

		if !session.waiting {
			session.waiting = true
			time.Sleep(5 * time.Second)
			session.setNextStep("description")
			c.Send("Додай дату, коли це було і короткий опис подій.")
		}
	}

	return nil
}

func handleStart(c telebot.Context) error {
	return c.Send("Вітаю!", menu)
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

		b.Forward(recipient, admMsg)
	}
}

func getAdminRecipient() telebot.Recipient {
	admChatId, _ := os.LookupEnv("OUTPUT_CHAT_ID")

	return AdminChat{
		chatId: admChatId,
	}
}
