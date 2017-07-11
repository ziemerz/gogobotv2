package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init(){
	flag.StringVar(&key, "t", "", "Bot Token")
	flag.Parse()
}

var key string

func main() {
	if key == "" {
		log.Fatal("Please provide a private key. Example: gogobotv2 -t mypriavtekey1233")
	}
	fmt.Println(key)

	bot := NewBot(key)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	bot.Close()
}
