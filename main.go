package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

var (
	b *telebot.Bot
	// Universal markup builders.
	menu = &telebot.ReplyMarkup{ResizeKeyboard: true, ForceReply: true}

	// Reply buttons.
	btnSendVid  = menu.Text("Прислати відео")
	btnSendImg  = menu.Text("Прислати фото")
	btnFeedback = menu.Text("Прислати ідею, пропозицію чи скаргу")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	initSessionStorage()
	token, tokenExists := os.LookupEnv("TOKEN")

	if !tokenExists || token == "" {
		panic("Telegram API token is absent.")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	b, _ = telebot.NewBot(pref)

	initMenu()

	b.Handle("/start", handleStart, checkIsAdminChat)

	b.Handle(&btnSendVid, func(c telebot.Context) error {
		return buttonHandler(c, "video", "Завантаж відео")
	}, checkIsAdminChat)
	b.Handle(&btnSendImg, func(c telebot.Context) error {
		return buttonHandler(c, "image", "Завантаж фото")
	}, checkIsAdminChat)
	b.Handle(&btnFeedback, func(c telebot.Context) error {
		return buttonHandler(c, "feedback", "Додайте ідею, пропозицію чи скаргу.")
	}, checkIsAdminChat)

	b.Handle(telebot.OnText, handleText, checkIsAdminChat, checkSession)
	b.Handle(telebot.OnVideo, handleVid, checkIsAdminChat, checkSession)
	b.Handle(telebot.OnPhoto, handleImg, checkIsAdminChat, checkSession)

	b.Start()

}

func initMenu() {
	menu.Reply(
		menu.Row(btnSendVid),
		menu.Row(btnSendImg),
		menu.Row(btnFeedback),
	)
}
