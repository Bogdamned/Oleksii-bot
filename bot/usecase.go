package bot

import (
	"BotLeha/Oleksii-bot/models"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, token string) error
	Get(ctx context.Context, id string) (*models.Bot, error)
	GetAll(ctx context.Context) ([]*models.Bot, error)
	Update(ctx context.Context, bot *models.Bot) error
	Delete(ctx context.Context, id string) error

	GetCfg(ctx context.Context, botID string) (*models.BotCfg, error)
	UpdateCfg(ctx context.Context, cfg *models.BotCfg, botID string) error

	Start()
	Stop()
	Restart()

	SendMsg()
	SendGroupMsg()
}
