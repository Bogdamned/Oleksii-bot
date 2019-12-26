package models

//Bot describes telegram bot
type Bot struct {
	ID     string
	Token  string
	Active bool
	Status bool
}
