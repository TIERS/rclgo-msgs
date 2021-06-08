/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_srvs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_srvs__rosidl_typesupport_c -lstd_srvs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_srvs/srv/trigger.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_srvs/Trigger_Response", Trigger_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewTrigger_Response
// function instead.
type Trigger_Response struct {
	Success bool `yaml:"success"`// indicate successful run of triggered service
	Message string `yaml:"message"`// informational, e.g. for error messages
}

// NewTrigger_Response creates a new Trigger_Response with default values.
func NewTrigger_Response() *Trigger_Response {
	self := Trigger_Response{}
	self.SetDefaults()
	return &self
}

func (t *Trigger_Response) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *Trigger_Response) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var Trigger_ResponseTypeSupport types.MessageTypeSupport = _Trigger_ResponseTypeSupport{}

type _Trigger_ResponseTypeSupport struct{}

func (t _Trigger_ResponseTypeSupport) New() types.Message {
	return NewTrigger_Response()
}

func (t _Trigger_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_srvs__srv__Trigger_Response
	return (unsafe.Pointer)(C.std_srvs__srv__Trigger_Response__create())
}

func (t _Trigger_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_srvs__srv__Trigger_Response__destroy((*C.std_srvs__srv__Trigger_Response)(pointer_to_free))
}

func (t _Trigger_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Trigger_Response)
	mem := (*C.std_srvs__srv__Trigger_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.message), m.Message)
}

func (t _Trigger_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Trigger_Response)
	mem := (*C.std_srvs__srv__Trigger_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.Message, unsafe.Pointer(&mem.message))
}

func (t _Trigger_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_srvs__srv__Trigger_Response())
}

type CTrigger_Response = C.std_srvs__srv__Trigger_Response
type CTrigger_Response__Sequence = C.std_srvs__srv__Trigger_Response__Sequence

func Trigger_Response__Sequence_to_Go(goSlice *[]Trigger_Response, cSlice CTrigger_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Trigger_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_srvs__srv__Trigger_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_srvs__srv__Trigger_Response * uintptr(i)),
		))
		Trigger_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Trigger_Response__Sequence_to_C(cSlice *CTrigger_Response__Sequence, goSlice []Trigger_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_srvs__srv__Trigger_Response)(C.malloc((C.size_t)(C.sizeof_struct_std_srvs__srv__Trigger_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_srvs__srv__Trigger_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_srvs__srv__Trigger_Response * uintptr(i)),
		))
		Trigger_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Trigger_Response__Array_to_Go(goSlice []Trigger_Response, cSlice []CTrigger_Response) {
	for i := 0; i < len(cSlice); i++ {
		Trigger_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Trigger_Response__Array_to_C(cSlice []CTrigger_Response, goSlice []Trigger_Response) {
	for i := 0; i < len(goSlice); i++ {
		Trigger_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}