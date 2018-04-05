import QtQuick 2.7
import QtQuick.Window 2.2
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

Item {
    width: parent.width
    height: 50
    id: <id>

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
                    id: <textfieldid>
                    text: <name>
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignLeft
                    padding: 5
                    elide: parent.parent.parent.containsMouse ? Text.ElideLeft : Text.ElideNone
                    Binding {
                        id: <textbindingid>
                        target: <textfieldid>
                        property: "objectName"
                    }
                }

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
            }
            Item {
                height: parent.height
                width: 10
            }
        }
    }
}