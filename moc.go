package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *QmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQmlBridge092563_Constructor
func callbackQmlBridge092563_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQmlBridge092563_Load
func callbackQmlBridge092563_Load(ptr unsafe.Pointer, blockNumber C.ulonglong, networkId C.struct_Moc_PackedString, peers C.ulonglong, gasPrice C.ulonglong, sync C.struct_Moc_PackedString, hashRate C.struct_Moc_PackedString, txInCurrentBlockNo C.int, pendingNodeTxNo C.int) {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		(*(*func(uint64, string, uint64, uint64, string, string, int, int))(signal))(uint64(blockNumber), cGoUnpackString(networkId), uint64(peers), uint64(gasPrice), cGoUnpackString(sync), cGoUnpackString(hashRate), int(int32(txInCurrentBlockNo)), int(int32(pendingNodeTxNo)))
	}

}

func (ptr *QmlBridge) ConnectLoad(f func(blockNumber uint64, networkId string, peers uint64, gasPrice uint64, sync string, hashRate string, txInCurrentBlockNo int, pendingNodeTxNo int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "load") {
			C.QmlBridge092563_ConnectLoad(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			f := func(blockNumber uint64, networkId string, peers uint64, gasPrice uint64, sync string, hashRate string, txInCurrentBlockNo int, pendingNodeTxNo int) {
				(*(*func(uint64, string, uint64, uint64, string, string, int, int))(signal))(blockNumber, networkId, peers, gasPrice, sync, hashRate, txInCurrentBlockNo, pendingNodeTxNo)
				f(blockNumber, networkId, peers, gasPrice, sync, hashRate, txInCurrentBlockNo, pendingNodeTxNo)
			}
			qt.ConnectSignal(ptr.Pointer(), "load", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectLoad() {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_DisconnectLoad(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QmlBridge) Load(blockNumber uint64, networkId string, peers uint64, gasPrice uint64, sync string, hashRate string, txInCurrentBlockNo int, pendingNodeTxNo int) {
	if ptr.Pointer() != nil {
		var networkIdC *C.char
		if networkId != "" {
			networkIdC = C.CString(networkId)
			defer C.free(unsafe.Pointer(networkIdC))
		}
		var syncC *C.char
		if sync != "" {
			syncC = C.CString(sync)
			defer C.free(unsafe.Pointer(syncC))
		}
		var hashRateC *C.char
		if hashRate != "" {
			hashRateC = C.CString(hashRate)
			defer C.free(unsafe.Pointer(hashRateC))
		}
		C.QmlBridge092563_Load(ptr.Pointer(), C.ulonglong(blockNumber), C.struct_Moc_PackedString{data: networkIdC, len: C.longlong(len(networkId))}, C.ulonglong(peers), C.ulonglong(gasPrice), C.struct_Moc_PackedString{data: syncC, len: C.longlong(len(sync))}, C.struct_Moc_PackedString{data: hashRateC, len: C.longlong(len(hashRate))}, C.int(int32(txInCurrentBlockNo)), C.int(int32(pendingNodeTxNo)))
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge092563_QmlBridge092563_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge092563_QmlBridge092563_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge092563_QmlBridge092563_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge092563_QmlBridge092563_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge092563_QmlBridge092563_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge092563_QmlBridge092563_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge092563_QmlBridge092563_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge092563_QmlBridge092563_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge092563___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return C.QmlBridge092563___children_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QmlBridge092563___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QmlBridge092563___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge092563___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.QmlBridge092563___findChildren_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge092563___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QmlBridge092563___findChildren_newList3(ptr.Pointer())
}

func (ptr *QmlBridge) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge092563___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __qFindChildren_newList2() unsafe.Pointer {
	return C.QmlBridge092563___qFindChildren_newList2(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	tmpValue := NewQmlBridgeFromPointer(C.QmlBridge092563_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge092563_DestroyQmlBridge
func callbackQmlBridge092563_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge092563_ChildEvent
func callbackQmlBridge092563_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge092563_ConnectNotify
func callbackQmlBridge092563_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge092563_CustomEvent
func callbackQmlBridge092563_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge092563_DeleteLater
func callbackQmlBridge092563_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge092563_Destroyed
func callbackQmlBridge092563_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQmlBridge092563_DisconnectNotify
func callbackQmlBridge092563_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge092563_Event
func callbackQmlBridge092563_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge092563_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQmlBridge092563_EventFilter
func callbackQmlBridge092563_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge092563_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQmlBridge092563_ObjectNameChanged
func callbackQmlBridge092563_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge092563_TimerEvent
func callbackQmlBridge092563_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge092563_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
