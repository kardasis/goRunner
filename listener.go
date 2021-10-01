package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type eventType int64

const (
	eventTick eventType = iota
	eventStartRun
	eventEndRun
)

type runEvent struct {
	eventType string
	tickData  time.Time
}

func tickListen(c chan time.Time) {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	pin := rpio.Pin(10)
	pin.Input() // Input mode

	for i := 0; i < 10; i++ {
		//////////////////////////////////////////////////
		///////  FAKE
		//////////////////////////////////////////////////
		time.Sleep(500 * time.Millisecond)
		c <- time.Now()
	}
}

func timeoutListen(c chan time.Time) {

}

func eventListen(eventChan chan runEvent) {
	tickChan := make(chan time.Time)
	clockChan := make(chan time.Time)
	go tickListen(tickChan)
	for {
		t := <-tickChan
		fmt.Println(t)
		go timeoutListen(clockChan)
	}
}
