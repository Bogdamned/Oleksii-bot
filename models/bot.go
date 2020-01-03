package models

//Bot describes telegram bot
type Bot struct {
	ID     string
	Token  string
	Active bool
	Status bool
	Config *BotCfg
}

type BotCfg struct {
	ID             string
	BotID          string
	HelloMsg       string
	AskQuestion    string
	Answer         string
	SubscribeTxt   string
	UnsubscribeTxt string
	MenuBlocks     []BotMenuBlock
}

type BotMenuBlock struct {
	CfgID       string
	Enabled     bool
	Menu        string
	LinkCaption string
	LinkUrl     string
	Msg         string
}
