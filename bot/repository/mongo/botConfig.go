package mongo

import (
	"BotLeha/Oleksii-bot/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type BotCfg struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	BotID          primitive.ObjectID `bson:"botId"`
	HelloMsg       string             `bson:"helloMsg"`
	AskQuestion    string             `bson:"askQuestion"`
	Answer         string             `bson:"answer"`
	SubscribeTxt   string             `bson:"subscribeTxt"`
	UnsubscribeTxt string             `bson:"unsubscribeTxt"`
	MenuBlocks     []BotMenuBlock     `bson:"menuBlocks"`
}

type BotMenuBlock struct {
	CfgID       primitive.ObjectID `bson:"cfgId"`
	Enabled     bool               `bson:"enabled"`
	Menu        string             `bson:"menu"`
	LinkCaption string             `bson:"linkCaption"`
	LinkURL     string             `bson:"linkUrl"`
	Msg         string             `bson:"msg"`
}

type BotCfgRepository struct {
	db *mongo.Collection
}

// NewBotCfgRepository creates and returns new BotCfgRepository
func NewBotCfgRepository(db *mongo.Database, collection string) *BotCfgRepository {
	return &BotCfgRepository{
		db: db.Collection(collection),
	}
}

// GetCfg from db
func (br *BotCfgRepository) GetCfg(ctx context.Context, botID string) (*models.BotCfg, error) {
	objID, _ := primitive.ObjectIDFromHex(botID)
	cfg := new(BotCfg)

	err := br.db.FindOne(ctx, bson.M{
		"botId": objID,
	}).Decode(cfg)

	if err != nil {
		return nil, err
	}

	return toModelBCfg(cfg), nil
}

// UpsertCfg is about Update or Insert cfg to DB
func (br *BotCfgRepository) UpsertCfg(ctx context.Context, botCfg *models.BotCfg, botID string) error {
	objID, _ := primitive.ObjectIDFromHex(botID)
	cfg := toMongoBCfg(botCfg)
	opt := options.Update()
	opt.SetUpsert(true)

	// Perform upsert
	_, err := br.db.UpdateOne(ctx, bson.M{"botId": objID}, cfg, opt)

	if err != nil {
		return err
	}

	return nil
}

func toMongoBCfg(bc *models.BotCfg) *BotCfg {
	botID, _ := primitive.ObjectIDFromHex(bc.BotID)
	return &BotCfg{
		BotID:          botID,
		HelloMsg:       bc.HelloMsg,
		AskQuestion:    bc.AskQuestion,
		Answer:         bc.Answer,
		SubscribeTxt:   bc.SubscribeTxt,
		UnsubscribeTxt: bc.UnsubscribeTxt,
		//MenuBlocks
	}
}

func toModelBCfg(bc *BotCfg) *models.BotCfg {
	return &models.BotCfg{
		ID:             bc.ID.Hex(),
		HelloMsg:       bc.HelloMsg,
		AskQuestion:    bc.AskQuestion,
		Answer:         bc.Answer,
		SubscribeTxt:   bc.SubscribeTxt,
		UnsubscribeTxt: bc.UnsubscribeTxt,
	}
}
