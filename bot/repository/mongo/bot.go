package mongo

import (
	"BotLeha/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Bot entity in MongoDB
type Bot struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Token  string             `bson:"token"`
	Active bool               `bson:"active"`
	Status bool               `bson:"status"`
	Removed bool			  `bon:"removed"`
}

// BotRepository implementation for general interface
type BotRepository struct {
	db *mongo.Collection
}

// NewBotRepository creates and returns new BotRepository
func NewBotRepository(db *mongo.Database, collection string) *UserRepository {
	return &BotRepository{
		db: db.Collection(collection),
	}
}

func (br *BotRepository) Add(ctx context.Context, bot *models.Bot) error {
	model := toMongoBot(bot)
	res, err := br.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bot.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (br *BotRepository) Get(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	bot := new(Bot)
	err := ur.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(bot)

	if err != nil {
		return nil, err
	}

	return toModel(bot), nil
}

func (br *BotRepository) GetAll(ctx context.Context, bot *models.Bot) ([]models.Bot, error) {
	cur, err := br.db.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	if err != nil{
		return nil, err
	}

	out:= make([]*Bot, 0)

	for cur.Next(ctx){
		bot := new(Bot)
		err := cur.Decode(bot)
		if err != nil{
			return nil, err
		}

		out = append(out, bot)
	}
	if err:= cur.Err(); err!= nil {
		return nil, err
	}

	return toMongoBots(out)
}

// TODO: доделать упдейт
// func (br *BotRepository) Update(ctx context.Context, bot *models.Bot, id string) error {
// 	objID, _:= primitive.ObjectIDFromHex(id)
// 	model := toMongoBot(bot)


// }

func (br *BotRepository) Delete(ctx context.Context, bot *models.Bot, id string) error {
	objID, _:= primitive.ObjectIDFromHex(id)

	_, err := br.db.DeleteOne(ctx, bson.M{"_id": objID}) 
	return err
}

func toMongoBot(b *models.Bot) *Bot {
	return &User{
		Token:  b.Token,
		Active: b.Active.
		Status: b.Status,	
	}
}

func toMongoBots(list []*Bot) []*models.Bot {
	out := make([]*models.Bot, len(list))

	for i, bot := range list {
		out[i] = toMongoBot(bot)
	}

	return out
}

func toModel(b *Bot) *models.Bot {
	return &models.User{
		ID:     b.ID.Hex(),
		Token:  b.Token,
		Active: b.Active.
		Status: b.Status,
	}
}
