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

#include <px4_msgs/msg/differential_pressure.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/DifferentialPressure", DifferentialPressureTypeSupport)
}

// Do not create instances of this type directly. Always use NewDifferentialPressure
// function instead.
type DifferentialPressure struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ErrorCount uint64 `yaml:"error_count"`// Number of errors detected by driver
	DifferentialPressureRawPa float32 `yaml:"differential_pressure_raw_pa"`// Raw differential pressure reading (may be negative)
	DifferentialPressureFilteredPa float32 `yaml:"differential_pressure_filtered_pa"`// Low pass filtered differential pressure reading
	Temperature float32 `yaml:"temperature"`// Temperature provided by sensor, -1000.0f if unknown
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
}

// NewDifferentialPressure creates a new DifferentialPressure with default values.
func NewDifferentialPressure() *DifferentialPressure {
	self := DifferentialPressure{}
	self.SetDefaults()
	return &self
}

func (t *DifferentialPressure) Clone() *DifferentialPressure {
	c := &DifferentialPressure{}
	c.Timestamp = t.Timestamp
	c.ErrorCount = t.ErrorCount
	c.DifferentialPressureRawPa = t.DifferentialPressureRawPa
	c.DifferentialPressureFilteredPa = t.DifferentialPressureFilteredPa
	c.Temperature = t.Temperature
	c.DeviceId = t.DeviceId
	return c
}

func (t *DifferentialPressure) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DifferentialPressure) SetDefaults() {
	t.Timestamp = 0
	t.ErrorCount = 0
	t.DifferentialPressureRawPa = 0
	t.DifferentialPressureFilteredPa = 0
	t.Temperature = 0
	t.DeviceId = 0
}

// CloneDifferentialPressureSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDifferentialPressureSlice(dst, src []DifferentialPressure) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DifferentialPressureTypeSupport types.MessageTypeSupport = _DifferentialPressureTypeSupport{}

type _DifferentialPressureTypeSupport struct{}

func (t _DifferentialPressureTypeSupport) New() types.Message {
	return NewDifferentialPressure()
}

func (t _DifferentialPressureTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__DifferentialPressure
	return (unsafe.Pointer)(C.px4_msgs__msg__DifferentialPressure__create())
}

func (t _DifferentialPressureTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__DifferentialPressure__destroy((*C.px4_msgs__msg__DifferentialPressure)(pointer_to_free))
}

func (t _DifferentialPressureTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DifferentialPressure)
	mem := (*C.px4_msgs__msg__DifferentialPressure)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.error_count = C.uint64_t(m.ErrorCount)
	mem.differential_pressure_raw_pa = C.float(m.DifferentialPressureRawPa)
	mem.differential_pressure_filtered_pa = C.float(m.DifferentialPressureFilteredPa)
	mem.temperature = C.float(m.Temperature)
	mem.device_id = C.uint32_t(m.DeviceId)
}

func (t _DifferentialPressureTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DifferentialPressure)
	mem := (*C.px4_msgs__msg__DifferentialPressure)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.ErrorCount = uint64(mem.error_count)
	m.DifferentialPressureRawPa = float32(mem.differential_pressure_raw_pa)
	m.DifferentialPressureFilteredPa = float32(mem.differential_pressure_filtered_pa)
	m.Temperature = float32(mem.temperature)
	m.DeviceId = uint32(mem.device_id)
}

func (t _DifferentialPressureTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__DifferentialPressure())
}

type CDifferentialPressure = C.px4_msgs__msg__DifferentialPressure
type CDifferentialPressure__Sequence = C.px4_msgs__msg__DifferentialPressure__Sequence

func DifferentialPressure__Sequence_to_Go(goSlice *[]DifferentialPressure, cSlice CDifferentialPressure__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DifferentialPressure, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__DifferentialPressure__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DifferentialPressure * uintptr(i)),
		))
		DifferentialPressureTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func DifferentialPressure__Sequence_to_C(cSlice *CDifferentialPressure__Sequence, goSlice []DifferentialPressure) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__DifferentialPressure)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__DifferentialPressure * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__DifferentialPressure)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DifferentialPressure * uintptr(i)),
		))
		DifferentialPressureTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func DifferentialPressure__Array_to_Go(goSlice []DifferentialPressure, cSlice []CDifferentialPressure) {
	for i := 0; i < len(cSlice); i++ {
		DifferentialPressureTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DifferentialPressure__Array_to_C(cSlice []CDifferentialPressure, goSlice []DifferentialPressure) {
	for i := 0; i < len(goSlice); i++ {
		DifferentialPressureTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
