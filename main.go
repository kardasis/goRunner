package main

import (
	"time"

	"github.com/joho/godotenv"
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

func main() {
	godotenv.Load()
	eventChan := make(chan runEvent)

	go eventListen(eventChan)
	for {
		event := <-eventChan
		switch event.eventType {
		case "tick":
		case "timeout":
		}
	}
	// putRunToS3(run)
}
