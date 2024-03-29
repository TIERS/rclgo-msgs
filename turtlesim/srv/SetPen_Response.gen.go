/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/set_pen.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/SetPen_Response", SetPen_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetPen_Response
// function instead.
type SetPen_Response struct {
}

// NewSetPen_Response creates a new SetPen_Response with default values.
func NewSetPen_Response() *SetPen_Response {
	self := SetPen_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetPen_Response) Clone() *SetPen_Response {
	c := &SetPen_Response{}
	return c
}

func (t *SetPen_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetPen_Response) SetDefaults() {
}

// CloneSetPen_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetPen_ResponseSlice(dst, src []SetPen_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetPen_ResponseTypeSupport types.MessageTypeSupport = _SetPen_ResponseTypeSupport{}

type _SetPen_ResponseTypeSupport struct{}

func (t _SetPen_ResponseTypeSupport) New() types.Message {
	return NewSetPen_Response()
}

func (t _SetPen_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__SetPen_Response
	return (unsafe.Pointer)(C.turtlesim__srv__SetPen_Response__create())
}

func (t _SetPen_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__SetPen_Response__destroy((*C.turtlesim__srv__SetPen_Response)(pointer_to_free))
}

func (t _SetPen_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _SetPen_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _SetPen_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__SetPen_Response())
}

type CSetPen_Response = C.turtlesim__srv__SetPen_Response
type CSetPen_Response__Sequence = C.turtlesim__srv__SetPen_Response__Sequence

func SetPen_Response__Sequence_to_Go(goSlice *[]SetPen_Response, cSlice CSetPen_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetPen_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__SetPen_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__SetPen_Response * uintptr(i)),
		))
		SetPen_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetPen_Response__Sequence_to_C(cSlice *CSetPen_Response__Sequence, goSlice []SetPen_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__SetPen_Response)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__SetPen_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__SetPen_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__SetPen_Response * uintptr(i)),
		))
		SetPen_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetPen_Response__Array_to_Go(goSlice []SetPen_Response, cSlice []CSetPen_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetPen_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetPen_Response__Array_to_C(cSlice []CSetPen_Response, goSlice []SetPen_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetPen_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
