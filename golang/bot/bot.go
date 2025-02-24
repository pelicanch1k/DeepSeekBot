package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github/pelicanch1k/DeepSeekBot/service"
)

func Run(botAPI *tgbotapi.BotAPI, messageService *service.DeepSeekService) {
	log.Print("Бот запущен!")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1

	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет")
				_, err := botAPI.Send(msg)
				if err != nil {
					log.Println("Ошибка при отправке сообщения:", err)
				}
			} else {
				// Передаем сообщение в контроллер
				answer, err := messageService.Request(update.Message.Text)
				if err != nil {
					log.Println("Ошибка при обработке сообщения:", err)

					sendMsg(botAPI, update, "Ошибка на сервере")
					continue
				} else {
					sendMsg(botAPI, update, answer)
				}
			}
		}
	}
}

func sendMsg(botAPI *tgbotapi.BotAPI, update tgbotapi.Update, msg string) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	_, err := botAPI.Send(message)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}
