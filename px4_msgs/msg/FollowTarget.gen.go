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

#include <px4_msgs/msg/follow_target.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/FollowTarget", FollowTargetTypeSupport)
}

// Do not create instances of this type directly. Always use NewFollowTarget
// function instead.
type FollowTarget struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Lat float64 `yaml:"lat"`// target position (deg * 1e7)
	Lon float64 `yaml:"lon"`// target position (deg * 1e7)
	Alt float32 `yaml:"alt"`// target position
	Vy float32 `yaml:"vy"`// target vel in y
	Vx float32 `yaml:"vx"`// target vel in x
	Vz float32 `yaml:"vz"`// target vel in z
	EstCap uint8 `yaml:"est_cap"`// target reporting capabilities
}

// NewFollowTarget creates a new FollowTarget with default values.
func NewFollowTarget() *FollowTarget {
	self := FollowTarget{}
	self.SetDefaults()
	return &self
}

func (t *FollowTarget) Clone() *FollowTarget {
	c := &FollowTarget{}
	c.Timestamp = t.Timestamp
	c.Lat = t.Lat
	c.Lon = t.Lon
	c.Alt = t.Alt
	c.Vy = t.Vy
	c.Vx = t.Vx
	c.Vz = t.Vz
	c.EstCap = t.EstCap
	return c
}

func (t *FollowTarget) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FollowTarget) SetDefaults() {
	t.Timestamp = 0
	t.Lat = 0
	t.Lon = 0
	t.Alt = 0
	t.Vy = 0
	t.Vx = 0
	t.Vz = 0
	t.EstCap = 0
}

// CloneFollowTargetSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFollowTargetSlice(dst, src []FollowTarget) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FollowTargetTypeSupport types.MessageTypeSupport = _FollowTargetTypeSupport{}

type _FollowTargetTypeSupport struct{}

func (t _FollowTargetTypeSupport) New() types.Message {
	return NewFollowTarget()
}

func (t _FollowTargetTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__FollowTarget
	return (unsafe.Pointer)(C.px4_msgs__msg__FollowTarget__create())
}

func (t _FollowTargetTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__FollowTarget__destroy((*C.px4_msgs__msg__FollowTarget)(pointer_to_free))
}

func (t _FollowTargetTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FollowTarget)
	mem := (*C.px4_msgs__msg__FollowTarget)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.lat = C.double(m.Lat)
	mem.lon = C.double(m.Lon)
	mem.alt = C.float(m.Alt)
	mem.vy = C.float(m.Vy)
	mem.vx = C.float(m.Vx)
	mem.vz = C.float(m.Vz)
	mem.est_cap = C.uint8_t(m.EstCap)
}

func (t _FollowTargetTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FollowTarget)
	mem := (*C.px4_msgs__msg__FollowTarget)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Lat = float64(mem.lat)
	m.Lon = float64(mem.lon)
	m.Alt = float32(mem.alt)
	m.Vy = float32(mem.vy)
	m.Vx = float32(mem.vx)
	m.Vz = float32(mem.vz)
	m.EstCap = uint8(mem.est_cap)
}

func (t _FollowTargetTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__FollowTarget())
}

type CFollowTarget = C.px4_msgs__msg__FollowTarget
type CFollowTarget__Sequence = C.px4_msgs__msg__FollowTarget__Sequence

func FollowTarget__Sequence_to_Go(goSlice *[]FollowTarget, cSlice CFollowTarget__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FollowTarget, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__FollowTarget__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__FollowTarget * uintptr(i)),
		))
		FollowTargetTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func FollowTarget__Sequence_to_C(cSlice *CFollowTarget__Sequence, goSlice []FollowTarget) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__FollowTarget)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__FollowTarget * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__FollowTarget)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__FollowTarget * uintptr(i)),
		))
		FollowTargetTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func FollowTarget__Array_to_Go(goSlice []FollowTarget, cSlice []CFollowTarget) {
	for i := 0; i < len(cSlice); i++ {
		FollowTargetTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FollowTarget__Array_to_C(cSlice []CFollowTarget, goSlice []FollowTarget) {
	for i := 0; i < len(goSlice); i++ {
		FollowTargetTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
