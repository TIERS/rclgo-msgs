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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/generator_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/GeneratorStatus", GeneratorStatusTypeSupport)
}
const (
	GeneratorStatus_STATUS_FLAG_OFF uint64 = 1// Generator is off.
	GeneratorStatus_STATUS_FLAG_READY uint64 = 2// Generator is ready to start generating power.
	GeneratorStatus_STATUS_FLAG_GENERATING uint64 = 4// Generator is generating power.
	GeneratorStatus_STATUS_FLAG_CHARGING uint64 = 8// Generator is charging the batteries (generating enough power to charge and provide the load).
	GeneratorStatus_STATUS_FLAG_REDUCED_POWER uint64 = 16// Generator is operating at a reduced maximum power.
	GeneratorStatus_STATUS_FLAG_MAXPOWER uint64 = 32// Generator is providing the maximum output.
	GeneratorStatus_STATUS_FLAG_OVERTEMP_WARNING uint64 = 64// Generator is near the maximum operating temperature, cooling is insufficient.
	GeneratorStatus_STATUS_FLAG_OVERTEMP_FAULT uint64 = 128// Generator hit the maximum operating temperature and shutdown.
	GeneratorStatus_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING uint64 = 256// Power electronics are near the maximum operating temperature, cooling is insufficient.
	GeneratorStatus_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT uint64 = 512// Power electronics hit the maximum operating temperature and shutdown.
	GeneratorStatus_STATUS_FLAG_ELECTRONICS_FAULT uint64 = 1024// Power electronics experienced a fault and shutdown.
	GeneratorStatus_STATUS_FLAG_POWERSOURCE_FAULT uint64 = 2048// The power source supplying the generator failed e.g. mechanical generator stopped, tether is no longer providing power, solar cell is in shade, hydrogen reaction no longer happening.
	GeneratorStatus_STATUS_FLAG_COMMUNICATION_WARNING uint64 = 4096// Generator controller having communication problems.
	GeneratorStatus_STATUS_FLAG_COOLING_WARNING uint64 = 8192// Power electronic or generator cooling system error.
	GeneratorStatus_STATUS_FLAG_POWER_RAIL_FAULT uint64 = 16384// Generator controller power rail experienced a fault.
	GeneratorStatus_STATUS_FLAG_OVERCURRENT_FAULT uint64 = 32768// Generator controller exceeded the overcurrent threshold and shutdown to prevent damage.
	GeneratorStatus_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT uint64 = 65536// Generator controller detected a high current going into the batteries and shutdown to prevent battery damage. |
	GeneratorStatus_STATUS_FLAG_OVERVOLTAGE_FAULT uint64 = 131072// Generator controller exceeded it's overvoltage threshold and shutdown to prevent it exceeding the voltage rating.
	GeneratorStatus_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT uint64 = 262144// Batteries are under voltage (generator will not start).
	GeneratorStatus_STATUS_FLAG_START_INHIBITED uint64 = 524288// Generator start is inhibited by e.g. a safety switch.
	GeneratorStatus_STATUS_FLAG_MAINTENANCE_REQUIRED uint64 = 1048576// Generator requires maintenance.
	GeneratorStatus_STATUS_FLAG_WARMING_UP uint64 = 2097152// Generator is not ready to generate yet.
	GeneratorStatus_STATUS_FLAG_IDLE uint64 = 4194304// Generator is idle.
)

// Do not create instances of this type directly. Always use NewGeneratorStatus
// function instead.
type GeneratorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Status uint64 `yaml:"status"`// Status flags
	BatteryCurrent float32 `yaml:"battery_current"`// [A] Current into/out of battery. Positive for out. Negative for in. NaN: field not provided.
	LoadCurrent float32 `yaml:"load_current"`// [A] Current going to the UAV. If battery current not available this is the DC current from the generator. Positive for out. Negative for in. NaN: field not provided
	PowerGenerated float32 `yaml:"power_generated"`// [W] The power being generated. NaN: field not provided
	BusVoltage float32 `yaml:"bus_voltage"`// [V] Voltage of the bus seen at the generator, or battery bus if battery bus is controlled by generator and at a different voltage to main bus.
	BatCurrentSetpoint float32 `yaml:"bat_current_setpoint"`// [A] The target battery current. Positive for out. Negative for in. NaN: field not provided
	Runtime uint32 `yaml:"runtime"`// [s] Seconds this generator has run since it was rebooted. UINT32_MAX: field not provided.
	TimeUntilMaintenance int32 `yaml:"time_until_maintenance"`// [s] Seconds until this generator requires maintenance.  A negative value indicates maintenance is past-due. INT32_MAX: field not provided.
	GeneratorSpeed uint16 `yaml:"generator_speed"`// [rpm] Speed of electrical generator or alternator. UINT16_MAX: field not provided.
	RectifierTemperature int16 `yaml:"rectifier_temperature"`// [degC] The temperature of the rectifier or power converter. INT16_MAX: field not provided.
	GeneratorTemperature int16 `yaml:"generator_temperature"`// [degC] The temperature of the mechanical motor, fuel cell core or generator. INT16_MAX: field not provided.
}

