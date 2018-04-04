package ui

import (
	"encoding/json"
	"log"
	"os"

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

	// Load the main qml file
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	root = window.Create(engine.RootContext())

	// Execute app
	gui.QGuiApplication_Exec()
}

func pushNotification(notification interpreter.Notification) {
	log.Println(notification.Content)
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
		qbController.SendToQml(hermes.ModeAddFromFile, Column_FileList, hermes.BuildAddModeJSON("filelistitem.qml", "name", "'"+filemanager.Current().Name()+"'"))
		// TODO: remove this testlines
		//qbController.SendToQml(hermes.ModeSet, Text_CurrentCmd, `{"text":"`+filemanager.Current().Name()+`"}`)
		//qbController.SendToQml(hermes.ModeRemove, "filesColumn", "")
	}
}
