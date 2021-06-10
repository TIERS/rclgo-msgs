/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/battery_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/BatteryStatus", BatteryStatusTypeSupport)
}
const (
	BatteryStatus_BATTERY_SOURCE_POWER_MODULE uint8 = 0
	BatteryStatus_BATTERY_SOURCE_EXTERNAL uint8 = 1
	BatteryStatus_BATTERY_SOURCE_ESCS uint8 = 2
	BatteryStatus_BATTERY_WARNING_NONE uint8 = 0// no battery low voltage warning active
	BatteryStatus_BATTERY_WARNING_LOW uint8 = 1// warning of low voltage
	BatteryStatus_BATTERY_WARNING_CRITICAL uint8 = 2// critical voltage, return / abort immediately
	BatteryStatus_BATTERY_WARNING_EMERGENCY uint8 = 3// immediate landing required
	BatteryStatus_BATTERY_WARNING_FAILED uint8 = 4// the battery has failed completely
	BatteryStatus_MAX_INSTANCES uint8 = 4
)

// Do not create instances of this type directly. Always use NewBatteryStatus
// function instead.
type BatteryStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	VoltageV float32 `yaml:"voltage_v"`// Battery voltage in volts, 0 if unknown
	VoltageFilteredV float32 `yaml:"voltage_filtered_v"`// Battery voltage in volts, filtered, 0 if unknown
	CurrentA float32 `yaml:"current_a"`// Battery current in amperes, -1 if unknown
	CurrentFilteredA float32 `yaml:"current_filtered_a"`// Battery current in amperes, filtered, 0 if unknown
	AverageCurrentA float32 `yaml:"average_current_a"`// Battery current average in amperes, -1 if unknown
	DischargedMah float32 `yaml:"discharged_mah"`// Discharged amount in mAh, -1 if unknown
	Remaining float32 `yaml:"remaining"`// From 1 to 0, -1 if unknown
	Scale float32 `yaml:"scale"`// Power scaling factor, >= 1, or -1 if unknown
	Temperature float32 `yaml:"temperature"`// temperature of the battery. NaN if unknown
	CellCount int32 `yaml:"cell_count"`// Number of cells
	Connected bool `yaml:"connected"`// Whether or not a battery is connected, based on a voltage threshold
	Source uint8 `yaml:"source"`// Battery source
	Priority uint8 `yaml:"priority"`// Zero based priority is the connection on the Power Controller V1..Vn AKA BrickN-1
	Capacity uint16 `yaml:"capacity"`// actual capacity of the battery
	CycleCount uint16 `yaml:"cycle_count"`// number of discharge cycles the battery has experienced
	RunTimeToEmpty uint16 `yaml:"run_time_to_empty"`// predicted remaining battery capacity based on the present rate of discharge in min
	AverageTimeToEmpty uint16 `yaml:"average_time_to_empty"`// predicted remaining battery capacity based on the average rate of discharge in min
	SerialNumber uint16 `yaml:"serial_number"`// serial number of the battery pack
	ManufactureDate uint16 `yaml:"manufacture_date"`// manufacture date, part of serial number of the battery pack. formated as: Day + Month×32 + (Year–1980)×512
	StateOfHealth uint16 `yaml:"state_of_health"`// state of health. FullChargeCapacity/DesignCapacity.
	MaxError uint16 `yaml:"max_error"`// max error, expected margin of error in % in the state-of-charge calculation with a range of 1 to 100%
	Id uint8 `yaml:"id"`// ID number of a battery. Should be unique and consistent for the lifetime of a vehicle. 1-indexed.
	InterfaceError uint16 `yaml:"interface_error"`// SMBUS interface error counter
	VoltageCellV [10]float32 `yaml:"voltage_cell_v"`// Battery individual cell voltages
	MaxCellVoltageDelta float32 `yaml:"max_cell_voltage_delta"`// Max difference between individual cell voltages
	IsPoweringOff bool `yaml:"is_powering_off"`// Power off event imminent indication, false if unknown
	Warning uint8 `yaml:"warning"`// current battery warning
}

