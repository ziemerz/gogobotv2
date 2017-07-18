package main

import (
	"github.com/iopred/discordgo"
	"log"
	"fmt"
	"strings"
	. "github.com/ziemerz/gogobotv2/gogotypes"
)
var name string = "discord.go:: "

type Discord struct {
	received chan *Message
	outgoing chan *Message
	session *discordgo.Session
	bot *Bot
}

// NewDiscord adds all channels to receive and send messages
// and adds the bot so it can be passed to the dispatcher
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

// Open opens a discord session and starts a listener (Send) for sending messages back to Discord
// whenever one of the commands are done handling the incoming message
func (d *Discord) Open() {
	err := d.session.Open()
	if err != nil {
		log.Fatal("Couldn't initiate session")
	}
	go d.Send()
}

// Send sends messages to the appropriate Discord channel. Consumes from the outgoing chan
func (d *Discord) Send() {
	for msg := range d.outgoing {
		d.session.ChannelMessageSend(msg.Channel, msg.Content)
	}
}

// receive adds a message to the received chan
func (d *Discord) receive(m *Message) {
	d.received <- m
	fmt.Println("receive called")
}

// MessageCreate is a handler for Discord whenever a message is sent.
func (d *Discord) MessageCreate(session *discordgo.Session, mc *discordgo.MessageCreate) {
	if strings.HasPrefix(mc.Content, "!gogo") {
		fmt.Println("MEssageCreate called")
		d.receive(&Message{
			Content: mc.Content,
			Channel: mc.ChannelID,
		})
	}
}