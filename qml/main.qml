import QtQuick 2.7
import QtQuick.Controls 2.1

ApplicationWindow {
    id: window
    visible: true
    title: "EduRM"
    minimumWidth: 400
    minimumHeight: 400
    width: 640
    height: 480

    ToolBar {
        id: toolBar
        position: ToolBar.Header
        height: 50
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.top: parent.top

        Row {
            anchors.fill: parent
            ToolButton {
                id: runButton
                text: qsTr("Run")
                font.capitalization: Font.MixedCase
            }
            ToolButton {
                id: stepButton
                text: qsTr("Step")
                font.capitalization: Font.MixedCase
            }
            ToolButton {
                id: pauseButton
                text: qsTr("Pause")
                font.capitalization: Font.MixedCase
            }
            ToolButton {
                id: stopButton
                text: qsTr("Stop")
                font.capitalization: Font.MixedCase
            }
            Slider {
                id: speedSlider
                width: 100
            }
        }
    }

    Row {
        anchors.topMargin: toolBar.height
        anchors.fill: parent
        Flickable {
            id: flick

            width: parent.width * .5
            height: parent.height
            flickableDirection: Flickable.VerticalFlick
            boundsBehavior: Flickable.StopAtBounds
            contentWidth: textEdit.paintedWidth
            contentHeight: textEdit.paintedHeight
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

            TextEdit {
                width: parent.parent.width
                height: parent.height
                id: textEdit
                focus: true
                wrapMode: TextEdit.Wrap
                onCursorRectangleChanged:  flick.ensureVisible(cursorRectangle)
            }
        }
        Flickable {
            clip: true
            width: parent.width * .5
            height: parent.height
            boundsBehavior: Flickable.StopAtBounds
            contentHeight: registerGrid.implicitHeight

            flickableDirection: Flickable.VerticalFlick
            Grid {
                id: registerGrid
                columns: 4
                width: parent.width
                Repeater {
                    id: registersRepeater
                    model: 100
                    delegate: Text {
                        width: parent.width / 4
                        height: width
                        text: "0"
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        padding: 5
                    }
                }
            }
        }
    }
}
