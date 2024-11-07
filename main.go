package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOTFATHERTOKEN"))
	if err != nil {
		log.Fatal("Failed to create new telegram bot from token:", err)
	}

	log.Printf("Bot started! Authorized on account %s", bot.Self.UserName)

	buyButton := tgbotapi.NewInlineKeyboardButtonData("Buy 📈", "EXACT_INPUT")
	sellButton := tgbotapi.NewInlineKeyboardButtonData("Sell 🤑", "EXACT_OUTPUT")
	buySellKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{buyButton, sellButton},
	)

	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	// the chatID for @uniswapx_tg_bot is 1910201650.

	updates := bot.GetUpdatesChan(updatesConfig)

	for update := range updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Text {
			case "/start", "/quote":
				msg.ReplyMarkup = buySellKeyboard
				msg.Text = "Do you want to buy or sell?"
			default:
				msg.Text = "Can you please write /start or /quote to begin :)"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Uh oh! You have an error which is %v", err)
			}
		}
	}

}

// type SwapRequest struct {
// 	Direction swapDirection `json:"type"`
// 	Amount    float64       `json:"amount,string"`
// 	TokenIn   string        `json:"tokenIn"`
// 	TokenOut  string        `json:"tokenOut"`
// }

// type swapDirection string

// const (
// 	EXACT_INPUT  swapDirection = "EXACT_INPUT"
// 	EXACT_OUTPUT swapDirection = "EXACT_OUTPUT"
// )
