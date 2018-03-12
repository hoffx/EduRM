package script

type Script struct {
	Name string
	Path string
	[]Instruction instructions
}

type Instructions struct {
	Identifier string
	Parameters []int
}
