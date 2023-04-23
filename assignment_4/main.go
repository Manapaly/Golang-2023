package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	unsplashAccessKey = os.Getenv("UNSPLASH_ACCESS_KEY")
	telegramBotToken  = os.Getenv("TELEGRAM_BOT_TOKEN")
)

func getRandomImage(count int) string {
	url := "https://api.unsplash.com/photos/random?client_id=" + unsplashAccessKey + "&count=" + strconv.Itoa(count)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching random image from Unsplash API: ", err)
	}
	defer resp.Body.Close()

	var images []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&images)

	image := images[0]["urls"].(map[string]interface{})["regular"].(string)

	return image
}

func handleImage(update tgbotapi.Update) {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal("Error initializing Telegram bot: ", err)
	}
	chatID := update.Message.Chat.ID
	message := update.Message.Text

	if strings.Contains(message, "/image") || strings.Contains(message, "image") {
		image := getRandomImage(1)

		photo := tgbotapi.NewPhotoUpload(chatID, image)
		bot.Send(photo)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal("Error initializing Telegram bot: ", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal("Error starting bot updates: ", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		handleImage(update)
	}
}
