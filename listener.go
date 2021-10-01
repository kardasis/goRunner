package main

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type listener struct {
	state State
}

func (l *listener) SetupPin() {
	l.state = NotRunning
}

func tickListen(c chan time.Time, pin rpio.Pin) {
	for i := 0; i < 10; i++ {
		//////////////////////////////////////////////////
		///////  FAKE
		//////////////////////////////////////////////////
		time.Sleep(500 * time.Millisecond)
		c <- time.Now()
	}
}
