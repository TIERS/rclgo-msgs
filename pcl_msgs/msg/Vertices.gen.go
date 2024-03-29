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
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pcl_msgs/msg/vertices.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pcl_msgs/Vertices", VerticesTypeSupport)
}

// Do not create instances of this type directly. Always use NewVertices
// function instead.
type Vertices struct {
	Vertices []uint32 `yaml:"vertices"`// List of point indices
}

// NewVertices creates a new Vertices with default values.
func NewVertices() *Vertices {
	self := Vertices{}
	self.SetDefaults()
	return &self
}

func (t *Vertices) Clone() *Vertices {
	c := &Vertices{}
	if t.Vertices != nil {
		c.Vertices = make([]uint32, len(t.Vertices))
		copy(c.Vertices, t.Vertices)
	}
	return c
}

func (t *Vertices) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Vertices) SetDefaults() {
	t.Vertices = nil
}

// CloneVerticesSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVerticesSlice(dst, src []Vertices) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VerticesTypeSupport types.MessageTypeSupport = _VerticesTypeSupport{}

type _VerticesTypeSupport struct{}

func (t _VerticesTypeSupport) New() types.Message {
	return NewVertices()
}

func (t _VerticesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__Vertices
	return (unsafe.Pointer)(C.pcl_msgs__msg__Vertices__create())
}

func (t _VerticesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__Vertices__destroy((*C.pcl_msgs__msg__Vertices)(pointer_to_free))
}

func (t _VerticesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vertices)
	mem := (*C.pcl_msgs__msg__Vertices)(dst)
	primitives.Uint32__Sequence_to_C((*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.vertices)), m.Vertices)
}

func (t _VerticesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vertices)
	mem := (*C.pcl_msgs__msg__Vertices)(ros2_message_buffer)
	primitives.Uint32__Sequence_to_Go(&m.Vertices, *(*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.vertices)))
}

func (t _VerticesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__Vertices())
}

type CVertices = C.pcl_msgs__msg__Vertices
type CVertices__Sequence = C.pcl_msgs__msg__Vertices__Sequence

func Vertices__Sequence_to_Go(goSlice *[]Vertices, cSlice CVertices__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vertices, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pcl_msgs__msg__Vertices__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__Vertices * uintptr(i)),
		))
		VerticesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Vertices__Sequence_to_C(cSlice *CVertices__Sequence, goSlice []Vertices) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__Vertices)(C.malloc((C.size_t)(C.sizeof_struct_pcl_msgs__msg__Vertices * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pcl_msgs__msg__Vertices)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__Vertices * uintptr(i)),
		))
		VerticesTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Vertices__Array_to_Go(goSlice []Vertices, cSlice []CVertices) {
	for i := 0; i < len(cSlice); i++ {
		VerticesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vertices__Array_to_C(cSlice []CVertices, goSlice []Vertices) {
	for i := 0; i < len(goSlice); i++ {
		VerticesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
