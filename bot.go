package main

type Bot struct {
	discord  *Discord
	commands map[string]Command
	messages chan Message
}

func NewBot(key string) *Bot {
	bot := new(Bot)
	testCmd := NewTestCommand()
	timerCmd := NewTimerCommand()
	bot.commands = make(map[string]Command)
	bot.commands[testCmd.Name()] = testCmd
	bot.commands[timerCmd.Name()] = timerCmd
	discord := NewDiscord(key, bot)
	discord.Open()
	return bot
}

func (b *Bot) Close() {
	b.discord.session.Close()
}
