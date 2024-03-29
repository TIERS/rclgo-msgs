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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/parameter_descriptor.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/ParameterDescriptor", ParameterDescriptorTypeSupport)
}

// Do not create instances of this type directly. Always use NewParameterDescriptor
// function instead.
type ParameterDescriptor struct {
	Name string `yaml:"name"`// The name of the parameter.
	Type uint8 `yaml:"type"`// Enum values are defined in the `ParameterType.msg` message.
	Description string `yaml:"description"`// Description of the parameter, visible from introspection tools.
	AdditionalConstraints string `yaml:"additional_constraints"`// Plain English description of additional constraints which cannot be expressedwith the available constraints, e.g. "only prime numbers".By convention, this should only be used to clarify constraints which cannotbe completely expressed with the parameter constraints below.
	ReadOnly bool `yaml:"read_only"`// If 'true' then the value cannot change after it has been initialized.
	FloatingPointRange []FloatingPointRange `yaml:"floating_point_range"`// FloatingPointRange consists of a from_value, a to_value, and a step.
	IntegerRange []IntegerRange `yaml:"integer_range"`// IntegerRange consists of a from_value, a to_value, and a step.
}

// NewParameterDescriptor creates a new ParameterDescriptor with default values.
func NewParameterDescriptor() *ParameterDescriptor {
	self := ParameterDescriptor{}
	self.SetDefaults()
	return &self
}

func (t *ParameterDescriptor) Clone() *ParameterDescriptor {
	c := &ParameterDescriptor{}
	c.Name = t.Name
	c.Type = t.Type
	c.Description = t.Description
	c.AdditionalConstraints = t.AdditionalConstraints
	c.ReadOnly = t.ReadOnly
	if t.FloatingPointRange != nil {
		c.FloatingPointRange = make([]FloatingPointRange, len(t.FloatingPointRange))
		CloneFloatingPointRangeSlice(c.FloatingPointRange, t.FloatingPointRange)
	}
	if t.IntegerRange != nil {
		c.IntegerRange = make([]IntegerRange, len(t.IntegerRange))
		CloneIntegerRangeSlice(c.IntegerRange, t.IntegerRange)
	}
	return c
}

func (t *ParameterDescriptor) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ParameterDescriptor) SetDefaults() {
	t.Name = ""
	t.Type = 0
	t.Description = ""
	t.AdditionalConstraints = ""
	t.ReadOnly = false
	t.FloatingPointRange = nil
	t.IntegerRange = nil
}

// CloneParameterDescriptorSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneParameterDescriptorSlice(dst, src []ParameterDescriptor) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ParameterDescriptorTypeSupport types.MessageTypeSupport = _ParameterDescriptorTypeSupport{}

type _ParameterDescriptorTypeSupport struct{}

func (t _ParameterDescriptorTypeSupport) New() types.Message {
	return NewParameterDescriptor()
}

func (t _ParameterDescriptorTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__ParameterDescriptor
	return (unsafe.Pointer)(C.rcl_interfaces__msg__ParameterDescriptor__create())
}

func (t _ParameterDescriptorTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__ParameterDescriptor__destroy((*C.rcl_interfaces__msg__ParameterDescriptor)(pointer_to_free))
}

func (t _ParameterDescriptorTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ParameterDescriptor)
	mem := (*C.rcl_interfaces__msg__ParameterDescriptor)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	mem._type = C.uint8_t(m.Type)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.description), m.Description)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.additional_constraints), m.AdditionalConstraints)
	mem.read_only = C.bool(m.ReadOnly)
	FloatingPointRange__Sequence_to_C(&mem.floating_point_range, m.FloatingPointRange)
	IntegerRange__Sequence_to_C(&mem.integer_range, m.IntegerRange)
}

func (t _ParameterDescriptorTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ParameterDescriptor)
	mem := (*C.rcl_interfaces__msg__ParameterDescriptor)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	m.Type = uint8(mem._type)
	primitives.StringAsGoStruct(&m.Description, unsafe.Pointer(&mem.description))
	primitives.StringAsGoStruct(&m.AdditionalConstraints, unsafe.Pointer(&mem.additional_constraints))
	m.ReadOnly = bool(mem.read_only)
	FloatingPointRange__Sequence_to_Go(&m.FloatingPointRange, mem.floating_point_range)
	IntegerRange__Sequence_to_Go(&m.IntegerRange, mem.integer_range)
}

func (t _ParameterDescriptorTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__ParameterDescriptor())
}

type CParameterDescriptor = C.rcl_interfaces__msg__ParameterDescriptor
type CParameterDescriptor__Sequence = C.rcl_interfaces__msg__ParameterDescriptor__Sequence

func ParameterDescriptor__Sequence_to_Go(goSlice *[]ParameterDescriptor, cSlice CParameterDescriptor__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParameterDescriptor, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__msg__ParameterDescriptor__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(i)),
		))
		ParameterDescriptorTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ParameterDescriptor__Sequence_to_C(cSlice *CParameterDescriptor__Sequence, goSlice []ParameterDescriptor) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__ParameterDescriptor)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__msg__ParameterDescriptor)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__ParameterDescriptor * uintptr(i)),
		))
		ParameterDescriptorTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ParameterDescriptor__Array_to_Go(goSlice []ParameterDescriptor, cSlice []CParameterDescriptor) {
	for i := 0; i < len(cSlice); i++ {
		ParameterDescriptorTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ParameterDescriptor__Array_to_C(cSlice []CParameterDescriptor, goSlice []ParameterDescriptor) {
	for i := 0; i < len(goSlice); i++ {
		ParameterDescriptorTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
