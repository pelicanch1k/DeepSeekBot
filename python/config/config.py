import os


class BotConfig:
    __botToken: str
    __jwtToken: str
    __url: str

    def __init__(self) -> None:
        self.__botToken = os.getenv("BOT_TOKEN")
        self.__jwtToken= os.getenv("JWT_TOKEN")
        self.__url = os.getenv("URL")

    def get_bot_token(self) -> str:
        return self.__botToken
    
    def get_jwt_token(self) -> str:
        return self.__jwtToken
    
    def get_url(self) -> str:
        return self.__url