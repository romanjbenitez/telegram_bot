package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	screaming = false
	bot       *tgbotapi.BotAPI
)

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {

	for {
		select {

		case <-ctx.Done():
			return
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	handleMessage(update.Message)
}

func handleMessage(message *tgbotapi.Message) {
	var err error
	user := message.From
	text := message.Text

	if user == nil {
		return
	}
	log.Println(message.Voice)

	envID, err := strconv.ParseInt(os.Getenv("user_id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	if user.ID != envID {
		_, _ = bot.Send(tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper("Toca de aca gato de leche")))
		return
	}

	log.Printf("%s wrote %s", user.FirstName, text)
	log.Println(user.ID)

	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else if screaming && len(text) > 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
		// To preserve markdown, we attach entities (bold, italic..)
		msg.Entities = message.Entities
		_, err = bot.Send(msg)
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func handleCommand(chatId int64, command string) error {
	var err error

	switch command {
	case "/scream":
		screaming = true
		break

	case EGRESO:
		screaming = false
		egresoCommand()
		break
	}
	return err
}
