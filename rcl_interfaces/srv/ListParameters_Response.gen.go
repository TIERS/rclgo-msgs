/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	rcl_interfaces_msg "github.com/TIERS/rclgo-msgs/rcl_interfaces/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/list_parameters.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/ListParameters_Response", ListParameters_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewListParameters_Response
// function instead.
type ListParameters_Response struct {
	Result rcl_interfaces_msg.ListParametersResult `yaml:"result"`// The list of parameter names and their prefixes.
}

// NewListParameters_Response creates a new ListParameters_Response with default values.
func NewListParameters_Response() *ListParameters_Response {
	self := ListParameters_Response{}
	self.SetDefaults()
	return &self
}

func (t *ListParameters_Response) Clone() *ListParameters_Response {
	c := &ListParameters_Response{}
	c.Result = *t.Result.Clone()
	return c
}

func (t *ListParameters_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ListParameters_Response) SetDefaults() {
	t.Result.SetDefaults()
}

// CloneListParameters_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneListParameters_ResponseSlice(dst, src []ListParameters_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ListParameters_ResponseTypeSupport types.MessageTypeSupport = _ListParameters_ResponseTypeSupport{}

type _ListParameters_ResponseTypeSupport struct{}

func (t _ListParameters_ResponseTypeSupport) New() types.Message {
	return NewListParameters_Response()
}

func (t _ListParameters_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__ListParameters_Response
	return (unsafe.Pointer)(C.rcl_interfaces__srv__ListParameters_Response__create())
}

func (t _ListParameters_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__ListParameters_Response__destroy((*C.rcl_interfaces__srv__ListParameters_Response)(pointer_to_free))
}

func (t _ListParameters_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ListParameters_Response)
	mem := (*C.rcl_interfaces__srv__ListParameters_Response)(dst)
	rcl_interfaces_msg.ListParametersResultTypeSupport.AsCStruct(unsafe.Pointer(&mem.result), &m.Result)
}

func (t _ListParameters_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ListParameters_Response)
	mem := (*C.rcl_interfaces__srv__ListParameters_Response)(ros2_message_buffer)
	rcl_interfaces_msg.ListParametersResultTypeSupport.AsGoStruct(&m.Result, unsafe.Pointer(&mem.result))
}

func (t _ListParameters_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__ListParameters_Response())
}

type CListParameters_Response = C.rcl_interfaces__srv__ListParameters_Response
type CListParameters_Response__Sequence = C.rcl_interfaces__srv__ListParameters_Response__Sequence

func ListParameters_Response__Sequence_to_Go(goSlice *[]ListParameters_Response, cSlice CListParameters_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ListParameters_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__ListParameters_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__ListParameters_Response * uintptr(i)),
		))
		ListParameters_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ListParameters_Response__Sequence_to_C(cSlice *CListParameters_Response__Sequence, goSlice []ListParameters_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__ListParameters_Response)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__ListParameters_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__ListParameters_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__ListParameters_Response * uintptr(i)),
		))
		ListParameters_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ListParameters_Response__Array_to_Go(goSlice []ListParameters_Response, cSlice []CListParameters_Response) {
	for i := 0; i < len(cSlice); i++ {
		ListParameters_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ListParameters_Response__Array_to_C(cSlice []CListParameters_Response, goSlice []ListParameters_Response) {
	for i := 0; i < len(goSlice); i++ {
		ListParameters_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
