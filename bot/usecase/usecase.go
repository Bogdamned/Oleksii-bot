package usecase

import (
	"BotLeha/Oleksii-bot/bot"
	"BotLeha/Oleksii-bot/models"
	"context"
)

type BotUseCase struct {
	//bot        models.Bot
	botRepo    bot.Repository
	botCfgRepo bot.CfgRepository
}

func NewBotUseCase(
	botRepo bot.Repository,
	botCfgRepo bot.CfgRepository) *BotUseCase {
	return &BotUseCase{
		botRepo:    botRepo,
		botCfgRepo: botCfgRepo,
	}
}

func (uc *BotUseCase) Create(ctx context.Context, token string) error {
	bot := new(models.Bot)
	bot.Token = token
	bot.Active = false
	bot.Status = false

	return uc.botRepo.Insert(ctx, bot)
}

func (uc *BotUseCase) Get(ctx context.Context, id string) (*models.Bot, error) {
	return uc.botRepo.Get(ctx, id)
}

func (uc *BotUseCase) GetAll(ctx context.Context) ([]*models.Bot, error) {
	return uc.botRepo.GetAll(ctx)
}

func (uc *BotUseCase) Update(ctx context.Context, bot *models.Bot) error {
	return uc.botRepo.Update(ctx, bot, bot.ID)
}

func (uc *BotUseCase) Delete(ctx context.Context, id string) error {
	return uc.botRepo.Delete(ctx, id)
}
