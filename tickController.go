package TickController

import (
	"fmt"
	"time"
)

/*
Tick controller, attempts to tick objects at a regular predefined interval and cleanup on program exit
*/

type TickController struct {
	tickJobs []Ticker
}

type Ticker interface {
	Initialize()
	Tick()
	CleanUp()
}

func (controller *TickController) StartTick(finish *bool, interval int64) {
	controller.Create()
	controller.Tick(finish, interval)
	controller.CleanUp()
}

func (controller *TickController) Create() {
	for _, job := range controller.tickJobs {
		job.Initialize()
	}
}

func (controller *TickController) Tick(finish *bool, interval int64) {
	var lastTick time.Time
	var delta time.Duration
	sleepDuration, _ := time.ParseDuration(fmt.Sprintf("%vns", interval/4))

	for *finish {
		delta = time.Now().Sub(lastTick)

		if delta.Nanoseconds() < interval {
			time.Sleep(sleepDuration)
			continue
		}

		for _, job := range controller.tickJobs {
			job.Tick()
		}
		lastTick = time.Now()
	}
}

func (controller *TickController) CleanUp() {
	for _, job := range controller.tickJobs {
		job.CleanUp()
	}
}

func (controller *TickController) AddTicker(job Ticker) {
	controller.tickJobs = append(controller.tickJobs, job)
}

func PerSecondInterval(perSecond int) int64 {
	return int64(1000000000 / perSecond)
}

/*
Need to test if adding a new ticker during a tick is allowed
*/
