/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/multirotor_motor_limits.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/MultirotorMotorLimits", MultirotorMotorLimitsTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultirotorMotorLimits
// function instead.
type MultirotorMotorLimits struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SaturationStatus uint16 `yaml:"saturation_status"`// Integer bit mask indicating which axes in the control mixer are saturated
}

// NewMultirotorMotorLimits creates a new MultirotorMotorLimits with default values.
func NewMultirotorMotorLimits() *MultirotorMotorLimits {
	self := MultirotorMotorLimits{}
	self.SetDefaults()
	return &self
}

func (t *MultirotorMotorLimits) Clone() *MultirotorMotorLimits {
	c := &MultirotorMotorLimits{}
	c.Timestamp = t.Timestamp
	c.SaturationStatus = t.SaturationStatus
	return c
}

func (t *MultirotorMotorLimits) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultirotorMotorLimits) SetDefaults() {
	t.Timestamp = 0
	t.SaturationStatus = 0
}

// CloneMultirotorMotorLimitsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultirotorMotorLimitsSlice(dst, src []MultirotorMotorLimits) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultirotorMotorLimitsTypeSupport types.MessageTypeSupport = _MultirotorMotorLimitsTypeSupport{}

type _MultirotorMotorLimitsTypeSupport struct{}

func (t _MultirotorMotorLimitsTypeSupport) New() types.Message {
	return NewMultirotorMotorLimits()
}

func (t _MultirotorMotorLimitsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__MultirotorMotorLimits
	return (unsafe.Pointer)(C.px4_msgs__msg__MultirotorMotorLimits__create())
}

func (t _MultirotorMotorLimitsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__MultirotorMotorLimits__destroy((*C.px4_msgs__msg__MultirotorMotorLimits)(pointer_to_free))
}

func (t _MultirotorMotorLimitsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultirotorMotorLimits)
	mem := (*C.px4_msgs__msg__MultirotorMotorLimits)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.saturation_status = C.uint16_t(m.SaturationStatus)
}

func (t _MultirotorMotorLimitsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultirotorMotorLimits)
	mem := (*C.px4_msgs__msg__MultirotorMotorLimits)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.SaturationStatus = uint16(mem.saturation_status)
}

func (t _MultirotorMotorLimitsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__MultirotorMotorLimits())
}

type CMultirotorMotorLimits = C.px4_msgs__msg__MultirotorMotorLimits
type CMultirotorMotorLimits__Sequence = C.px4_msgs__msg__MultirotorMotorLimits__Sequence

func MultirotorMotorLimits__Sequence_to_Go(goSlice *[]MultirotorMotorLimits, cSlice CMultirotorMotorLimits__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultirotorMotorLimits, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__MultirotorMotorLimits__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MultirotorMotorLimits * uintptr(i)),
		))
		MultirotorMotorLimitsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MultirotorMotorLimits__Sequence_to_C(cSlice *CMultirotorMotorLimits__Sequence, goSlice []MultirotorMotorLimits) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__MultirotorMotorLimits)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__MultirotorMotorLimits * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__MultirotorMotorLimits)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MultirotorMotorLimits * uintptr(i)),
		))
		MultirotorMotorLimitsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MultirotorMotorLimits__Array_to_Go(goSlice []MultirotorMotorLimits, cSlice []CMultirotorMotorLimits) {
	for i := 0; i < len(cSlice); i++ {
		MultirotorMotorLimitsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultirotorMotorLimits__Array_to_C(cSlice []CMultirotorMotorLimits, goSlice []MultirotorMotorLimits) {
	for i := 0; i < len(goSlice); i++ {
		MultirotorMotorLimitsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
