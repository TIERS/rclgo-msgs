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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/get_map.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/GetMap_Request", GetMap_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetMap_Request
// function instead.
type GetMap_Request struct {
}

// NewGetMap_Request creates a new GetMap_Request with default values.
func NewGetMap_Request() *GetMap_Request {
	self := GetMap_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetMap_Request) Clone() *GetMap_Request {
	c := &GetMap_Request{}
	return c
}

func (t *GetMap_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetMap_Request) SetDefaults() {
}

// CloneGetMap_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetMap_RequestSlice(dst, src []GetMap_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetMap_RequestTypeSupport types.MessageTypeSupport = _GetMap_RequestTypeSupport{}

type _GetMap_RequestTypeSupport struct{}

func (t _GetMap_RequestTypeSupport) New() types.Message {
	return NewGetMap_Request()
}

func (t _GetMap_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__GetMap_Request
	return (unsafe.Pointer)(C.nav_msgs__srv__GetMap_Request__create())
}

func (t _GetMap_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__GetMap_Request__destroy((*C.nav_msgs__srv__GetMap_Request)(pointer_to_free))
}

func (t _GetMap_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _GetMap_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _GetMap_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__GetMap_Request())
}

type CGetMap_Request = C.nav_msgs__srv__GetMap_Request
type CGetMap_Request__Sequence = C.nav_msgs__srv__GetMap_Request__Sequence

func GetMap_Request__Sequence_to_Go(goSlice *[]GetMap_Request, cSlice CGetMap_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetMap_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__GetMap_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetMap_Request * uintptr(i)),
		))
		GetMap_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetMap_Request__Sequence_to_C(cSlice *CGetMap_Request__Sequence, goSlice []GetMap_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__GetMap_Request)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__GetMap_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__GetMap_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetMap_Request * uintptr(i)),
		))
		GetMap_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetMap_Request__Array_to_Go(goSlice []GetMap_Request, cSlice []CGetMap_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetMap_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetMap_Request__Array_to_C(cSlice []CGetMap_Request, goSlice []GetMap_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetMap_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
