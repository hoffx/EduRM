import QtQuick 2.7
import QtQuick.Window 2.2
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3
ApplicationWindow {
    id: window
    visible: true
    title: "EduRM"
    minimumWidth: 640
    minimumHeight: 480
    width: Screen.desktopAvailableWidth
    height: Screen.desktopAvailableHeight

    Connections
    {
        target: qmlBridge
        onSendToQml:
        {
            switch(target) {
            case "topToolBar.Row2.currentCmdText":
                switch(action) {
                case "write":
                    currentCmdText.text = datacontent
                    break;
                case "delete":
                    currentCmdText.text = ""
                    break;
                default:
                    currentCmdText.text = ""
                    break;
                }
                break;
            case "filemanagement":
                switch(action) {
                case "add":
                    filepath.text = ""
                    Qt.createQmlObject(openFile('filelistitem.qml').replace("datacontent", '"' + datacontent + '"'), filesColumn, 'filelistitem.qml')
                }
                break;
            default:
                currentCmdText.text = ""
                break;
            }
        }
    }

    function openFile(fileUrl) {
        var request = new XMLHttpRequest();
        request.open("GET", fileUrl, false);
        request.send(null);
        return request.responseText;
    }

    ToolBar {
        objectName: "topToolBar"
        id: toolBar
        position: ToolBar.Header
        height: 50
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.top: parent.top

        Row {
            objectName: "topToolBar.Row"
            anchors.leftMargin: 10
            anchors.left: parent.left
            height: parent.height
            width: parent.width * .7
            ToolButton {
                objectName: "topToolBar.Row.loadButton"
                id: loadButton
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/load.png"
                }
                onClicked: qmlBridge.sendToGo(this.objectName, "click", "reload")                
            }
            ToolButton {
                id: runButton
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/run.png"
                }
            }
            ToolButton {
                id: stepButton
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/step.png"
                }
            }
            ToolButton {
                id: pauseButton
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/pause.png"
                }
            }
            ToolButton {
                id: stopButton
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/stop.png"
                }
            }
            Item {
                height: parent.height
                width: 40
            }
            Slider {
                id: speedSlider
                width: 100
            }
            Text {
                objectName: "topToolBar.Row.sliderText"
                height: parent.height
                id: sliderText
                text: (speedSlider.value * 5).toLocaleString(Qt.locale("en_US"), "f",1) + " s"
                color: "#ffffff"
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }            
        }
        Row {
            objectName: "topToolBar.Row2"
            anchors.right: parent.right
            height: parent.height
            width: parent.width * .3
            Text {
                height: parent.height
                id: instructionCounterText
                color: "#ffffff"
                text: qsTr("25")
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pointSize: 18
            }
            Text {
                padding: 5
                height: parent.height
                color: "#ffffff"
                text: qsTr(":")
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pointSize: 18
            }
            Text {
                padding: 5
                objectName: "topToolBar.Row2.currentCmdText"
                id: currentCmdText
                //text: qsTr("LOAD 4")
                height: parent.height
                color: "#ffffff"
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pointSize: 18
            }
        }
        Text {
            anchors.right: parent.right
            anchors.rightMargin: 15
            height: parent.height
            id: accumulatorText
            text: qsTr("-12")
            color: "#ffffff"
            styleColor: "#ffffff"
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pointSize: 18
        }
    }

   Row {
        anchors.topMargin: toolBar.height
        anchors.fill: parent
        anchors.bottomMargin: bpBar.height

        Flickable {
            width: parent.width * .2
            height: parent.height
            flickableDirection: Flickable.VerticalFlick
            boundsBehavior: Flickable.StopAtBounds
            contentHeight: filesColumn.implicitHeight
            clip: true

            Column{
                id: filesColumn
                width: parent.width
                Row{
                    width: parent.width
                    height: 50
                    padding: 15

                    TextField {
                        padding: 5
                        id: filepath
                        objectName: "filepath"
                        placeholderText: "filepath"
                        width: parent.width - 30 - 2 * height
                        height: parent.height
                        verticalAlignment: Text.AlignVCenter
                        layer.enabled: true
                        font.pointSize: 14
                    }
                    ToolButton {
                        objectName: "addFileFromFilepath"
                        height: parent.height
                        width: height
                        Image{
                            anchors.fill: parent
                            scale: 0.5
                            source: "img/add.png"
                        }
                        onClicked: qmlBridge.sendToGo(this.objectName, "add", filepath.text)
                    }
                    ToolButton {
                        objectName: "openFile"
                        height: parent.height
                        width: height
                        Image{
                            anchors.fill: parent
                            scale: 0.5
                            source: "img/open.png"
                        }
                    }
                }
                Item{
                    width: parent.width
                    height: 30
                }
            }
            Item{
                width: parent.width
                height: 30
            }

            ScrollIndicator.vertical: ScrollIndicator{}
        }


        Flickable {
            id: flick
            width: parent.width * .5
            height: parent.height
            flickableDirection: Flickable.VerticalFlick
            boundsBehavior: Flickable.StopAtBounds


            TextArea.flickable: TextArea {
                objectName: "BodyRow.TextArea"
                font.pointSize: 13
                width: parent.parent.width
                height: parent.parent.height
                selectByMouse: true
                selectByKeyboard: true
                id: textEdit
                focus: true
                wrapMode: TextArea.Wrap
                padding: 15
                background: null
                font.family: "Menlo, Monaco, 'Courier New', monospace"
                text: "baum"

                MouseArea {
                    enabled: false
                    cursorShape: Qt.IBeamCursor
                    anchors.top: parent.top
                    anchors.topMargin: parent.padding
                    anchors.bottomMargin: parent.padding
                    height: parent.paintedHeight
                    width: parent.width
                }
            }

            ScrollBar.vertical: ScrollBar {}
            //ScrollIndicator.vertical: ScrollIndicator{}
        }

        Flickable {
            clip: true
            width: parent.width * .30
            height: parent.height
            boundsBehavior: Flickable.StopAtBounds
            contentHeight: registerGrid.implicitHeight + 30

            flickableDirection: Flickable.VerticalFlick
            Grid {
                objectName: "body.registerGrid"
                id: registerGrid
                columns: width / 85
                width: parent.width
                Repeater {
                    objectName: "body.registerGrid.Repeater"
                    id: registerRepeater
                    model: 99
                    delegate: Column{
                        width: parent.width / parent.columns
                        height: width

                        Text {
                            color: "#3f51b5"
                            objectName: "registerName" + index
                            width: parent.width
                            height: 3 * parent.height / 5
                            text: "R"+index
                            verticalAlignment: Text.AlignBottom
                            horizontalAlignment: Text.AlignHCenter
                            padding: 5
                        }
                        Text {
                            objectName: "register" + index
                            width: parent.width
                            height: 2 * parent.height / 5
                            text: index
                            verticalAlignment: Text.AlignVCenter
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 18
                        }

                    }
                }
            }

            ScrollIndicator.vertical: ScrollIndicator {}
        }
    }



   ToolBar {
       id: bpBar
       position: ToolBar.Footer
       height: 50
       anchors.left: parent.left
       anchors.bottom: parent.bottom
       anchors.right: parent.right

       Row{
           anchors.fill: parent
           anchors.leftMargin: 10
           anchors.rightMargin: 10

           Switch {
               id: bpSwitch
               objectName: "bpSwitch"
               scale: 0.8
               checked: true
               height: parent.height
           }
           Item {
               height: parent.height
               width: 40
           }
           TextField {
               objectName: "bpInput"
               id: bpInput
               height: parent.height
               verticalAlignment: Text.AlignVCenter
               color: "#ffffff"
               validator: IntValidator{bottom: 1;}
               placeholderText: "instruction counter"

           }
           ToolButton {
               objectName: "addBreakPoint"
               height: parent.height
               width: height
               Image{
                   anchors.fill: parent
                   scale: 0.5
                   source: "img/add.png"
               }
           }
           Item {
               height: parent.height
               width: 40
           }
           Flickable {
               clip: true
               width: parent.width - 80 - parent.height - bpInput.width - bpSwitch.width - 10
               height: parent.height
               boundsBehavior: Flickable.StopAtBounds
               contentWidth: bpList.implicitWidth
               flickableDirection: Flickable.HorizontalFlick

               Row {
                   id: bpList
                   height: parent.height
                   Repeater {
                       height: parent.height
                       id: bpRepeater
                       model: 20
                       delegate: ToolButton{
                            height: parent.height

                            contentItem: Text {
                                text: "Breakpoint" + index
                                height: parent.height
                                horizontalAlignment: Text.AlignHCenter
                                verticalAlignment: Text.AlignVCenter
                                color: parent.hovered ? "#e91e63" : "#ffffff"
                            }

                       }
                   }
               }
           }

       }
   }
}
