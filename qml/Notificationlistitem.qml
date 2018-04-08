import QtQuick 2.7
import QtQuick.Controls 2.1

Item {
    property string content: ""
    property string ic: ""
    property string type: ""
    height: notificationText.paintedHeight + 30
    width: parent.width - 20

    MouseArea {
        anchors.fill: parent
        hoverEnabled: true

        onExited: {
            notificationTimer.restart()
        }

        Timer {
            id: notificationTimer
            interval: 5000
            running: true
            repeat: false
            onTriggered: {
                if (!parent.containsMouse) {
                    parent.parent.destroy()
                }
            }
        }

        Rectangle {
            width: parent.width
            height: parent.height
            color: "#3f51b5"
            CustomBorder
            {
                commonBorder: false
                lBorderWidth: 0
                rBorderWidth: 10
                tBorderWidth: 0
                bBorderWidth: 0
                borderColor: type === "0" ? "#21D551" : type === "1" ? "#E8E80C" : "#E91E63"
            }

            ToolButton {
                anchors.right: parent.right
                anchors.top: parent.top
                width: 40
                height: 40
                onClicked: parent.parent.parent.destroy()

                Image{
                    anchors.fill: parent
                    scale: 0.5
                    source: 'img/closewhite.png'
                }
            } 

            Text {
                id: notificationText
                anchors.fill: parent
                anchors.rightMargin: 30
                padding: 15
                color: "#ffffff"
                text: ic === "-1" ? content : content + " at instruction " + ic
                wrapMode: Text.Wrap
            }
        }
    }
}
