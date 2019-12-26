package bot

import (
	"BotLeha/models"
	"context"
)

//BotRepository is an abstraction to work with DBs
type BotRepository interface {
	Add(ctx context.Context, bot *models.Bot) error
	Get(ctx context.Context, id string) (*models.Bot, error)
	GetAll(ctx context.Context, bot *models.Bot) ([]models.Bot, error)
	Update(ctx context.Context, bot *models.Bot, id string) error
	Delete(ctx context.Context, id string) error
}
