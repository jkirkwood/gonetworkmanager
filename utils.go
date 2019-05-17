package gonetworkmanager

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/godbus/dbus"
)

const (
	dbusMethodAddMatch = "org.freedesktop.DBus.AddMatch"
)

type dbusBase struct {
	conn *dbus.Conn
	obj  dbus.BusObject
}

func (d *dbusBase) init(iface string, objectPath dbus.ObjectPath) error {
	var err error

	d.conn, err = dbus.SystemBus()
	if err != nil {
		return err
	}

	d.obj = d.conn.Object(iface, objectPath)

	return nil
}

func (d *dbusBase) callAndPanic(value interface{}, method string, args ...interface{}) {
	err := d.callWithReturn(value, method, args...)
	if err != nil {
		panic(err)
	}
}

func (d *dbusBase) call(method string, args ...interface{}) error {
	return d.obj.Call(method, 0, args...).Err
}

func (d *dbusBase) callWithReturn(ret interface{}, method string, args ...interface{}) error {
	return d.obj.Call(method, 0, args...).Store(ret)
}

func (d *dbusBase) callWithReturn2(ret1 interface{}, ret2 interface{}, method string, args ...interface{}) error {
	return d.obj.Call(method, 0, args...).Store(ret1, ret2)
}

func (d *dbusBase) subscribe(iface, member string) {
	rule := fmt.Sprintf("type='signal',interface='%s',path='%s',member='%s'",
		iface, d.obj.Path(), NetworkManagerInterface)
	d.conn.BusObject().Call(dbusMethodAddMatch, 0, rule)
}

func (d *dbusBase) subscribeNamespace(namespace string) {
	rule := fmt.Sprintf("type='signal',path_namespace='%s'", namespace)
	d.conn.BusObject().Call(dbusMethodAddMatch, 0, rule)
}

func (d *dbusBase) getProperty(iface string) interface{} {
	variant, err := d.obj.GetProperty(iface)
	if err != nil {
		panic(err)
	}
	return variant.Value()
}

func (d *dbusBase) getObjectProperty(iface string) dbus.ObjectPath {
	value, ok := d.getProperty(iface).(dbus.ObjectPath)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceObjectProperty(iface string) []dbus.ObjectPath {
	value, ok := d.getProperty(iface).([]dbus.ObjectPath)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getBoolProperty(iface string) bool {
	value, ok := d.getProperty(iface).(bool)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getStringProperty(iface string) string {
	value, ok := d.getProperty(iface).(string)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceStringProperty(iface string) []string {
	value, ok := d.getProperty(iface).([]string)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceSliceByteProperty(iface string) [][]byte {
	value, ok := d.getProperty(iface).([][]byte)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getMapStringVariantProperty(iface string) map[string]dbus.Variant {
	value, ok := d.getProperty(iface).(map[string]dbus.Variant)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getUint8Property(iface string) uint8 {
	value, ok := d.getProperty(iface).(uint8)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getUint32Property(iface string) uint32 {
	value, ok := d.getProperty(iface).(uint32)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceUint32Property(iface string) []uint32 {
	value, ok := d.getProperty(iface).([]uint32)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceSliceUint32Property(iface string) [][]uint32 {
	value, ok := d.getProperty(iface).([][]uint32)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceMapStringVariantProperty(iface string) []map[string]dbus.Variant {
	value, ok := d.getProperty(iface).([]map[string]dbus.Variant)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func (d *dbusBase) getSliceByteProperty(iface string) []byte {
	value, ok := d.getProperty(iface).([]byte)
	if !ok {
		panic(makeErrVariantType(iface))
	}
	return value
}

func makeErrVariantType(iface string) error {
	return fmt.Errorf("unexpected variant type for '%s'", iface)
}

func ip4ToString(ip uint32) string {
	bs := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(bs, ip)
	return net.IP(bs).String()
}
