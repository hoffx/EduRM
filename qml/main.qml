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
                id: loadButton
                text: qsTr("Load")
                font.capitalization: Font.MixedCase
            }

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
            Text {
                height: parent.height
                id: sliderText
                text: (speedSlider.value * 2).toLocaleString(Qt.locale("en_US"), "f",2) + " s"
                color: "#ffffff"
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
            Text {
                height: parent.height
                id: currentCmdText
                color: "#ffffff"
                styleColor: "#ffffff"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
            Text {
                height: parent.height
                id: instructionCounterText
                color: "#ffffff"
                styleColor: "#ffffff"
                text: backend.instructionCounter
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
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
                height: parent.parent.height
                selectByMouse: true
                selectByKeyboard: true
                id: textEdit
                focus: true
                wrapMode: TextEdit.Wrap
                onCursorRectangleChanged:  flick.ensureVisible(cursorRectangle)
                padding: 15
                MouseArea {
                    enabled: false
                    cursorShape: Qt.IBeamCursor
                    anchors.fill: parent
                    anchors.margins: parent.padding
                }
            }

            ScrollIndicator.vertical: ScrollIndicator{}
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

            ScrollIndicator.vertical: ScrollIndicator {}
        }
    }
}
