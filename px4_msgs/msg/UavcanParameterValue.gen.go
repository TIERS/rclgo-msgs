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

#include <px4_msgs/msg/uavcan_parameter_value.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/UavcanParameterValue", UavcanParameterValueTypeSupport)
}

// Do not create instances of this type directly. Always use NewUavcanParameterValue
// function instead.
type UavcanParameterValue struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). UAVCAN-MAVLink parameter bridge response type
	NodeId uint8 `yaml:"node_id"`// UAVCAN node ID mapped from MAVLink component ID. UAVCAN-MAVLink parameter bridge response type
	ParamId [17]byte `yaml:"param_id"`// MAVLink/UAVCAN parameter name. UAVCAN-MAVLink parameter bridge response type
	ParamIndex int16 `yaml:"param_index"`// parameter index, if known. UAVCAN-MAVLink parameter bridge response type
	ParamCount uint16 `yaml:"param_count"`// number of parameters exposed by the node. UAVCAN-MAVLink parameter bridge response type
	ParamType uint8 `yaml:"param_type"`// MAVLink parameter type. UAVCAN-MAVLink parameter bridge response type
	IntValue int64 `yaml:"int_value"`// current value if param_type is int-like. UAVCAN-MAVLink parameter bridge response type
	RealValue float32 `yaml:"real_value"`// current value if param_type is float-like. UAVCAN-MAVLink parameter bridge response type
}

// NewUavcanParameterValue creates a new UavcanParameterValue with default values.
func NewUavcanParameterValue() *UavcanParameterValue {
	self := UavcanParameterValue{}
	self.SetDefaults()
	return &self
}

func (t *UavcanParameterValue) Clone() *UavcanParameterValue {
	c := &UavcanParameterValue{}
	c.Timestamp = t.Timestamp
	c.NodeId = t.NodeId
	c.ParamId = t.ParamId
	c.ParamIndex = t.ParamIndex
	c.ParamCount = t.ParamCount
	c.ParamType = t.ParamType
	c.IntValue = t.IntValue
	c.RealValue = t.RealValue
	return c
}

func (t *UavcanParameterValue) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UavcanParameterValue) SetDefaults() {
	t.Timestamp = 0
	t.NodeId = 0
	t.ParamId = [17]byte{}
	t.ParamIndex = 0
	t.ParamCount = 0
	t.ParamType = 0
	t.IntValue = 0
	t.RealValue = 0
}

// CloneUavcanParameterValueSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUavcanParameterValueSlice(dst, src []UavcanParameterValue) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UavcanParameterValueTypeSupport types.MessageTypeSupport = _UavcanParameterValueTypeSupport{}

type _UavcanParameterValueTypeSupport struct{}

func (t _UavcanParameterValueTypeSupport) New() types.Message {
	return NewUavcanParameterValue()
}

func (t _UavcanParameterValueTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__UavcanParameterValue
	return (unsafe.Pointer)(C.px4_msgs__msg__UavcanParameterValue__create())
}

func (t _UavcanParameterValueTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__UavcanParameterValue__destroy((*C.px4_msgs__msg__UavcanParameterValue)(pointer_to_free))
}

func (t _UavcanParameterValueTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UavcanParameterValue)
	mem := (*C.px4_msgs__msg__UavcanParameterValue)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.node_id = C.uint8_t(m.NodeId)
	cSlice_param_id := mem.param_id[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_param_id)), m.ParamId[:])
	mem.param_index = C.int16_t(m.ParamIndex)
	mem.param_count = C.uint16_t(m.ParamCount)
	mem.param_type = C.uint8_t(m.ParamType)
	mem.int_value = C.int64_t(m.IntValue)
	mem.real_value = C.float(m.RealValue)
}

func (t _UavcanParameterValueTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UavcanParameterValue)
	mem := (*C.px4_msgs__msg__UavcanParameterValue)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.NodeId = uint8(mem.node_id)
	cSlice_param_id := mem.param_id[:]
	primitives.Char__Array_to_Go(m.ParamId[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_param_id)))
	m.ParamIndex = int16(mem.param_index)
	m.ParamCount = uint16(mem.param_count)
	m.ParamType = uint8(mem.param_type)
	m.IntValue = int64(mem.int_value)
	m.RealValue = float32(mem.real_value)
}

func (t _UavcanParameterValueTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__UavcanParameterValue())
}

type CUavcanParameterValue = C.px4_msgs__msg__UavcanParameterValue
type CUavcanParameterValue__Sequence = C.px4_msgs__msg__UavcanParameterValue__Sequence

func UavcanParameterValue__Sequence_to_Go(goSlice *[]UavcanParameterValue, cSlice CUavcanParameterValue__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UavcanParameterValue, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__UavcanParameterValue__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UavcanParameterValue * uintptr(i)),
		))
		UavcanParameterValueTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UavcanParameterValue__Sequence_to_C(cSlice *CUavcanParameterValue__Sequence, goSlice []UavcanParameterValue) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__UavcanParameterValue)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__UavcanParameterValue * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__UavcanParameterValue)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__UavcanParameterValue * uintptr(i)),
		))
		UavcanParameterValueTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UavcanParameterValue__Array_to_Go(goSlice []UavcanParameterValue, cSlice []CUavcanParameterValue) {
	for i := 0; i < len(cSlice); i++ {
		UavcanParameterValueTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UavcanParameterValue__Array_to_C(cSlice []CUavcanParameterValue, goSlice []UavcanParameterValue) {
	for i := 0; i < len(goSlice); i++ {
		UavcanParameterValueTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
