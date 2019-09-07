import QtQuick 2.7          //ApplicationWindow
import QtQuick.Controls 2.1 //Dialog

ApplicationWindow {
    id: window
    width: 1280
    height: 800

    visible: true
    title: "Blocknalytics"
    minimumWidth: 1280
    minimumHeight: 800
    color:"black"
    Row {
        id:headerRow
        x: 1166
        y: 14
        width: 0
        height: 3

        Button {
            id: button
            text: qsTr("Node Settings")
        }
    }

    Row {
        id:mainRow
        x: 14
        y: 64
        width: 1266
        height: 73
        spacing:40

        Column {
            id: blockHeightCol
            Label {
                text: "Block Height"
                font.pixelSize: 20
                color:"white"
            }
            Text {
                id: blockHeight
                text: "20000000"
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
                text:"61"
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
                id: hashRate
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
            Label {
                id: pendingTransactions
                text: "2"
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
        id:graphRow
        topPadding: 10
        anchors.top: mainRow.bottom
        anchors.horizontalCenter: mainRow.horizontalCenter
        spacing:40
        Column {
            id: hash99
            Label {
                padding: 10
                text:"Hash rate for the last 99 blocks"
                font.pixelSize: 18
                color:"white"
            }
        }

        Column {
            id: count99
            Label {
                padding: 10
                text:"Transaction count last 99 blocks"
                font.pixelSize: 18
                color:"white"
            }
        }

        Column {
            id: gas99
            Label {
                padding: 10
                text:"Gas Used Last 99 blocks"
                font.pixelSize: 18
                color:"white"
            }

        }

        Column {
            id: uncle99
            Label {
                padding: 10
                text:"Uncles Last 99 blocks"
                font.pixelSize: 18
                color:"white"
            }

        }

    }

    Connections {
        target: QmlBridge
        onLoad: {
            blockHeight.text = blockNumber
            networkID.text = networkId
            peerCount.text = peers
            syncStatus.text
            gas.text = gasPrice + "GWEI"
            syncStatus.text = sync
            hashRate.text = hashrate

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
