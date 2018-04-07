package filemanager

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

// File holds a assembly file and further relevant information information
type File struct {
	id          string
	name        string
	path        string
	text        string
	breakpoints map[uint]bool
}

func NewFile(name, path, text string, breakpoints map[uint]bool) *File {
	return &File{
		id:          strconv.Itoa(rand.Intn(2147483647)),
		name:        name,
		path:        path,
		text:        text,
		breakpoints: breakpoints,
	}
}

func (f *File) ID() string { return f.id }

func (f *File) Name() string { return f.name }

func (f *File) Path() string { return f.path }

func (f *File) Text() string { return f.text }

func (f *File) Breakpoints() map[uint]bool { return f.breakpoints }

func (f *File) Save() (err error) {
	return ioutil.WriteFile(f.path, []byte(f.text), os.ModePerm)
}

func (f *File) SetBreakpoints(breakpoints map[uint]bool) { f.breakpoints = breakpoints }

func (f *File) AddBreakpoint(instructionCounter uint) {
	if f.breakpoints != nil {
		f.breakpoints[instructionCounter] = true
	}
}

func (f *File) DeleteBreakpoint(instructionCounter uint) {
	if f.breakpoints != nil {
		delete(f.breakpoints, instructionCounter)
	}
}

func (f *File) SetText(text string) { f.text = text }
