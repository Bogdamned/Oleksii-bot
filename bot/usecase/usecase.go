package usecase

import (
	"BotLeha/Oleksii-bot/bot"
	"BotLeha/Oleksii-bot/bot/engine"
	"BotLeha/Oleksii-bot/models"
	"context"
)

type BotUseCase struct {
	botRepo    bot.Repository
	botCfgRepo bot.CfgRepository

	engines *engine.EngineCache
}

func NewBotUseCase(
	botRepo bot.Repository,
	botCfgRepo bot.CfgRepository) *BotUseCase {
	return &BotUseCase{
		botRepo:    botRepo,
		botCfgRepo: botCfgRepo,
		engines:    engine.NewEngineCache(),
	}
}

func (uc *BotUseCase) CreateBot(ctx context.Context, token string) error {
	bot := new(models.Bot)
	bot.Token = token
	bot.Active = false
	bot.Status = false

	return uc.botRepo.Insert(ctx, bot)
}

func (uc *BotUseCase) GetBot(ctx context.Context, id string) (*models.Bot, error) {
	return uc.botRepo.Get(ctx, id)
}

func (uc *BotUseCase) GetBots(ctx context.Context) ([]*models.Bot, error) {
	return uc.botRepo.GetAll(ctx)
}

func (uc *BotUseCase) UpdateBot(ctx context.Context, bot *models.Bot) error {
	return uc.botRepo.Update(ctx, bot, bot.ID)
}

func (uc *BotUseCase) DeleteBot(ctx context.Context, id string) error {
	return uc.botRepo.Delete(ctx, id)
}

func (uc *BotUseCase) GetCfg(ctx context.Context, botID string) (*models.BotCfg, error) {
	return uc.botCfgRepo.GetCfg(ctx, botID)
}

func (uc *BotUseCase) UpdateCfg(ctx context.Context, cfg *models.BotCfg, botID string) error {
	return uc.botCfgRepo.UpsertCfg(ctx, cfg, botID)
}

func (uc *BotUseCase) Start(ctx context.Context, botID string) error {
	bot, err := uc.botRepo.Get(ctx, botID)
	if err != nil {
		return err
	}

	return uc.engines.Start(bot)
}

func (uc *BotUseCase) Stop(ctx context.Context, botID string) error {
	bot, err := uc.botRepo.Get(ctx, botID)
	if err != nil {
		return err
	}

	return uc.engines.Stop(bot)
}

func (uc *BotUseCase) Restart(ctx context.Context, botID string) error {
	bot, err := uc.botRepo.Get(ctx, botID)
	if err != nil {
		return err
	}

	return uc.engines.Restart(bot)
}
