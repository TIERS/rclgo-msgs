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

#include <px4_msgs/msg/position_controller_landing_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/PositionControllerLandingStatus", PositionControllerLandingStatusTypeSupport)
}

// Do not create instances of this type directly. Always use NewPositionControllerLandingStatus
// function instead.
type PositionControllerLandingStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	HorizontalSlopeDisplacement float32 `yaml:"horizontal_slope_displacement"`
	SlopeAngleRad float32 `yaml:"slope_angle_rad"`
	FlareLength float32 `yaml:"flare_length"`
	AbortLanding bool `yaml:"abort_landing"`// true if landing should be aborted
}

// NewPositionControllerLandingStatus creates a new PositionControllerLandingStatus with default values.
func NewPositionControllerLandingStatus() *PositionControllerLandingStatus {
	self := PositionControllerLandingStatus{}
	self.SetDefaults()
	return &self
}

func (t *PositionControllerLandingStatus) Clone() *PositionControllerLandingStatus {
	c := &PositionControllerLandingStatus{}
	c.Timestamp = t.Timestamp
	c.HorizontalSlopeDisplacement = t.HorizontalSlopeDisplacement
	c.SlopeAngleRad = t.SlopeAngleRad
	c.FlareLength = t.FlareLength
	c.AbortLanding = t.AbortLanding
	return c
}

func (t *PositionControllerLandingStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PositionControllerLandingStatus) SetDefaults() {
	t.Timestamp = 0
	t.HorizontalSlopeDisplacement = 0
	t.SlopeAngleRad = 0
	t.FlareLength = 0
	t.AbortLanding = false
}

// ClonePositionControllerLandingStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePositionControllerLandingStatusSlice(dst, src []PositionControllerLandingStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PositionControllerLandingStatusTypeSupport types.MessageTypeSupport = _PositionControllerLandingStatusTypeSupport{}

type _PositionControllerLandingStatusTypeSupport struct{}

func (t _PositionControllerLandingStatusTypeSupport) New() types.Message {
	return NewPositionControllerLandingStatus()
}

func (t _PositionControllerLandingStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__PositionControllerLandingStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__PositionControllerLandingStatus__create())
}

func (t _PositionControllerLandingStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__PositionControllerLandingStatus__destroy((*C.px4_msgs__msg__PositionControllerLandingStatus)(pointer_to_free))
}

func (t _PositionControllerLandingStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PositionControllerLandingStatus)
	mem := (*C.px4_msgs__msg__PositionControllerLandingStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.horizontal_slope_displacement = C.float(m.HorizontalSlopeDisplacement)
	mem.slope_angle_rad = C.float(m.SlopeAngleRad)
	mem.flare_length = C.float(m.FlareLength)
	mem.abort_landing = C.bool(m.AbortLanding)
}

func (t _PositionControllerLandingStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PositionControllerLandingStatus)
	mem := (*C.px4_msgs__msg__PositionControllerLandingStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.HorizontalSlopeDisplacement = float32(mem.horizontal_slope_displacement)
	m.SlopeAngleRad = float32(mem.slope_angle_rad)
	m.FlareLength = float32(mem.flare_length)
	m.AbortLanding = bool(mem.abort_landing)
}

func (t _PositionControllerLandingStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__PositionControllerLandingStatus())
}

type CPositionControllerLandingStatus = C.px4_msgs__msg__PositionControllerLandingStatus
type CPositionControllerLandingStatus__Sequence = C.px4_msgs__msg__PositionControllerLandingStatus__Sequence

func PositionControllerLandingStatus__Sequence_to_Go(goSlice *[]PositionControllerLandingStatus, cSlice CPositionControllerLandingStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PositionControllerLandingStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__PositionControllerLandingStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionControllerLandingStatus * uintptr(i)),
		))
		PositionControllerLandingStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PositionControllerLandingStatus__Sequence_to_C(cSlice *CPositionControllerLandingStatus__Sequence, goSlice []PositionControllerLandingStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__PositionControllerLandingStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__PositionControllerLandingStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__PositionControllerLandingStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionControllerLandingStatus * uintptr(i)),
		))
		PositionControllerLandingStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PositionControllerLandingStatus__Array_to_Go(goSlice []PositionControllerLandingStatus, cSlice []CPositionControllerLandingStatus) {
	for i := 0; i < len(cSlice); i++ {
		PositionControllerLandingStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PositionControllerLandingStatus__Array_to_C(cSlice []CPositionControllerLandingStatus, goSlice []PositionControllerLandingStatus) {
	for i := 0; i < len(goSlice); i++ {
		PositionControllerLandingStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
