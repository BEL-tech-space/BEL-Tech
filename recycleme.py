from telegram import InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext import Application, CommandHandler, CallbackQueryHandler
from pymongo import MongoClient
import random
import string
import logging

# Set up logging to track errors
logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s', level=logging.INFO)
logger = logging.getLogger(__name__)

# Replace 'YOUR_BOT_TOKEN' with your actual bot token
TOKEN = '7526280917:AAFF-610OuJybkwy6wamiuHIZ1szXllLOnk'

# MongoDB connection
client = MongoClient('mongodb://localhost:27017/')
db = client['waste_collection_db']
users_collection = db['users']

# Async function to handle the /start command
async def start(update, context):
    keyboard = [
        [InlineKeyboardButton("Deposit Waste", callback_data='deposit')],
        [InlineKeyboardButton("Balance", callback_data='balance')],
        [InlineKeyboardButton("Withdraw", callback_data='withdraw')]
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await update.message.reply_text("Please choose:", reply_markup=reply_markup)

# Async function to handle button presses
async def button(update, context):
    query = update.callback_query
    await query.answer()

    if query.data == 'deposit':
        await query.edit_message_text(text="Select your location: UI, Unilag, LASU")
        # Implement location selection here
    elif query.data == 'balance':
        user_id = query.message.chat_id
        user = users_collection.find_one({'user_id': user_id})
        balance = user.get('waste_deposited', 0) if user else 0
        await query.edit_message_text(text=f"Your total balance is {balance} KG.")
    elif query.data == 'withdraw':
        await query.edit_message_text(text="Withdrawal not available yet.")
        # Implement withdrawal logic here

# Function to generate a random verification code
def generate_code():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=6))

# Async function to handle /deposit command
async def deposit(update, context):
    user_id = update.message.chat_id
    try:
        amount = float(context.args[0])
    except (IndexError, ValueError):
        await update.message.reply_text("Please provide a valid amount of waste to deposit.")
        return

    code = generate_code()

    # Store deposit and code in the database
    users_collection.update_one(
        {'user_id': user_id},
        {'$inc': {'waste_deposited': amount}, '$set': {'code': code}},
        upsert=True
    )

    # Send code to user
    await update.message.reply_text(f"{amount} KG of waste deposited! Verification code: {code}")

# Main function to start the bot
async def main():
    # Create the application and pass the bot token
    application = Application.builder().token(TOKEN).build()

    # Add handlers
    application.add_handler(CommandHandler("start", start))
    application.add_handler(CallbackQueryHandler(button))
    application.add_handler(CommandHandler("deposit", deposit))

   # Main function to start the bot
async def main():
    # Create the application and pass the bot token
    application = Application.builder().token(TOKEN).build()

    # Add handlers
    application.add_handler(CommandHandler("start", start))
    application.add_handler(CallbackQueryHandler(button))
    application.add_handler(CommandHandler("deposit", deposit))

    # Start polling for updates
    await application.start()

    # Block until you send a stop signal
    await application.idle()

# Start the bot
if __name__ == '__main__':
    import asyncio
    asyncio.run(main())
