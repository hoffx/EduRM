package ui

import (
	"fmt"
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
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
	obj := window.Create(engine.RootContext())
	for _, o := range obj.Children()[2].Children()[0].Children() {
		fmt.Println("---", o.ObjectName())
		o.DumpObjectInfo()
	}
	//obj.FindChild("textEdit", core.Qt__FindChildrenRecursively).SetProperty("text", core.NewQVariant14("hi"))

	go func() {
		time.Sleep(5 * time.Second)
		obj.FindChild("QQuickTextEdit", core.Qt__FindChildrenRecursively).DumpObjectInfo() //.SetProperty("text", core.NewQVariant14("delay"))
	}()

	// Execute app
	gui.QGuiApplication_Exec()
}
