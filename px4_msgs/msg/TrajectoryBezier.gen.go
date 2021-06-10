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

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/trajectory_bezier.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/TrajectoryBezier", TrajectoryBezierTypeSupport)
}

// Do not create instances of this type directly. Always use NewTrajectoryBezier
// function instead.
type TrajectoryBezier struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Position [3]float32 `yaml:"position"`// local position x,y,z (metres)
	Yaw float32 `yaml:"yaw"`// yaw angle (rad)
	Delta float32 `yaml:"delta"`// time it should take to get to this waypoint, if this is the final waypoint (seconds)
}

// NewTrajectoryBezier creates a new TrajectoryBezier with default values.
func NewTrajectoryBezier() *TrajectoryBezier {
	self := TrajectoryBezier{}
	self.SetDefaults()
	return &self
}

func (t *TrajectoryBezier) Clone() *TrajectoryBezier {
	c := &TrajectoryBezier{}
	c.Timestamp = t.Timestamp
	c.Position = t.Position
	c.Yaw = t.Yaw
	c.Delta = t.Delta
	return c
}

func (t *TrajectoryBezier) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TrajectoryBezier) SetDefaults() {
	t.Timestamp = 0
	t.Position = [3]float32{}
	t.Yaw = 0
	t.Delta = 0
}

// CloneTrajectoryBezierSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTrajectoryBezierSlice(dst, src []TrajectoryBezier) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TrajectoryBezierTypeSupport types.MessageTypeSupport = _TrajectoryBezierTypeSupport{}

type _TrajectoryBezierTypeSupport struct{}

func (t _TrajectoryBezierTypeSupport) New() types.Message {
	return NewTrajectoryBezier()
}

func (t _TrajectoryBezierTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TrajectoryBezier
	return (unsafe.Pointer)(C.px4_msgs__msg__TrajectoryBezier__create())
}

func (t _TrajectoryBezierTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TrajectoryBezier__destroy((*C.px4_msgs__msg__TrajectoryBezier)(pointer_to_free))
}

func (t _TrajectoryBezierTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TrajectoryBezier)
	mem := (*C.px4_msgs__msg__TrajectoryBezier)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_position := mem.position[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_position)), m.Position[:])
	mem.yaw = C.float(m.Yaw)
	mem.delta = C.float(m.Delta)
}

func (t _TrajectoryBezierTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TrajectoryBezier)
	mem := (*C.px4_msgs__msg__TrajectoryBezier)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_position := mem.position[:]
	primitives.Float32__Array_to_Go(m.Position[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_position)))
	m.Yaw = float32(mem.yaw)
	m.Delta = float32(mem.delta)
}

func (t _TrajectoryBezierTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TrajectoryBezier())
}

type CTrajectoryBezier = C.px4_msgs__msg__TrajectoryBezier
type CTrajectoryBezier__Sequence = C.px4_msgs__msg__TrajectoryBezier__Sequence

func TrajectoryBezier__Sequence_to_Go(goSlice *[]TrajectoryBezier, cSlice CTrajectoryBezier__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TrajectoryBezier, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TrajectoryBezier__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectoryBezier * uintptr(i)),
		))
		TrajectoryBezierTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TrajectoryBezier__Sequence_to_C(cSlice *CTrajectoryBezier__Sequence, goSlice []TrajectoryBezier) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TrajectoryBezier)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TrajectoryBezier * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TrajectoryBezier)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectoryBezier * uintptr(i)),
		))
		TrajectoryBezierTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TrajectoryBezier__Array_to_Go(goSlice []TrajectoryBezier, cSlice []CTrajectoryBezier) {
	for i := 0; i < len(cSlice); i++ {
		TrajectoryBezierTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TrajectoryBezier__Array_to_C(cSlice []CTrajectoryBezier, goSlice []TrajectoryBezier) {
	for i := 0; i < len(goSlice); i++ {
		TrajectoryBezierTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
