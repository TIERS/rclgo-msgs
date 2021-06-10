/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pendulum_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpendulum_msgs__rosidl_typesupport_c -lpendulum_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pendulum_msgs/msg/joint_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pendulum_msgs/JointState", JointStateTypeSupport)
}

// Do not create instances of this type directly. Always use NewJointState
// function instead.
type JointState struct {
	Position float64 `yaml:"position"`
	Velocity float64 `yaml:"velocity"`
	Effort float64 `yaml:"effort"`
}

// NewJointState creates a new JointState with default values.
func NewJointState() *JointState {
	self := JointState{}
	self.SetDefaults()
	return &self
}

func (t *JointState) Clone() *JointState {
	c := &JointState{}
	c.Position = t.Position
	c.Velocity = t.Velocity
	c.Effort = t.Effort
	return c
}

func (t *JointState) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JointState) SetDefaults() {
	t.Position = 0
	t.Velocity = 0
	t.Effort = 0
}

// CloneJointStateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJointStateSlice(dst, src []JointState) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JointStateTypeSupport types.MessageTypeSupport = _JointStateTypeSupport{}

type _JointStateTypeSupport struct{}

func (t _JointStateTypeSupport) New() types.Message {
	return NewJointState()
}

func (t _JointStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pendulum_msgs__msg__JointState
	return (unsafe.Pointer)(C.pendulum_msgs__msg__JointState__create())
}

func (t _JointStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pendulum_msgs__msg__JointState__destroy((*C.pendulum_msgs__msg__JointState)(pointer_to_free))
}

func (t _JointStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JointState)
	mem := (*C.pendulum_msgs__msg__JointState)(dst)
	mem.position = C.double(m.Position)
	mem.velocity = C.double(m.Velocity)
	mem.effort = C.double(m.Effort)
}

func (t _JointStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JointState)
	mem := (*C.pendulum_msgs__msg__JointState)(ros2_message_buffer)
	m.Position = float64(mem.position)
	m.Velocity = float64(mem.velocity)
	m.Effort = float64(mem.effort)
}

func (t _JointStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pendulum_msgs__msg__JointState())
}

type CJointState = C.pendulum_msgs__msg__JointState
type CJointState__Sequence = C.pendulum_msgs__msg__JointState__Sequence

func JointState__Sequence_to_Go(goSlice *[]JointState, cSlice CJointState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JointState, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pendulum_msgs__msg__JointState__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__JointState * uintptr(i)),
		))
		JointStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func JointState__Sequence_to_C(cSlice *CJointState__Sequence, goSlice []JointState) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pendulum_msgs__msg__JointState)(C.malloc((C.size_t)(C.sizeof_struct_pendulum_msgs__msg__JointState * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pendulum_msgs__msg__JointState)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__JointState * uintptr(i)),
		))
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func JointState__Array_to_Go(goSlice []JointState, cSlice []CJointState) {
	for i := 0; i < len(cSlice); i++ {
		JointStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JointState__Array_to_C(cSlice []CJointState, goSlice []JointState) {
	for i := 0; i < len(goSlice); i++ {
		JointStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
