package controller

import (
	"strings"
	"time"

	"github.com/hoffx/EduRM/interpreter"
	"github.com/hoffx/EduRM/script"
)

// Controller modes
const (
	run = iota
	step
	pause
	stop
)

// Controller warnings
const (
	WarStoppedByUser       = "the process was successfully killed"
	WarStoppedInfiniteLoop = "the process was killed, because it ran into an infinite loop"
)

type Controller struct {
	ContextChan              chan interpreter.Context
	context                  interpreter.Context
	modeChan                 chan int
	delayChan                chan int
	breakpointActivateChan   chan uint
	breakpointDeactivateChan chan uint
	delay                    int
	breakpoints              map[uint]bool
}

// NewController returns a Controller and an error if occurred
func NewController(text string, registerAmount int) (*Controller, error) {
	s, err := script.Parse("", strings.NewReader(text))
	if err != nil {
		return nil, err
	}
	return &Controller{
		ContextChan:              make(chan interpreter.Context),
		context:                  *interpreter.NewInterpreterContext(*s, registerAmount),
		modeChan:                 make(chan int),
		delayChan:                make(chan int),
		breakpointActivateChan:   make(chan uint),
		breakpointDeactivateChan: make(chan uint),
		delay:       0,
		breakpoints: make(map[uint]bool, 0),
	}, nil
}

// Process interprets the interpreter context contained by the Controller c
// Process should run as a separate goroutine
func (c *Controller) Process() {
	var mode int = pause

	var lastInstructionCounter uint = 0
	var infiniteLoopCandidates = make([]interpreter.Context, 0)

	c.ContextChan <- c.context

	for {
		select {
		case m := <-c.modeChan:
			mode = m
		case d := <-c.delayChan:
			c.delay = d
		case bp := <-c.breakpointActivateChan:
			c.breakpoints[bp] = true
		case bp := <-c.breakpointDeactivateChan:
			delete(c.breakpoints, bp)
		default:
			if mode != pause {
				if mode == stop {
					c.context.Output = append(c.context.Output, interpreter.Notification{interpreter.Warning, WarStoppedByUser, int(c.context.InstructionCounter)})
					c.context.Status = interpreter.Failure
					c.ContextChan <- c.context
				} else {
					c.context.Next()
					if mode == step {
						mode = pause
					}
					// check if script ran into infinite loop
					if lastInstructionCounter+1 != c.context.InstructionCounter {
						// check if current context is identical to a previous candidate
						for _, iLCtx := range infiniteLoopCandidates {
							if iLCtx.InstructionCounter == c.context.InstructionCounter && c.context.Accumulator == iLCtx.Accumulator {
								var isNotInfinite bool
								for i := range c.context.Registers {
									if c.context.Registers[i] != iLCtx.Registers[i] {
										isNotInfinite = true
										break
									}
								}
								if !isNotInfinite {
									// script ran into infinite loop
									c.context.Output = append(c.context.Output, interpreter.Notification{interpreter.Warning, WarStoppedInfiniteLoop, int(c.context.InstructionCounter)})
									c.context.Status = interpreter.Failure
									break
								}
							}
						}
						oldRegisters := make([]int, len(c.context.Registers))
						copy(oldRegisters, c.context.Registers)
						oldContext := interpreter.Context{
							InstructionCounter: c.context.InstructionCounter,
							Accumulator:        c.context.Accumulator,
							Script:             script.Script{},
							Output:             []interpreter.Notification{},
							Status:             c.context.Status,
							Instruction:        c.context.Instruction,
							Registers:          oldRegisters,
						}
						infiniteLoopCandidates = append(infiniteLoopCandidates, oldContext)
					}
					if lastInstructionCounter == 0 {
						time.Sleep(time.Duration(c.delay) * time.Millisecond)
					}
					c.ContextChan <- c.context
					if mode == run {
						// don't delay if in step mode
						time.Sleep(time.Duration(c.delay) * time.Millisecond)
					}

					c.context.Output = make([]interpreter.Notification, 0)
					lastInstructionCounter = c.context.InstructionCounter
				}

			}

			// check for breakpoint
			if c.breakpoints[c.context.InstructionCounter] {
				mode = pause
			}

			// exit if stopped or terminated otherwise
			if c.context.Status != interpreter.Running {
				return
			}
		}
	}
}

func (c *Controller) Run() {
	c.modeChan <- run
}

func (c *Controller) Step() {
	c.modeChan <- step
}

func (c *Controller) Pause() {
	c.modeChan <- pause
}

func (c *Controller) Stop() {
	c.modeChan <- stop
}

func (c *Controller) SetDelay(duration int) {
	c.delayChan <- duration
}

func (c *Controller) AddBreakpoint(position uint) {
	c.breakpointActivateChan <- position
}

func (c *Controller) DeleteBreakpoint(position uint) {
	c.breakpointDeactivateChan <- position
}
