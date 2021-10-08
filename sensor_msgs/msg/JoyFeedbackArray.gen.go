/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joy_feedback_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/JoyFeedbackArray", JoyFeedbackArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewJoyFeedbackArray
// function instead.
type JoyFeedbackArray struct {
	Array []JoyFeedback `yaml:"array"`// This message publishes values for multiple feedback at once.
}

// NewJoyFeedbackArray creates a new JoyFeedbackArray with default values.
func NewJoyFeedbackArray() *JoyFeedbackArray {
	self := JoyFeedbackArray{}
	self.SetDefaults()
	return &self
}

func (t *JoyFeedbackArray) Clone() *JoyFeedbackArray {
	c := &JoyFeedbackArray{}
	if t.Array != nil {
		c.Array = make([]JoyFeedback, len(t.Array))
		CloneJoyFeedbackSlice(c.Array, t.Array)
	}
	return c
}

func (t *JoyFeedbackArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JoyFeedbackArray) SetDefaults() {
	t.Array = nil
}

// CloneJoyFeedbackArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJoyFeedbackArraySlice(dst, src []JoyFeedbackArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JoyFeedbackArrayTypeSupport types.MessageTypeSupport = _JoyFeedbackArrayTypeSupport{}

type _JoyFeedbackArrayTypeSupport struct{}

func (t _JoyFeedbackArrayTypeSupport) New() types.Message {
	return NewJoyFeedbackArray()
}

func (t _JoyFeedbackArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__JoyFeedbackArray
	return (unsafe.Pointer)(C.sensor_msgs__msg__JoyFeedbackArray__create())
}

func (t _JoyFeedbackArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__JoyFeedbackArray__destroy((*C.sensor_msgs__msg__JoyFeedbackArray)(pointer_to_free))
}

func (t _JoyFeedbackArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JoyFeedbackArray)
	mem := (*C.sensor_msgs__msg__JoyFeedbackArray)(dst)
	JoyFeedback__Sequence_to_C(&mem.array, m.Array)
}

func (t _JoyFeedbackArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JoyFeedbackArray)
	mem := (*C.sensor_msgs__msg__JoyFeedbackArray)(ros2_message_buffer)
	JoyFeedback__Sequence_to_Go(&m.Array, mem.array)
}

func (t _JoyFeedbackArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__JoyFeedbackArray())
}

type CJoyFeedbackArray = C.sensor_msgs__msg__JoyFeedbackArray
type CJoyFeedbackArray__Sequence = C.sensor_msgs__msg__JoyFeedbackArray__Sequence

func JoyFeedbackArray__Sequence_to_Go(goSlice *[]JoyFeedbackArray, cSlice CJoyFeedbackArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JoyFeedbackArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__JoyFeedbackArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JoyFeedbackArray * uintptr(i)),
		))
		JoyFeedbackArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func JoyFeedbackArray__Sequence_to_C(cSlice *CJoyFeedbackArray__Sequence, goSlice []JoyFeedbackArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__JoyFeedbackArray)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__JoyFeedbackArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__JoyFeedbackArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JoyFeedbackArray * uintptr(i)),
		))
		JoyFeedbackArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func JoyFeedbackArray__Array_to_Go(goSlice []JoyFeedbackArray, cSlice []CJoyFeedbackArray) {
	for i := 0; i < len(cSlice); i++ {
		JoyFeedbackArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JoyFeedbackArray__Array_to_C(cSlice []CJoyFeedbackArray, goSlice []JoyFeedbackArray) {
	for i := 0; i < len(goSlice); i++ {
		JoyFeedbackArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
