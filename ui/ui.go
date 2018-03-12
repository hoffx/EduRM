package ui

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func Run() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	mainWidget := buildMainWidget()

	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}

func buildMainWidget() *QWidget {
	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	ipPath := widgets.NewQLineEdit(nil)
	ipPath.SetPlaceholderText("filepath")
	layout.AddWidget(ipPath, 0, 0)

	// Create a button and add it to the layout
	btPath := widgets.NewQPushButton2("submit", nil)
	layout.AddWidget(btPath, 0, 0)

	// Connect event for button
	ptPath.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "OK", ipPath.Text(),
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
}
