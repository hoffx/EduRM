import QtQuick 2.7
import QtQuick.Controls 2.3

Item {
    property string hide: ""
    
    height: 50
    anchors.left: parent.left
    anchors.bottom: parent.bottom
    anchors.right: parent.right

    ToolBar {
        position: ToolBar.Footer
        anchors.fill: parent
        visible: hide === "true" ? true : false
    }
}