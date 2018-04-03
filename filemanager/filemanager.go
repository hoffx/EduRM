package filemanager

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// filemanager errors
const (
	ErrInvalidFilepath   string = "invalid filepath"
	ErrFileAlreadyLoaded string = "file already loaded"
)

var (
	files   map[string]*File
	current string
)

func init() {
	files = make(map[string]*File, 0)
}

// AddFile loads the file from the given path
// and sets it as the currently opened file
func AddFile(path string) (err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}
	path = filepath.Clean(path)
	path = filepath.ToSlash(path)
	text, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	name := filepath.Base(path)
	if name == "." || name == "/" {
		return errors.New(ErrInvalidFilepath)
	}

	// create a unique and user-friendly names for all files
	// with same base
	var equalNamedFiles []File
	for k, _ := range files {
		if strings.Contains(k, name) {
			equalNamedFiles = append(equalNamedFiles, *files[k])
		}
	}

	for _, f := range equalNamedFiles {
		if f.path == path {
			return errors.New(ErrFileAlreadyLoaded)
		} else {
			var oldsname, newsname string
			old := strings.Split(f.path, "/")
			new := strings.Split(path, "/")
			length := len(old)
			if len(new) < len(old) {
				length = len(new)
			}
			diff := len(new) - len(old)
			for i := length; i >= 0; i-- {
				o := i - 1
				n := i - 1
				if diff < 0 {
					o -= diff
				} else {
					n += diff
				}
				if old[o] != new[n] {
					newsname = buildName(new[n:])
					oldsname = buildName(old[o:])
					break
				}
			}
			name = newsname
			delete(files, f.name)
			f.name = oldsname
			files[f.name] = &f
		}

	}

	files[name] = &File{
		name:        name,
		path:        path,
		script:      nil,
		text:        string(text),
		breakpoints: make(map[uint]bool, 0),
	}
	current = name
	return
}

// Current returns the currently opened file. If there are
// no files yet, nil is returned
func Current() *File {
	if current == "" {
		return nil
	}
	return files[current]
}

// GetByName returns the file connected to the given identifier
func GetByName(name string) *File {
	return files[name]
}

// GetAll returns a map of all files. The key is the file's name
func GetAll() map[string]*File {
	return files
}

// Remove deletes the according file from the program's
// list, but not from the disk
func Remove(name string) {
	if name == current {
		current = ""
	}
	delete(files, name)
}

// Erase deletes the according file from the program's
// list and from the disk
func Erase(name string) (err error) {
	err = os.Remove(files[name].path)
	if err != nil {
		return
	}
	Remove(name)
	return
}

// SetCurrent sets the currently opened file by its name
func SetCurrent(name string) {
	current = name
}

func buildName(pathelements []string) (name string) {
	name += "..."
	for _, e := range pathelements {
		name += "/" + e
	}
	return
}
