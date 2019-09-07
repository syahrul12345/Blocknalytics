

#pragma once

#ifndef GO_MOC_092563_H
#define GO_MOC_092563_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge092563;
void QmlBridge092563_QmlBridge092563_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge092563_ConnectLoad(void* ptr);
void QmlBridge092563_DisconnectLoad(void* ptr);
void QmlBridge092563_Load(void* ptr, unsigned long long blockNumber, struct Moc_PackedString networkId, unsigned long long peers, unsigned long long gasPrice, struct Moc_PackedString sync, unsigned long long hashrate);
int QmlBridge092563_QmlBridge092563_QRegisterMetaType();
int QmlBridge092563_QmlBridge092563_QRegisterMetaType2(char* typeName);
int QmlBridge092563_QmlBridge092563_QmlRegisterType();
int QmlBridge092563_QmlBridge092563_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge092563___children_atList(void* ptr, int i);
void QmlBridge092563___children_setList(void* ptr, void* i);
void* QmlBridge092563___children_newList(void* ptr);
void* QmlBridge092563___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge092563___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge092563___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge092563___findChildren_atList(void* ptr, int i);
void QmlBridge092563___findChildren_setList(void* ptr, void* i);
void* QmlBridge092563___findChildren_newList(void* ptr);
void* QmlBridge092563___findChildren_atList3(void* ptr, int i);
void QmlBridge092563___findChildren_setList3(void* ptr, void* i);
void* QmlBridge092563___findChildren_newList3(void* ptr);
void* QmlBridge092563___qFindChildren_atList2(void* ptr, int i);
void QmlBridge092563___qFindChildren_setList2(void* ptr, void* i);
void* QmlBridge092563___qFindChildren_newList2(void* ptr);
void* QmlBridge092563_NewQmlBridge(void* parent);
void QmlBridge092563_DestroyQmlBridge(void* ptr);
void QmlBridge092563_DestroyQmlBridgeDefault(void* ptr);
void QmlBridge092563_ChildEventDefault(void* ptr, void* event);
void QmlBridge092563_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge092563_CustomEventDefault(void* ptr, void* event);
void QmlBridge092563_DeleteLaterDefault(void* ptr);
void QmlBridge092563_DisconnectNotifyDefault(void* ptr, void* sign);
char QmlBridge092563_EventDefault(void* ptr, void* e);
char QmlBridge092563_EventFilterDefault(void* ptr, void* watched, void* event);
;
void QmlBridge092563_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif