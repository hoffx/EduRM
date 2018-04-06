package filemanager

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/hoffx/EduRM/script"
)

// File holds a assembly file and further relevant information information
type File struct {
	id          string
	name        string
	path        string
	script      *script.Script
	text        string
	breakpoints map[uint]bool
}

func NewFile(name, path, text string, script *script.Script, breakpoints map[uint]bool) *File {
	return &File{
		id:          strconv.Itoa(rand.Intn(2147483647)),
		name:        name,
		path:        path,
		script:      script,
		text:        text,
		breakpoints: breakpoints,
	}
}

func (f *File) ID() string { return f.id }

func (f *File) Name() string { return f.name }

func (f *File) Path() string { return f.path }

func (f *File) Script() *script.Script { return f.script }

func (f *File) Text() string { return f.text }

func (f *File) Breakpoints() map[uint]bool { return f.breakpoints }

func (f *File) Save() (err error) {
	return ioutil.WriteFile(f.path, []byte(f.text), os.ModePerm)
}

func (f *File) SetBreakpints(breakpoints map[uint]bool) { f.breakpoints = breakpoints }

func (f *File) ParseFile() (err error) {
	f.script, err = script.ParseFile(f.path)
	return
}

func (f *File) SetText(text string) { f.text = text }
