package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stianeikeland/go-rpio/v4"
)

// State is the state
type State int64

const (
	// Running is running
	Running State = iota
	// NotRunning is not running
	NotRunning
)

// Tick is a moment when the sensor ticks
type Tick struct {
	Timestamp time.Time `json:"timestamp"`
}

func newTick() Tick {
	return Tick{Timestamp: time.Now()}
}

// mutates run by adding a tick
func addTick(run *Run, timestamp time.Time) {
	run.Ticks = append(run.Ticks, Tick{timestamp})
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

func main() {
	godotenv.Load()

	c := make(chan time.Time)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()
	pin := rpio.Pin(10)
	pin.Input() // Input mode

	var run Run

	go tickListen(c, pin)
	for {
		timestamp := <-c
		switch state {
		case Running:
			addTick(&run, timestamp)
			state = NotRunning
		case NotRunning:
			run = newRun()
			state = Running
		}
	}
	// putRunToS3(run)
}
