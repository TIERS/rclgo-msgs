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

#include <px4_msgs/msg/debug_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/DebugArray", DebugArrayTypeSupport)
}
const (
	DebugArray_ARRAY_SIZE uint8 = 58
)

// Do not create instances of this type directly. Always use NewDebugArray
// function instead.
type DebugArray struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Id uint16 `yaml:"id"`// unique ID of debug array, used to discriminate between arrays
	Name [10]byte `yaml:"name"`// name of the debug array (max. 10 characters)
	Data [58]float32 `yaml:"data"`// data
}

// NewDebugArray creates a new DebugArray with default values.
func NewDebugArray() *DebugArray {
	self := DebugArray{}
	self.SetDefaults()
	return &self
}

func (t *DebugArray) Clone() *DebugArray {
	c := &DebugArray{}
	c.Timestamp = t.Timestamp
	c.Id = t.Id
	c.Name = t.Name
	c.Data = t.Data
	return c
}

func (t *DebugArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DebugArray) SetDefaults() {
	t.Timestamp = 0
	t.Id = 0
	t.Name = [10]byte{}
	t.Data = [58]float32{}
}

// CloneDebugArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDebugArraySlice(dst, src []DebugArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DebugArrayTypeSupport types.MessageTypeSupport = _DebugArrayTypeSupport{}

type _DebugArrayTypeSupport struct{}

func (t _DebugArrayTypeSupport) New() types.Message {
	return NewDebugArray()
}

func (t _DebugArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__DebugArray
	return (unsafe.Pointer)(C.px4_msgs__msg__DebugArray__create())
}

func (t _DebugArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__DebugArray__destroy((*C.px4_msgs__msg__DebugArray)(pointer_to_free))
}

func (t _DebugArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DebugArray)
	mem := (*C.px4_msgs__msg__DebugArray)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.id = C.uint16_t(m.Id)
	cSlice_name := mem.name[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_name)), m.Name[:])
	cSlice_data := mem.data[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_data)), m.Data[:])
}

func (t _DebugArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DebugArray)
	mem := (*C.px4_msgs__msg__DebugArray)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Id = uint16(mem.id)
	cSlice_name := mem.name[:]
	primitives.Char__Array_to_Go(m.Name[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_name)))
	cSlice_data := mem.data[:]
	primitives.Float32__Array_to_Go(m.Data[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_data)))
}

func (t _DebugArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__DebugArray())
}

type CDebugArray = C.px4_msgs__msg__DebugArray
type CDebugArray__Sequence = C.px4_msgs__msg__DebugArray__Sequence

func DebugArray__Sequence_to_Go(goSlice *[]DebugArray, cSlice CDebugArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DebugArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__DebugArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugArray * uintptr(i)),
		))
		DebugArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func DebugArray__Sequence_to_C(cSlice *CDebugArray__Sequence, goSlice []DebugArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__DebugArray)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__DebugArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__DebugArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__DebugArray * uintptr(i)),
		))
		DebugArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func DebugArray__Array_to_Go(goSlice []DebugArray, cSlice []CDebugArray) {
	for i := 0; i < len(cSlice); i++ {
		DebugArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DebugArray__Array_to_C(cSlice []CDebugArray, goSlice []DebugArray) {
	for i := 0; i < len(goSlice); i++ {
		DebugArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
