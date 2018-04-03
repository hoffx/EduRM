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

	_ func(data string)                 `signal:"sendToQml"`
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

	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendToQml(t.Format(time.ANSIC))
		}
	}()

	// Execute app
	gui.QGuiApplication_Exec()
}

func interpretQmlCommand(source, action, data string) {
	log.Println(source + action + data)
	if source == "topToolBar.Row.loadButton" && action == "click" && data == "reload" {
		/*textarea := root.FindChild("BodyRow.TextArea", core.Qt__FindChildrenRecursively)
		textarea.DumpObjectInfo()
		textAreaText := textarea.Property("text")
		text := "baum"
		textAreaText = &core.QVariant{unsafe.Pointer{}}*/
	}
}

/*
//go:generate qtmoc
type QmlBridge struct {
	core.QObject

	_ func(source, action, data string) `signal:"sendToQml"`
	_ func(source, action, data string) `slot:"sendToGo"`
}

var qmlBridge *QmlBridge

func main() {
	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	app.NewQHBox

	layout.AddWidget(newQmlWidget(), 0, 0)

	// Execute app
	gui.QGuiApplication_Exec()
}

func newQmlWidget() *quick.QQuickWidget {
	var quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	//nitQmlContext(quickWidget)
	initQmlBridge(quickWidget)

	quickWidget.SetSource(core.NewQUrl3("qml/bridge.qml", 0))

	return quickWidget
}

func initQmlBridge(quickWidget *quick.QQuickWidget) {

	qmlBridge = NewQmlBridge(nil)
	quickWidget.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	qmlBridge.ConnectSendToGo(func(source, action, data string) {
		log.Println(source + action + data)
	})
}
*/
