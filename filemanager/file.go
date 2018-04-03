package filemanager

import (
	"os"

	"github.com/hoffx/EduRM/script"
)

// File holds a assembly file and further relevant information information
type File struct {
	name        string
	path        string
	script      *script.Script
	text        string
	breakpoints map[uint]bool
}

func (f *File) Name() string { return f.name }

func (f *File) Path() string { return f.path }

func (f *File) Script() *script.Script { return f.script }

func (f *File) Text() string { return f.text }

func (f *File) Breakpoints() map[uint]bool { return f.breakpoints }

func (f *File) Save() (err error) {
	file, err := os.Open(f.path)
	if err != nil {
		return err
	}
	_, err = file.WriteString(f.text)
	return err
}

func (f *File) SetBreakpints(breakpoints map[uint]bool) { f.breakpoints = breakpoints }

func (f *File) ParseFile() (err error) {
	f.script, err = script.ParseFile(f.path)
	return
}
