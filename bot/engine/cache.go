package engine

import (
	"BotLeha/Oleksii-bot/models"
	"sync"
)

type EngineCache struct {
	engCfgID map[string]*BotEngine
	*sync.Mutex
}

func NewEngineCache() *EngineCache {
	return &EngineCache{
		engCfgID: make(map[string]*BotEngine),
		Mutex:    new(sync.Mutex),
	}
}
func (c *EngineCache) Start(bot *models.Bot) error {
	c.Lock()
	defer c.Unlock()

	eng, ok := c.engCfgID[bot.ID]
	if ok {
		if eng.Active {
			return nil
		}

		go eng.Start()
		return nil
	}

	eng = newEngine(bot)
	go eng.Start()

	c.engCfgID[bot.ID] = eng
	return nil
}

func (c *EngineCache) Stop(bot *models.Bot) error {
	c.Lock()
	defer c.Unlock()
	eng, ok := c.engCfgID[bot.ID]
	if !ok {
		return nil
	}

	err := eng.Stop()
	delete(c.engCfgID, bot.ID)

	return err
}

func (c *EngineCache) Restart(bot *models.Bot) error {
	c.Lock()
	defer c.Unlock()
	eng, ok := c.engCfgID[bot.ID]
	if !ok {
		eng = newEngine(bot)
		go eng.Start()

		return nil
	}

	eng.Stop()
	go eng.Start()

	return nil
}
