package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

var (
	// Universal markup builders.
	menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

	// Reply buttons.
	btnSendVid = menu.Text("Надіслати відео і опис")
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

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	initMenu()

	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello!", menu)
	})

	b.Handle(&btnSendVid, sendVid)

	b.Start()

}

func initMenu() {
	menu.Reply(
		menu.Row(btnSendVid),
	)
}

func sendVid(c telebot.Context) error {
	return c.Send("Відео надіслано")
}
