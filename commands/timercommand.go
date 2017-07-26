package commands

import (
	"time"
	. "github.com/ziemerz/gogobotv2"
	"strings"
	"strconv"
	"fmt"
)

type TimerCommand struct {
	messageChannel chan *Message
	channelIDs map[string] *TimerEntry
}

func NewTimerCommand() *TimerCommand {
	ids := make(map[string] *TimerEntry)
	tmc := new(TimerCommand)
	tmc.channelIDs = ids
	return tmc
}

func (tc *TimerCommand) Name() string{
	return "timer"
}

func (tc *TimerCommand) AddChannel(msgChan chan *Message) {
	tc.messageChannel = msgChan
}

func (tc *TimerCommand) Fire(msg *Message, out chan *Message) {
	timer := &TimerEntry{
		channel: msg.Channel,
		out: out,
		signal:make(chan bool),
	}

	command := strings.Split(msg.Content, " ")

	if len(command) > 2 {
		subcmd := command[2];
		if subcmd ==  "in" {
			tc.in(timer, command[3])
		}

		if subcmd == "stop" {
			tc.stop(msg.Channel)
		}
	}
}

// startTimer starts a timer.
func (tc *TimerCommand) startTimer(timer *TimerEntry) {
	go timer.Start()
}

// Interface to make sure the TimerEntry has a start command

type Timer interface {
	Start()
}

type TimerEntry struct {
	Timer
	channel string
	out chan *Message
	signal chan bool
	duration time.Duration
}


func (t *TimerEntry) Start() {

	fmt.Println("Starting timer")
	//ttime := (time.Hour * t.hour) + (time.Minute * t.minute) + (time.Second * t.second)
	totalTime := t.duration
	notif2 := totalTime - (time.Second * 5)
	notif1 := totalTime - (time.Second * 30)
	notif2chan := time.NewTimer(notif2).C
	notif1chan := time.NewTimer(notif1).C
	upchan := time.NewTimer(totalTime).C

	donechan := make(chan bool)

	go func() {
		for {
			select {
			case <-t.signal:
				donechan <- true
				return

			case <- notif1chan:
				t.out <- &Message{
					Content: "1 hour remaining",
					Channel: t.channel,
				}
			case <- notif2chan:
				t.out <- &Message{
					Content: "15 minutes remaining",
					Channel: t.channel,
				}
			case <- upchan:
				t.out <- &Message{
					Content: "Time's up",
					Channel: t.channel,
				}
				donechan <- true
				return
			}
		}
	}()
	<- donechan
	fmt.Println("Timer done/stopped")
}

func (tc *TimerCommand) in(timer *TimerEntry, duration string) {

	fmt.Println("In called")
	split := strings.Split(duration, ":")
	var h, m, s time.Duration
	var hi, mi, si int
	var err error
	if len(split) >= 1 {
		hi, err = strconv.Atoi(split[0])
		if len(split) >= 2 {
			mi, err = strconv.Atoi(split[1]);
			if len(split) == 3 {
				si, err  = strconv.Atoi(split[2]);
			}
		}
	}

	if err != nil {
		fmt.Println("Wrong formatting")
	}

	h = time.Duration(hi)
	m = time.Duration(mi)
	s = time.Duration(si)

	// Set the duration for the timer.
	timer.duration = (time.Hour * h) + (time.Minute * m) + (time.Second * s)

	if tc.channelIDs[timer.channel] == nil {
		fmt.Println("If channel in list is nil")
		tc.channelIDs[timer.channel] = timer
		go tc.startTimer(tc.channelIDs[timer.channel])
	} else {
		fmt.Println("If channel in list is != nil")
		tc.channelIDs[timer.channel].signal <- true
		tc.channelIDs[timer.channel] = nil
		tc.channelIDs[timer.channel] = timer
		go tc.startTimer(tc.channelIDs[timer.channel])
	}
	fmt.Println("In done")
}

func (tc *TimerCommand) stop(channel string) {
	t := tc.channelIDs[channel]
	if t != nil {
		t.signal <- true
		tc.channelIDs[channel] = nil
		fmt.Println("Stopped")
	}
}

