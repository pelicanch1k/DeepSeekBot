package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DeepSeekService struct {
	client *http.Client
}

func InitDeepSeekService() *DeepSeekService {
	return &DeepSeekService{client: &http.Client{}}
}

func (s *DeepSeekService) Request(userMessage string) (string, error) {
	requestBody, err := json.Marshal(map[string]any{
		"model": "deepseek/deepseek-r1:free",
		"messages": []map[string]string{{
			"role":    "user",
			"content": userMessage,
		}},
	})
	if err != nil {
		return "", fmt.Errorf("ошибка при маршалинге JSON: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("URL"), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("JWT_TOKEN"))

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	if resp.StatusCode == 200 {
		// Преобразование JSON в мапу
		var result map[string]interface{}
		if json.Unmarshal(body, &result) != nil {
			return "", fmt.Errorf("ошибка при парсинге JSON: %v", err)
		}

		// Достаем строку по пути ["choices"][0]["message"]["content"]
		content, err := getAnswerFromJSON(result, "choices", 0, "message", "content")
		if err != nil {
			return "", fmt.Errorf("ошибка: %v", err)
		}

		return content, nil
	}

	log.Print(string(body))

	return "Ошибка на сервере", nil
}

// getStringFromJSON извлекает ответ из вложенной структуры JSON
func getAnswerFromJSON(data map[string]interface{}, keys ...interface{}) (string, error) {
	var current interface{} = data

	for _, key := range keys {
		switch k := key.(type) {
		case string:
			// Если ключ — строка, предполагаем, что это мапа
			m, ok := current.(map[string]interface{})
			if !ok {
				return "", fmt.Errorf("ожидалась мапа, но получен другой тип")
			}
			current = m[k]
		case int:
			// Если ключ — число, предполагаем, что это срез
			s, ok := current.([]interface{})
			if !ok {
				return "", fmt.Errorf("ожидался срез, но получен другой тип")
			}
			if k < 0 || k >= len(s) {
				return "", fmt.Errorf("индекс %d вне диапазона", k)
			}
			current = s[k]
		default:
			return "", fmt.Errorf("неподдерживаемый тип ключа: %T", key)
		}
	}

	// Проверяем, что итоговое значение — строка
	str, ok := current.(string)
	if !ok {
		return "", fmt.Errorf("ожидалась строка, но получен другой тип")
	}

	return str, nil
}
