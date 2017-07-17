package main

import (
	"github.com/iopred/discordgo"
	"log"
	"fmt"
	"strings"
)
var name string = "discord.go:: "

type Discord struct {
	received chan *Message
	outgoing chan *Message
	session *discordgo.Session
	bot *Bot
}

func NewDiscord(key string, bot *Bot) *Discord{
	var err error
	disc := new(Discord)
	disc.received = make(chan *Message, 100)
	disc.outgoing = make(chan *Message, 100)
	disc.session, err = discordgo.New("Bot " + key)

	disc.bot = bot

	disc.session.AddHandler(disc.MessageCreate)
	if err != nil {
		log.Fatal(name + "Couldn't create discord session :(")
	}

	RunDispatcher(disc.received, disc.outgoing, disc.bot)
	return disc
}

func (d *Discord) AddHandler(command Command) {
	fmt.Println("Added handler ", command.Name())
}

func (d *Discord) Open() {
	err := d.session.Open()
	if err != nil {
		log.Fatal("Couldn't initiate session")
	}
	go d.Send()
}

func (d *Discord) Send() {
	for msg := range d.outgoing {
		d.session.ChannelMessageSend(msg.channel, msg.content)
	}
}

func (d *Discord) receive(m *Message) {
	d.received <- m
	fmt.Println("receive called")
}

func (d *Discord) MessageCreate(session *discordgo.Session, mc *discordgo.MessageCreate) {
	if strings.HasPrefix(mc.Content, "!gogo") {
		fmt.Println("MEssageCreate called")
		d.receive(&Message{
			content: mc.Content,
			channel: mc.ChannelID,
		})
	}
}