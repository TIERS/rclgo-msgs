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
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/debug_key_value.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/DebugKeyValue", DebugKeyValueTypeSupport)
}

// Do not create instances of this type directly. Always use NewDebugKeyValue
// function instead.
type DebugKeyValue struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Key [10]byte `yaml:"key"`// max. 10 characters as key / name
	Value float32 `yaml:"value"`// the value to send as debug output
}

// NewDebugKeyValue creates a new DebugKeyValue with default values.
func NewDebugKeyValue() *DebugKeyValue {
	self := DebugKeyValue{}
	self.SetDefaults()
	return &self
}

func (t *DebugKeyValue) Clone() *DebugKeyValue {
	c := &DebugKeyValue{}
	c.Timestamp = t.Timestamp
	c.Key = t.Key
	c.Value = t.Value
	return c
}

func (t *DebugKeyValue) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DebugKeyValue) SetDefaults() {
	t.Timestamp = 0
	t.Key = [10]byte{}
	t.Value = 0
}

// CloneDebugKeyValueSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDebugKeyValueSlice(dst, src []DebugKeyValue) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DebugKeyValueTypeSupport types.MessageTypeSupport = _DebugKeyValueTypeSupport{}

type _DebugKeyValueTypeSupport struct{}

func (t _DebugKeyValueTypeSupport) New() types.Message {
	return NewDebugKeyValue()
}

func (t _DebugKeyValueTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__DebugKeyValue
	return (unsafe.Pointer)(C.px4_msgs__msg__DebugKeyValue__create())
}

func (t _DebugKeyValueTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__DebugKeyValue__destroy((*C.px4_msgs__msg__DebugKeyValue)(pointer_to_free))
}

func (t _DebugKeyValueTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DebugKeyValue)
	mem := (*C.px4_msgs__msg__DebugKeyValue)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_key := mem.key[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_key)), m.Key[:])
	mem.value = C.float(m.Value)
}

func (t _DebugKeyValueTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DebugKeyValue)
	mem := (*C.px4_msgs__msg__DebugKeyValue)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_key := mem.key[:]
	primitives.Char__Array_to_Go(m.Key[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_key)))
	m.Value = float32(mem.value)
}

func (t _DebugKeyValueTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__DebugKeyValue())
}

type CDebugKeyValue = C.px4_msgs__msg__DebugKeyValue
type CDebugKeyValue__Sequence = C.px4_msgs__msg__DebugKeyValue__Sequence

func DebugKeyValue__Sequence_to_Go(goSlice *[]DebugKeyValue, cSlice CDebugKeyValue__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DebugKeyValue, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__DebugKeyValue__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugKeyValue * uintptr(i)),
		))
		DebugKeyValueTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func DebugKeyValue__Sequence_to_C(cSlice *CDebugKeyValue__Sequence, goSlice []DebugKeyValue) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__DebugKeyValue)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__DebugKeyValue * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__DebugKeyValue)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugKeyValue * uintptr(i)),
		))
		DebugKeyValueTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func DebugKeyValue__Array_to_Go(goSlice []DebugKeyValue, cSlice []CDebugKeyValue) {
	for i := 0; i < len(cSlice); i++ {
		DebugKeyValueTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DebugKeyValue__Array_to_C(cSlice []CDebugKeyValue, goSlice []DebugKeyValue) {
	for i := 0; i < len(goSlice); i++ {
		DebugKeyValueTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
