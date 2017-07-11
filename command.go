package main

//import "github.com/iopred/discordgo"

type CommandEntry struct {
	Command
	messageChannel chan *Message
}
type Command interface {
	Name() string
	Fire(msg *Message, out chan *Message)
}
