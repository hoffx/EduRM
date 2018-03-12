package script

import (
	"fmt"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	fmt.Println(ParseInstruction("12: MOV 1"))
}
