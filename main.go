package main

import (
	"os"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/core"
	"github.com/syahrul12345/Blocknalytics/packages"
	
)

var (
	// qmlObjects = make(map[string]*core.QObject)
	qmlBridge *QmlBridge
)

type QmlBridge struct {
	core.QObject
	_ func(blockNumber uint64,networkId string,peers uint64, gasPrice uint64,sync string,hashRate string,txInCurrentBlockNo int,pendingNodeTxNo int) `signal:"load"`
}


func main() {
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML
	gui.NewQGuiApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	// create the qml application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	// load the embedded qml file
	// created by either qtrcc or qtdeploy
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	// you can also load a local file like this instead:
	//engine.Load(core.QUrl_FromLocalFile("./qml/main.qml"))

	//creates the bridge to connect to our go application
	var qmlBridge = NewQmlBridge(nil)
	engine.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	
	go func() {
		for{
			blockNumber,networkId,peers,gasPrice,syncStatus,hashRate,txInCurrentBlock,pendingNodeTx := load.Start()
			qmlBridge.Load(blockNumber,networkId,peers,gasPrice,syncStatus,hashRate,len(txInCurrentBlock),len(pendingNodeTx))
		}
		
	}()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gui.QGuiApplication_Exec()
}

