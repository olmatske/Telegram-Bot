package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)
// Send any text message to the bot after the bot has been started
// 'token:=' is the assignment of the correct type, otherwise it would be e.g. 'token string = ...'

var userState = make(map[int64]string)

func main() {
	token:= os.Getenv("TOKEN")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	
	if update.Message == nil {
		return
	}

	chatID:= update.Message.Chat.ID
	text:=    update.Message.Text

	switch userState[chatID]{
	case "":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text: "Hi, what's your Name?",
		})
		userState[chatID] = "await.name"

	case "await.name":
		name := text
		if name == "Noel" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text: "Hope you're having a great day, I love you! ❤️",
			})
		} else if name == "Jonas" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text: "Servus Tudn 😎",
			})
		} else if name == "Milena" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text: "Wanna grab a coffee sometime? ☕️",
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text: "Hi " + name + "!",
			})
		}
		userState[chatID] = ""
	}
}


	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   "",
	// })
// }



// package main

// import "fmt"

// func main() {
// 	fmt.Println("Welcome")
// }
