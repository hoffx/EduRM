ToolButton{
    id: <breakpointFlickableRowButton.id>
    height: parent.height

    contentItem: Text {
        id: <breakpointFlickableRowButtonText.id>
        text: <breakpointFlickableRowButtonText.text>
        height: parent.height
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        color: parent.hovered ? "#e91e63" : "#ffffff"
    }
}