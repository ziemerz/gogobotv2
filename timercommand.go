package gogobotv2

import (
	"time"
	"fmt"
)

type TimerCommand struct {
	messageChannel chan *Message
	channelIDs map[string] interface{}
	signal chan string
}

func NewTimerCommand() *TimerCommand {
	ids := make(map[string] interface{})
	tmc := new(TimerCommand)
	tmc.channelIDs = ids
	tmc.signal = make(chan string)
	return tmc
}

func (tc *TimerCommand) Name() string{
	return "timer"
}

func (tc *TimerCommand) AddChannel(msgChan chan *Message) {
	tc.messageChannel = msgChan
}

func (tc *TimerCommand) Fire(msg *Message, out chan *Message) {

	if tc.channelIDs[msg.Channel] == nil {
		tc.channelIDs[msg.Channel] = tc.startTimer
		go tc.channelIDs[msg.Channel].(func(string, chan *Message))(msg.Channel, out)
	} else {
		fmt.Println("Putting in stop signal")
		tc.signal <- "stop"
		tc.channelIDs[msg.Channel] = tc.startTimer
		go tc.channelIDs[msg.Channel].(func(string, chan *Message))(msg.Channel, out)
	}
}

func (tc *TimerCommand)startTimer(channel string, out chan *Message) {
	fmt.Println("Timer called")
	totalTime := time.Second * 30
	notif15 := totalTime - (time.Second * 5)
	notif30 := totalTime - (time.Second * 15)
	notif15chan := time.NewTimer(notif15).C
	notif1chan := time.NewTimer(notif30).C
	upchan := time.NewTimer(totalTime).C

	donechan := make(chan bool)

	go func() {
		for {
			select {
			case <-tc.signal:
				fmt.Println("Hit case of signal")
				donechan <- true
				return

			case <- notif1chan:
				out <- &Message{
					Content: "30 minutes remaining",
					Channel: channel,
				}
			case <- notif15chan:
				out <- &Message{
					Content: "15 minutes remaining",
					Channel: channel,
				}
			case <- upchan:
				out <- &Message{
					Content: "Time's up",
					Channel: channel,
				}
				donechan <- true
			}
		}
	}()
	<- donechan
	fmt.Println("DOne now!")
	// Clean up and make room for a new timer.
	tc.channelIDs[channel] = nil
}