package ioengine

import (
	"time"
)

type event struct {
	callback func()
}

type IOEngine struct {
	on bool
	pending []event
	queue chan event
}

func (self *IOEngine) AddCallback(callback func()) {
	evt := event{callback}

	if self.on {
		self.queue <- evt
	} else {
		self.pending = append(self.pending, evt)
	}
}

func (self *IOEngine) Start() {
	self.queue = make(chan event)
	self.on = true
	go self.loop()

	for len(self.pending) > 0 {
		self.queue <- self.pending[0]
		self.pending = self.pending[1:]
	}

	for self.on {
		time.Sleep(100 * time.Millisecond)
	}
}

func (self *IOEngine) Status() bool {
	return self.on
}

func (self *IOEngine) Stop() {
	self.on = false
}

func (self *IOEngine) loop() {
	if self.on {
		event := <-self.queue
		go event.callback()
		go self.loop()
	}
}
