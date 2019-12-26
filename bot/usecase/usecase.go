package usecase

import (
	"BotLeha/models"
)

type BotUseCase struct {
	bot models.Bot
}

func NewBotUseCase() *BotUseCase {
	return &BotUseCase{}
}
