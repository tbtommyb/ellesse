import QtQuick 2.3
import QtQuick.Window 2.2
import QtQuick.Controls 1.4

// main window
Window {	
    visible: true
    width: 400; height: 600

    Column {
        id: "column"
        spacing: 5
        width: 400; height: 600

        Button {
            id: "back"
            width: 30; height: 30
            text: ".."
            onClicked: _Ellesse.change("..")
        }

        ListView {
            width: 400; height: column.height - back.height
            model: _Ellesse.fileList
            delegate: Rectangle {
                height: 25
                width: column.width - 20
                MouseArea {
                    anchors.fill: parent
                    onClicked: _Ellesse.change(modelData)
                }
                Text { text: modelData }
            }
        }
    }
}
