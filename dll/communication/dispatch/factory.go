package dispatch

import (
	"../context"
)

func CreateDispacther(ctx *context.DispatchContext) *Dispatcher {
	// create the disptacher
	disptacher := &Dispatcher{
		context: ctx,
		closing: make(chan bool),
	}

	// start the disptach loop
	go disptacher.dispatch()

	return disptacher
}
