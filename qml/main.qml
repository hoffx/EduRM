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

    ToolBar {
        id: toolBar
        position: ToolBar.Header
        height: 50
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.top: parent.top

        Row {
            anchors.leftMargin: 10
            anchors.left: parent.left
            height: parent.height
            width: parent.width * .7
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
                height: parent.height
                id: sliderText
                text: (speedSlider.value * 2).toLocaleString(Qt.locale("en_US"), "f",2) + " s"
                color: "#ffffff"
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
            Item {
                height: parent.height
                width: 40
            }
            Switch {
                id: bpSwitch
                objectName: "bpSwitch"
                scale: 0.8
                checked: true
            }
        }
        Row {
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
                id: currentCmdText
                text: qsTr("LOAD 4")
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

            function ensureVisible(r) {
                if (contentX >= r.x)
                    contentX = r.x
                else if (contentX + width <= r.x + r.width)
                    contentX = r.x + r.width - width
                if (contentY >= r.y)
                    contentY = r.y
                else if (contentY + height <= r.y + r.height)
                    contentY = r.y + r.height - height
            }

            Column{
                id: filesColumn
                width: parent.width

                Repeater {
                    id: filesRepeater
                    model: 25
                    delegate:Item {
                        width: parent.width
                        height: 50

                        MouseArea {
                            anchors.fill: parent
                            hoverEnabled: true

                            RowLayout{
                                anchors.fill: parent

                                Item {
                                    height: parent.height
                                    width: 10
                                }

                                Button {
                                    objectName: "fileName" + index
                                    height: parent.height
                                    //width: text.length > 15 ? parent.width - 30 - 2 * height : undefined
                                    Layout.fillWidth: true
                                    flat: true
                                    font.capitalization: Font.MixedCase

                                    Text {
                                        anchors.fill: parent
                                        text: "filename" + index
                                        verticalAlignment: Text.AlignVCenter
                                        horizontalAlignment: Text.AlignLeft
                                        padding: 5
                                        elide: parent.parent.parent.containsMouse ? Text.ElideLeft : Text.ElideNone
                                    }

                                }

                                ToolButton {
                                    objectName: "saveFile" + index
                                    height: parent.height
                                    width: height
                                    opacity: parent.parent.containsMouse ? 1 : 0
                                    Image{
                                        anchors.fill: parent
                                        scale: 0.5
                                        source: "img/save.png"
                                    }
                                }
                                ToolButton {
                                    objectName: "closeFile" + index
                                    height: parent.height
                                    width: height
                                    opacity: parent.parent.containsMouse ? 1 : 0
                                    Image{
                                        anchors.fill: parent
                                        scale: 0.5
                                        source: "img/close.png"
                                    }
                                }
                                Item {
                                    height: parent.height
                                    width: 10
                                }
                            }
                        }
                    }
                }
                Item{
                    width: parent.width
                    height: 30
                }
                Row{
                    width: parent.width
                    height: 50
                    padding: 15

                    TextField {
                        padding: 5
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

            ScrollIndicator.vertical: ScrollIndicator{}
        }


        Flickable {
            id: flick
            width: parent.width * .5
            height: parent.height
            flickableDirection: Flickable.VerticalFlick
            boundsBehavior: Flickable.StopAtBounds


            TextArea.flickable: TextArea {
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
                id: registerGrid
                columns: width / 85
                width: parent.width
                Repeater {
                    id: registersRepeater
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

   }
}