// NewBatteryStatus creates a new BatteryStatus with default values.
func NewBatteryStatus() *BatteryStatus {
	self := BatteryStatus{}
	self.SetDefaults()
	return &self
}

func (t *BatteryStatus) Clone() *BatteryStatus {
	c := &BatteryStatus{}
	c.Timestamp = t.Timestamp
	c.VoltageV = t.VoltageV
	c.VoltageFilteredV = t.VoltageFilteredV
	c.CurrentA = t.CurrentA
	c.CurrentFilteredA = t.CurrentFilteredA
	c.AverageCurrentA = t.AverageCurrentA
	c.DischargedMah = t.DischargedMah
	c.Remaining = t.Remaining
	c.Scale = t.Scale
	c.Temperature = t.Temperature
	c.CellCount = t.CellCount
	c.Connected = t.Connected
	c.Source = t.Source
	c.Priority = t.Priority
	c.Capacity = t.Capacity
	c.CycleCount = t.CycleCount
	c.RunTimeToEmpty = t.RunTimeToEmpty
	c.AverageTimeToEmpty = t.AverageTimeToEmpty
	c.SerialNumber = t.SerialNumber
	c.ManufactureDate = t.ManufactureDate
	c.StateOfHealth = t.StateOfHealth
	c.MaxError = t.MaxError
	c.Id = t.Id
	c.InterfaceError = t.InterfaceError
	c.VoltageCellV = t.VoltageCellV
	c.MaxCellVoltageDelta = t.MaxCellVoltageDelta
	c.IsPoweringOff = t.IsPoweringOff
	c.Warning = t.Warning
	return c
}

func (t *BatteryStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BatteryStatus) SetDefaults() {
	t.Timestamp = 0
	t.VoltageV = 0
	t.VoltageFilteredV = 0
	t.CurrentA = 0
	t.CurrentFilteredA = 0
	t.AverageCurrentA = 0
	t.DischargedMah = 0
	t.Remaining = 0
	t.Scale = 0
	t.Temperature = 0
	t.CellCount = 0
	t.Connected = false
	t.Source = 0
	t.Priority = 0
	t.Capacity = 0
	t.CycleCount = 0
	t.RunTimeToEmpty = 0
	t.AverageTimeToEmpty = 0
	t.SerialNumber = 0
	t.ManufactureDate = 0
	t.StateOfHealth = 0
	t.MaxError = 0
	t.Id = 0
	t.InterfaceError = 0
	t.VoltageCellV = [10]float32{}
	t.MaxCellVoltageDelta = 0
	t.IsPoweringOff = false
	t.Warning = 0
}

// CloneBatteryStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBatteryStatusSlice(dst, src []BatteryStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BatteryStatusTypeSupport types.MessageTypeSupport = _BatteryStatusTypeSupport{}

type _BatteryStatusTypeSupport struct{}

func (t _BatteryStatusTypeSupport) New() types.Message {
	return NewBatteryStatus()
}

func (t _BatteryStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__BatteryStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__BatteryStatus__create())
}

func (t _BatteryStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__BatteryStatus__destroy((*C.px4_msgs__msg__BatteryStatus)(pointer_to_free))
}

