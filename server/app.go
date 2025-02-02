package server

import (
	"BotLeha/Oleksii-bot/bot"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"BotLeha/Oleksii-bot/auth"

	authhttp "BotLeha/Oleksii-bot/auth/delivery/http"
	authmongo "BotLeha/Oleksii-bot/auth/repository/mongo"
	authusecase "BotLeha/Oleksii-bot/auth/usecase"

	botmongo "BotLeha/Oleksii-bot/bot/repository/mongo"
	botusecase "BotLeha/Oleksii-bot/bot/usecase"
)

type App struct {
	httpServer *http.Server

	//bookmarkUC bookmark.UseCase
	authUC auth.UseCase
	botUC  bot.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	botRepo := botmongo.NewBotRepository(db, viper.GetString("mongo.bot_collection"))
	botCfgRepo := botmongo.NewBotCfgRepository(db, viper.GetString("mongo.bot_cfg_collection"))
	//bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
		botUC: botusecase.NewBotUseCase(
			botRepo,
			botCfgRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	//authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	// api := router.Group("/api", authMiddleware)

	//bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	fmt.Println("Initializing DB")
	fmt.Println("mongo.uri: ", viper.GetString("mongo.uri"))
	fmt.Println("mongo.DBname: ", viper.GetString("mongo.name"))

	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
