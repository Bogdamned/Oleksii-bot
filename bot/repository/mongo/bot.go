package mongo

import (
	"BotLeha/Oleksii-bot/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Bot entity in MongoDB
type Bot struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Token   string             `bson:"token"`
	Active  bool               `bson:"active"`
	Status  bool               `bson:"status"`
	Removed bool               `bson:"removed"`
	*BotCfg
}

// BotRepository implementation for general interface
type BotRepository struct {
	db *mongo.Collection
}

// NewBotRepository creates and returns new BotRepository
func NewBotRepository(db *mongo.Database, collection string) *BotRepository {
	return &BotRepository{
		db: db.Collection(collection),
	}
}

// Insert new bot
func (br *BotRepository) Insert(ctx context.Context, bot *models.Bot) error {
	model := toMongoBot(bot)
	res, err := br.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bot.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// Get bot from mongo db by ID
func (br *BotRepository) Get(ctx context.Context, id string) (*models.Bot, error) {
	botID, _ := primitive.ObjectIDFromHex(id)

	bot := new(Bot)
	err := br.db.FindOne(ctx, bson.M{
		"_id": botID,
	}).Decode(bot)

	if err != nil {
		return nil, err
	}

	return toModel(bot), nil
}

// GetAll bots from mongo db
func (br *BotRepository) GetAll(ctx context.Context) ([]*models.Bot, error) {
	cur, err := br.db.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Bot, 0)

	for cur.Next(ctx) {
		bot := new(Bot)
		err := cur.Decode(bot)
		if err != nil {
			return nil, err
		}

		out = append(out, bot)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toMongoBots(out), nil
}

// Update bot entity in mongo DB
func (br *BotRepository) Update(ctx context.Context, bot *models.Bot, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	model := toMongoBot(bot)

	_, err := br.db.UpdateOne(ctx, bson.M{"_id": objID}, model)

	return err
}

// Delete bot from mongo db
func (br *BotRepository) Delete(ctx context.Context, bot *models.Bot, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := br.db.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func toMongoBot(b *models.Bot) *Bot {
	ret := &Bot{
		Token:  b.Token,
		Active: b.Active,
		Status: b.Status,
	}

	return ret
}

func toMongoBots(list []*Bot) []*models.Bot {
	out := make([]*models.Bot, len(list))

	for i, bot := range list {
		out[i] = toModel(bot)
	}

	return out
}

func toModel(b *Bot) *models.Bot {
	return &models.Bot{
		ID:     b.ID.Hex(),
		Token:  b.Token,
		Active: b.Active,
		Status: b.Status,
	}
}
