package ioengine

type event struct {
	callback func()
	payload interface{}
}

type IOEngine struct {
	queue chan event
	on bool
}

func (self *IOEngine) AddCallback(callback func(), payload interface{}) {
	self.queue <- event{
		callback: callback,
		payload: payload,
	}
}

func (self *IOEngine) loop() bool {
	for self.on {
		event <- self.queue
		go event.callback(event.payload)
	}

	return true
}

func (self *IOEngine) Stop() {
	self.on = false
}

func (self *IOEngine) Start() {
	self.on = true
	go self.loop()
}

func (self IOEngine) Status() bool {
	return self.on
}
