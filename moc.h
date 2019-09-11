

#pragma once

#ifndef GO_MOC_33383e_H
#define GO_MOC_33383e_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge33383e;
void QmlBridge33383e_QmlBridge33383e_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge33383e_ConnectLoad(void* ptr);
void QmlBridge33383e_DisconnectLoad(void* ptr);
void QmlBridge33383e_Load(void* ptr, unsigned long long blockNumber, struct Moc_PackedString networkId, unsigned long long peers, unsigned long long gasPrice, struct Moc_PackedString sync, struct Moc_PackedString hashRate, int txInCurrentBlockNo, int pendingNodeTxNo);
int QmlBridge33383e_QmlBridge33383e_QRegisterMetaType();
int QmlBridge33383e_QmlBridge33383e_QRegisterMetaType2(char* typeName);
int QmlBridge33383e_QmlBridge33383e_QmlRegisterType();
int QmlBridge33383e_QmlBridge33383e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge33383e___children_atList(void* ptr, int i);
void QmlBridge33383e___children_setList(void* ptr, void* i);
void* QmlBridge33383e___children_newList(void* ptr);
void* QmlBridge33383e___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge33383e___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge33383e___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge33383e___findChildren_atList(void* ptr, int i);
void QmlBridge33383e___findChildren_setList(void* ptr, void* i);
void* QmlBridge33383e___findChildren_newList(void* ptr);
void* QmlBridge33383e___findChildren_atList3(void* ptr, int i);
void QmlBridge33383e___findChildren_setList3(void* ptr, void* i);
void* QmlBridge33383e___findChildren_newList3(void* ptr);
void* QmlBridge33383e___qFindChildren_atList2(void* ptr, int i);
void QmlBridge33383e___qFindChildren_setList2(void* ptr, void* i);
void* QmlBridge33383e___qFindChildren_newList2(void* ptr);
void* QmlBridge33383e_NewQmlBridge(void* parent);
void QmlBridge33383e_DestroyQmlBridge(void* ptr);
void QmlBridge33383e_DestroyQmlBridgeDefault(void* ptr);
void QmlBridge33383e_ChildEventDefault(void* ptr, void* event);
void QmlBridge33383e_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge33383e_CustomEventDefault(void* ptr, void* event);
void QmlBridge33383e_DeleteLaterDefault(void* ptr);
void QmlBridge33383e_DisconnectNotifyDefault(void* ptr, void* sign);
char QmlBridge33383e_EventDefault(void* ptr, void* e);
char QmlBridge33383e_EventFilterDefault(void* ptr, void* watched, void* event);
;
void QmlBridge33383e_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif