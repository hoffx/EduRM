package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/hoffx/EduRM/ui"
	"runtime"
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
	if !path.IsAbs(GuiPath) && runtime.GOOS != "windows" {
		GuiPath = filepath.Dir(os.Args[0]) + "/" + GuiPath
	}
	ui.Run(GuiPath, Version)
}
