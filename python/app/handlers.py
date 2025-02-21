from aiogram import F, Router
from aiogram.types import Message, CallbackQuery
from aiogram.filters import CommandStart, Command
from aiogram.fsm.state import State, StatesGroup
from aiogram.fsm.context import FSMContext

router = Router()

@router.message(CommandStart())
async def cmd_start(message: Message):
    await message.reply(f"Привет, {message.from_user.first_name}\nЗадай свой вопрос")

@router.message()
async def question_from_user(message: Message):
    pass