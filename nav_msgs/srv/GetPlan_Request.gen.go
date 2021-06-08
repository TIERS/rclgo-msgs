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
	geometry_msgs_msg "github.com/tiiuae/rclgo-msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/get_plan.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/GetPlan_Request", GetPlan_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetPlan_Request
// function instead.
type GetPlan_Request struct {
	Start geometry_msgs_msg.PoseStamped `yaml:"start"`// The start pose for the plan
	Goal geometry_msgs_msg.PoseStamped `yaml:"goal"`// The final pose of the goal position
	Tolerance float32 `yaml:"tolerance"`// If the goal is obstructed, how many meters the planner canrelax the constraint in x and y before failing.
}

// NewGetPlan_Request creates a new GetPlan_Request with default values.
func NewGetPlan_Request() *GetPlan_Request {
	self := GetPlan_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetPlan_Request) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *GetPlan_Request) SetDefaults() {
	t.Start.SetDefaults()
	t.Goal.SetDefaults()
	
}

// Modifying this variable is undefined behavior.
var GetPlan_RequestTypeSupport types.MessageTypeSupport = _GetPlan_RequestTypeSupport{}

type _GetPlan_RequestTypeSupport struct{}

func (t _GetPlan_RequestTypeSupport) New() types.Message {
	return NewGetPlan_Request()
}

func (t _GetPlan_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__GetPlan_Request
	return (unsafe.Pointer)(C.nav_msgs__srv__GetPlan_Request__create())
}

func (t _GetPlan_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__GetPlan_Request__destroy((*C.nav_msgs__srv__GetPlan_Request)(pointer_to_free))
}

func (t _GetPlan_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetPlan_Request)
	mem := (*C.nav_msgs__srv__GetPlan_Request)(dst)
	geometry_msgs_msg.PoseStampedTypeSupport.AsCStruct(unsafe.Pointer(&mem.start), &m.Start)
	geometry_msgs_msg.PoseStampedTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal), &m.Goal)
	mem.tolerance = C.float(m.Tolerance)
}

func (t _GetPlan_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetPlan_Request)
	mem := (*C.nav_msgs__srv__GetPlan_Request)(ros2_message_buffer)
	geometry_msgs_msg.PoseStampedTypeSupport.AsGoStruct(&m.Start, unsafe.Pointer(&mem.start))
	geometry_msgs_msg.PoseStampedTypeSupport.AsGoStruct(&m.Goal, unsafe.Pointer(&mem.goal))
	m.Tolerance = float32(mem.tolerance)
}

func (t _GetPlan_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__GetPlan_Request())
}

type CGetPlan_Request = C.nav_msgs__srv__GetPlan_Request
type CGetPlan_Request__Sequence = C.nav_msgs__srv__GetPlan_Request__Sequence

func GetPlan_Request__Sequence_to_Go(goSlice *[]GetPlan_Request, cSlice CGetPlan_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetPlan_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__GetPlan_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(i)),
		))
		GetPlan_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetPlan_Request__Sequence_to_C(cSlice *CGetPlan_Request__Sequence, goSlice []GetPlan_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__GetPlan_Request)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__GetPlan_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetPlan_Request * uintptr(i)),
		))
		GetPlan_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetPlan_Request__Array_to_Go(goSlice []GetPlan_Request, cSlice []CGetPlan_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetPlan_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetPlan_Request__Array_to_C(cSlice []CGetPlan_Request, goSlice []GetPlan_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetPlan_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}