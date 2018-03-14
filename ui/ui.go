package ui

import (
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func Run() {
	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	// Load the main qml file
	engine.Load(core.NewQUrl3("qml/main.qml", 0))

	go func() {
		time.Sleep(5 * time.Second)
		//engine.FindChild("textEdit", core.Qt__FindChildrenRecursively). //SetProperty("text", core.NewQVariant14("hi"))
		engine.DumpObjectTree()
	}()

	// Execute app
	gui.QGuiApplication_Exec()
}
