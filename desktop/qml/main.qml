import QtQuick 2.10          //ApplicationWindow
import QtQuick.Controls 2.1 //Dialog
import QtQuick.Controls 1.4	 //TableView
import CustomQmlTypes 1.0 //CustomTableModel
import QtQuick.Layouts 1.3 //ColumnLayout

ApplicationWindow {
    id: window
    width: 1280
    height: 900

    visible: true
    title: "Blocknalytics"
    minimumWidth: 1280
    minimumHeight: 900
    color:"black"

    Dialog {
        title: qsTr("Node Settings")
        id: nodeDialog
        Label {
            text: "Change your node endpoints"
        }
    }

    Row {
        id:headerRow
        x: 1166
        y: 14
        width: 0
        height: 3

        Button {
            id: button
            text: qsTr("Add to table")
            onClicked: {
                tableview.model.add(["john", "doe"])
            }
        }
    }

    Row {
        id:mainRow
        x: 14
        y: 64
        width: 995
        height: 73
        spacing:40
        anchors.horizontalCenter: parent.horizontalCenter
        Column {
            id: blockHeightCol
            Label {
                text: "Block Height"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: blockHeight
                text: ""
                font.pixelSize: 40
                color:"white"
            }

        }

        Column {
            id: chainIDCol
            Label {
                text: "networkID"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id:networkID
                text:""
                font.pixelSize: 40
                color:"white"
            }
        }

        Column {
            id: synCol
            Label {
                text: "Syncing"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: syncStatus
                text: ""
                font.pixelSize: 40
                color:"white"
            }
        }


        Column {
            id: gasPriceCol
            Label {
                text: "Gas Price"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: gas
                text: ""
                font.pixelSize: 40
                color:"white"
            }
        }

        Column {
            id: hashRateCol
            Label {
                text: "Network Hashrate"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: hashRateText
                text: ""
                font.pixelSize: 40
                color:"white"
            }
        }

        Column {
            id: pendingCol
            Label {
                text: "Pending Transactions"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: pendingTransactions
                text: ""
                font.pixelSize: 40
                color:"white"
            }
        }

        Column {
            id: peersCol
            Label {
                text: "Peers"
                font.pixelSize: 18
                color:"white"
            }
            Text {
                id: peerCount
                text: ""
                font.pixelSize: 40
                color:"white"
            }

        }
    }
    Row {
        ColumnLayout {
            anchors.fill: parent

		TableView {
			id: tableview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: CustomTableModel{}

			TableViewColumn {
				role: "FirstName"
				title: role
			}

			TableViewColumn {
				role: "LastName"
				title: role
			}
		}
        }
    }

    Dialog {
        id: someDialog

        x: (window.width - width) * 0.5
        y: (window.height - height) * 0.5

        contentWidth: window.width * 0.5
        contentHeight: window.height * 0.25
        standardButtons: Dialog.Ok
        function show() {
            someDialog.open()
        }

    }

}
