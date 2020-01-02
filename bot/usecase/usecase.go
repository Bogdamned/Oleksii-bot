package usecase

import (
	"BotLeha/Oleksii-bot/models"
)

type BotUseCase struct {
	bot models.Bot
}

func NewBotUseCase() *BotUseCase {
	return &BotUseCase{}
}
