package interpreter

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
	ErrInvalidInstructionAddress = "the called instruction does not exist"
)

// Interpreter warnings
const ()

// Interpreter messages
const (
	MessTerminatedWithSuccess = "the script has terminated with success"
)

type Notification struct {
	Type    int
	Content string
	// Instruction is -1 if Notification is not instruction-specific
	Instruction int
}

type Context struct {
	InstructionCounter uint
	Accumulator        int
	Script             Script
	Registers          []int
	Output             []Notification
	Status             int
}

func NewInterpreterContext(Script script, int registerAmount) *Context {
	return &Context{0, 0, script, make([]int, registerAmount), make([]Notification, 0), Running}
}

func (ctx *Context) Next() {
	if ctx.Status != Running {
		ctx.Output = append(ctx.Output, Notification{Error, ErrNotRunning, -1})
		return
	}
	if ctx.InstructionCounter > len(ctx.Script.Instructions)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrInvalidInstructionAddress, ctx.InstructionCounter})
		return
	}
	Interpret(ctx)
	switch ctx.Status {
	case Success:
		ctx.Output = append(ctx.Output, Notification{Message, MessTerminatedWithSuccess, ctx.InstructionCounter})
	case Failure:
		ctx.Output = append(ctx.Output, Notification{Error, MessTerminatedWithSuccess, ctx.InstructionCounter})
	}
}
