package com.example.DeepSeekBot.service;

import com.example.DeepSeekBot.config.BotConfig;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class DeepSeekService {
    private final BotConfig botConfig;
    private static final String url = "https://openrouter.ai/api/v1/chat/completions";

    @Autowired
    public DeepSeekService(BotConfig botConfig) {
        this.botConfig = botConfig;
    }

    public String getAnswer(String userMessage) throws IOException, ParseException {
        RestTemplate restTemplate = new RestTemplate();

        HttpEntity<Map<String, Object>> httpEntity = new HttpEntity<>(getRequestBody(userMessage), getHeaders());

        String responce = restTemplate.postForObject(url, httpEntity, String.class);

        JsonNode jsonNode = new ObjectMapper().readTree(responce);

        String answer = "Ошибка, попробуйте еще раз";

        try {
            answer = jsonNode.get("choices").get(0).get("message").get("content").toString();
            answer = formatAnswer(answer); // Форматируем ответ
        } catch (NullPointerException ignored) {}

        return answer;
    }

    private Map<String, Object> getRequestBody(String userMessage){
        // Создаем тело запроса с помощью Map
        Map<String, Object> requestBody = new HashMap<>();
        requestBody.put("model", "deepseek/deepseek-r1:free");

        List<Map<String, String>> messages = new ArrayList<>();
        Map<String, String> message = new HashMap<>();
        message.put("role", "user");
        message.put("content", userMessage);
        messages.add(message);

        requestBody.put("messages", messages);

        return requestBody;
    }

    private HttpHeaders getHeaders(){
        HttpHeaders headers = new HttpHeaders();

        headers.setContentType(MediaType.APPLICATION_JSON);
        headers.add("Authorization", "Bearer " + botConfig.getJwtToken());

        return headers;
    }

    private String formatAnswer(String answer) {
        // Удаляем лишние кавычки
        answer = answer.replace("\"", "");

        // Заменяем escape-последовательности на переносы строк
        answer = answer.replace("\\n", "\n");

        // Добавляем переносы строк после знаков препинания
        answer = answer.replace(". ", ".\n");
        answer = answer.replace("! ", "!\n");
        answer = answer.replace("? ", "?\n");

        // Удаляем HTML-теги, если они есть
        answer = answer.replaceAll("<[^>]*>", "");

        return answer;
    }
}
