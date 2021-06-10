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

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <diagnostic_msgs/srv/add_diagnostics.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("diagnostic_msgs/AddDiagnostics_Request", AddDiagnostics_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewAddDiagnostics_Request
// function instead.
type AddDiagnostics_Request struct {
	LoadNamespace string `yaml:"load_namespace"`// The load_namespace parameter defines the namespace where parameters for theinitialization of analyzers in the diagnostic aggregator have been loaded. Thevalue should be a global name (i.e. /my/name/space), not a relative(my/name/space) or private (~my/name/space) name. Analyzers will not be addedif a non-global name is used. The call will also fail if the namespacecontains parameters that follow a namespace structure that does not conform tothat expected by the analyzer definitions. Seehttp://wiki.ros.org/diagnostics/Tutorials/Configuring%20Diagnostic%20Aggregatorsand http://wiki.ros.org/diagnostics/Tutorials/Using%20the%20GenericAnalyzerfor examples of the structure of yaml files which are expected to have beenloaded into the namespace.
}

// NewAddDiagnostics_Request creates a new AddDiagnostics_Request with default values.
func NewAddDiagnostics_Request() *AddDiagnostics_Request {
	self := AddDiagnostics_Request{}
	self.SetDefaults()
	return &self
}

func (t *AddDiagnostics_Request) Clone() *AddDiagnostics_Request {
	c := &AddDiagnostics_Request{}
	c.LoadNamespace = t.LoadNamespace
	return c
}

func (t *AddDiagnostics_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AddDiagnostics_Request) SetDefaults() {
	t.LoadNamespace = ""
}

// CloneAddDiagnostics_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAddDiagnostics_RequestSlice(dst, src []AddDiagnostics_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AddDiagnostics_RequestTypeSupport types.MessageTypeSupport = _AddDiagnostics_RequestTypeSupport{}

type _AddDiagnostics_RequestTypeSupport struct{}

func (t _AddDiagnostics_RequestTypeSupport) New() types.Message {
	return NewAddDiagnostics_Request()
}

func (t _AddDiagnostics_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__srv__AddDiagnostics_Request
	return (unsafe.Pointer)(C.diagnostic_msgs__srv__AddDiagnostics_Request__create())
}

func (t _AddDiagnostics_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__srv__AddDiagnostics_Request__destroy((*C.diagnostic_msgs__srv__AddDiagnostics_Request)(pointer_to_free))
}

func (t _AddDiagnostics_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AddDiagnostics_Request)
	mem := (*C.diagnostic_msgs__srv__AddDiagnostics_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.load_namespace), m.LoadNamespace)
}

func (t _AddDiagnostics_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AddDiagnostics_Request)
	mem := (*C.diagnostic_msgs__srv__AddDiagnostics_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.LoadNamespace, unsafe.Pointer(&mem.load_namespace))
}

func (t _AddDiagnostics_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__srv__AddDiagnostics_Request())
}

type CAddDiagnostics_Request = C.diagnostic_msgs__srv__AddDiagnostics_Request
type CAddDiagnostics_Request__Sequence = C.diagnostic_msgs__srv__AddDiagnostics_Request__Sequence

func AddDiagnostics_Request__Sequence_to_Go(goSlice *[]AddDiagnostics_Request, cSlice CAddDiagnostics_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AddDiagnostics_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.diagnostic_msgs__srv__AddDiagnostics_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__AddDiagnostics_Request * uintptr(i)),
		))
		AddDiagnostics_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func AddDiagnostics_Request__Sequence_to_C(cSlice *CAddDiagnostics_Request__Sequence, goSlice []AddDiagnostics_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.diagnostic_msgs__srv__AddDiagnostics_Request)(C.malloc((C.size_t)(C.sizeof_struct_diagnostic_msgs__srv__AddDiagnostics_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.diagnostic_msgs__srv__AddDiagnostics_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__srv__AddDiagnostics_Request * uintptr(i)),
		))
		AddDiagnostics_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func AddDiagnostics_Request__Array_to_Go(goSlice []AddDiagnostics_Request, cSlice []CAddDiagnostics_Request) {
	for i := 0; i < len(cSlice); i++ {
		AddDiagnostics_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AddDiagnostics_Request__Array_to_C(cSlice []CAddDiagnostics_Request, goSlice []AddDiagnostics_Request) {
	for i := 0; i < len(goSlice); i++ {
		AddDiagnostics_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
