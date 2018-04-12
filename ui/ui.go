package ui

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hoffx/EduRM/controller"
	"github.com/hoffx/EduRM/filemanager"
	"github.com/hoffx/EduRM/interpreter"
	"github.com/theMomax/hermes"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

var (
	// ui
	root        *core.QObject
	hController *hermes.Controller

	// backend
	c              *controller.Controller
	registerCount  int = -1
	delay          int
	status         int
	useBreakpoints bool = true
	ready          bool = false
)

func Run(path, version string) {
	core.QCoreApplication_SetOrganizationName("HoffX")
	core.QCoreApplication_SetApplicationName("EduRM")
	core.QCoreApplication_SetApplicationVersion(version)

	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app.ConnectEvent(event)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	hController = hermes.NewBridgeController(engine)
	hController.AddEventListener(Event_WindowLoaded, windowLoaded)
	hController.AddEventListener(Event_AddFile, addFileToSystem)
	hController.AddEventListener(Event_RemoveFile, removeFile)
	hController.AddEventListener(Event_SaveFile, saveFile)
	hController.AddEventListener(Event_SaveAllFiles, saveAllFiles)
	hController.AddEventListener(Event_ShowFile, showFile)
	hController.AddEventListener(Event_AddBreakpoint, addBreakpoint)
	hController.AddEventListener(Event_RemoveBreakpoint, removeBreakpoint)
	hController.AddEventListener(Event_AddRegister, addRegister)
	hController.AddEventListener(Event_RemoveRegister, removeRegister)
	hController.AddEventListener(Event_SliderMoved, sliderMoved)
	hController.AddEventListener(Event_Run, run)
	hController.AddEventListener(Event_Step, step)
	hController.AddEventListener(Event_Pause, pause)
	hController.AddEventListener(Event_Stop, stop)
	hController.AddEventListener(Event_ToggleBreakpoints, toggleBreakpoints)
	hController.AddEventListener(Event_SaveTempFile, saveTempFileAs)

	// Load the main qml file
	window := qml.NewQQmlComponent5(engine, core.NewQUrl3(path, 0), nil)
	root = window.Create(engine.RootContext())

	// Execute app
	gui.QGuiApplication_Exec()
}

// qt event handler
func event(event *core.QEvent) bool {
	// handle macOS file opening
	if event.Type() == core.QEvent__FileOpen {
		for !ready {
			time.Sleep(50 * time.Millisecond)
		}
		file := gui.NewQFileOpenEventFromPointer(event.Pointer())
		addFileToSystem("", hermes.BuildSetModeJSON("path", file.File(), "text", ""))
		event.Accept()
		return true
	}
	return root.EventDefault(event)
}

// event listeners:

func windowLoaded(source, jsondata string) {
	for i := registerCount; i < 15; i++ {
		addRegister("", "")
	}
	// handle windows & linux file opening
	for i := 1; i < len(os.Args); i++ {
		if _, err := os.Stat(os.Args[i]); err == nil {
			addFileToSystem("", hermes.BuildSetModeJSON("path", os.Args[i], "text", ""))
		}
	}
	ready = true
}

var idToSave string

func saveTempFileAs(source, jsondata string) {
	type FileInfo struct {
		Path string
		Text string
	}
	var fi FileInfo
	err := json.Unmarshal([]byte(jsondata), &fi)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
		return
	}
	var file *filemanager.File
	if filemanager.Current() != nil && filemanager.Current().ID() == strings.TrimPrefix(source, "file") {
		file = filemanager.Current()
		storeFileContent(file.ID(), fi.Text)
	} else {
		file = filemanager.GetByID(strings.TrimPrefix(source, "file"))
		if file == nil {
			log.Fatal("illegal state")
		}
	}
	cleanedPath := filemanager.CleanPath(fi.Path)
	if cleanedPath == "" {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     "illegal filepath",
			Instruction: -1,
		})
		return
	}
	err = ioutil.WriteFile(cleanedPath, []byte(file.Text()), 0644)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
		return
	}
	filemanager.Remove(file.Name())
	hController.RemoveFromQml(source)
	addFileToSystem("", jsondata)
}

func addFileToSystem(source, jsondata string) {
	type FileInfo struct {
		Path string
		Text string
	}
	var fi FileInfo
	err := json.Unmarshal([]byte(jsondata), &fi)
	if err != nil {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Error,
			Content:     err.Error(),
			Instruction: -1,
		})
		return
	}

	if filemanager.Current() != nil {
		storeFileContent(filemanager.Current().ID(), fi.Text)
		hController.SetInQml("file"+filemanager.Current().ID(), hermes.BuildSetModeJSON("current", "false"))
		deleteAllBreakpoints(filemanager.Current().Breakpoints())
	}
	if fi.Path == "" {
		filemanager.AddTempFile("")
		hController.AddToQmlFromFile(Column_FileList, hermes.BuildAddModeJSON(`tmplFile.qml`, "id", "file"+filemanager.Current().ID(), "filename", filemanager.Current().Name(), "temp", "true"))
	} else {
		err = filemanager.AddFile(fi.Path)
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     err.Error(),
				Instruction: -1,
			})
			return
		}
		hController.AddToQmlFromFile(Column_FileList, hermes.BuildAddModeJSON(`tmplFile.qml`, "id", "file"+filemanager.Current().ID(), "filename", filemanager.Current().Name(), "temp", "false"))
	}
	hController.SetInQml(TextField_FilePath, hermes.BuildSetModeJSON("text", ""))
	displayFile(strings.TrimPrefix(filemanager.Current().ID(), "file"))
	for _, f := range filemanager.GetAll() {
		if f.Name() != filemanager.Current().Name() {
			hController.SetInQml("file"+f.ID(), hermes.BuildSetModeJSON("filename", f.Name()))
		}
	}
}

func showFile(source, jsondata string) {
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil {
		log.Fatal(err)
	} else {
		// set the old file's content and delete its breakpoints
		if filemanager.Current() != nil {
			filemanager.Current().SetText(f.Text)
			deleteAllBreakpoints(filemanager.Current().Breakpoints())
			hController.SetInQml("file"+filemanager.Current().ID(), hermes.BuildSetModeJSON("current", "false"))
		}

		displayFile(strings.TrimPrefix(source, "file"))
	}

}

func removeFile(source, jsondata string) {
	file := filemanager.GetByID(strings.TrimPrefix(source, "file"))
	if file != nil {
		if file == filemanager.Current() {
			hController.SetInQml(TextArea_FileContent, hermes.BuildSetModeJSON("text", "", "hidden", "true"))
		}
		deleteAllBreakpoints(file.Breakpoints())
		filemanager.Remove(file.Name())
	}
	hController.RemoveFromQml(source)
}

func saveFile(source, jsondata string) {
	var id string
	if source == "" {
		id = filemanager.Current().ID()
	} else {
		id = strings.TrimPrefix(source, "file")
	}
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil && jsondata != "" {
		log.Fatal(err)
	}
	if id == filemanager.Current().ID() {
		filemanager.Current().SetText(f.Text)
		err = filemanager.Current().Save()
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     err.Error(),
				Instruction: -1,
			})
			return
		}
	} else {
		// save non-current file (no changes, since not opened)
		err = filemanager.GetByID(id).Save()
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     err.Error(),
				Instruction: -1,
			})
			return
		}
	}
}

func saveAllFiles(source, jsondata string) {
	type File struct {
		Text string `json:"text"`
	}
	var f File
	err := json.Unmarshal([]byte(jsondata), &f)
	if err != nil && jsondata != "" {
		log.Fatal(err)
	}

	if filemanager.Current() != nil {
		filemanager.Current().SetText(f.Text)
	}
	for _, f := range filemanager.GetAll() {
		if f.IsTemp() {
			saveTempFileAs("file"+f.ID(), hermes.BuildSetModeJSON("text", f.Text(), "path", ""))
		} else {
			err = f.Save()
			if err != nil {
				pushNotification(interpreter.Notification{
					Type:        interpreter.Error,
					Content:     err.Error(),
					Instruction: -1,
				})
				err = nil
			}
		}
	}
}

func addBreakpoint(source, jsondata string) {
	type Breakpoint struct {
		InstructionCounter string `json:"instructioncounter"`
	}
	var bp Breakpoint
	err := json.Unmarshal([]byte(jsondata), &bp)
	if err == nil {
		if filemanager.Current() != nil {
			ic, err := strconv.Atoi(bp.InstructionCounter)
			if err == nil {
				filemanager.Current().AddBreakpoint(uint(ic))
				displayBreakpoint(uint(ic))
				hController.SetInQml(TextField_BreakpointIC, hermes.BuildSetModeJSON("text", ""))
			}
		}
	}
}

func removeBreakpoint(source, jsondata string) {
	ic, err := strconv.Atoi(strings.TrimPrefix(source, "breakpoint"))
	if err != nil {
		log.Fatal(err)
	}
	if filemanager.Current() != nil {
		filemanager.Current().DeleteBreakpoint(uint(ic))
	}
	hController.RemoveFromQml(source)
}

func toggleBreakpoints(source, jsondata string) {
	if status != interpreter.Running {
		useBreakpoints = !useBreakpoints
	} else {
		hController.SetInQml(Switch_Breakpoints, hermes.BuildSetModeJSON("checked", strconv.FormatBool(useBreakpoints)))
	}
}

func sliderMoved(source, jsondata string) {
	type SliderInfo struct {
		Value float64 `json:"value"`
	}
	var si SliderInfo
	err := json.Unmarshal([]byte(jsondata), &si)
	if err != nil {
		log.Fatal(err)
	}
	delay = int(1000 * 5 * si.Value)
	if c != nil {
		go c.SetDelay(delay)
	}
}

