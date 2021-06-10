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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/camera_trigger_secondary.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/CameraTriggerSecondary", CameraTriggerSecondaryTypeSupport)
}

// Do not create instances of this type directly. Always use NewCameraTriggerSecondary
// function instead.
type CameraTriggerSecondary struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampUtc uint64 `yaml:"timestamp_utc"`// UTC timestamp
	Seq uint32 `yaml:"seq"`// Image sequence number
	Feedback bool `yaml:"feedback"`// Trigger feedback from camera
}

// NewCameraTriggerSecondary creates a new CameraTriggerSecondary with default values.
func NewCameraTriggerSecondary() *CameraTriggerSecondary {
	self := CameraTriggerSecondary{}
	self.SetDefaults()
	return &self
}

func (t *CameraTriggerSecondary) Clone() *CameraTriggerSecondary {
	c := &CameraTriggerSecondary{}
	c.Timestamp = t.Timestamp
	c.TimestampUtc = t.TimestampUtc
	c.Seq = t.Seq
	c.Feedback = t.Feedback
	return c
}

func (t *CameraTriggerSecondary) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CameraTriggerSecondary) SetDefaults() {
	t.Timestamp = 0
	t.TimestampUtc = 0
	t.Seq = 0
	t.Feedback = false
}

// CloneCameraTriggerSecondarySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCameraTriggerSecondarySlice(dst, src []CameraTriggerSecondary) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CameraTriggerSecondaryTypeSupport types.MessageTypeSupport = _CameraTriggerSecondaryTypeSupport{}

type _CameraTriggerSecondaryTypeSupport struct{}

func (t _CameraTriggerSecondaryTypeSupport) New() types.Message {
	return NewCameraTriggerSecondary()
}

func (t _CameraTriggerSecondaryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CameraTriggerSecondary
	return (unsafe.Pointer)(C.px4_msgs__msg__CameraTriggerSecondary__create())
}

func (t _CameraTriggerSecondaryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CameraTriggerSecondary__destroy((*C.px4_msgs__msg__CameraTriggerSecondary)(pointer_to_free))
}

func (t _CameraTriggerSecondaryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CameraTriggerSecondary)
	mem := (*C.px4_msgs__msg__CameraTriggerSecondary)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_utc = C.uint64_t(m.TimestampUtc)
	mem.seq = C.uint32_t(m.Seq)
	mem.feedback = C.bool(m.Feedback)
}

func (t _CameraTriggerSecondaryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CameraTriggerSecondary)
	mem := (*C.px4_msgs__msg__CameraTriggerSecondary)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampUtc = uint64(mem.timestamp_utc)
	m.Seq = uint32(mem.seq)
	m.Feedback = bool(mem.feedback)
}

func (t _CameraTriggerSecondaryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CameraTriggerSecondary())
}

type CCameraTriggerSecondary = C.px4_msgs__msg__CameraTriggerSecondary
type CCameraTriggerSecondary__Sequence = C.px4_msgs__msg__CameraTriggerSecondary__Sequence

func CameraTriggerSecondary__Sequence_to_Go(goSlice *[]CameraTriggerSecondary, cSlice CCameraTriggerSecondary__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CameraTriggerSecondary, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CameraTriggerSecondary__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraTriggerSecondary * uintptr(i)),
		))
		CameraTriggerSecondaryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CameraTriggerSecondary__Sequence_to_C(cSlice *CCameraTriggerSecondary__Sequence, goSlice []CameraTriggerSecondary) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CameraTriggerSecondary)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CameraTriggerSecondary * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CameraTriggerSecondary)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CameraTriggerSecondary * uintptr(i)),
		))
		CameraTriggerSecondaryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CameraTriggerSecondary__Array_to_Go(goSlice []CameraTriggerSecondary, cSlice []CCameraTriggerSecondary) {
	for i := 0; i < len(cSlice); i++ {
		CameraTriggerSecondaryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CameraTriggerSecondary__Array_to_C(cSlice []CCameraTriggerSecondary, goSlice []CameraTriggerSecondary) {
	for i := 0; i < len(goSlice); i++ {
		CameraTriggerSecondaryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
