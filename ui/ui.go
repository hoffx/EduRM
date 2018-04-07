package ui

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hoffx/EduRM/filemanager"
	"github.com/hoffx/EduRM/interpreter"
	"github.com/theMomax/hermes"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

var (
	root        *core.QObject
	hController *hermes.Controller

	alphanumRegex *regexp.Regexp
)

func init() {
	alphanumRegex, _ = regexp.Compile("[^a-zA-Z0-9]+")
}

func Run() {

	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("HoffX")
	core.QCoreApplication_SetApplicationName("EduRM")
	core.QCoreApplication_SetApplicationVersion("development")
	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	hController = hermes.NewBridgeController(engine)
	hController.AddEventListener(Event_AddFile, addFileToSystem)
	hController.AddEventListener(Event_RemoveFile, removeFile)
	hController.AddEventListener(Event_SaveFile, saveFile)
	hController.AddEventListener(Event_StoreFileContent, storeFileContent)
	hController.AddEventListener(Event_ShowFile, showFile)
	hController.AddEventListener(Event_AddBreakpoint, addBreakpoint)
	hController.AddEventListener(Event_RemoveBreakpoint, removeBreakpoint)
	hermes.DoLog = true

	// Load the main qml file
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	root = window.Create(engine.RootContext())

	// Execute app
	gui.QGuiApplication_Exec()
}

// event listeners:

func receivePropertyInfo(source, jsondata string) {
	log.Println(jsondata)
}

func addFileToSystem(source, jsondata string) {
	type FileInfo struct {
		Path string
		Text string
	}
	var fi FileInfo
	err := json.Unmarshal([]byte(jsondata), &fi)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
	}
	if filemanager.Current() != nil {
		storeFileContent("file"+filemanager.Current().ID(), hermes.BuildSetModeJSON("text", fi.Text))
		hController.SetInQml("file"+filemanager.Current().ID(), hermes.BuildSetModeJSON("current", "false"))
		deleteAllBreakpoints(filemanager.Current().Breakpoints())
	}
	err = filemanager.AddFile(fi.Path)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
	} else {
		hController.AddToQmlFromFile(Column_FileList, hermes.BuildAddModeJSON(`tmplFile.qml`, "id", "file"+filemanager.Current().ID(), "filename", filemanager.Current().Name()))
		hController.SetInQml(TextField_FilePath, hermes.BuildSetModeJSON("text", ""))
		displayFile(strings.TrimPrefix(filemanager.Current().ID(), "file"))
		for _, f := range filemanager.GetAll() {
			if f.Name() != filemanager.Current().Name() {
				hController.SetInQml("file"+f.ID(), hermes.BuildSetModeJSON("filename", f.Name()))
			}
		}
	}
}

func showFile(source, jsondata string) {
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil {
		log.Fatal(err)
	} else {
		// set the old file's content and delete its breakpoints
		if filemanager.Current() != nil {
			filemanager.Current().SetText(regexp.MustCompile(`\\n`).ReplaceAllString(regexp.MustCompile(`\\r\\n`).ReplaceAllString(f.Text, "\r\n"), "\n"))
			deleteAllBreakpoints(filemanager.Current().Breakpoints())
			hController.SetInQml("file"+filemanager.Current().ID(), hermes.BuildSetModeJSON("current", "false"))
		}

		displayFile(strings.TrimPrefix(source, "file"))
	}

}

func removeFile(source, jsondata string) {
	file := filemanager.GetByID(strings.TrimPrefix(source, "file"))

	if file != nil {
		filemanager.Remove(file.Name())
		deleteAllBreakpoints(file.Breakpoints())
	}
	hController.RemoveFromQml(source)
	hController.SetInQml(TextArea_FileContent, hermes.BuildSetModeJSON("text", ""))
}

func saveFile(source, jsondata string) {
	id := strings.TrimPrefix(source, "file")
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil && jsondata != "" {
		log.Fatal(err)
	}

	f.Text = regexp.MustCompile(`\\n`).ReplaceAllString(regexp.MustCompile(`\\r\\n`).ReplaceAllString(f.Text, "\r\n"), "\n")
	if id == filemanager.Current().ID() {
		log.Println(f.Text)
		filemanager.Current().SetText(f.Text)
		err = filemanager.Current().Save()
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     err.Error(),
				Instruction: -1,
			})
		}
	} else {
		// save non-current file (no changes, since not opened)
		filemanager.GetByID(id).Save()
	}
}

func storeFileContent(source, jsondata string) {
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil {
		log.Fatal(err)
	}
	file := filemanager.GetByID(strings.TrimPrefix(source, "file"))
	if file != nil {
		file.SetText(f.Text)
	}
}

func addBreakpoint(source, jsondata string) {
	type Breakpoint struct {
		InstructionCounter string `json:"instructioncounter"`
	}
	var bp Breakpoint
	err := json.Unmarshal([]byte(jsondata), &bp)
	if err == nil {
		if filemanager.Current() != nil {
			ic, err := strconv.Atoi(bp.InstructionCounter)
			if err == nil {
				filemanager.Current().AddBreakpoint(uint(ic))
				displayBreakpoint(uint(ic))
				hController.SetInQml(TextField_BreakpointIC, hermes.BuildSetModeJSON("text", ""))
			}
		}
	}
}

func removeBreakpoint(source, jsondata string) {
	ic, err := strconv.Atoi(strings.TrimPrefix(source, "breakpoint"))
	if err != nil {
		log.Fatal(err)
	}
	if filemanager.Current() != nil {
		filemanager.Current().DeleteBreakpoint(uint(ic))
	}
	hController.RemoveFromQml(source)
}

// helper functions:

func displayBreakpoint(ic uint) {
	bpString := strconv.Itoa(int(ic))
	hController.AddToQmlFromFile(Row_Breakpoints, hermes.BuildAddModeJSON("tmplBreakpoint.qml", "id", "breakpoint"+bpString, "idtext", "breakpoint"+bpString, "buttontext", bpString))
}

func displayFile(id string) {
	// set current file to requested id
	filemanager.SetCurrent(filemanager.GetByID(id).Name())
	// set filelistitem's current property to true
	hController.SetInQml("file"+id, hermes.BuildSetModeJSON("current", "true"))
	// set textarea content to current file text
	hController.SetInQml(TextArea_FileContent, hermes.BuildSetModeJSON("text", filemanager.GetByID(id).Text()))
	// load breakpoints
	for bp := range filemanager.Current().Breakpoints() {
		displayBreakpoint(bp)
	}
}

func deleteAllBreakpoints(bp map[uint]bool) {
	for ic := range bp {
		hController.RemoveFromQml("breakpoint" + strconv.Itoa(int(ic)))
	}
}

func pushNotification(notification interpreter.Notification) {
	log.Println(notification.Content)
}
