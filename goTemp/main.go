package main

import (
	"encoding/json"
	"fmt"
	"line_local/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	fmt.Println("Beginning to connect...")

	// Get environment variables
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CHANNEL_TOKEN := envs["CHANNEL_TOKEN"]
	CHANNEL_SECRET := envs["CHANNEL_SECRET"]
	PORT := envs["PORT"]
	fmt.Println("Got these envs ", CHANNEL_TOKEN, " ", CHANNEL_SECRET)
	// Initialize LINE BOT
	bot, err := linebot.New(
		CHANNEL_SECRET,
		CHANNEL_TOKEN,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Handler
	messageHandler := service.NewHandler()

	// Setup HTTP Server for receiving requests from LINE platform
	fmt.Println("Setting up HTTP Server")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		type Response struct {
			Message string `json:"message"`
		}
		// Create a response object
		response := Response{
			Message: "Successful GET request!",
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Encode the response object as JSON
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Fatal(err)
		}
	})

	// Our route for Line OA webhook URL
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					response := messageHandler.MapResponse(message.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, *response).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	// TODO: For actual use, we must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	fmt.Println("Listening and serving on port ", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