// NewGeneratorStatus creates a new GeneratorStatus with default values.
func NewGeneratorStatus() *GeneratorStatus {
	self := GeneratorStatus{}
	self.SetDefaults()
	return &self
}

func (t *GeneratorStatus) Clone() *GeneratorStatus {
	c := &GeneratorStatus{}
	c.Timestamp = t.Timestamp
	c.Status = t.Status
	c.BatteryCurrent = t.BatteryCurrent
	c.LoadCurrent = t.LoadCurrent
	c.PowerGenerated = t.PowerGenerated
	c.BusVoltage = t.BusVoltage
	c.BatCurrentSetpoint = t.BatCurrentSetpoint
	c.Runtime = t.Runtime
	c.TimeUntilMaintenance = t.TimeUntilMaintenance
	c.GeneratorSpeed = t.GeneratorSpeed
	c.RectifierTemperature = t.RectifierTemperature
	c.GeneratorTemperature = t.GeneratorTemperature
	return c
}

func (t *GeneratorStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GeneratorStatus) SetDefaults() {
	t.Timestamp = 0
	t.Status = 0
	t.BatteryCurrent = 0
	t.LoadCurrent = 0
	t.PowerGenerated = 0
	t.BusVoltage = 0
	t.BatCurrentSetpoint = 0
	t.Runtime = 0
	t.TimeUntilMaintenance = 0
	t.GeneratorSpeed = 0
	t.RectifierTemperature = 0
	t.GeneratorTemperature = 0
}

// CloneGeneratorStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGeneratorStatusSlice(dst, src []GeneratorStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GeneratorStatusTypeSupport types.MessageTypeSupport = _GeneratorStatusTypeSupport{}

type _GeneratorStatusTypeSupport struct{}

func (t _GeneratorStatusTypeSupport) New() types.Message {
	return NewGeneratorStatus()
}

func (t _GeneratorStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GeneratorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__GeneratorStatus__create())
}

func (t _GeneratorStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GeneratorStatus__destroy((*C.px4_msgs__msg__GeneratorStatus)(pointer_to_free))
}

func (t _GeneratorStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GeneratorStatus)
	mem := (*C.px4_msgs__msg__GeneratorStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.status = C.uint64_t(m.Status)
	mem.battery_current = C.float(m.BatteryCurrent)
	mem.load_current = C.float(m.LoadCurrent)
	mem.power_generated = C.float(m.PowerGenerated)
	mem.bus_voltage = C.float(m.BusVoltage)
	mem.bat_current_setpoint = C.float(m.BatCurrentSetpoint)
	mem.runtime = C.uint32_t(m.Runtime)
	mem.time_until_maintenance = C.int32_t(m.TimeUntilMaintenance)
	mem.generator_speed = C.uint16_t(m.GeneratorSpeed)
	mem.rectifier_temperature = C.int16_t(m.RectifierTemperature)
	mem.generator_temperature = C.int16_t(m.GeneratorTemperature)
}

func (t _GeneratorStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GeneratorStatus)
	mem := (*C.px4_msgs__msg__GeneratorStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Status = uint64(mem.status)
	m.BatteryCurrent = float32(mem.battery_current)
	m.LoadCurrent = float32(mem.load_current)
	m.PowerGenerated = float32(mem.power_generated)
	m.BusVoltage = float32(mem.bus_voltage)
	m.BatCurrentSetpoint = float32(mem.bat_current_setpoint)
	m.Runtime = uint32(mem.runtime)
	m.TimeUntilMaintenance = int32(mem.time_until_maintenance)
	m.GeneratorSpeed = uint16(mem.generator_speed)
	m.RectifierTemperature = int16(mem.rectifier_temperature)
	m.GeneratorTemperature = int16(mem.generator_temperature)
}

func (t _GeneratorStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GeneratorStatus())
}

type CGeneratorStatus = C.px4_msgs__msg__GeneratorStatus
type CGeneratorStatus__Sequence = C.px4_msgs__msg__GeneratorStatus__Sequence

func GeneratorStatus__Sequence_to_Go(goSlice *[]GeneratorStatus, cSlice CGeneratorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GeneratorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GeneratorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GeneratorStatus * uintptr(i)),
		))
		GeneratorStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GeneratorStatus__Sequence_to_C(cSlice *CGeneratorStatus__Sequence, goSlice []GeneratorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GeneratorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GeneratorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GeneratorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GeneratorStatus * uintptr(i)),
		))
		GeneratorStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GeneratorStatus__Array_to_Go(goSlice []GeneratorStatus, cSlice []CGeneratorStatus) {
	for i := 0; i < len(cSlice); i++ {
		GeneratorStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GeneratorStatus__Array_to_C(cSlice []CGeneratorStatus, goSlice []GeneratorStatus) {
	for i := 0; i < len(goSlice); i++ {
		GeneratorStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
