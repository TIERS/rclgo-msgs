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

#include <px4_msgs/msg/vehicle_trajectory_bezier.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleTrajectoryBezier", VehicleTrajectoryBezierTypeSupport)
}
const (
	VehicleTrajectoryBezier_POINT_0 uint8 = 0
	VehicleTrajectoryBezier_POINT_1 uint8 = 1
	VehicleTrajectoryBezier_POINT_2 uint8 = 2
	VehicleTrajectoryBezier_POINT_3 uint8 = 3
	VehicleTrajectoryBezier_POINT_4 uint8 = 4
	VehicleTrajectoryBezier_NUMBER_POINTS uint8 = 5
)

// Do not create instances of this type directly. Always use NewVehicleTrajectoryBezier
// function instead.
type VehicleTrajectoryBezier struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ControlPoints [5]TrajectoryBezier `yaml:"control_points"`
	BezierOrder uint8 `yaml:"bezier_order"`
}

// NewVehicleTrajectoryBezier creates a new VehicleTrajectoryBezier with default values.
func NewVehicleTrajectoryBezier() *VehicleTrajectoryBezier {
	self := VehicleTrajectoryBezier{}
	self.SetDefaults()
	return &self
}

func (t *VehicleTrajectoryBezier) Clone() *VehicleTrajectoryBezier {
	c := &VehicleTrajectoryBezier{}
	c.Timestamp = t.Timestamp
	CloneTrajectoryBezierSlice(c.ControlPoints[:], t.ControlPoints[:])
	c.BezierOrder = t.BezierOrder
	return c
}

func (t *VehicleTrajectoryBezier) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleTrajectoryBezier) SetDefaults() {
	t.Timestamp = 0
	for i := range t.ControlPoints {
		t.ControlPoints[i].SetDefaults()
	}
	t.BezierOrder = 0
}

// CloneVehicleTrajectoryBezierSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleTrajectoryBezierSlice(dst, src []VehicleTrajectoryBezier) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleTrajectoryBezierTypeSupport types.MessageTypeSupport = _VehicleTrajectoryBezierTypeSupport{}

type _VehicleTrajectoryBezierTypeSupport struct{}

func (t _VehicleTrajectoryBezierTypeSupport) New() types.Message {
	return NewVehicleTrajectoryBezier()
}

func (t _VehicleTrajectoryBezierTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleTrajectoryBezier
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleTrajectoryBezier__create())
}

func (t _VehicleTrajectoryBezierTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleTrajectoryBezier__destroy((*C.px4_msgs__msg__VehicleTrajectoryBezier)(pointer_to_free))
}

func (t _VehicleTrajectoryBezierTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleTrajectoryBezier)
	mem := (*C.px4_msgs__msg__VehicleTrajectoryBezier)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	TrajectoryBezier__Array_to_C(mem.control_points[:], m.ControlPoints[:])
	mem.bezier_order = C.uint8_t(m.BezierOrder)
}

func (t _VehicleTrajectoryBezierTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleTrajectoryBezier)
	mem := (*C.px4_msgs__msg__VehicleTrajectoryBezier)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	TrajectoryBezier__Array_to_Go(m.ControlPoints[:], mem.control_points[:])
	m.BezierOrder = uint8(mem.bezier_order)
}

func (t _VehicleTrajectoryBezierTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleTrajectoryBezier())
}

type CVehicleTrajectoryBezier = C.px4_msgs__msg__VehicleTrajectoryBezier
type CVehicleTrajectoryBezier__Sequence = C.px4_msgs__msg__VehicleTrajectoryBezier__Sequence

func VehicleTrajectoryBezier__Sequence_to_Go(goSlice *[]VehicleTrajectoryBezier, cSlice CVehicleTrajectoryBezier__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleTrajectoryBezier, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryBezier__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryBezier * uintptr(i)),
		))
		VehicleTrajectoryBezierTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleTrajectoryBezier__Sequence_to_C(cSlice *CVehicleTrajectoryBezier__Sequence, goSlice []VehicleTrajectoryBezier) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleTrajectoryBezier)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryBezier * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleTrajectoryBezier)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleTrajectoryBezier * uintptr(i)),
		))
		VehicleTrajectoryBezierTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleTrajectoryBezier__Array_to_Go(goSlice []VehicleTrajectoryBezier, cSlice []CVehicleTrajectoryBezier) {
	for i := 0; i < len(cSlice); i++ {
		VehicleTrajectoryBezierTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleTrajectoryBezier__Array_to_C(cSlice []CVehicleTrajectoryBezier, goSlice []VehicleTrajectoryBezier) {
	for i := 0; i < len(goSlice); i++ {
		VehicleTrajectoryBezierTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
