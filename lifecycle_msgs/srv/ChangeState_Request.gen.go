/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package lifecycle_msgs_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	lifecycle_msgs_msg "github.com/TIERS/rclgo-msgs/lifecycle_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/srv/change_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/ChangeState_Request", ChangeState_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewChangeState_Request
// function instead.
type ChangeState_Request struct {
	Transition lifecycle_msgs_msg.Transition `yaml:"transition"`// The requested transition.This change state service will fail if the transition is not possible.
}

// NewChangeState_Request creates a new ChangeState_Request with default values.
func NewChangeState_Request() *ChangeState_Request {
	self := ChangeState_Request{}
	self.SetDefaults()
	return &self
}

func (t *ChangeState_Request) Clone() *ChangeState_Request {
	c := &ChangeState_Request{}
	c.Transition = *t.Transition.Clone()
	return c
}

func (t *ChangeState_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ChangeState_Request) SetDefaults() {
	t.Transition.SetDefaults()
}

// CloneChangeState_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneChangeState_RequestSlice(dst, src []ChangeState_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ChangeState_RequestTypeSupport types.MessageTypeSupport = _ChangeState_RequestTypeSupport{}

type _ChangeState_RequestTypeSupport struct{}

func (t _ChangeState_RequestTypeSupport) New() types.Message {
	return NewChangeState_Request()
}

func (t _ChangeState_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__srv__ChangeState_Request
	return (unsafe.Pointer)(C.lifecycle_msgs__srv__ChangeState_Request__create())
}

func (t _ChangeState_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__srv__ChangeState_Request__destroy((*C.lifecycle_msgs__srv__ChangeState_Request)(pointer_to_free))
}

func (t _ChangeState_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ChangeState_Request)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Request)(dst)
	lifecycle_msgs_msg.TransitionTypeSupport.AsCStruct(unsafe.Pointer(&mem.transition), &m.Transition)
}

func (t _ChangeState_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ChangeState_Request)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Request)(ros2_message_buffer)
	lifecycle_msgs_msg.TransitionTypeSupport.AsGoStruct(&m.Transition, unsafe.Pointer(&mem.transition))
}

func (t _ChangeState_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__srv__ChangeState_Request())
}

type CChangeState_Request = C.lifecycle_msgs__srv__ChangeState_Request
type CChangeState_Request__Sequence = C.lifecycle_msgs__srv__ChangeState_Request__Sequence

func ChangeState_Request__Sequence_to_Go(goSlice *[]ChangeState_Request, cSlice CChangeState_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ChangeState_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.lifecycle_msgs__srv__ChangeState_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Request * uintptr(i)),
		))
		ChangeState_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ChangeState_Request__Sequence_to_C(cSlice *CChangeState_Request__Sequence, goSlice []ChangeState_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.lifecycle_msgs__srv__ChangeState_Request)(C.malloc((C.size_t)(C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.lifecycle_msgs__srv__ChangeState_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Request * uintptr(i)),
		))
		ChangeState_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ChangeState_Request__Array_to_Go(goSlice []ChangeState_Request, cSlice []CChangeState_Request) {
	for i := 0; i < len(cSlice); i++ {
		ChangeState_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ChangeState_Request__Array_to_C(cSlice []CChangeState_Request, goSlice []ChangeState_Request) {
	for i := 0; i < len(goSlice); i++ {
		ChangeState_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
