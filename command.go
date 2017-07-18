package main

import (
	. "github.com/ziemerz/gogobotv2/gogotypes"
)

// CommandEntry struct implements the Command interface
type CommandEntry struct {
	Command
	messageChannel chan *Message
}

// Command interface makes sure a command has a Name and Fire method
type Command interface {
	Name() string
	Fire(msg *Message, out chan *Message)
}
