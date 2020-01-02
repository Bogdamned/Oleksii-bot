package mongo

import (
	"BotLeha/Oleksii-bot/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represents user entity in MongoDB
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

// UserRepository implementation for general interface
type UserRepository struct {
	db *mongo.Collection
}

// NewUserRepository creates and returns new UserRepositorycle
func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

// CreateUser inserts user in DB
func (ur UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toMongoUser(user)
	res, err := ur.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// GetUser selects user from DB
func (ur UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User)
	err := ur.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toMongoUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}
