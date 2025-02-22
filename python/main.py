from dotenv import load_dotenv
import os

import asyncio
from aiogram import Bot, Dispatcher

from app.handlers import router

load_dotenv()

async def main():
    bot = Bot(token=os.getenv("BOT_TOKEN"))
    dp = Dispatcher()
    dp.include_router(router)
    await dp.start_polling(bot)


if __name__ == '__main__':
    try:
        print('Бот включен')
        asyncio.run(main())
    except KeyboardInterrupt:
        print('Бот выключен')