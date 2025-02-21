package com.example.DeepSeekBot;

import com.example.DeepSeekBot.config.BotConfig;
import com.example.DeepSeekBot.service.DeepSeekService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.telegram.telegrambots.bots.TelegramLongPollingBot;
import org.telegram.telegrambots.meta.api.methods.send.SendMessage;
import org.telegram.telegrambots.meta.api.objects.Update;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;

import java.io.IOException;
import java.text.ParseException;

@Component
public class TelegramBot extends TelegramLongPollingBot {
    private final BotConfig botConfig;
    private final DeepSeekService deepSeekService;

    @Autowired
    public TelegramBot(BotConfig botConfig, DeepSeekService deepSeekService) {
        this.botConfig = botConfig;
        this.deepSeekService = deepSeekService;
    }

    @Override
    public String getBotUsername() {
        return botConfig.getBotName();
    }

    @Override
    public String getBotToken() {
        return botConfig.getToken();
    }

    @Override
    public void onUpdateReceived(Update update) {
        if(update.hasMessage() && update.getMessage().hasText()){
            String messageText = update.getMessage().getText();
            long chatId = update.getMessage().getChatId();

            switch (messageText){
                case "/start":

                    startCommandReceived(chatId, update.getMessage().getChat().getFirstName());
                    break;
                default:
                    try {
                        String answer = deepSeekService.getAnswer(messageText);
                        sendMessage(chatId, answer);

                    } catch (IOException e) {
                        sendMessage(chatId, "Not found");
                    } catch (ParseException e) {
                        throw new RuntimeException("Unable to parse date");
                    }
            }
        }

    }

    private void startCommandReceived(Long chatId, String name) {
        String answer = "Привет, " + name + "\n" +
                "Задай свой вопрос";
                sendMessage(chatId, answer);
    }

    private void sendMessage(Long chatId, String textToSend){
        SendMessage message = new SendMessage();
        message.setChatId(chatId);
        message.setText(textToSend);
        message.enableMarkdown(true);
        try {
            execute(message);
        } catch (TelegramApiException e) {

        }
    }
}
