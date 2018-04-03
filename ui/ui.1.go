// package ui

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/therecipe/qt/core"
// 	"github.com/therecipe/qt/gui"
// 	"github.com/therecipe/qt/qml"
// 	"github.com/therecipe/qt/quickcontrols2"
// )

// func Run() {

// 	// Create application
// 	app := gui.NewQGuiApplication(len(os.Args), os.Args)
// 	core.QCoreApplication_SetOrganizationName("HoffX")
// 	core.QCoreApplication_SetApplicationName("EduRM")
// 	core.QCoreApplication_SetApplicationVersion("development")
// 	// Enable high DPI scaling
// 	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

// 	// Use the material style for qml
// 	quickcontrols2.QQuickStyle_SetStyle("material")

// 	// Create a QML application engine
// 	engine := qml.NewQQmlApplicationEngine(nil)

// 	// Load the main qml file

// 	window := qml.NewQQmlComponent5(engine, core.NewQUrl3("qml/main.qml", 0), nil)
// 	obj := window.Create(engine.RootContext())

// 	go func() {
// 		time.Sleep(time.Second)
// 		fmt.Println("test")
// 		// Access registers
// 		//children := obj.FindChildren
// 		/*for _, r := range obj.FindChildren2(core.NewQRegExp2("register*", core.Qt__CaseSensitive, core.QRegExp__Wildcard), core.Qt__FindChildrenRecursively) {
// 			fmt.Println("-------")
// 			r.DumpObjectInfo()
// 		}*/
// 		loadButton := obj.FindChild("topToolBar.Row.loadButton", core.Qt__FindChildrenRecursively)
// 		loadButton.DumpObjectInfo()

// 		fmt.Println("-------")
// 		sliderText := obj.FindChild("topToolBar.Row.sliderText", core.Qt__FindChildrenRecursively)
// 		sliderText.DumpObjectInfo()
// 		fmt.Println("-------")
// 		sliderTextText := sliderText.Property("text")
// 		fmt.Println(sliderTextText.ToString())
// 		/*loadButton.ConnectEvent(func(e *core.QEvent) bool {
// 			fmt.Println(e)
// 			return true
// 		})*/
// 	}()

// 	// Execute app
// 	gui.QGuiApplication_Exec()

// 	/*
// 		widgets.NewQApplication(len(os.Args), os.Args)

// 		var layout = widgets.NewQHBoxLayout()
// 		layout.AddWidget(newCppWidget(), 0, 0)
// 		layout.AddWidget(newSeperator(), 0, 0)
// 		layout.AddWidget(newQmlWidget(), 0, 0)

// 		var window = widgets.NewQMainWindow(nil, 0)

// 		var centralWidget = widgets.NewQWidget(window, 0)
// 		centralWidget.SetLayout(layout)
// 		window.SetCentralWidget(centralWidget)

// 		window.Show()

// 		widgets.QApplication_Exec()
// 	*/
// }
