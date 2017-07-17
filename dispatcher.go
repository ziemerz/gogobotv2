package main

import (
	"strings"
	"fmt"
)

type Dispatcher struct {
	incoming chan *Message
	outgoing chan *Message
	bot *Bot
}

func RunDispatcher(inc, out chan *Message, bot *Bot) {
	disp := new(Dispatcher)
	disp.outgoing = out
	disp.incoming = inc
	disp.bot = bot
	go disp.dispatch()
}

func (disp *Dispatcher) dispatch(){
	for msg := range disp.incoming {
		//msg := <-disp.incoming
		msg.content = strings.Split(msg.content, " ")[1]
		disp.bot.commands[msg.content].Fire(msg, disp.outgoing)
		fmt.Println("Message handled, returning")
	}
}