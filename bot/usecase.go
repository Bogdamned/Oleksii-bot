package bot

import (
	"BotLeha/Oleksii-bot/models"
	"context"
)

type UseCase interface {
	// Bot CRUD
	CreateBot(ctx context.Context, token string) error
	GetBot(ctx context.Context, id string) (*models.Bot, error)
	GetBots(ctx context.Context) ([]*models.Bot, error)
	UpdateBot(ctx context.Context, bot *models.Bot) error
	DeleteBot(ctx context.Context, id string) error

	// Configuration CRUD
	GetCfg(ctx context.Context, botID string) (*models.BotCfg, error)
	UpdateCfg(ctx context.Context, cfg *models.BotCfg, botID string) error

	// Engine
	Start(ctx context.Context, botID string) error
	Stop(ctx context.Context, botID string) error
	Restart(ctx context.Context, botID string) error

	SendMsg(ctx context.Context, botID string) error
	SendGroupMsg(ctx context.Context, botID string) error
}
