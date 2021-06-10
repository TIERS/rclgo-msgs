/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package action_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	action_msgs_msg "github.com/tiiuae/rclgo-msgs/action_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -laction_msgs__rosidl_typesupport_c -laction_msgs__rosidl_generator_c
#cgo LDFLAGS: -laction_msgs__rosidl_typesupport_c -laction_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/srv/cancel_goal.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/CancelGoal_Request", CancelGoal_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewCancelGoal_Request
// function instead.
type CancelGoal_Request struct {
	GoalInfo action_msgs_msg.GoalInfo `yaml:"goal_info"`// Goal info describing the goals to cancel, see above.
}

// NewCancelGoal_Request creates a new CancelGoal_Request with default values.
func NewCancelGoal_Request() *CancelGoal_Request {
	self := CancelGoal_Request{}
	self.SetDefaults()
	return &self
}

func (t *CancelGoal_Request) Clone() *CancelGoal_Request {
	c := &CancelGoal_Request{}
	c.GoalInfo = *t.GoalInfo.Clone()
	return c
}

func (t *CancelGoal_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CancelGoal_Request) SetDefaults() {
	t.GoalInfo.SetDefaults()
}

// CloneCancelGoal_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCancelGoal_RequestSlice(dst, src []CancelGoal_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CancelGoal_RequestTypeSupport types.MessageTypeSupport = _CancelGoal_RequestTypeSupport{}

type _CancelGoal_RequestTypeSupport struct{}

func (t _CancelGoal_RequestTypeSupport) New() types.Message {
	return NewCancelGoal_Request()
}

func (t _CancelGoal_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__srv__CancelGoal_Request
	return (unsafe.Pointer)(C.action_msgs__srv__CancelGoal_Request__create())
}

func (t _CancelGoal_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__srv__CancelGoal_Request__destroy((*C.action_msgs__srv__CancelGoal_Request)(pointer_to_free))
}

func (t _CancelGoal_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CancelGoal_Request)
	mem := (*C.action_msgs__srv__CancelGoal_Request)(dst)
	action_msgs_msg.GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_info), &m.GoalInfo)
}

func (t _CancelGoal_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CancelGoal_Request)
	mem := (*C.action_msgs__srv__CancelGoal_Request)(ros2_message_buffer)
	action_msgs_msg.GoalInfoTypeSupport.AsGoStruct(&m.GoalInfo, unsafe.Pointer(&mem.goal_info))
}

func (t _CancelGoal_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__srv__CancelGoal_Request())
}

type CCancelGoal_Request = C.action_msgs__srv__CancelGoal_Request
type CCancelGoal_Request__Sequence = C.action_msgs__srv__CancelGoal_Request__Sequence

func CancelGoal_Request__Sequence_to_Go(goSlice *[]CancelGoal_Request, cSlice CCancelGoal_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CancelGoal_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.action_msgs__srv__CancelGoal_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_action_msgs__srv__CancelGoal_Request * uintptr(i)),
		))
		CancelGoal_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CancelGoal_Request__Sequence_to_C(cSlice *CCancelGoal_Request__Sequence, goSlice []CancelGoal_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.action_msgs__srv__CancelGoal_Request)(C.malloc((C.size_t)(C.sizeof_struct_action_msgs__srv__CancelGoal_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.action_msgs__srv__CancelGoal_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_action_msgs__srv__CancelGoal_Request * uintptr(i)),
		))
		CancelGoal_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CancelGoal_Request__Array_to_Go(goSlice []CancelGoal_Request, cSlice []CCancelGoal_Request) {
	for i := 0; i < len(cSlice); i++ {
		CancelGoal_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CancelGoal_Request__Array_to_C(cSlice []CCancelGoal_Request, goSlice []CancelGoal_Request) {
	for i := 0; i < len(goSlice); i++ {
		CancelGoal_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
