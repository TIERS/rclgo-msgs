/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/trigger.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Trigger_Request", Trigger_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewTrigger_Request
// function instead.
type Trigger_Request struct {
}

// NewTrigger_Request creates a new Trigger_Request with default values.
func NewTrigger_Request() *Trigger_Request {
	self := Trigger_Request{}
	self.SetDefaults()
	return &self
}

func (t *Trigger_Request) Clone() *Trigger_Request {
	c := &Trigger_Request{}
	return c
}

func (t *Trigger_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Trigger_Request) SetDefaults() {
}

// CloneTrigger_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTrigger_RequestSlice(dst, src []Trigger_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Trigger_RequestTypeSupport types.MessageTypeSupport = _Trigger_RequestTypeSupport{}

type _Trigger_RequestTypeSupport struct{}

func (t _Trigger_RequestTypeSupport) New() types.Message {
	return NewTrigger_Request()
}

func (t _Trigger_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__Trigger_Request
	return (unsafe.Pointer)(C.example_interfaces__srv__Trigger_Request__create())
}

func (t _Trigger_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__Trigger_Request__destroy((*C.example_interfaces__srv__Trigger_Request)(pointer_to_free))
}

func (t _Trigger_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Trigger_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Trigger_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__Trigger_Request())
}

type CTrigger_Request = C.example_interfaces__srv__Trigger_Request
type CTrigger_Request__Sequence = C.example_interfaces__srv__Trigger_Request__Sequence

func Trigger_Request__Sequence_to_Go(goSlice *[]Trigger_Request, cSlice CTrigger_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Trigger_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__srv__Trigger_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Request * uintptr(i)),
		))
		Trigger_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Trigger_Request__Sequence_to_C(cSlice *CTrigger_Request__Sequence, goSlice []Trigger_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__Trigger_Request)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__srv__Trigger_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__srv__Trigger_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__srv__Trigger_Request * uintptr(i)),
		))
		Trigger_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Trigger_Request__Array_to_Go(goSlice []Trigger_Request, cSlice []CTrigger_Request) {
	for i := 0; i < len(cSlice); i++ {
		Trigger_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Trigger_Request__Array_to_C(cSlice []CTrigger_Request, goSlice []Trigger_Request) {
	for i := 0; i < len(goSlice); i++ {
		Trigger_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
