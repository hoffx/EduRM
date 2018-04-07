import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

Item {
    width: parent.width
    height: 50
    property string filename: ""
    property string idtext: ""
    property string current: "false"

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
                height: parent.height
                Layout.fillWidth: true
                flat: true
                font.capitalization: Font.MixedCase

                Text {
                    anchors.fill: parent
                    text: filename
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignLeft
                    padding: 5
                    elide: parent.parent.parent.containsMouse ? Text.ElideLeft : Text.ElideNone
                    color: current == "true" ? "#3f51b5" : "#000000"
                }
                onClicked: hermes.sendToGo("event_showfile", idtext, '{"text":"'+textEdit.text+'"}')
            }
            ToolButton {
                height: parent.height
                width: height
                opacity: parent.parent.containsMouse ? 1 : 0
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/save.png"
                }
                onClicked: hermes.sendToGo("event_savefile", idtext, '{"text":"'+textEdit.text+'"}')
            }
            ToolButton {
                height: parent.height
                width: height
                opacity: parent.parent.containsMouse ? 1 : 0
                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: "img/close.png"
                }
                onClicked: hermes.sendToGo("event_removefile", idtext, "")
            }
            Item {
                height: parent.height
                width: 10
            }
        }
    }
}