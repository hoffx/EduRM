package ui

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/hoffx/EduRM/filemanager"
	"github.com/hoffx/EduRM/interpreter"
	"github.com/hoffx/hermes"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

var (
	root         *core.QObject
	qbController *hermes.Controller
)

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

	qbController = hermes.NewBridgeController(engine)
	qbController.AddEventListener(Event_AddFile, addFileToSystem)
	qbController.AddEventListener("read", receivePropertyInfo)

	// Load the main qml file
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	root = window.Create(engine.RootContext())

	go func() {
		time.Sleep(time.Second)
		qbController.ReadQml("kevin", hermes.BuildReadModeJSON("read", "filename"))
	}()

	// Execute app
	gui.QGuiApplication_Exec()
}

func pushNotification(notification interpreter.Notification) {
	log.Println(notification.Content)
}

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
	err = filemanager.AddFile(fi.Path)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
	} else {
		qbController.AddToQmlFromFile(Column_FileList, hermes.BuildAddModeJSON(`tmplFile.qml`, "id", "file1", "filename", "'"+filemanager.Current().Name()+"'"))
		//qbController.SendToQml(hermes.ModeAddFromFile, Column_FileList, hermes.BuildAddModeJSON("filelistitem.qml", "name", "'"+filemanager.Current().Name()+"'", "id", "file"+filemanager.Current().Name(), "filename", filemanager.Current().Name()))
		// TODO: remove this testlines
		//qbController.SendToQml(hermes.ModeSet, Text_CurrentCmd, `{"text":"`+filemanager.Current().Name()+`"}`)
		time.Sleep(time.Second)
		qbController.SetInQml("file1", hermes.BuildSetModeJSON("filename", "'tree'"))
		qbController.RemoveFromQml("sliderText")
	}
}
