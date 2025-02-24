package main

import (
	"github/pelicanch1k/DeepSeekBot/bot"
	"github/pelicanch1k/DeepSeekBot/service"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	// Инициализация бота
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Print(os.Getenv("BOT_TOKEN"))
		log.Panic(err)
	}

	// Инициализация сервиса для HTTP-запросов
	httpService := service.InitDeepSeekService()

	// Запуск бота
	bot.Run(botAPI, httpService)
}
