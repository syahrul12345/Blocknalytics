package main

import (
	"os"

	"fmt"

	"github.com/syahrul12345/Blocknalytics/desktop/packages/load"
	"github.com/syahrul12345/Blocknalytics/desktop/packages/plot"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func init() { CustomTableModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTableModel") }

const (
	FirstName = int(core.Qt__UserRole) + 1<<iota
	LastName
)

var (
	qmlBridge *QmlBridge
)

type TableItem struct {
	firstName string
	lastName  string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_         func()                                  `signal:"remove,auto"`
	_         func(item []*core.QVariant)             `signal:"add,auto"`
	_         func(firstName string, lastName string) `signal:"edit,auto"`
	modelData []TableItem
}

type QmlBridge struct {
	core.QObject
	_ func(blockNumber uint64, networkId string, peers uint64, gasPrice uint64, sync string, hashRate string, txInCurrentBlockNo int, pendingNodeTxNo int) `signal:"load"`
	_ func(transaction99Map map[uint64]int, gasUsed99Map map[uint64]uint64)                                                                                `signal:"plot"`
}

func (m *CustomTableModel) init() {
	m.modelData = []TableItem{{"john", "doe"}, {"john", "bob"}}
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		FirstName: core.NewQByteArray2("FirstName", -1),
		LastName:  core.NewQByteArray2("LastName", -1),
	}
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case FirstName:
		return core.NewQVariant1(item.firstName)
	case LastName:
		return core.NewQVariant1(item.lastName)
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item []*core.QVariant) {
	fmt.Println("Added function activated")
	fmt.Println("adding")
	fmt.Println(item)
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, TableItem{item[0].ToString(), item[1].ToString()})
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = TableItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()), []int{FirstName, LastName})
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	quickcontrols2.QQuickStyle_SetStyle("Material")
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	var qmlBridge = NewQmlBridge(nil)
	engine.RootContext().SetContextProperty("qmlBridge", qmlBridge)
	fmt.Println("context set succesfully")

	go func() {
		for {
			fmt.Println("getting data from ethereum node...")
			blockNumber, networkId, peers, gasPrice, syncStatus, hashRate, txInCurrentBlock, pendingNodeTx := load.Start()
			fmt.Println(load.Start())
			qmlBridge.Load(blockNumber, networkId, peers, gasPrice, syncStatus, hashRate, len(txInCurrentBlock), len(pendingNodeTx))
		}

	}()

	go func() {
		for {
			fmt.Println("Getting transaction & gasused history data...")
			transaction99Map, gasUsed99Map, err := plot.Plot()
			fmt.Println(plot.Plot())
			if err != nil {
				fmt.Println(err)
			}
			qmlBridge.Plot(transaction99Map, gasUsed99Map)
		}
	}()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gui.QGuiApplication_Exec()
}
