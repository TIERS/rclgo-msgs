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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/msg/integer_range.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/IntegerRange", IntegerRangeTypeSupport)
}

// Do not create instances of this type directly. Always use NewIntegerRange
// function instead.
type IntegerRange struct {
	FromValue int64 `yaml:"from_value"`// Start value for valid values, inclusive.
	ToValue int64 `yaml:"to_value"`// End value for valid values, inclusive.
	Step uint64 `yaml:"step"`// Size of valid steps between the from and to bound.A step value of zero implies a continuous range of values. Ideally, the stepwould be less than or equal to the distance between the bounds, as well as aneven multiple of the distance between the bounds, but neither are required.If the absolute value of the step is larger than or equal to the distancebetween the two bounds, then the bounds will be the only valid values. e.g. ifthe range is defined as {from_value: 1, to_value: 2, step: 5} then the validvalues will be 1 and 2.If the step is less than the distance between the bounds, but the distance isnot a multiple of the step, then the "to" bound will always be a valid value,e.g. if the range is defined as {from_value: 2, to_value: 5, step: 2} thenthe valid values will be 2, 4, and 5.
}

// NewIntegerRange creates a new IntegerRange with default values.
func NewIntegerRange() *IntegerRange {
	self := IntegerRange{}
	self.SetDefaults()
	return &self
}

func (t *IntegerRange) Clone() *IntegerRange {
	c := &IntegerRange{}
	c.FromValue = t.FromValue
	c.ToValue = t.ToValue
	c.Step = t.Step
	return c
}

func (t *IntegerRange) CloneMsg() types.Message {
	return t.Clone()
}

func (t *IntegerRange) SetDefaults() {
	t.FromValue = 0
	t.ToValue = 0
	t.Step = 0
}

// CloneIntegerRangeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneIntegerRangeSlice(dst, src []IntegerRange) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var IntegerRangeTypeSupport types.MessageTypeSupport = _IntegerRangeTypeSupport{}

type _IntegerRangeTypeSupport struct{}

func (t _IntegerRangeTypeSupport) New() types.Message {
	return NewIntegerRange()
}

func (t _IntegerRangeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__msg__IntegerRange
	return (unsafe.Pointer)(C.rcl_interfaces__msg__IntegerRange__create())
}

func (t _IntegerRangeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__msg__IntegerRange__destroy((*C.rcl_interfaces__msg__IntegerRange)(pointer_to_free))
}

func (t _IntegerRangeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*IntegerRange)
	mem := (*C.rcl_interfaces__msg__IntegerRange)(dst)
	mem.from_value = C.int64_t(m.FromValue)
	mem.to_value = C.int64_t(m.ToValue)
	mem.step = C.uint64_t(m.Step)
}

func (t _IntegerRangeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*IntegerRange)
	mem := (*C.rcl_interfaces__msg__IntegerRange)(ros2_message_buffer)
	m.FromValue = int64(mem.from_value)
	m.ToValue = int64(mem.to_value)
	m.Step = uint64(mem.step)
}

func (t _IntegerRangeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__msg__IntegerRange())
}

type CIntegerRange = C.rcl_interfaces__msg__IntegerRange
type CIntegerRange__Sequence = C.rcl_interfaces__msg__IntegerRange__Sequence

func IntegerRange__Sequence_to_Go(goSlice *[]IntegerRange, cSlice CIntegerRange__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]IntegerRange, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__msg__IntegerRange__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__IntegerRange * uintptr(i)),
		))
		IntegerRangeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func IntegerRange__Sequence_to_C(cSlice *CIntegerRange__Sequence, goSlice []IntegerRange) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__msg__IntegerRange)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__msg__IntegerRange * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__msg__IntegerRange)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__msg__IntegerRange * uintptr(i)),
		))
		IntegerRangeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func IntegerRange__Array_to_Go(goSlice []IntegerRange, cSlice []CIntegerRange) {
	for i := 0; i < len(cSlice); i++ {
		IntegerRangeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func IntegerRange__Array_to_C(cSlice []CIntegerRange, goSlice []IntegerRange) {
	for i := 0; i < len(goSlice); i++ {
		IntegerRangeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
