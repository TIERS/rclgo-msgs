/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package composition_interfaces_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lcomposition_interfaces__rosidl_typesupport_c -lcomposition_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <composition_interfaces/srv/load_node.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("composition_interfaces/LoadNode_Response", LoadNode_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewLoadNode_Response
// function instead.
type LoadNode_Response struct {
	Success bool `yaml:"success"`// key/value arguments that are specific to a type of container process.True if the node was successfully loaded.
	ErrorMessage string `yaml:"error_message"`// Human readable error message if success is false, else empty string.
	FullNodeName string `yaml:"full_node_name"`// Name of the loaded composable node (including namespace).
	UniqueId uint64 `yaml:"unique_id"`// A unique identifier for the loaded node.
}

// NewLoadNode_Response creates a new LoadNode_Response with default values.
func NewLoadNode_Response() *LoadNode_Response {
	self := LoadNode_Response{}
	self.SetDefaults()
	return &self
}

func (t *LoadNode_Response) Clone() *LoadNode_Response {
	c := &LoadNode_Response{}
	c.Success = t.Success
	c.ErrorMessage = t.ErrorMessage
	c.FullNodeName = t.FullNodeName
	c.UniqueId = t.UniqueId
	return c
}

func (t *LoadNode_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *LoadNode_Response) SetDefaults() {
	t.Success = false
	t.ErrorMessage = ""
	t.FullNodeName = ""
	t.UniqueId = 0
}

// CloneLoadNode_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneLoadNode_ResponseSlice(dst, src []LoadNode_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var LoadNode_ResponseTypeSupport types.MessageTypeSupport = _LoadNode_ResponseTypeSupport{}

type _LoadNode_ResponseTypeSupport struct{}

func (t _LoadNode_ResponseTypeSupport) New() types.Message {
	return NewLoadNode_Response()
}

func (t _LoadNode_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.composition_interfaces__srv__LoadNode_Response
	return (unsafe.Pointer)(C.composition_interfaces__srv__LoadNode_Response__create())
}

func (t _LoadNode_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.composition_interfaces__srv__LoadNode_Response__destroy((*C.composition_interfaces__srv__LoadNode_Response)(pointer_to_free))
}

func (t _LoadNode_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*LoadNode_Response)
	mem := (*C.composition_interfaces__srv__LoadNode_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.error_message), m.ErrorMessage)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.full_node_name), m.FullNodeName)
	mem.unique_id = C.uint64_t(m.UniqueId)
}

func (t _LoadNode_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*LoadNode_Response)
	mem := (*C.composition_interfaces__srv__LoadNode_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.ErrorMessage, unsafe.Pointer(&mem.error_message))
	primitives.StringAsGoStruct(&m.FullNodeName, unsafe.Pointer(&mem.full_node_name))
	m.UniqueId = uint64(mem.unique_id)
}

func (t _LoadNode_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__composition_interfaces__srv__LoadNode_Response())
}

type CLoadNode_Response = C.composition_interfaces__srv__LoadNode_Response
type CLoadNode_Response__Sequence = C.composition_interfaces__srv__LoadNode_Response__Sequence

func LoadNode_Response__Sequence_to_Go(goSlice *[]LoadNode_Response, cSlice CLoadNode_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LoadNode_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.composition_interfaces__srv__LoadNode_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__LoadNode_Response * uintptr(i)),
		))
		LoadNode_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func LoadNode_Response__Sequence_to_C(cSlice *CLoadNode_Response__Sequence, goSlice []LoadNode_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.composition_interfaces__srv__LoadNode_Response)(C.malloc((C.size_t)(C.sizeof_struct_composition_interfaces__srv__LoadNode_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.composition_interfaces__srv__LoadNode_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__LoadNode_Response * uintptr(i)),
		))
		LoadNode_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func LoadNode_Response__Array_to_Go(goSlice []LoadNode_Response, cSlice []CLoadNode_Response) {
	for i := 0; i < len(cSlice); i++ {
		LoadNode_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func LoadNode_Response__Array_to_C(cSlice []CLoadNode_Response, goSlice []LoadNode_Response) {
	for i := 0; i < len(goSlice); i++ {
		LoadNode_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
