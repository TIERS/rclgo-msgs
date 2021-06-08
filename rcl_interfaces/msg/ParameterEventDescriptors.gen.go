/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/parameter_event_descriptors.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/ParameterEventDescriptors", ParameterEventDescriptorsTypeSupport)
}

// Do not create instances of this type directly. Always use NewParameterEventDescriptors
// function instead.
type ParameterEventDescriptors struct {
	NewParameters []ParameterDescriptor `yaml:"new_parameters"`
	ChangedParameters []ParameterDescriptor `yaml:"changed_parameters"`
	DeletedParameters []ParameterDescriptor `yaml:"deleted_parameters"`
}

// NewParameterEventDescriptors creates a new ParameterEventDescriptors with default values.
func NewParameterEventDescriptors() *ParameterEventDescriptors {
	self := ParameterEventDescriptors{}
	self.SetDefaults()
	return &self
}

func (t *ParameterEventDescriptors) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *ParameterEventDescriptors) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var ParameterEventDescriptorsTypeSupport types.MessageTypeSupport = _ParameterEventDescriptorsTypeSupport{}

type _ParameterEventDescriptorsTypeSupport struct{}

func (t _ParameterEventDescriptorsTypeSupport) New() types.Message {
	return NewParameterEventDescriptors()
}

func (t _ParameterEventDescriptorsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__ParameterEventDescriptors
	return (unsafe.Pointer)(C.rcl_interfaces__msg__ParameterEventDescriptors__create())
}

func (t _ParameterEventDescriptorsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__ParameterEventDescriptors__destroy((*C.rcl_interfaces__msg__ParameterEventDescriptors)(pointer_to_free))
}

func (t _ParameterEventDescriptorsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ParameterEventDescriptors)
	mem := (*C.rcl_interfaces__msg__ParameterEventDescriptors)(dst)
	ParameterDescriptor__Sequence_to_C(&mem.new_parameters, m.NewParameters)
	ParameterDescriptor__Sequence_to_C(&mem.changed_parameters, m.ChangedParameters)
	ParameterDescriptor__Sequence_to_C(&mem.deleted_parameters, m.DeletedParameters)
}

func (t _ParameterEventDescriptorsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ParameterEventDescriptors)
	mem := (*C.rcl_interfaces__msg__ParameterEventDescriptors)(ros2_message_buffer)
	ParameterDescriptor__Sequence_to_Go(&m.NewParameters, mem.new_parameters)
	ParameterDescriptor__Sequence_to_Go(&m.ChangedParameters, mem.changed_parameters)
	ParameterDescriptor__Sequence_to_Go(&m.DeletedParameters, mem.deleted_parameters)
}

func (t _ParameterEventDescriptorsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__ParameterEventDescriptors())
}

type CParameterEventDescriptors = C.rcl_interfaces__msg__ParameterEventDescriptors
type CParameterEventDescriptors__Sequence = C.rcl_interfaces__msg__ParameterEventDescriptors__Sequence

func ParameterEventDescriptors__Sequence_to_Go(goSlice *[]ParameterEventDescriptors, cSlice CParameterEventDescriptors__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParameterEventDescriptors, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__msg__ParameterEventDescriptors__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterEventDescriptors * uintptr(i)),
		))
		ParameterEventDescriptorsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ParameterEventDescriptors__Sequence_to_C(cSlice *CParameterEventDescriptors__Sequence, goSlice []ParameterEventDescriptors) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__ParameterEventDescriptors)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__msg__ParameterEventDescriptors * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__msg__ParameterEventDescriptors)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterEventDescriptors * uintptr(i)),
		))
		ParameterEventDescriptorsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ParameterEventDescriptors__Array_to_Go(goSlice []ParameterEventDescriptors, cSlice []CParameterEventDescriptors) {
	for i := 0; i < len(cSlice); i++ {
		ParameterEventDescriptorsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ParameterEventDescriptors__Array_to_C(cSlice []CParameterEventDescriptors, goSlice []ParameterEventDescriptors) {
	for i := 0; i < len(goSlice); i++ {
		ParameterEventDescriptorsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}