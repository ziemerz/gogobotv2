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

func RunDispatcher(inc, out chan *Message, bot *Bot) *Dispatcher {
	disp := new(Dispatcher)
	disp.outgoing = out
	disp.incoming = inc
	disp.bot = bot
	go disp.dispatch()

	return disp
}

func (disp *Dispatcher) dispatch(){
	msg := <-disp.incoming
	msg.content = strings.Split(msg.content, " ")[1]
	disp.bot.commands[msg.content].Fire(msg, disp.outgoing)
	fmt.Println("Message handled, returning")
	<-disp.incoming
}