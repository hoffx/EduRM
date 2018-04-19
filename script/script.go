package script

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var instructionRegex = regexp.MustCompile(`^([0-9]+)[:\s]+([a-zA-Z]+)(\s+(-?[0-9]+))?\s*(--.*)?$`)

// Script represents a script written in assembly ready to be interpreted
type Script struct {
	Name         string
	Path         string
	Instructions map[int]Instruction
}

// Instruction represents a single assembly instruction
type Instruction struct {
	Number     int
	Identifier string
	Parameters []int
}

// NewScript creates a Script by name
func NewScript(name string) *Script {
	return &Script{
		Name:         name,
		Instructions: make(map[int]Instruction),
	}
}

// AppendInstruction adds a given Instruction to the Script
func (s *Script) AppendInstruction(instruction Instruction) {
	s.Instructions[instruction.Number] = instruction
}

// ParseInstruction adds an Instruction given as string to the Script
func (s *Script) ParseInstruction(instruction string) error {
	i, err := ParseInstruction(instruction)
	if err != nil {
		return err
	}
	s.AppendInstruction(i)
	return nil
}

// ParseInstruction parses the string representation of an Instruction and returns the Instruction
func ParseInstruction(instruction string) (Instruction, error) {
	if !instructionRegex.MatchString(instruction) {
		return Instruction{}, errors.New("no match")
	}
	s := instructionRegex.FindStringSubmatch(instruction)
	if s == nil || len(s) != 6 {
		return Instruction{}, errors.New("error while matching")
	}
	id, _ := strconv.Atoi(s[1])
	params := make([]int, 0)
	param, err := strconv.Atoi(s[4])
	if err == nil {
		params = append(params, param)
	}
	return Instruction{
		Number:     id,
		Identifier: strings.ToUpper(s[2]),
		Parameters: params,
	}, nil
}

// Parse parses the given reader and returns it's contents as a ready assembly script
func Parse(name string, r io.Reader) (*Script, error) {
	s := NewScript(name)
	rd := bufio.NewReader(r)
	for {
		l, _, err := rd.ReadLine()
		if err == io.EOF && len(l) == 0 {
			break
		} else if err != nil && err != io.EOF {
			return nil, err
		}
		err = s.ParseInstruction(string(l))
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ParseFile parses the file in the given location and returns it's contents as a ready assembly script
func ParseFile(path string) (*Script, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	s, err := Parse(f.Name(), f)
	s.Path = path
	return s, err
}
