/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/string.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/String", StringTypeSupport)
}

// Do not create instances of this type directly. Always use NewString
// function instead.
type String struct {
	Data string `yaml:"data"`// This is an example message of using a primitive datatype, string.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewString creates a new String with default values.
func NewString() *String {
	self := String{}
	self.SetDefaults()
	return &self
}

func (t *String) Clone() *String {
	c := &String{}
	c.Data = t.Data
	return c
}

func (t *String) CloneMsg() types.Message {
	return t.Clone()
}

func (t *String) SetDefaults() {
	t.Data = ""
}

// CloneStringSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStringSlice(dst, src []String) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StringTypeSupport types.MessageTypeSupport = _StringTypeSupport{}

type _StringTypeSupport struct{}

func (t _StringTypeSupport) New() types.Message {
	return NewString()
}

func (t _StringTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__String
	return (unsafe.Pointer)(C.example_interfaces__msg__String__create())
}

func (t _StringTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__String__destroy((*C.example_interfaces__msg__String)(pointer_to_free))
}

func (t _StringTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*String)
	mem := (*C.example_interfaces__msg__String)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.data), m.Data)
}

func (t _StringTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*String)
	mem := (*C.example_interfaces__msg__String)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Data, unsafe.Pointer(&mem.data))
}

func (t _StringTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__String())
}

type CString = C.example_interfaces__msg__String
type CString__Sequence = C.example_interfaces__msg__String__Sequence

func String__Sequence_to_Go(goSlice *[]String, cSlice CString__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]String, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__String__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__String * uintptr(i)),
		))
		StringTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func String__Sequence_to_C(cSlice *CString__Sequence, goSlice []String) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__String)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__String * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__String)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__String * uintptr(i)),
		))
		StringTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func String__Array_to_Go(goSlice []String, cSlice []CString) {
	for i := 0; i < len(cSlice); i++ {
		StringTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func String__Array_to_C(cSlice []CString, goSlice []String) {
	for i := 0; i < len(goSlice); i++ {
		StringTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
