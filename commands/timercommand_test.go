package commands_test

import (
	. "github.com/ziemerz/gogobotv2/commands"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ziemerz/gogobotv2"
	"time"
)

var _ = Describe("Timercommand", func() {
	var (
		timerCommand *TimerCommand
		out chan *Message
		// timerEntry TimerEntry
	)

	BeforeEach(func(){
		timerCommand = NewTimerCommand()
		out = make(chan *Message)
	})

	Describe("Fire. Timer should put a message in the out channel", func() {
		Context("After 10 seconds", func(){
			It("Should have a message after 10 seconds", func(){
				msg := &Message{
					Content:"!gogo timer in 0:0:10",
					Channel: "mockchannel",
				}
				timerCommand.Fire(msg, out)

				time.Sleep(time.Second * 11)
				outMsg := <- out
				Expect(outMsg.Content).To(Equal("Time's up"))
			})
		})
	})


})
