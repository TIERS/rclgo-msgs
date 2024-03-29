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

#include <turtlesim/srv/teleport_absolute.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/TeleportAbsolute_Response", TeleportAbsolute_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewTeleportAbsolute_Response
// function instead.
type TeleportAbsolute_Response struct {
}

// NewTeleportAbsolute_Response creates a new TeleportAbsolute_Response with default values.
func NewTeleportAbsolute_Response() *TeleportAbsolute_Response {
	self := TeleportAbsolute_Response{}
	self.SetDefaults()
	return &self
}

func (t *TeleportAbsolute_Response) Clone() *TeleportAbsolute_Response {
	c := &TeleportAbsolute_Response{}
	return c
}

func (t *TeleportAbsolute_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TeleportAbsolute_Response) SetDefaults() {
}

// CloneTeleportAbsolute_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTeleportAbsolute_ResponseSlice(dst, src []TeleportAbsolute_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TeleportAbsolute_ResponseTypeSupport types.MessageTypeSupport = _TeleportAbsolute_ResponseTypeSupport{}

type _TeleportAbsolute_ResponseTypeSupport struct{}

func (t _TeleportAbsolute_ResponseTypeSupport) New() types.Message {
	return NewTeleportAbsolute_Response()
}

func (t _TeleportAbsolute_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__TeleportAbsolute_Response
	return (unsafe.Pointer)(C.turtlesim__srv__TeleportAbsolute_Response__create())
}

func (t _TeleportAbsolute_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__TeleportAbsolute_Response__destroy((*C.turtlesim__srv__TeleportAbsolute_Response)(pointer_to_free))
}

func (t _TeleportAbsolute_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _TeleportAbsolute_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _TeleportAbsolute_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__TeleportAbsolute_Response())
}

type CTeleportAbsolute_Response = C.turtlesim__srv__TeleportAbsolute_Response
type CTeleportAbsolute_Response__Sequence = C.turtlesim__srv__TeleportAbsolute_Response__Sequence

func TeleportAbsolute_Response__Sequence_to_Go(goSlice *[]TeleportAbsolute_Response, cSlice CTeleportAbsolute_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TeleportAbsolute_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__TeleportAbsolute_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Response * uintptr(i)),
		))
		TeleportAbsolute_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TeleportAbsolute_Response__Sequence_to_C(cSlice *CTeleportAbsolute_Response__Sequence, goSlice []TeleportAbsolute_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__TeleportAbsolute_Response)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__TeleportAbsolute_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Response * uintptr(i)),
		))
		TeleportAbsolute_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TeleportAbsolute_Response__Array_to_Go(goSlice []TeleportAbsolute_Response, cSlice []CTeleportAbsolute_Response) {
	for i := 0; i < len(cSlice); i++ {
		TeleportAbsolute_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TeleportAbsolute_Response__Array_to_C(cSlice []CTeleportAbsolute_Response, goSlice []TeleportAbsolute_Response) {
	for i := 0; i < len(goSlice); i++ {
		TeleportAbsolute_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
