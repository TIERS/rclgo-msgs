/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/msg/pose.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/Pose", PoseTypeSupport)
}

// Do not create instances of this type directly. Always use NewPose
// function instead.
type Pose struct {
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Theta float32 `yaml:"theta"`
	LinearVelocity float32 `yaml:"linear_velocity"`
	AngularVelocity float32 `yaml:"angular_velocity"`
}

// NewPose creates a new Pose with default values.
func NewPose() *Pose {
	self := Pose{}
	self.SetDefaults()
	return &self
}

func (t *Pose) Clone() *Pose {
	c := &Pose{}
	c.X = t.X
	c.Y = t.Y
	c.Theta = t.Theta
	c.LinearVelocity = t.LinearVelocity
	c.AngularVelocity = t.AngularVelocity
	return c
}

func (t *Pose) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Pose) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Theta = 0
	t.LinearVelocity = 0
	t.AngularVelocity = 0
}

// ClonePoseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseSlice(dst, src []Pose) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseTypeSupport types.MessageTypeSupport = _PoseTypeSupport{}

type _PoseTypeSupport struct{}

func (t _PoseTypeSupport) New() types.Message {
	return NewPose()
}

func (t _PoseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__msg__Pose
	return (unsafe.Pointer)(C.turtlesim__msg__Pose__create())
}

func (t _PoseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__msg__Pose__destroy((*C.turtlesim__msg__Pose)(pointer_to_free))
}

func (t _PoseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Pose)
	mem := (*C.turtlesim__msg__Pose)(dst)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.theta = C.float(m.Theta)
	mem.linear_velocity = C.float(m.LinearVelocity)
	mem.angular_velocity = C.float(m.AngularVelocity)
}

func (t _PoseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Pose)
	mem := (*C.turtlesim__msg__Pose)(ros2_message_buffer)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Theta = float32(mem.theta)
	m.LinearVelocity = float32(mem.linear_velocity)
	m.AngularVelocity = float32(mem.angular_velocity)
}

func (t _PoseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__msg__Pose())
}

type CPose = C.turtlesim__msg__Pose
type CPose__Sequence = C.turtlesim__msg__Pose__Sequence

func Pose__Sequence_to_Go(goSlice *[]Pose, cSlice CPose__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Pose, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__msg__Pose__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__msg__Pose * uintptr(i)),
		))
		PoseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Pose__Sequence_to_C(cSlice *CPose__Sequence, goSlice []Pose) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__msg__Pose)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__msg__Pose * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__msg__Pose)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__msg__Pose * uintptr(i)),
		))
		PoseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Pose__Array_to_Go(goSlice []Pose, cSlice []CPose) {
	for i := 0; i < len(cSlice); i++ {
		PoseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Pose__Array_to_C(cSlice []CPose, goSlice []Pose) {
	for i := 0; i < len(goSlice); i++ {
		PoseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
