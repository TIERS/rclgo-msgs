/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	diagnostic_msgs_msg "github.com/TIERS/rclgo-msgs/diagnostic_msgs/msg"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <diagnostic_msgs/srv/self_test.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("diagnostic_msgs/SelfTest_Response", SelfTest_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSelfTest_Response
// function instead.
type SelfTest_Response struct {
	Id string `yaml:"id"`
	Passed byte `yaml:"passed"`
	Status []diagnostic_msgs_msg.DiagnosticStatus `yaml:"status"`
}

// NewSelfTest_Response creates a new SelfTest_Response with default values.
func NewSelfTest_Response() *SelfTest_Response {
	self := SelfTest_Response{}
	self.SetDefaults()
	return &self
}

func (t *SelfTest_Response) Clone() *SelfTest_Response {
	c := &SelfTest_Response{}
	c.Id = t.Id
	c.Passed = t.Passed
	if t.Status != nil {
		c.Status = make([]diagnostic_msgs_msg.DiagnosticStatus, len(t.Status))
		diagnostic_msgs_msg.CloneDiagnosticStatusSlice(c.Status, t.Status)
	}
	return c
}

func (t *SelfTest_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SelfTest_Response) SetDefaults() {
	t.Id = ""
	t.Passed = 0
	t.Status = nil
}

// CloneSelfTest_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSelfTest_ResponseSlice(dst, src []SelfTest_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SelfTest_ResponseTypeSupport types.MessageTypeSupport = _SelfTest_ResponseTypeSupport{}

type _SelfTest_ResponseTypeSupport struct{}

func (t _SelfTest_ResponseTypeSupport) New() types.Message {
	return NewSelfTest_Response()
}

func (t _SelfTest_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__srv__SelfTest_Response
	return (unsafe.Pointer)(C.diagnostic_msgs__srv__SelfTest_Response__create())
}

func (t _SelfTest_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__srv__SelfTest_Response__destroy((*C.diagnostic_msgs__srv__SelfTest_Response)(pointer_to_free))
}

func (t _SelfTest_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SelfTest_Response)
	mem := (*C.diagnostic_msgs__srv__SelfTest_Response)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.id), m.Id)
	mem.passed = C.uint8_t(m.Passed)
	diagnostic_msgs_msg.DiagnosticStatus__Sequence_to_C((*diagnostic_msgs_msg.CDiagnosticStatus__Sequence)(unsafe.Pointer(&mem.status)), m.Status)
}

func (t _SelfTest_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SelfTest_Response)
	mem := (*C.diagnostic_msgs__srv__SelfTest_Response)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Id, unsafe.Pointer(&mem.id))
	m.Passed = byte(mem.passed)
	diagnostic_msgs_msg.DiagnosticStatus__Sequence_to_Go(&m.Status, *(*diagnostic_msgs_msg.CDiagnosticStatus__Sequence)(unsafe.Pointer(&mem.status)))
}

func (t _SelfTest_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__srv__SelfTest_Response())
}

type CSelfTest_Response = C.diagnostic_msgs__srv__SelfTest_Response
type CSelfTest_Response__Sequence = C.diagnostic_msgs__srv__SelfTest_Response__Sequence

func SelfTest_Response__Sequence_to_Go(goSlice *[]SelfTest_Response, cSlice CSelfTest_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SelfTest_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.diagnostic_msgs__srv__SelfTest_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(i)),
		))
		SelfTest_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SelfTest_Response__Sequence_to_C(cSlice *CSelfTest_Response__Sequence, goSlice []SelfTest_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.diagnostic_msgs__srv__SelfTest_Response)(C.malloc((C.size_t)(C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.diagnostic_msgs__srv__SelfTest_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__SelfTest_Response * uintptr(i)),
		))
		SelfTest_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SelfTest_Response__Array_to_Go(goSlice []SelfTest_Response, cSlice []CSelfTest_Response) {
	for i := 0; i < len(cSlice); i++ {
		SelfTest_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SelfTest_Response__Array_to_C(cSlice []CSelfTest_Response, goSlice []SelfTest_Response) {
	for i := 0; i < len(goSlice); i++ {
		SelfTest_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
