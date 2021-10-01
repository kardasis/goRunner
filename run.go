package main

import "time"

// Run is an instance of a run
type Run struct {
	StartTime time.Time `json:"start_ime"`
	Ticks     []Tick    `json:"ticks"`
}

func newRun() Run {
	run := Run{
		StartTime: time.Now(),
		Ticks:     []Tick{{Timestamp: time.Now()}},
	}
	return run
}
