package whenPutlerDies

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func WhenPutlerDies(tgBot *tgbotapi.BotAPI, update tgbotapi.Update) {

	//command := strings.Split(update.Message.Text, " ")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, " NuLL ")

	randSource := rand.NewSource(time.Now().UnixNano())
	randValue := rand.New(randSource)
	deathTime := randValue.Intn(100) //creating random int up to 100

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(" путін сдохне через %d секунд", deathTime))
	//msg.ReplyToMessageID = update.Message.MessageID
	tgBot.Send(msg)

}
