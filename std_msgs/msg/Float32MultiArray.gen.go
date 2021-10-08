/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/float32_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Float32MultiArray", Float32MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat32MultiArray
// function instead.
type Float32MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []float32 `yaml:"data"`// array of data
}

// NewFloat32MultiArray creates a new Float32MultiArray with default values.
func NewFloat32MultiArray() *Float32MultiArray {
	self := Float32MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Float32MultiArray) Clone() *Float32MultiArray {
	c := &Float32MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]float32, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Float32MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Float32MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// CloneFloat32MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFloat32MultiArraySlice(dst, src []Float32MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Float32MultiArrayTypeSupport types.MessageTypeSupport = _Float32MultiArrayTypeSupport{}

type _Float32MultiArrayTypeSupport struct{}

func (t _Float32MultiArrayTypeSupport) New() types.Message {
	return NewFloat32MultiArray()
}

func (t _Float32MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Float32MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__Float32MultiArray__create())
}

func (t _Float32MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Float32MultiArray__destroy((*C.std_msgs__msg__Float32MultiArray)(pointer_to_free))
}

func (t _Float32MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float32MultiArray)
	mem := (*C.std_msgs__msg__Float32MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Float32MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float32MultiArray)
	mem := (*C.std_msgs__msg__Float32MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Float32__Sequence_to_Go(&m.Data, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Float32MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Float32MultiArray())
}

type CFloat32MultiArray = C.std_msgs__msg__Float32MultiArray
type CFloat32MultiArray__Sequence = C.std_msgs__msg__Float32MultiArray__Sequence

func Float32MultiArray__Sequence_to_Go(goSlice *[]Float32MultiArray, cSlice CFloat32MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float32MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Float32MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Float32MultiArray * uintptr(i)),
		))
		Float32MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Float32MultiArray__Sequence_to_C(cSlice *CFloat32MultiArray__Sequence, goSlice []Float32MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Float32MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Float32MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Float32MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Float32MultiArray * uintptr(i)),
		))
		Float32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Float32MultiArray__Array_to_Go(goSlice []Float32MultiArray, cSlice []CFloat32MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Float32MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float32MultiArray__Array_to_C(cSlice []CFloat32MultiArray, goSlice []Float32MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Float32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
