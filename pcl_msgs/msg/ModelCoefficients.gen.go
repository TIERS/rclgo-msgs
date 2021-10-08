/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pcl_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/TIERS/rclgo-msgs/std_msgs/msg"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pcl_msgs/msg/model_coefficients.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pcl_msgs/ModelCoefficients", ModelCoefficientsTypeSupport)
}

// Do not create instances of this type directly. Always use NewModelCoefficients
// function instead.
type ModelCoefficients struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Values []float32 `yaml:"values"`
}

// NewModelCoefficients creates a new ModelCoefficients with default values.
func NewModelCoefficients() *ModelCoefficients {
	self := ModelCoefficients{}
	self.SetDefaults()
	return &self
}

func (t *ModelCoefficients) Clone() *ModelCoefficients {
	c := &ModelCoefficients{}
	c.Header = *t.Header.Clone()
	if t.Values != nil {
		c.Values = make([]float32, len(t.Values))
		copy(c.Values, t.Values)
	}
	return c
}

func (t *ModelCoefficients) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ModelCoefficients) SetDefaults() {
	t.Header.SetDefaults()
	t.Values = nil
}

// CloneModelCoefficientsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneModelCoefficientsSlice(dst, src []ModelCoefficients) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ModelCoefficientsTypeSupport types.MessageTypeSupport = _ModelCoefficientsTypeSupport{}

type _ModelCoefficientsTypeSupport struct{}

func (t _ModelCoefficientsTypeSupport) New() types.Message {
	return NewModelCoefficients()
}

func (t _ModelCoefficientsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__ModelCoefficients
	return (unsafe.Pointer)(C.pcl_msgs__msg__ModelCoefficients__create())
}

func (t _ModelCoefficientsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__ModelCoefficients__destroy((*C.pcl_msgs__msg__ModelCoefficients)(pointer_to_free))
}

func (t _ModelCoefficientsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ModelCoefficients)
	mem := (*C.pcl_msgs__msg__ModelCoefficients)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.values)), m.Values)
}

func (t _ModelCoefficientsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ModelCoefficients)
	mem := (*C.pcl_msgs__msg__ModelCoefficients)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.Float32__Sequence_to_Go(&m.Values, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.values)))
}

func (t _ModelCoefficientsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__ModelCoefficients())
}

type CModelCoefficients = C.pcl_msgs__msg__ModelCoefficients
type CModelCoefficients__Sequence = C.pcl_msgs__msg__ModelCoefficients__Sequence

func ModelCoefficients__Sequence_to_Go(goSlice *[]ModelCoefficients, cSlice CModelCoefficients__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ModelCoefficients, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pcl_msgs__msg__ModelCoefficients__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(i)),
		))
		ModelCoefficientsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ModelCoefficients__Sequence_to_C(cSlice *CModelCoefficients__Sequence, goSlice []ModelCoefficients) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__ModelCoefficients)(C.malloc((C.size_t)(C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pcl_msgs__msg__ModelCoefficients)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__ModelCoefficients * uintptr(i)),
		))
		ModelCoefficientsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ModelCoefficients__Array_to_Go(goSlice []ModelCoefficients, cSlice []CModelCoefficients) {
	for i := 0; i < len(cSlice); i++ {
		ModelCoefficientsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ModelCoefficients__Array_to_C(cSlice []CModelCoefficients, goSlice []ModelCoefficients) {
	for i := 0; i < len(goSlice); i++ {
		ModelCoefficientsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
