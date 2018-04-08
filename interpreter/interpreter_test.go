package interpreter

import (
	"log"
	"testing"

	"github.com/hoffx/EduRM/script"
)

func TestInterpreter(t *testing.T) {
	s, err := script.ParseFile("test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	ctx := NewInterpreterContext(*s, 16)
	for ctx.Status == Running {
		ctx.Next()
		log.Println(ctx.Output)
		log.Println(ctx.Status)
	}
}
