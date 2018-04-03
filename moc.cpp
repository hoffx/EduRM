

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class QmlBridge: public QObject
{
Q_OBJECT
public:
	QmlBridge(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge_QmlBridge_QRegisterMetaType();QmlBridge_QmlBridge_QRegisterMetaTypes();callbackQmlBridge_Constructor(this);};
	void Signal_SendToQml(QString target, QString action, QString data) { QByteArray t0e8a3a = target.toUtf8(); Moc_PackedString targetPacked = { const_cast<char*>(t0e8a3a.prepend("WHITESPACE").constData()+10), t0e8a3a.size()-10 };QByteArray t34eb4c = action.toUtf8(); Moc_PackedString actionPacked = { const_cast<char*>(t34eb4c.prepend("WHITESPACE").constData()+10), t34eb4c.size()-10 };QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridge_SendToQml(this, targetPacked, actionPacked, dataPacked); };
	 ~QmlBridge() { callbackQmlBridge_DestroyQmlBridge(this); };
	bool event(QEvent * e) { return callbackQmlBridge_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlBridge_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge_TimerEvent(this, event); };
	
signals:
	void sendToQml(QString target, QString action, QString data);
public slots:
	void sendToGo(QString source, QString action, QString data) { QByteArray t828d33 = source.toUtf8(); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33.prepend("WHITESPACE").constData()+10), t828d33.size()-10 };QByteArray t34eb4c = action.toUtf8(); Moc_PackedString actionPacked = { const_cast<char*>(t34eb4c.prepend("WHITESPACE").constData()+10), t34eb4c.size()-10 };QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridge_SendToGo(this, sourcePacked, actionPacked, dataPacked); };
private:
};

Q_DECLARE_METATYPE(QmlBridge*)


void QmlBridge_QmlBridge_QRegisterMetaTypes() {
}

void QmlBridge_ConnectSendToQml(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::Signal_SendToQml));
}

void QmlBridge_DisconnectSendToQml(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::Signal_SendToQml));
}

void QmlBridge_SendToQml(void* ptr, struct Moc_PackedString target, struct Moc_PackedString action, struct Moc_PackedString data)
{
	static_cast<QmlBridge*>(ptr)->sendToQml(QString::fromUtf8(target.data, target.len), QString::fromUtf8(action.data, action.len), QString::fromUtf8(data.data, data.len));
}

void QmlBridge_SendToGo(void* ptr, struct Moc_PackedString source, struct Moc_PackedString action, struct Moc_PackedString data)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "sendToGo", Q_ARG(QString, QString::fromUtf8(source.data, source.len)), Q_ARG(QString, QString::fromUtf8(action.data, action.len)), Q_ARG(QString, QString::fromUtf8(data.data, data.len)));
}

int QmlBridge_QmlBridge_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge*>();
}

int QmlBridge_QmlBridge_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge*>(const_cast<const char*>(typeName));
}

int QmlBridge_QmlBridge_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>();
#else
	return 0;
#endif
}

int QmlBridge_QmlBridge_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlBridge___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlBridge___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlBridge___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlBridge_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge(static_cast<QObject*>(parent));
	}
}

void QmlBridge_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->~QmlBridge();
}

void QmlBridge_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlBridge_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridge_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->QObject::deleteLater();
}

void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
