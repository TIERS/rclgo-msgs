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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpendulum_msgs__rosidl_typesupport_c -lpendulum_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pendulum_msgs/msg/joint_command.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pendulum_msgs/JointCommand", JointCommandTypeSupport)
}

// Do not create instances of this type directly. Always use NewJointCommand
// function instead.
type JointCommand struct {
	Position float64 `yaml:"position"`
}

// NewJointCommand creates a new JointCommand with default values.
func NewJointCommand() *JointCommand {
	self := JointCommand{}
	self.SetDefaults()
	return &self
}

func (t *JointCommand) Clone() *JointCommand {
	c := &JointCommand{}
	c.Position = t.Position
	return c
}

func (t *JointCommand) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JointCommand) SetDefaults() {
	t.Position = 0
}

// CloneJointCommandSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJointCommandSlice(dst, src []JointCommand) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JointCommandTypeSupport types.MessageTypeSupport = _JointCommandTypeSupport{}

type _JointCommandTypeSupport struct{}

func (t _JointCommandTypeSupport) New() types.Message {
	return NewJointCommand()
}

func (t _JointCommandTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pendulum_msgs__msg__JointCommand
	return (unsafe.Pointer)(C.pendulum_msgs__msg__JointCommand__create())
}

func (t _JointCommandTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pendulum_msgs__msg__JointCommand__destroy((*C.pendulum_msgs__msg__JointCommand)(pointer_to_free))
}

func (t _JointCommandTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JointCommand)
	mem := (*C.pendulum_msgs__msg__JointCommand)(dst)
	mem.position = C.double(m.Position)
}

func (t _JointCommandTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JointCommand)
	mem := (*C.pendulum_msgs__msg__JointCommand)(ros2_message_buffer)
	m.Position = float64(mem.position)
}

func (t _JointCommandTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pendulum_msgs__msg__JointCommand())
}

type CJointCommand = C.pendulum_msgs__msg__JointCommand
type CJointCommand__Sequence = C.pendulum_msgs__msg__JointCommand__Sequence

func JointCommand__Sequence_to_Go(goSlice *[]JointCommand, cSlice CJointCommand__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JointCommand, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pendulum_msgs__msg__JointCommand__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__JointCommand * uintptr(i)),
		))
		JointCommandTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func JointCommand__Sequence_to_C(cSlice *CJointCommand__Sequence, goSlice []JointCommand) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pendulum_msgs__msg__JointCommand)(C.malloc((C.size_t)(C.sizeof_struct_pendulum_msgs__msg__JointCommand * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pendulum_msgs__msg__JointCommand)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__JointCommand * uintptr(i)),
		))
		JointCommandTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func JointCommand__Array_to_Go(goSlice []JointCommand, cSlice []CJointCommand) {
	for i := 0; i < len(cSlice); i++ {
		JointCommandTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JointCommand__Array_to_C(cSlice []CJointCommand, goSlice []JointCommand) {
	for i := 0; i < len(goSlice); i++ {
		JointCommandTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
