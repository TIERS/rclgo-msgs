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

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	lifecycle_msgs_msg "github.com/tiiuae/rclgo-msgs/lifecycle_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/srv/get_available_transitions.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/GetAvailableTransitions_Response", GetAvailableTransitions_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetAvailableTransitions_Response
// function instead.
type GetAvailableTransitions_Response struct {
	AvailableTransitions []lifecycle_msgs_msg.TransitionDescription `yaml:"available_transitions"`// An array of the possible start_state-goal_state transitions
}

// NewGetAvailableTransitions_Response creates a new GetAvailableTransitions_Response with default values.
func NewGetAvailableTransitions_Response() *GetAvailableTransitions_Response {
	self := GetAvailableTransitions_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetAvailableTransitions_Response) Clone() *GetAvailableTransitions_Response {
	c := &GetAvailableTransitions_Response{}
	if t.AvailableTransitions != nil {
		c.AvailableTransitions = make([]lifecycle_msgs_msg.TransitionDescription, len(t.AvailableTransitions))
		lifecycle_msgs_msg.CloneTransitionDescriptionSlice(c.AvailableTransitions, t.AvailableTransitions)
	}
	return c
}

func (t *GetAvailableTransitions_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetAvailableTransitions_Response) SetDefaults() {
	t.AvailableTransitions = nil
}

// CloneGetAvailableTransitions_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetAvailableTransitions_ResponseSlice(dst, src []GetAvailableTransitions_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetAvailableTransitions_ResponseTypeSupport types.MessageTypeSupport = _GetAvailableTransitions_ResponseTypeSupport{}

type _GetAvailableTransitions_ResponseTypeSupport struct{}

func (t _GetAvailableTransitions_ResponseTypeSupport) New() types.Message {
	return NewGetAvailableTransitions_Response()
}

func (t _GetAvailableTransitions_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__srv__GetAvailableTransitions_Response
	return (unsafe.Pointer)(C.lifecycle_msgs__srv__GetAvailableTransitions_Response__create())
}

func (t _GetAvailableTransitions_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__srv__GetAvailableTransitions_Response__destroy((*C.lifecycle_msgs__srv__GetAvailableTransitions_Response)(pointer_to_free))
}

func (t _GetAvailableTransitions_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetAvailableTransitions_Response)
	mem := (*C.lifecycle_msgs__srv__GetAvailableTransitions_Response)(dst)
	lifecycle_msgs_msg.TransitionDescription__Sequence_to_C((*lifecycle_msgs_msg.CTransitionDescription__Sequence)(unsafe.Pointer(&mem.available_transitions)), m.AvailableTransitions)
}

func (t _GetAvailableTransitions_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetAvailableTransitions_Response)
	mem := (*C.lifecycle_msgs__srv__GetAvailableTransitions_Response)(ros2_message_buffer)
	lifecycle_msgs_msg.TransitionDescription__Sequence_to_Go(&m.AvailableTransitions, *(*lifecycle_msgs_msg.CTransitionDescription__Sequence)(unsafe.Pointer(&mem.available_transitions)))
}

func (t _GetAvailableTransitions_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__srv__GetAvailableTransitions_Response())
}

type CGetAvailableTransitions_Response = C.lifecycle_msgs__srv__GetAvailableTransitions_Response
type CGetAvailableTransitions_Response__Sequence = C.lifecycle_msgs__srv__GetAvailableTransitions_Response__Sequence

func GetAvailableTransitions_Response__Sequence_to_Go(goSlice *[]GetAvailableTransitions_Response, cSlice CGetAvailableTransitions_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetAvailableTransitions_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.lifecycle_msgs__srv__GetAvailableTransitions_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__GetAvailableTransitions_Response * uintptr(i)),
		))
		GetAvailableTransitions_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetAvailableTransitions_Response__Sequence_to_C(cSlice *CGetAvailableTransitions_Response__Sequence, goSlice []GetAvailableTransitions_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.lifecycle_msgs__srv__GetAvailableTransitions_Response)(C.malloc((C.size_t)(C.sizeof_struct_lifecycle_msgs__srv__GetAvailableTransitions_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.lifecycle_msgs__srv__GetAvailableTransitions_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__GetAvailableTransitions_Response * uintptr(i)),
		))
		GetAvailableTransitions_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetAvailableTransitions_Response__Array_to_Go(goSlice []GetAvailableTransitions_Response, cSlice []CGetAvailableTransitions_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetAvailableTransitions_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetAvailableTransitions_Response__Array_to_C(cSlice []CGetAvailableTransitions_Response, goSlice []GetAvailableTransitions_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetAvailableTransitions_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
