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
	nav_msgs_msg "github.com/TIERS/rclgo-msgs/nav_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/srv/get_map.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/GetMap_Response", GetMap_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetMap_Response
// function instead.
type GetMap_Response struct {
	Map nav_msgs_msg.OccupancyGrid `yaml:"map"`// Get the map as a nav_msgs/OccupancyGridThe current map hosted by this map service.
}

// NewGetMap_Response creates a new GetMap_Response with default values.
func NewGetMap_Response() *GetMap_Response {
	self := GetMap_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetMap_Response) Clone() *GetMap_Response {
	c := &GetMap_Response{}
	c.Map = *t.Map.Clone()
	return c
}

func (t *GetMap_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetMap_Response) SetDefaults() {
	t.Map.SetDefaults()
}

// CloneGetMap_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetMap_ResponseSlice(dst, src []GetMap_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetMap_ResponseTypeSupport types.MessageTypeSupport = _GetMap_ResponseTypeSupport{}

type _GetMap_ResponseTypeSupport struct{}

func (t _GetMap_ResponseTypeSupport) New() types.Message {
	return NewGetMap_Response()
}

func (t _GetMap_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__srv__GetMap_Response
	return (unsafe.Pointer)(C.nav_msgs__srv__GetMap_Response__create())
}

func (t _GetMap_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__srv__GetMap_Response__destroy((*C.nav_msgs__srv__GetMap_Response)(pointer_to_free))
}

func (t _GetMap_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetMap_Response)
	mem := (*C.nav_msgs__srv__GetMap_Response)(dst)
	nav_msgs_msg.OccupancyGridTypeSupport.AsCStruct(unsafe.Pointer(&mem._map), &m.Map)
}

func (t _GetMap_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetMap_Response)
	mem := (*C.nav_msgs__srv__GetMap_Response)(ros2_message_buffer)
	nav_msgs_msg.OccupancyGridTypeSupport.AsGoStruct(&m.Map, unsafe.Pointer(&mem._map))
}

func (t _GetMap_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__srv__GetMap_Response())
}

type CGetMap_Response = C.nav_msgs__srv__GetMap_Response
type CGetMap_Response__Sequence = C.nav_msgs__srv__GetMap_Response__Sequence

func GetMap_Response__Sequence_to_Go(goSlice *[]GetMap_Response, cSlice CGetMap_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetMap_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__srv__GetMap_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetMap_Response * uintptr(i)),
		))
		GetMap_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetMap_Response__Sequence_to_C(cSlice *CGetMap_Response__Sequence, goSlice []GetMap_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__srv__GetMap_Response)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__srv__GetMap_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__srv__GetMap_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__srv__GetMap_Response * uintptr(i)),
		))
		GetMap_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetMap_Response__Array_to_Go(goSlice []GetMap_Response, cSlice []CGetMap_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetMap_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetMap_Response__Array_to_C(cSlice []CGetMap_Response, goSlice []GetMap_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetMap_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
