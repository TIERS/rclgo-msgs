/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	nav_msgs_msg "github.com/tiiuae/rclgo-msgs/nav_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/get_plan.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/GetPlan_Response", GetPlan_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetPlan_Response
// function instead.
type GetPlan_Response struct {
	Plan nav_msgs_msg.Path `yaml:"plan"`// If the goal is obstructed, how many meters the planner canrelax the constraint in x and y before failing.Array of poses from start to goal if one was successfully found.
}

// NewGetPlan_Response creates a new GetPlan_Response with default values.
func NewGetPlan_Response() *GetPlan_Response {
	self := GetPlan_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetPlan_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GetPlan_Response) SetDefaults() {
	t.Plan.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var GetPlan_ResponseTypeSupport types.MessageTypeSupport = _GetPlan_ResponseTypeSupport{}

type _GetPlan_ResponseTypeSupport struct{}

func (t _GetPlan_ResponseTypeSupport) New() types.Message {
	return NewGetPlan_Response()
}

func (t _GetPlan_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__GetPlan_Response
	return (unsafe.Pointer)(C.nav_msgs__srv__GetPlan_Response__create())
}

func (t _GetPlan_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__GetPlan_Response__destroy((*C.nav_msgs__srv__GetPlan_Response)(pointer_to_free))
}

func (t _GetPlan_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetPlan_Response)
	mem := (*C.nav_msgs__srv__GetPlan_Response)(dst)
	nav_msgs_msg.PathTypeSupport.AsCStruct(unsafe.Pointer(&mem.plan), &m.Plan)
}

func (t _GetPlan_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetPlan_Response)
	mem := (*C.nav_msgs__srv__GetPlan_Response)(ros2_message_buffer)
	nav_msgs_msg.PathTypeSupport.AsGoStruct(&m.Plan, unsafe.Pointer(&mem.plan))
}

func (t _GetPlan_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__GetPlan_Response())
}

type CGetPlan_Response = C.nav_msgs__srv__GetPlan_Response
type CGetPlan_Response__Sequence = C.nav_msgs__srv__GetPlan_Response__Sequence

func GetPlan_Response__Sequence_to_Go(goSlice *[]GetPlan_Response, cSlice CGetPlan_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetPlan_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__GetPlan_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Response * uintptr(i)),
		))
		GetPlan_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetPlan_Response__Sequence_to_C(cSlice *CGetPlan_Response__Sequence, goSlice []GetPlan_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__GetPlan_Response)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__GetPlan_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__GetPlan_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Response * uintptr(i)),
		))
		GetPlan_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetPlan_Response__Array_to_Go(goSlice []GetPlan_Response, cSlice []CGetPlan_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetPlan_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetPlan_Response__Array_to_C(cSlice []CGetPlan_Response, goSlice []GetPlan_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetPlan_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}