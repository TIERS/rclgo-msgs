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
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/trajectory_setpoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/TrajectorySetpoint", TrajectorySetpointTypeSupport)
}

// Do not create instances of this type directly. Always use NewTrajectorySetpoint
// function instead.
type TrajectorySetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	X float32 `yaml:"x"`// in meters NED
	Y float32 `yaml:"y"`// in meters NED
	Z float32 `yaml:"z"`// in meters NED
	Yaw float32 `yaml:"yaw"`// in radians NED -PI..+PI
	Yawspeed float32 `yaml:"yawspeed"`// in radians/sec
	Vx float32 `yaml:"vx"`// in meters/sec
	Vy float32 `yaml:"vy"`// in meters/sec
	Vz float32 `yaml:"vz"`// in meters/sec
	Acceleration [3]float32 `yaml:"acceleration"`// in meters/sec^2
	Jerk [3]float32 `yaml:"jerk"`// in meters/sec^3
	Thrust [3]float32 `yaml:"thrust"`// normalized thrust vector in NED
}

// NewTrajectorySetpoint creates a new TrajectorySetpoint with default values.
func NewTrajectorySetpoint() *TrajectorySetpoint {
	self := TrajectorySetpoint{}
	self.SetDefaults()
	return &self
}

func (t *TrajectorySetpoint) Clone() *TrajectorySetpoint {
	c := &TrajectorySetpoint{}
	c.Timestamp = t.Timestamp
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.Yaw = t.Yaw
	c.Yawspeed = t.Yawspeed
	c.Vx = t.Vx
	c.Vy = t.Vy
	c.Vz = t.Vz
	c.Acceleration = t.Acceleration
	c.Jerk = t.Jerk
	c.Thrust = t.Thrust
	return c
}

func (t *TrajectorySetpoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TrajectorySetpoint) SetDefaults() {
	t.Timestamp = 0
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.Yaw = 0
	t.Yawspeed = 0
	t.Vx = 0
	t.Vy = 0
	t.Vz = 0
	t.Acceleration = [3]float32{}
	t.Jerk = [3]float32{}
	t.Thrust = [3]float32{}
}

// CloneTrajectorySetpointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTrajectorySetpointSlice(dst, src []TrajectorySetpoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TrajectorySetpointTypeSupport types.MessageTypeSupport = _TrajectorySetpointTypeSupport{}

type _TrajectorySetpointTypeSupport struct{}

func (t _TrajectorySetpointTypeSupport) New() types.Message {
	return NewTrajectorySetpoint()
}

func (t _TrajectorySetpointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TrajectorySetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__TrajectorySetpoint__create())
}

func (t _TrajectorySetpointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TrajectorySetpoint__destroy((*C.px4_msgs__msg__TrajectorySetpoint)(pointer_to_free))
}

func (t _TrajectorySetpointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TrajectorySetpoint)
	mem := (*C.px4_msgs__msg__TrajectorySetpoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	mem.yaw = C.float(m.Yaw)
	mem.yawspeed = C.float(m.Yawspeed)
	mem.vx = C.float(m.Vx)
	mem.vy = C.float(m.Vy)
	mem.vz = C.float(m.Vz)
	cSlice_acceleration := mem.acceleration[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_acceleration)), m.Acceleration[:])
	cSlice_jerk := mem.jerk[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_jerk)), m.Jerk[:])
	cSlice_thrust := mem.thrust[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_thrust)), m.Thrust[:])
}

func (t _TrajectorySetpointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TrajectorySetpoint)
	mem := (*C.px4_msgs__msg__TrajectorySetpoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	m.Yaw = float32(mem.yaw)
	m.Yawspeed = float32(mem.yawspeed)
	m.Vx = float32(mem.vx)
	m.Vy = float32(mem.vy)
	m.Vz = float32(mem.vz)
	cSlice_acceleration := mem.acceleration[:]
	primitives.Float32__Array_to_Go(m.Acceleration[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_acceleration)))
	cSlice_jerk := mem.jerk[:]
	primitives.Float32__Array_to_Go(m.Jerk[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_jerk)))
	cSlice_thrust := mem.thrust[:]
	primitives.Float32__Array_to_Go(m.Thrust[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_thrust)))
}

func (t _TrajectorySetpointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TrajectorySetpoint())
}

type CTrajectorySetpoint = C.px4_msgs__msg__TrajectorySetpoint
type CTrajectorySetpoint__Sequence = C.px4_msgs__msg__TrajectorySetpoint__Sequence

func TrajectorySetpoint__Sequence_to_Go(goSlice *[]TrajectorySetpoint, cSlice CTrajectorySetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TrajectorySetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TrajectorySetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(i)),
		))
		TrajectorySetpointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TrajectorySetpoint__Sequence_to_C(cSlice *CTrajectorySetpoint__Sequence, goSlice []TrajectorySetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TrajectorySetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TrajectorySetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TrajectorySetpoint * uintptr(i)),
		))
		TrajectorySetpointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TrajectorySetpoint__Array_to_Go(goSlice []TrajectorySetpoint, cSlice []CTrajectorySetpoint) {
	for i := 0; i < len(cSlice); i++ {
		TrajectorySetpointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TrajectorySetpoint__Array_to_C(cSlice []CTrajectorySetpoint, goSlice []TrajectorySetpoint) {
	for i := 0; i < len(goSlice); i++ {
		TrajectorySetpointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
