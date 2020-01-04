package bot

import (
	"BotLeha/Oleksii-bot/models"
	"context"
)

// Repository is an abstraction to work with DBs
type Repository interface {
	Insert(ctx context.Context, bot *models.Bot) error
	Get(ctx context.Context, id string) (*models.Bot, error)
	GetAll(ctx context.Context) ([]*models.Bot, error)
	Update(ctx context.Context, bot *models.Bot, id string) error
	Delete(ctx context.Context, id string) error
}

//CfgRepository abstraction to work with bot settings
type CfgRepository interface {
	GetCfg(ctx context.Context, botID string) (*models.BotCfg, error)
	UpsertCfg(ctx context.Context, botCfg *models.BotCfg, botID string) error
}
