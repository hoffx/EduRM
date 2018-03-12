package script

type Script struct {
	Name         string
	Path         string
	Instructions Instructions
}

type Instructions struct {
	Identifier string
	Parameters []int
}
