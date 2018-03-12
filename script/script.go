package script

import (
	"errors"
	"regexp"
	"strconv"
)

var InstructionRegex = regexp.MustCompile(`^([0-9]+)[:\s]+([A-Z]+)\s+([0-9]+)`)

type Script struct {
	Name         string
	Path         string
	Instructions map[int]Instruction
}

type Instruction struct {
	Number     int
	Identifier string
	Parameters []int
}

func NewScript(name string) *Script {
	return &Script{
		Name:         name,
		Instructions: make(map[int]Instruction),
	}
}

func (s *Script) AppendInstruction(instruction Instruction) {
	s.Instructions[instruction.Number] = instruction
}

func ParseInstruction(instruction string) (Instruction, error) {
	if !InstructionRegex.MatchString(instruction) {
		return Instruction{}, errors.New("no match")
	}
	s := InstructionRegex.FindStringSubmatch(instruction)
	if s == nil || len(s) != 4 {
		return Instruction{}, errors.New("error while matching")
	}
	id, _ := strconv.Atoi(s[1])
	param, _ := strconv.Atoi(s[3])
	return Instruction{
		Number:     id,
		Identifier: s[2],
		Parameters: []int{param},
	}, nil
}
