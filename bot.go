package main

import (
	"github.com/ziemerz/gogobotv2/commands"
	. "github.com/ziemerz/gogobotv2/gogotypes"
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
	testCmd := commands.NewTestCommand()
	timerCmd := commands.NewTimerCommand()
	bot.commands = make(map[string]Command)
	bot.commands[testCmd.Name()] = testCmd
	bot.commands[timerCmd.Name()] = timerCmd
	discord := NewDiscord(key, bot)
	discord.Open()
	return bot
}

// Close closes the Discord session
func (b *Bot) Close() {
	b.discord.session.Close()
}
