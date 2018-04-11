import QtQuick 2.7
import QtQuick.Controls 2.1

ToolButton{
    height: parent.height
    property string buttontext: ""
    property string idtext: ""
    
    ToolTip.visible: hovered
    ToolTip.delay: 1000
    ToolTip.text: "Click here to delete this breakpoint"

    contentItem: Text {
        text: buttontext
        height: parent.height
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        color: parent.hovered ? "#e91e63" : "#ffffff"
    }
    onClicked: hermes.sendToGo("event_removebreakpoint", idtext, "")
}