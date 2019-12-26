package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {

}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error connecting to mongo db")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err, client.Connect(ctx)
	if err != nil {
		log.Fatalf(err)
	}

	err= client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf(err)
	}

	retur client.Database(viper.GetString("mongo.name"))
}