func (t _BatteryStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BatteryStatus)
	mem := (*C.px4_msgs__msg__BatteryStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.voltage_v = C.float(m.VoltageV)
	mem.voltage_filtered_v = C.float(m.VoltageFilteredV)
	mem.current_a = C.float(m.CurrentA)
	mem.current_filtered_a = C.float(m.CurrentFilteredA)
	mem.average_current_a = C.float(m.AverageCurrentA)
	mem.discharged_mah = C.float(m.DischargedMah)
	mem.remaining = C.float(m.Remaining)
	mem.scale = C.float(m.Scale)
	mem.temperature = C.float(m.Temperature)
	mem.cell_count = C.int32_t(m.CellCount)
	mem.connected = C.bool(m.Connected)
	mem.source = C.uint8_t(m.Source)
	mem.priority = C.uint8_t(m.Priority)
	mem.capacity = C.uint16_t(m.Capacity)
	mem.cycle_count = C.uint16_t(m.CycleCount)
	mem.run_time_to_empty = C.uint16_t(m.RunTimeToEmpty)
	mem.average_time_to_empty = C.uint16_t(m.AverageTimeToEmpty)
	mem.serial_number = C.uint16_t(m.SerialNumber)
	mem.manufacture_date = C.uint16_t(m.ManufactureDate)
	mem.state_of_health = C.uint16_t(m.StateOfHealth)
	mem.max_error = C.uint16_t(m.MaxError)
	mem.id = C.uint8_t(m.Id)
	mem.interface_error = C.uint16_t(m.InterfaceError)
	cSlice_voltage_cell_v := mem.voltage_cell_v[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_voltage_cell_v)), m.VoltageCellV[:])
	mem.max_cell_voltage_delta = C.float(m.MaxCellVoltageDelta)
	mem.is_powering_off = C.bool(m.IsPoweringOff)
	mem.warning = C.uint8_t(m.Warning)
}

func (t _BatteryStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BatteryStatus)
	mem := (*C.px4_msgs__msg__BatteryStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.VoltageV = float32(mem.voltage_v)
	m.VoltageFilteredV = float32(mem.voltage_filtered_v)
	m.CurrentA = float32(mem.current_a)
	m.CurrentFilteredA = float32(mem.current_filtered_a)
	m.AverageCurrentA = float32(mem.average_current_a)
	m.DischargedMah = float32(mem.discharged_mah)
	m.Remaining = float32(mem.remaining)
	m.Scale = float32(mem.scale)
	m.Temperature = float32(mem.temperature)
	m.CellCount = int32(mem.cell_count)
	m.Connected = bool(mem.connected)
	m.Source = uint8(mem.source)
	m.Priority = uint8(mem.priority)
	m.Capacity = uint16(mem.capacity)
	m.CycleCount = uint16(mem.cycle_count)
	m.RunTimeToEmpty = uint16(mem.run_time_to_empty)
	m.AverageTimeToEmpty = uint16(mem.average_time_to_empty)
	m.SerialNumber = uint16(mem.serial_number)
	m.ManufactureDate = uint16(mem.manufacture_date)
	m.StateOfHealth = uint16(mem.state_of_health)
	m.MaxError = uint16(mem.max_error)
	m.Id = uint8(mem.id)
	m.InterfaceError = uint16(mem.interface_error)
	cSlice_voltage_cell_v := mem.voltage_cell_v[:]
	primitives.Float32__Array_to_Go(m.VoltageCellV[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_voltage_cell_v)))
	m.MaxCellVoltageDelta = float32(mem.max_cell_voltage_delta)
	m.IsPoweringOff = bool(mem.is_powering_off)
	m.Warning = uint8(mem.warning)
}

func (t _BatteryStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__BatteryStatus())
}

type CBatteryStatus = C.px4_msgs__msg__BatteryStatus
type CBatteryStatus__Sequence = C.px4_msgs__msg__BatteryStatus__Sequence

func BatteryStatus__Sequence_to_Go(goSlice *[]BatteryStatus, cSlice CBatteryStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BatteryStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__BatteryStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__BatteryStatus * uintptr(i)),
		))
		BatteryStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func BatteryStatus__Sequence_to_C(cSlice *CBatteryStatus__Sequence, goSlice []BatteryStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__BatteryStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__BatteryStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__BatteryStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__BatteryStatus * uintptr(i)),
		))
		BatteryStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func BatteryStatus__Array_to_Go(goSlice []BatteryStatus, cSlice []CBatteryStatus) {
	for i := 0; i < len(cSlice); i++ {
		BatteryStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BatteryStatus__Array_to_C(cSlice []CBatteryStatus, goSlice []BatteryStatus) {
	for i := 0; i < len(goSlice); i++ {
		BatteryStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
