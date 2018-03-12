package interpreter

// Interpretion errors
const (
	ErrWrongArgAmount     = "wrong amount of arguments"
	ErrIllegalArgument    = "illegal argument"
	ErrUnknownInstruction = "unknown instruction"
)

func Interpret(ctx *Context) {
	switch ctx.Script.Instructions[ctx.InstructionCounter].Identifier {
	case "LOAD":
		load(ctx)
	case "DLOAD":
		dload(ctx)
	case "STORE":
		store(ctx)
	case "ADD":
		add(ctx)
	case "SUB":
		sub(ctx)
	case "MULT":
		mult(ctx)
	case "DIV":
		div(ctx)
	case "JUMP":
		jump(ctx)
	case "JGE":
		jge(ctx)
	case "JGT":
		jgt(ctx)
	case "JLE":
		jle(ctx)
	case "JLT":
		jlt(ctx)
	case "JEQ":
		jeq(ctx)
	case "JNE":
		jne(ctx)
	default:
		ctx.Output = append(ctx.Output, Notification{Error, ErrUnknownInstuction, ctx.InstructionCounter})
		ctx.Status = Failure
	}
}

func load(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator = ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]]
	ctx.InstructionCounter++
}

func dload(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	ctx.InstructionCounter++
}

func store(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]] = ctx.Accumulator
	ctx.InstructionCounter++
}

func add(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator += ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]]
	ctx.InstructionCounter++
}

func sub(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator -= ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]]
	ctx.InstructionCounter++
}

func mult(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator *= ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]]
	ctx.InstructionCounter++
}

func div(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] < 0 || ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0] > len(ctx.Registers)-1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrIllegalArgument, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.Accumulator /= ctx.Registers[ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]]
	ctx.InstructionCounter++
}

func jump(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
}

func jge(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator >= 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func jgt(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator > 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func jle(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator <= 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func jlt(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator < 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func geq(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator == 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func jne(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 1 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	if ctx.Accumulator != 0 {
		ctx.InstructionCounter = ctx.Script.Instructions[ctx.InstructionCounter].Parameters[0]
	} else {
		ctx.InstructionCounter++
	}
}

func end(ctx *Context) {
	if len(ctx.Script.Instructions[ctx.InstructionCounter].Parameters) != 0 {
		ctx.Output = append(ctx.Output, Notification{Error, ErrWrongArgAmount, ctx.InstructionCounter})
		ctx.Status = Failure
		return
	}
	ctx.InstructionCounter++
	ctx.Status = Success
}
