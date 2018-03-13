package controller

import (
	"time"

	"github.com/hoffx/EduRM/interpreter"
	"github.com/hoffx/EduRM/script"
)

// Controller modes
const (
	Run = iota
	Step
	Pause
	Stop
)

// Controller warnings
const (
	WarStoppedByUser       = "the process was successfully killed"
	WarStoppedInfiniteLoop = "the process was killed, because if ran into an infinite loop"
)

type Controller struct {
	ContextChan chan (interpreter.Context)
	Context     interpreter.Context
	Mode        chan (int)
	Delay       chan (int)
}

// NewController returns a Controller and an error if occurred
func NewController(filepath string, registerAmount int) (*Controller, error) {
	s, err := script.ParseFile(filepath)
	if err != nil {
		return nil, err
	}
	return &Controller{make(chan (interpreter.Context)), *interpreter.NewInterpreterContext(*s, registerAmount), make(chan (int)), make(chan (int))}, nil
}

// Process interprets the interpreter context contained by the Controller c
// Process should run as a separate goroutine
func (c *Controller) Process() {
	var delay int = 0
	var mode int = Pause

	var lastInstructionCounter uint = 0
	var infiniteLoopCandidates []interpreter.Context = make([]interpreter.Context, 0)

	c.ContextChan <- c.Context

	for {
		select {
		case m := <-c.Mode:
			mode = m
		case d := <-c.Delay:
			delay = d
		default:
			if mode != Pause {
				if mode == Stop {
					c.Context.Output = append(c.Context.Output, interpreter.Notification{interpreter.Warning, WarStoppedByUser, int(c.Context.InstructionCounter)})
					c.Context.Status = interpreter.Failure
				} else {
					if mode == Run {
						time.Sleep(time.Duration(delay) * time.Millisecond)
						c.Context.Next()
					} else {
						c.Context.Next()
						mode = Pause
					}
					// check if script ran into infinite loop
					if lastInstructionCounter+1 != c.Context.InstructionCounter {
						// check if current context is identical to a previous candidate
						for _, iLCtx := range infiniteLoopCandidates {
							if iLCtx.InstructionCounter == c.Context.InstructionCounter && c.Context.Accumulator == iLCtx.Accumulator {
								for i := range c.Context.Registers {
									if c.Context.Registers[i] != iLCtx.Registers[i] {
										break
									}
									// script ran into infinite loop
									c.Context.Output = append(c.Context.Output, interpreter.Notification{interpreter.Warning, WarStoppedInfiniteLoop, int(c.Context.InstructionCounter)})
									c.Context.Status = interpreter.Failure
								}
							}
						}
						infiniteLoopCandidates = append(infiniteLoopCandidates, c.Context)
					}
					lastInstructionCounter = c.Context.InstructionCounter
				}
				c.ContextChan <- c.Context
				c.Context.Output = make([]interpreter.Notification, 0)
			}

			// exit if stopped or terminated otherwise
			if c.Context.Status != interpreter.Running {
				return
			}
		}
	}
}

func (c *Controller) Run() {
	c.Mode <- Run
}

func (c *Controller) Step() {
	c.Mode <- Step
}

func (c *Controller) Pause() {
	c.Mode <- Pause
}

func (c *Controller) Stop() {
	c.Mode <- Stop
}

func (c *Controller) SetDelay(duration int) {
	c.Delay <- duration
}
