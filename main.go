package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/hoffx/EduRM/ui"
)

// Version holds the version-string. It is set during the
// build process via the "make deploy" command.
var Version string

// GitCommit holds the comment-id of the latest commit at
// build time. It is set during the build process via
// the "make deploy" command.
var GitCommit string

// GuiPath holds the path to the main.qml file.
// It is set during the build process via the "make
// deploy" command.
var GuiPath string

func main() {
	if Version == "" {
		Version = GitCommit
	}
	if GuiPath == "" {
		GuiPath = "qml/main.qml"
	}
	if !path.IsAbs(GuiPath) {
		GuiPath = filepath.Dir(os.Args[0]) + "/" + GuiPath
	}
	var args string
	for _, a := range os.Args {
		args += a + " | "
	}
	ioutil.WriteFile("/Volumes/Data/theMomax/edurm.txt", []byte(args), os.ModePerm)
	ui.Run(GuiPath, Version)
}
