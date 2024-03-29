/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/get_map_roi.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/GetMapROI_Request", GetMapROI_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetMapROI_Request
// function instead.
type GetMapROI_Request struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	LX float64 `yaml:"l_x"`
	LY float64 `yaml:"l_y"`
}

// NewGetMapROI_Request creates a new GetMapROI_Request with default values.
func NewGetMapROI_Request() *GetMapROI_Request {
	self := GetMapROI_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetMapROI_Request) Clone() *GetMapROI_Request {
	c := &GetMapROI_Request{}
	c.X = t.X
	c.Y = t.Y
	c.LX = t.LX
	c.LY = t.LY
	return c
}

func (t *GetMapROI_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetMapROI_Request) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.LX = 0
	t.LY = 0
}

// CloneGetMapROI_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetMapROI_RequestSlice(dst, src []GetMapROI_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetMapROI_RequestTypeSupport types.MessageTypeSupport = _GetMapROI_RequestTypeSupport{}

type _GetMapROI_RequestTypeSupport struct{}

func (t _GetMapROI_RequestTypeSupport) New() types.Message {
	return NewGetMapROI_Request()
}

func (t _GetMapROI_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__GetMapROI_Request
	return (unsafe.Pointer)(C.map_msgs__srv__GetMapROI_Request__create())
}

func (t _GetMapROI_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__GetMapROI_Request__destroy((*C.map_msgs__srv__GetMapROI_Request)(pointer_to_free))
}

func (t _GetMapROI_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetMapROI_Request)
	mem := (*C.map_msgs__srv__GetMapROI_Request)(dst)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.l_x = C.double(m.LX)
	mem.l_y = C.double(m.LY)
}

func (t _GetMapROI_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetMapROI_Request)
	mem := (*C.map_msgs__srv__GetMapROI_Request)(ros2_message_buffer)
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.LX = float64(mem.l_x)
	m.LY = float64(mem.l_y)
}

func (t _GetMapROI_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__GetMapROI_Request())
}

type CGetMapROI_Request = C.map_msgs__srv__GetMapROI_Request
type CGetMapROI_Request__Sequence = C.map_msgs__srv__GetMapROI_Request__Sequence

func GetMapROI_Request__Sequence_to_Go(goSlice *[]GetMapROI_Request, cSlice CGetMapROI_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetMapROI_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__GetMapROI_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__GetMapROI_Request * uintptr(i)),
		))
		GetMapROI_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetMapROI_Request__Sequence_to_C(cSlice *CGetMapROI_Request__Sequence, goSlice []GetMapROI_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__GetMapROI_Request)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__GetMapROI_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__GetMapROI_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__GetMapROI_Request * uintptr(i)),
		))
		GetMapROI_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetMapROI_Request__Array_to_Go(goSlice []GetMapROI_Request, cSlice []CGetMapROI_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetMapROI_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetMapROI_Request__Array_to_C(cSlice []CGetMapROI_Request, goSlice []GetMapROI_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetMapROI_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
