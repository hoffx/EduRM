import QtQuick 2.7
import QtQuick.Controls 2.1
TextArea {

    property string hidden: "true"

    visible: hidden === "false" ? true : false
    font.pointSize: 13
    width: parent.parent.width
    height: parent.parent.height
    selectByMouse: true
    selectByKeyboard: true
    wrapMode: TextArea.Wrap
    padding: 15
    background: null
    text: "1: "
    font.family: "Menlo, Monaco, 'Courier New', monospace"

    Keys.onPressed: {
        if (event.key == Qt.Key_Return) {
            var nextIC = ""
            this.text = this.text.replace(/.+/g, function(match, offset){
                if (cursorPosition === offset + match.length) {
                    match.replace(/([0-9]+)[:\s]+.*/g, function(match, p1) {
                        nextIC = (parseInt(p1)+1).toString() + ": "
                        event.accepted = true
                        return match
                    })
                    return match
                }
                return match
            })
            if (nextIC != "") {
                this.insert(cursorPosition, "\n"+nextIC)
            }
        }
    }

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