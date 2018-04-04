package filemanager_test

import (
	"log"
	"testing"

	"github.com/hoffx/EduRM/filemanager"
)

func TestAddingFiles(t *testing.T) {
	err := filemanager.AddFile("test.txt")
	if err != nil {
		t.Error(err)
	}
	err = filemanager.AddFile("file:///Volumes/Data/themomax/Development/go/src/github.com/hoffx/EduRM/test.txt")
	if err != nil {
		t.Error(err)
	}
	err = filemanager.AddFile("../script/test.txt")
	if err != nil {
		t.Error(err)
	}
	err = filemanager.AddFile("../controller/test.txt")
	if err != nil {
		t.Error(err)
	}
	files := filemanager.GetAll()
	for k := range files {
		log.Println(k)
	}
}
