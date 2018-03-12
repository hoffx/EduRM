package script

import (
	"fmt"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	fmt.Println(ParseInstruction("12: MOV 1"))
}

func TestParseFile(t *testing.T) {
	fmt.Println(ParseFile("test.txt"))
}
