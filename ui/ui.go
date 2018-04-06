package ui

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hoffx/EduRM/filemanager"
	"github.com/hoffx/EduRM/interpreter"
	"github.com/hoffx/hermes"

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
	hController.AddEventListener("read", receivePropertyInfo)
	hermes.DoLog = true

	// Load the main qml file
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	root = window.Create(engine.RootContext())

	// Execute app
	gui.QGuiApplication_Exec()
}

// filesystem event listeners:

func receivePropertyInfo(source, jsondata string) {
	log.Println(jsondata)
}

func addFileToSystem(source, jsondata string) {
	type FileInfo struct {
		Path string
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
	storeFileContent("", "")
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

var fileToShow string

func showFile(source, jsondata string) {
	id := strings.TrimPrefix(source, "file")
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil && jsondata != "" {
		log.Fatal(err)
	} else if jsondata == "" {
		// request current file's content
		fileToShow = id
		hController.ReadQml(TextArea_FileContent, hermes.BuildReadModeJSON(Event_ShowFile, "text"))
	} else {
		// set file content
		filemanager.Current().SetText(regexp.MustCompile(`\\n`).ReplaceAllString(regexp.MustCompile(`\\r\\n`).ReplaceAllString(f.Text, "\r\n"), "\n"))
		displayFile(fileToShow)
	}

}

func removeFile(source, jsondata string) {
	file := filemanager.GetByID(strings.TrimPrefix(source, "file"))
	if file != nil {
		filemanager.Remove(file.Name())
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
	if source == "" {
		// request file content
		hController.ReadQml(TextArea_FileContent, hermes.BuildReadModeJSON(Event_StoreFileContent, "text"))
	} else {
		// store file content
		type File struct {
			Text string `json:"text"`
		}
		var f File
		err := json.Unmarshal([]byte(jsondata), &f)
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     err.Error(),
				Instruction: -1,
			})
		} else {
			file := filemanager.GetByID(strings.TrimPrefix(source, "file"))
			if file != nil {
				file.SetText(f.Text)
			}
		}
	}
}

// help functions:

func displayFile(id string) {
	// set textarea content to current file text
	filemanager.SetCurrent(filemanager.GetByID(id).Name())
	hController.SetInQml(TextArea_FileContent, hermes.BuildSetModeJSON("text", filemanager.GetByID(id).Text()))
}

func pushNotification(notification interpreter.Notification) {
	log.Println(notification.Content)
}
