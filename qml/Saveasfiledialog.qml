import QtQuick 2.7
import QtQuick.Dialogs 1.0

FileDialog {
    property string idToSave: ""
    visible: false
    title: qsTr("Define a filepath")
    selectMultiple: false
    selectFolder: false
    selectExisting: false
    nameFilters: [ "Assembly Files (*.spasm *.spaen)", "All files (*)" ]
    selectedNameFilter: "Assembly Files (*.spasm *.spaen)"
    sidebarVisible: true
    onAccepted: {
        hermes.sendToGo("event_savetempfile", idToSave, '{ "path": "' + fileUrl + '", "text":"'+textEdit.text.replace(/"/g, '\\"').replace(/\t/g,"\\t")+'"}')
    }
    onRejected: {}
}