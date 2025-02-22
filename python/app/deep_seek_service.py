import json
from pprint import pprint

import requests
from config.config import BotConfig


class DeepSeekService:
    __bot_config: BotConfig

    def __init__(self):
        self.__bot_config = BotConfig()

    def find_answer(self, user_message: str) -> str:
        r: dict  = requests.post(self.__bot_config.get_url(), 
                        data=self.__get_request_body(user_message),
                        headers=self.__get_headers()
                        )
        
        answer:str

        try:
            answer = r.json()["choices"][0]["message"]["content"]
        except Exception as e:
            answer = "Произошла ошибка"

        print("bot answer: " + answer)
        return answer

    def __get_headers(self) -> dict[str:str]:
        return {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + self.__bot_config.get_jwt_token()
        }

    def __get_request_body(self, user_message: str):
        return json.dumps({
            "model": "deepseek/deepseek-r1:free",
            "messages": [
                {
                    "role": "user",
                    "content": user_message
                }
            ]})