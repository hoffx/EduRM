package main

import (
	"log"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

//go:generate qtmoc
type QmlBridge struct {
	core.QObject

	_ func(target, action, data string) `signal:"sendToQml"`
	_ func(source, action, data string) `slot:"sendToGo"`
}

var qmlBridge *QmlBridge
var root *core.QObject

func main() {

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

	// Load the main qml file
	qmlBridge = NewQmlBridge(nil)
	engine.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	qmlBridge.ConnectSendToGo(interpretQmlCommand)
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	root = window.Create(engine.RootContext())

	go doBackgroundTasks()

	// Execute app
	gui.QGuiApplication_Exec()
}

func interpretQmlCommand(source, action, data string) {
	log.Println(source + action + data)

	/*switch source {
	case :
		switch action {
		case :

		}
	}*/
}

func doBackgroundTasks() {

	for t := range time.NewTicker(time.Second * 1).C {
		option := time.Now().Second() % 3
		if option == 0 {
			qmlBridge.SendToQml("topToolBar.Row2.currentCmdText", "write", t.Format(time.ANSIC))
		} else if option == 1 {
			qmlBridge.SendToQml("topToolBar.Row2.currentCmdText", "delete", "")
		} else {
			qmlBridge.SendToQml("topToolBar.Row2.currentCmdText", "delete", "")
		}
	}

}
