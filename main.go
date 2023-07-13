package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

var b *telebot.Bot
var (
	// Universal markup builders.
	menu = &telebot.ReplyMarkup{ResizeKeyboard: true, ForceReply: true}

	// Reply buttons.
	btnSendVid  = menu.Text("Надіслати відео і опис")
	btnSendImg  = menu.Text("Надіслати фото і опис")
	btnComplain = menu.Text("Поскаржитися на контент")
	btnRequest  = menu.Text("Що цікаво знати більше про війну?")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
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

	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("Вітаю!", menu)
	})

	b.Handle(&btnSendVid, getVid)
	b.Handle(&btnSendImg, sendImg)
	b.Handle(&btnComplain, sendComplain)
	b.Handle(&btnRequest, sendContentRequest)

	b.Handle(telebot.OnText, handleText, checkSession)
	b.Handle(telebot.OnVideo, handleVid, checkSession)
	b.Handle(telebot.OnPhoto, handleImg, checkSession)

	b.Start()

}

func initMenu() {
	menu.Reply(
		menu.Row(btnSendVid),
		menu.Row(btnSendImg),
		menu.Row(btnComplain),
		menu.Row(btnRequest),
	)
}
