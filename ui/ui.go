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

	go func() {
		time.Sleep(time.Second)
		fmt.Println("test")
		// Access registers
		for _, r := range obj.FindChildren2(core.NewQRegExp2("register*", core.Qt__CaseSensitive, core.QRegExp__Wildcard), core.Qt__FindChildrenRecursively) {
			fmt.Println("-------")
			r.DumpObjectInfo()
		}
	}()

	// Execute app
	gui.QGuiApplication_Exec()

	/*
		widgets.NewQApplication(len(os.Args), os.Args)

		var layout = widgets.NewQHBoxLayout()
		layout.AddWidget(newCppWidget(), 0, 0)
		layout.AddWidget(newSeperator(), 0, 0)
		layout.AddWidget(newQmlWidget(), 0, 0)

		var window = widgets.NewQMainWindow(nil, 0)

		var centralWidget = widgets.NewQWidget(window, 0)
		centralWidget.SetLayout(layout)
		window.SetCentralWidget(centralWidget)

		window.Show()

		widgets.QApplication_Exec()
	*/
}
