package script

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	s, err := ParseFile("test.txt")
	if err != nil {
		t.Error(err, s)
	}
}

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		instruction string
		want        Instruction
		wantErr     bool
	}{
		{"1: JUMP 17", Instruction{1, "JUMP", []int{17}}, false},
		{"1: JUMP 17 -- comment", Instruction{1, "JUMP", []int{17}}, false},
		{"13 END", Instruction{13, "END", []int{}}, false},
		{"13 END--rehkirz", Instruction{13, "END", []int{}}, false},
		{"-3: JUMP 18", Instruction{}, true},
		{"-3: JUMP 18 -- prayforhoffi", Instruction{}, true},
		{"14: END - hohoho", Instruction{}, true},
		{"3: JUM1P 18", Instruction{}, true},
		{"a JUM1P 18", Instruction{}, true},
		{"1 JUMP a", Instruction{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.instruction, func(t *testing.T) {
			got, err := ParseInstruction(tt.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstruction() ret = %v error = %v, wantErr %v", got, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
