package gogobotv2

import (
	"errors"
)

type Bot struct {
	discord  *Discord
	commands map[string]Command
	messages chan Message
}

// NewBot creates a new bot and sets it up with all the commands it needs.
// It will also make sure to open up the Discord connection
func NewBot(key string) *Bot {
	bot := new(Bot)
	bot.commands = make(map[string]Command)
	discord := NewDiscord(key, bot)
	bot.discord = discord
	return bot
}

//Start opens the Discord session
func (b *Bot) Start() {
	b.discord.Open()
}

// Close closes the Discord session
func (b *Bot) Close() {
	b.discord.session.Close()
}

// AddCommand if the passed command does not already exist, register it.
func (b *Bot) AddCommand(cmd Command) error {
	if b.commands[cmd.Name()] == nil {
		b.commands[cmd.Name()] = cmd
		return nil
	}
	return errors.New("Command already registered")
}
