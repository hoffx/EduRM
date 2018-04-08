import QtQuick 2.7

Rectangle
{

    property bool commonBorder : true

    property int lBorderWidth : 1
    property int rBorderWidth : 1
    property int tBorderWidth : 1
    property int bBorderWidth : 1

    property int commonBorderWidth : 1

    z : -1

    property string borderColor : "white"

    color: borderColor

    anchors
    {
        left: parent.left
        right: parent.right
        top: parent.top
        bottom: parent.bottom

        topMargin    : commonBorder ? -commonBorderWidth : -tBorderWidth
        bottomMargin : commonBorder ? -commonBorderWidth : -bBorderWidth
        leftMargin   : commonBorder ? -commonBorderWidth : -lBorderWidth
        rightMargin  : commonBorder ? -commonBorderWidth : -rBorderWidth
    }
}