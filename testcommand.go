package gogobotv2

import (
)

type TestCommand struct {
	messageChannel chan *Message
}

func NewTestCommand() *TestCommand {
	return new(TestCommand)
}

func (tc *TestCommand) Name() string{
	return "test"
}

func (tc *TestCommand) AddChannel(msgChan chan *Message) {
	tc.messageChannel = msgChan
}

func (tc *TestCommand) Fire(msg *Message, out chan *Message) {
	m := Message{
		Content: "Hejsa mate",
		Channel: msg.Channel,
	}

	out <- &m
}

