package engine

import (
	"BotLeha/Oleksii-bot/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// BotEngine is to controll the bot
type BotEngine struct {
	*tgbotapi.BotAPI
	*models.Bot
	quit       chan bool
	refreshCfg chan bool
	active     bool
}

func newEngine(bot *models.Bot) *BotEngine {
	return &BotEngine{
		BotAPI:     &tgbotapi.BotAPI{},
		Bot:        bot,
		quit:       make(chan bool, 1),
		refreshCfg: make(chan bool, 1),
		active:     false,
	}
}

func (e *BotEngine) initBot() error {
	var err error
	e.BotAPI, err = tgbotapi.NewBotAPI(e.Bot.Token)
	if err != nil {
		log.Println("[ERROR] Failed to connect to bot :", err.Error())
		return err
	}

	e.BotAPI.Debug = false

	//e.stateMng.initStateMng()
	//e.messages.initMsg()

	log.Printf("Authorized on account %s", e.BotAPI.Self.UserName)
	return nil
}

//Start the bot server
func (e *BotEngine) Start() error {
	// If already running do nothing
	if e.active == true {
		return nil
	}

	// Set active
	e.active = true
	// Deactivate
	defer func() {
		e.active = false
	}()

	if err := e.initBot(); err != nil {
		return err
	}

	// Main loop
	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60
	updates, err := e.BotAPI.GetUpdatesChan(config)
	if err != nil {
		log.Println("[ERROR] Unable to get updates: ", err)
	}

	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				e.processMsg(&update)
			} else {
				e.processCallback(&update)
			}
		case <-e.quit:
			log.Println("[WARNING] Stopping the bot: ")
			return nil
		}
	}
}

// Stop the bot server
func (e *BotEngine) Stop() error {
	if !e.active {
		return nil
	}

	log.Println("[WARNING] Sending stop signal to bot: ")
	e.quit <- true

	return nil
}

// Restart the bot server
func (e *BotEngine) Restart() error {
	err := e.Stop()
	if err != nil {
		return err
	}

	go e.Start()

	return nil
}

func (e *BotEngine) processMsg(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Text {
	case "/start":
		// TODO: get keyboard
		//msg.ReplyMarkup = getKeyboard(update.Message.Chat.ID)
		msg.Text = e.Bot.Config.HelloMsg
		e.BotAPI.Send(msg)
	}
}

func (e *BotEngine) processCallback(update *tgbotapi.Update) {

	// default:
	// 	msg.Text = "/start"
	// }

	// bot.reply(update, msg)
	// bot.logActionCSV(update.CallbackQuery.Message.Chat.ID, btnLabel)
	// bot.logActionXLS(update.CallbackQuery.Message, btnLabel)
}
