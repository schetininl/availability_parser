package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot ...
type Bot interface {
	SendMessage(string) error
}

// NewBotAPIClient ...
func NewBotAPIClient(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil

}

type bot struct {
	botAPI *tgbotapi.BotAPI
	chatID int64
}

// SendMessage ...
func (b *bot) SendMessage(text string) (err error) {
	msg := tgbotapi.NewMessage(b.chatID, text)
	_, err = b.botAPI.Send(msg)
	return
}

// New ...
func New(token string, chatID int64) (Bot, error) {
	b, err := NewBotAPIClient(token)
	if err != nil {
		return nil, err
	}
	return &bot{
		botAPI: b,
		chatID: chatID,
	}, nil
}
