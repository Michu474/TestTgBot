package main

import (
	"log"
	killputler "https://github.com/Michu474/GoTgBot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5406481213:AAG8UZ1wNMB27VLU-IIQNIUHY4Q-dX_6Lxw") // creating new bot
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	botUpdate := tgbotapi.NewUpdate(0)
	botUpdate.Timeout = 60000 // bot remaining time until it stops

	updates, err := bot.GetUpdatesChan(botUpdate)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text) //logging message data

			WhenPutlerDies(bot, update)

		}
	}
}
