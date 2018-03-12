package interpreter

import "github.com/hoffx/EduRM/script"

// Notification types
const (
	Message = iota
	Warning
	Error
)

// Status values
const (
	Running = iota - 1
	Success
	Failure
)

// Interpreter errors
const (
	ErrNotRunning                = "the script has terminated already"
	ErrCritical                  = "the execution was stopped after a critical interpretion error occurred"
	ErrInvalidInstructionAddress = "the called instruction-address does not exist"
	MessTerminatedOnFailure      = "the script failed"
)

// Interpreter warnings
const ()

// Interpreter messages
const (
	MessTerminatedWithSuccess = "the script has terminated with success"
)

// Notification holds a notification like error messages, warnings, ...
type Notification struct {
	Type    int
	Content string
	// Instruction is -1 if Notification is not instruction-specific
	Instruction int
}

// Context holds everything that describes an interpreter's state
type Context struct {
	InstructionCounter uint
	Accumulator        int
	Script             script.Script
	Registers          []int
	Output             []Notification
	Status             int
}

// NewInterpreterContext returns a context for an interpreter ready to start execution
func NewInterpreterContext(script script.Script, registerAmount int) *Context {
	return &Context{1, 0, script, make([]int, registerAmount), make([]Notification, 0), Running}
}

// Next interprets the current instruction
func (ctx *Context) Next() {
	if ctx.Status != Running {
		ctx.Output = append(ctx.Output, Notification{Error, ErrNotRunning, -1})
		ctx.Status = Failure
		return
	}

	if _, ok := ctx.Script.Instructions[int(ctx.InstructionCounter)]; !ok {
		ctx.Output = append(ctx.Output, Notification{Error, ErrInvalidInstructionAddress, int(ctx.InstructionCounter)})
		ctx.Status = Failure
		return
	}
	Interpret(ctx)
	switch ctx.Status {
	case Success:
		ctx.Output = append(ctx.Output, Notification{Message, MessTerminatedWithSuccess, int(ctx.InstructionCounter)})
	case Failure:
		ctx.Output = append(ctx.Output, Notification{Error, MessTerminatedOnFailure, int(ctx.InstructionCounter)})
	}
}
