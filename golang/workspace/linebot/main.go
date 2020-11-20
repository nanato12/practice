package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

// EnvLoad load .env
func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	EnvLoad()

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		// botのインスタンス化
		bot, err := linebot.New(
			os.Getenv("CHANNEL_SECRET"),
			os.Getenv("CHANNEL_ACCESS_TOKEN"),
		)
		// reqをパースしてevents取得
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		// eventごとに処理を書く
		for _, event := range events {
			log.Print(event)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if event.ReplyToken == "00000000000000000000000000000000" {
						return
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
