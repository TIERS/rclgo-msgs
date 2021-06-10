/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package trajectory_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltrajectory_msgs__rosidl_typesupport_c -ltrajectory_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <trajectory_msgs/msg/multi_dof_joint_trajectory.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("trajectory_msgs/MultiDOFJointTrajectory", MultiDOFJointTrajectoryTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultiDOFJointTrajectory
// function instead.
type MultiDOFJointTrajectory struct {
	Header std_msgs_msg.Header `yaml:"header"`// The header is used to specify the coordinate frame and the reference time for the trajectory durations
	JointNames []string `yaml:"joint_names"`
	Points []MultiDOFJointTrajectoryPoint `yaml:"points"`
}

// NewMultiDOFJointTrajectory creates a new MultiDOFJointTrajectory with default values.
func NewMultiDOFJointTrajectory() *MultiDOFJointTrajectory {
	self := MultiDOFJointTrajectory{}
	self.SetDefaults()
	return &self
}

func (t *MultiDOFJointTrajectory) Clone() *MultiDOFJointTrajectory {
	c := &MultiDOFJointTrajectory{}
	c.Header = *t.Header.Clone()
	if t.JointNames != nil {
		c.JointNames = make([]string, len(t.JointNames))
		copy(c.JointNames, t.JointNames)
	}
	if t.Points != nil {
		c.Points = make([]MultiDOFJointTrajectoryPoint, len(t.Points))
		CloneMultiDOFJointTrajectoryPointSlice(c.Points, t.Points)
	}
	return c
}

func (t *MultiDOFJointTrajectory) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiDOFJointTrajectory) SetDefaults() {
	t.Header.SetDefaults()
	t.JointNames = nil
	t.Points = nil
}

// CloneMultiDOFJointTrajectorySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiDOFJointTrajectorySlice(dst, src []MultiDOFJointTrajectory) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiDOFJointTrajectoryTypeSupport types.MessageTypeSupport = _MultiDOFJointTrajectoryTypeSupport{}

type _MultiDOFJointTrajectoryTypeSupport struct{}

func (t _MultiDOFJointTrajectoryTypeSupport) New() types.Message {
	return NewMultiDOFJointTrajectory()
}

func (t _MultiDOFJointTrajectoryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.trajectory_msgs__msg__MultiDOFJointTrajectory
	return (unsafe.Pointer)(C.trajectory_msgs__msg__MultiDOFJointTrajectory__create())
}

func (t _MultiDOFJointTrajectoryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.trajectory_msgs__msg__MultiDOFJointTrajectory__destroy((*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(pointer_to_free))
}

func (t _MultiDOFJointTrajectoryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiDOFJointTrajectory)
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.joint_names)), m.JointNames)
	MultiDOFJointTrajectoryPoint__Sequence_to_C(&mem.points, m.Points)
}

func (t _MultiDOFJointTrajectoryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiDOFJointTrajectory)
	mem := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.String__Sequence_to_Go(&m.JointNames, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.joint_names)))
	MultiDOFJointTrajectoryPoint__Sequence_to_Go(&m.Points, mem.points)
}

func (t _MultiDOFJointTrajectoryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__trajectory_msgs__msg__MultiDOFJointTrajectory())
}

type CMultiDOFJointTrajectory = C.trajectory_msgs__msg__MultiDOFJointTrajectory
type CMultiDOFJointTrajectory__Sequence = C.trajectory_msgs__msg__MultiDOFJointTrajectory__Sequence

func MultiDOFJointTrajectory__Sequence_to_Go(goSlice *[]MultiDOFJointTrajectory, cSlice CMultiDOFJointTrajectory__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiDOFJointTrajectory, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(i)),
		))
		MultiDOFJointTrajectoryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MultiDOFJointTrajectory__Sequence_to_C(cSlice *CMultiDOFJointTrajectory__Sequence, goSlice []MultiDOFJointTrajectory) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(C.malloc((C.size_t)(C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.trajectory_msgs__msg__MultiDOFJointTrajectory)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_trajectory_msgs__msg__MultiDOFJointTrajectory * uintptr(i)),
		))
		MultiDOFJointTrajectoryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MultiDOFJointTrajectory__Array_to_Go(goSlice []MultiDOFJointTrajectory, cSlice []CMultiDOFJointTrajectory) {
	for i := 0; i < len(cSlice); i++ {
		MultiDOFJointTrajectoryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiDOFJointTrajectory__Array_to_C(cSlice []CMultiDOFJointTrajectory, goSlice []MultiDOFJointTrajectory) {
	for i := 0; i < len(goSlice); i++ {
		MultiDOFJointTrajectoryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
