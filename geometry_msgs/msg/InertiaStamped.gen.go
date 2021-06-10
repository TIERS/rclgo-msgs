/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/inertia_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/InertiaStamped", InertiaStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewInertiaStamped
// function instead.
type InertiaStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Inertia Inertia `yaml:"inertia"`
}

// NewInertiaStamped creates a new InertiaStamped with default values.
func NewInertiaStamped() *InertiaStamped {
	self := InertiaStamped{}
	self.SetDefaults()
	return &self
}

func (t *InertiaStamped) Clone() *InertiaStamped {
	c := &InertiaStamped{}
	c.Header = *t.Header.Clone()
	c.Inertia = *t.Inertia.Clone()
	return c
}

func (t *InertiaStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InertiaStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Inertia.SetDefaults()
}

// CloneInertiaStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInertiaStampedSlice(dst, src []InertiaStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InertiaStampedTypeSupport types.MessageTypeSupport = _InertiaStampedTypeSupport{}

type _InertiaStampedTypeSupport struct{}

func (t _InertiaStampedTypeSupport) New() types.Message {
	return NewInertiaStamped()
}

func (t _InertiaStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__InertiaStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__InertiaStamped__create())
}

func (t _InertiaStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__InertiaStamped__destroy((*C.geometry_msgs__msg__InertiaStamped)(pointer_to_free))
}

func (t _InertiaStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InertiaStamped)
	mem := (*C.geometry_msgs__msg__InertiaStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	InertiaTypeSupport.AsCStruct(unsafe.Pointer(&mem.inertia), &m.Inertia)
}

func (t _InertiaStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InertiaStamped)
	mem := (*C.geometry_msgs__msg__InertiaStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	InertiaTypeSupport.AsGoStruct(&m.Inertia, unsafe.Pointer(&mem.inertia))
}

func (t _InertiaStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__InertiaStamped())
}

type CInertiaStamped = C.geometry_msgs__msg__InertiaStamped
type CInertiaStamped__Sequence = C.geometry_msgs__msg__InertiaStamped__Sequence

func InertiaStamped__Sequence_to_Go(goSlice *[]InertiaStamped, cSlice CInertiaStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InertiaStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__InertiaStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__InertiaStamped * uintptr(i)),
		))
		InertiaStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InertiaStamped__Sequence_to_C(cSlice *CInertiaStamped__Sequence, goSlice []InertiaStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__InertiaStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__InertiaStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__InertiaStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__InertiaStamped * uintptr(i)),
		))
		InertiaStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InertiaStamped__Array_to_Go(goSlice []InertiaStamped, cSlice []CInertiaStamped) {
	for i := 0; i < len(cSlice); i++ {
		InertiaStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InertiaStamped__Array_to_C(cSlice []CInertiaStamped, goSlice []InertiaStamped) {
	for i := 0; i < len(goSlice); i++ {
		InertiaStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
