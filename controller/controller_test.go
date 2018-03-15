package controller

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/hoffx/EduRM/interpreter"
)

func TestController(t *testing.T) {
	c, err := NewController("../interpreter/test.txt", 16)
	if err != nil {
		t.Error(err)
	}
	go func() {
		for {

		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		c.Process()
		wg.Done()
	}()
	go func() {
		var status int = -1
		for status == -1 {
			select {
			case ctx := <-c.ContextChan:
				status = ctx.Status
				visualize(ctx)
			}
		}
		wg.Done()
	}()
	c.AddBreakpoint(3)
	c.SetDelay(200)
	c.Run()
	time.Sleep(10 * time.Millisecond)
	c.Pause()
	c.AddBreakpoint(4)
	time.Sleep(1000 * time.Millisecond)
	c.Step()
	c.DeleteBreakpoint(3)
	time.Sleep(1000 * time.Millisecond)
	c.Run()
	time.Sleep(100 * time.Millisecond)
	c.Stop()
	wg.Wait()
}

func visualize(ctx interpreter.Context) {
	var status string
	if ctx.Status == interpreter.Running {
		status = "RUNNING"
	} else if ctx.Status == interpreter.Success {
		status = "SUCCESS"
	} else {
		status = "FAILURE"
	}
	fmt.Printf("\r%v:  Instrcution Counter: %v   \t Accumulator: %v \t Registers: %v \t Output: %v", status, ctx.InstructionCounter, ctx.Accumulator, ctx.Registers, ctx.Output)
}
