/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/temperature.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Temperature", TemperatureTypeSupport)
}

// Do not create instances of this type directly. Always use NewTemperature
// function instead.
type Temperature struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp is the time the temperature was measured
	Temperature float64 `yaml:"temperature"`// Measurement of the Temperature in Degrees Celsius.
	Variance float64 `yaml:"variance"`// 0 is interpreted as variance unknown.
}

// NewTemperature creates a new Temperature with default values.
func NewTemperature() *Temperature {
	self := Temperature{}
	self.SetDefaults()
	return &self
}

func (t *Temperature) Clone() *Temperature {
	c := &Temperature{}
	c.Header = *t.Header.Clone()
	c.Temperature = t.Temperature
	c.Variance = t.Variance
	return c
}

func (t *Temperature) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Temperature) SetDefaults() {
	t.Header.SetDefaults()
	t.Temperature = 0
	t.Variance = 0
}

// CloneTemperatureSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTemperatureSlice(dst, src []Temperature) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TemperatureTypeSupport types.MessageTypeSupport = _TemperatureTypeSupport{}

type _TemperatureTypeSupport struct{}

func (t _TemperatureTypeSupport) New() types.Message {
	return NewTemperature()
}

func (t _TemperatureTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Temperature
	return (unsafe.Pointer)(C.sensor_msgs__msg__Temperature__create())
}

func (t _TemperatureTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Temperature__destroy((*C.sensor_msgs__msg__Temperature)(pointer_to_free))
}

func (t _TemperatureTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Temperature)
	mem := (*C.sensor_msgs__msg__Temperature)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.temperature = C.double(m.Temperature)
	mem.variance = C.double(m.Variance)
}

func (t _TemperatureTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Temperature)
	mem := (*C.sensor_msgs__msg__Temperature)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Temperature = float64(mem.temperature)
	m.Variance = float64(mem.variance)
}

func (t _TemperatureTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Temperature())
}

type CTemperature = C.sensor_msgs__msg__Temperature
type CTemperature__Sequence = C.sensor_msgs__msg__Temperature__Sequence

func Temperature__Sequence_to_Go(goSlice *[]Temperature, cSlice CTemperature__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Temperature, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Temperature__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Temperature * uintptr(i)),
		))
		TemperatureTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Temperature__Sequence_to_C(cSlice *CTemperature__Sequence, goSlice []Temperature) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Temperature)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Temperature * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Temperature)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Temperature * uintptr(i)),
		))
		TemperatureTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Temperature__Array_to_Go(goSlice []Temperature, cSlice []CTemperature) {
	for i := 0; i < len(cSlice); i++ {
		TemperatureTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Temperature__Array_to_C(cSlice []CTemperature, goSlice []Temperature) {
	for i := 0; i < len(goSlice); i++ {
		TemperatureTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