func addRegister(source, jsondata string) {
	if status != interpreter.Running {
		registerCount++
		hController.AddToQmlFromFile(Grid_RegisterList, hermes.BuildAddModeJSON("tmplRegister.qml", "id", "register"+strconv.Itoa(registerCount), "number", strconv.Itoa(registerCount), "value", "0"))
	} else {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Warning,
			Content:     "invalid status",
			Instruction: -1,
		})
	}
}

func removeRegister(source, jsondata string) {
	if status != interpreter.Running {
		if registerCount != -1 {
			hController.RemoveFromQml("register" + strconv.Itoa(registerCount))
			registerCount--
		}
	} else {
		pushNotification(interpreter.Notification{
			Type:        interpreter.Warning,
			Content:     "invalid status",
			Instruction: -1,
		})
	}
}

func run(source, jsondata string) {
	if c == nil {
		type File struct {
			Text string `json:"text"`
		}
		var f File
		err := json.Unmarshal([]byte(jsondata), &f)
		if err != nil {
			log.Fatal(err)
		}
		c, err = controller.NewController(f.Text, registerCount+1)
		if err != nil {
			pushNotification(interpreter.Notification{
				Type:        interpreter.Error,
				Content:     "parse failed: " + err.Error(),
				Instruction: -1,
			})
			return
		}
		go func() {
			hController.SetInQml(ToolBar_BreakpointsBarHider, hermes.BuildSetModeJSON("hide", "true"))
			c.Process()
			hController.SetInQml(ToolBar_BreakpointsBarHider, hermes.BuildSetModeJSON("hide", "false"))
		}()
		go c.SetDelay(delay)
		if useBreakpoints {
			if filemanager.Current() != nil {
				for bp := range filemanager.Current().Breakpoints() {
					go c.AddBreakpoint(bp)
				}
			}
		}
		go func() {
			status = interpreter.Running
			for status == interpreter.Running {
				select {
				case ctx := <-c.ContextChan:
					status = ctx.Status
					displayContext(ctx)
				}
			}
			c = nil
		}()
	}
	if c != nil {
		go c.Run()
	}
}

func step(source, jsondata string) {
	if c != nil {
		go c.Step()
	}
}

func pause(source, jsondata string) {
	if c != nil {
		go c.Pause()
	}
}

func stop(source, jsondata string) {
	if c != nil {
		go c.Stop()
	}
}

// helper functions:

func displayContext(ctx interpreter.Context) {
	hController.SetInQml(Text_InstructionCounter, hermes.BuildSetModeJSON("text", strconv.Itoa(int(ctx.InstructionCounter))))
	hController.SetInQml(Text_Instruction, hermes.BuildSetModeJSON("text", ctx.Instruction))
	hController.SetInQml(Text_Accumulator, hermes.BuildSetModeJSON("text", strconv.Itoa(ctx.Accumulator)))
	for i, r := range ctx.Registers {
		hController.SetInQml("register"+strconv.Itoa(i), hermes.BuildSetModeJSON("value", strconv.Itoa(r)))
	}
	for _, n := range ctx.Output {
		pushNotification(n)
	}
}

func displayBreakpoint(ic uint) {
	bpString := strconv.Itoa(int(ic))
	hController.AddToQmlFromFile(Row_Breakpoints, hermes.BuildAddModeJSON("tmplBreakpoint.qml", "id", "breakpoint"+bpString, "idtext", "breakpoint"+bpString, "buttontext", bpString))
}

func displayFile(id string) {
	file := filemanager.GetByID(id)
	// set current file to requested id
	filemanager.SetCurrent(file.Name())
	// set filelistitem's current property to true
	hController.SetInQml("file"+id, hermes.BuildSetModeJSON("current", "true"))
	// set textarea content to current file text
	hController.SetInQml(TextArea_FileContent, hermes.BuildSetModeJSON("text", filemanager.GetByID(id).Text(), "hidden", "false"))
	// load breakpoints
	for bp := range filemanager.Current().Breakpoints() {
		displayBreakpoint(bp)
	}
}

func deleteAllBreakpoints(bp map[uint]bool) {
	for ic := range bp {
		hController.RemoveFromQml("breakpoint" + strconv.Itoa(int(ic)))
	}
}

func pushNotification(notification interpreter.Notification) {
	hController.AddToQmlFromFile(Column_Notifications, hermes.BuildAddModeJSON("tmplNotification.qml", "content", notification.Content, "ic", strconv.Itoa(notification.Instruction), "type", strconv.Itoa(notification.Type)))
}

func storeFileContent(id, content string) {
	file := filemanager.GetByID(id)
	if file != nil {
		file.SetText(content)
	}
}
