from aiogram import F, Router
from aiogram.types import Message, CallbackQuery
from aiogram.filters import CommandStart, Command
from aiogram.fsm.state import State, StatesGroup
from aiogram.fsm.context import FSMContext

from .deep_seek_service import DeepSeekService

router = Router()

@router.message(CommandStart())
async def cmd_start(message: Message):
    await message.reply(f"Привет, {message.from_user.first_name}\nЗадай свой вопрос")

@router.message()
async def question_from_user(message: Message):
    print("user message: " + message.text)
    answer = DeepSeekService().find_answer(message.text)
    await message.reply(answer)