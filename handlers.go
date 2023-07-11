package main

import "gopkg.in/telebot.v3"

func getVid(c telebot.Context) error {
	//c.Chat().ID
	//Create session
	//https://pkg.go.dev/github.com/go-session/session
	//Middleware checks session
	//If session exists, making next step.
	//Else ignore

	//Steps
	//Request video
	//.onVideo <- Request description
	//.onText <- Request date
	//.onText <- Create message and send to admin chat

	return c.Send("Додайте відео і надайте короткий опис.")
}

//func sendVid(c telebot.Context) error {
//	return nil
//}

func sendImg(c telebot.Context) error {
	return nil
}

func sendComplain(c telebot.Context) error {
	return nil
}

func sendContentRequest(c telebot.Context) error {
	return nil
}
