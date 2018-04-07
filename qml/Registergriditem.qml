import QtQuick 2.7

Column{
    property string number: ""
    property string value: ""

    width: parent.width / parent.columns
    height: width

    Text {
        color: "#3f51b5"
        width: parent.width
        height: 3 * parent.height / 5
        text: "R"+number
        verticalAlignment: Text.AlignBottom
        horizontalAlignment: Text.AlignHCenter
        padding: 5
    }
    Text {
        width: parent.width
        height: 2 * parent.height / 5
        text: value
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        font.pointSize: 18
    }

}