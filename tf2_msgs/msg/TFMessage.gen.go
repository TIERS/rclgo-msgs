/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package tf2_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "github.com/TIERS/rclgo-msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltf2_msgs__rosidl_typesupport_c -ltf2_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <tf2_msgs/msg/tf_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("tf2_msgs/TFMessage", TFMessageTypeSupport)
}

// Do not create instances of this type directly. Always use NewTFMessage
// function instead.
type TFMessage struct {
	Transforms []geometry_msgs_msg.TransformStamped `yaml:"transforms"`
}

// NewTFMessage creates a new TFMessage with default values.
func NewTFMessage() *TFMessage {
	self := TFMessage{}
	self.SetDefaults()
	return &self
}

func (t *TFMessage) Clone() *TFMessage {
	c := &TFMessage{}
	if t.Transforms != nil {
		c.Transforms = make([]geometry_msgs_msg.TransformStamped, len(t.Transforms))
		geometry_msgs_msg.CloneTransformStampedSlice(c.Transforms, t.Transforms)
	}
	return c
}

func (t *TFMessage) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TFMessage) SetDefaults() {
	t.Transforms = nil
}

// CloneTFMessageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTFMessageSlice(dst, src []TFMessage) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TFMessageTypeSupport types.MessageTypeSupport = _TFMessageTypeSupport{}

type _TFMessageTypeSupport struct{}

func (t _TFMessageTypeSupport) New() types.Message {
	return NewTFMessage()
}

func (t _TFMessageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.tf2_msgs__msg__TFMessage
	return (unsafe.Pointer)(C.tf2_msgs__msg__TFMessage__create())
}

func (t _TFMessageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.tf2_msgs__msg__TFMessage__destroy((*C.tf2_msgs__msg__TFMessage)(pointer_to_free))
}

func (t _TFMessageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TFMessage)
	mem := (*C.tf2_msgs__msg__TFMessage)(dst)
	geometry_msgs_msg.TransformStamped__Sequence_to_C((*geometry_msgs_msg.CTransformStamped__Sequence)(unsafe.Pointer(&mem.transforms)), m.Transforms)
}

func (t _TFMessageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TFMessage)
	mem := (*C.tf2_msgs__msg__TFMessage)(ros2_message_buffer)
	geometry_msgs_msg.TransformStamped__Sequence_to_Go(&m.Transforms, *(*geometry_msgs_msg.CTransformStamped__Sequence)(unsafe.Pointer(&mem.transforms)))
}

func (t _TFMessageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__tf2_msgs__msg__TFMessage())
}

type CTFMessage = C.tf2_msgs__msg__TFMessage
type CTFMessage__Sequence = C.tf2_msgs__msg__TFMessage__Sequence

func TFMessage__Sequence_to_Go(goSlice *[]TFMessage, cSlice CTFMessage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TFMessage, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.tf2_msgs__msg__TFMessage__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__msg__TFMessage * uintptr(i)),
		))
		TFMessageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TFMessage__Sequence_to_C(cSlice *CTFMessage__Sequence, goSlice []TFMessage) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.tf2_msgs__msg__TFMessage)(C.malloc((C.size_t)(C.sizeof_struct_tf2_msgs__msg__TFMessage * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.tf2_msgs__msg__TFMessage)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_tf2_msgs__msg__TFMessage * uintptr(i)),
		))
		TFMessageTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TFMessage__Array_to_Go(goSlice []TFMessage, cSlice []CTFMessage) {
	for i := 0; i < len(cSlice); i++ {
		TFMessageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TFMessage__Array_to_C(cSlice []CTFMessage, goSlice []TFMessage) {
	for i := 0; i < len(goSlice); i++ {
		TFMessageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
