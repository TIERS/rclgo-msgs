/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/msg/projected_map_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/ProjectedMapInfo", ProjectedMapInfoTypeSupport)
}

// Do not create instances of this type directly. Always use NewProjectedMapInfo
// function instead.
type ProjectedMapInfo struct {
	FrameId string `yaml:"frame_id"`
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Width float64 `yaml:"width"`
	Height float64 `yaml:"height"`
	MinZ float64 `yaml:"min_z"`
	MaxZ float64 `yaml:"max_z"`
}

// NewProjectedMapInfo creates a new ProjectedMapInfo with default values.
func NewProjectedMapInfo() *ProjectedMapInfo {
	self := ProjectedMapInfo{}
	self.SetDefaults()
	return &self
}

func (t *ProjectedMapInfo) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *ProjectedMapInfo) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var ProjectedMapInfoTypeSupport types.MessageTypeSupport = _ProjectedMapInfoTypeSupport{}

type _ProjectedMapInfoTypeSupport struct{}

func (t _ProjectedMapInfoTypeSupport) New() types.Message {
	return NewProjectedMapInfo()
}

func (t _ProjectedMapInfoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__msg__ProjectedMapInfo
	return (unsafe.Pointer)(C.map_msgs__msg__ProjectedMapInfo__create())
}

func (t _ProjectedMapInfoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__msg__ProjectedMapInfo__destroy((*C.map_msgs__msg__ProjectedMapInfo)(pointer_to_free))
}

func (t _ProjectedMapInfoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ProjectedMapInfo)
	mem := (*C.map_msgs__msg__ProjectedMapInfo)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.frame_id), m.FrameId)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.width = C.double(m.Width)
	mem.height = C.double(m.Height)
	mem.min_z = C.double(m.MinZ)
	mem.max_z = C.double(m.MaxZ)
}

func (t _ProjectedMapInfoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ProjectedMapInfo)
	mem := (*C.map_msgs__msg__ProjectedMapInfo)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.FrameId, unsafe.Pointer(&mem.frame_id))
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.Width = float64(mem.width)
	m.Height = float64(mem.height)
	m.MinZ = float64(mem.min_z)
	m.MaxZ = float64(mem.max_z)
}

func (t _ProjectedMapInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__msg__ProjectedMapInfo())
}

type CProjectedMapInfo = C.map_msgs__msg__ProjectedMapInfo
type CProjectedMapInfo__Sequence = C.map_msgs__msg__ProjectedMapInfo__Sequence

func ProjectedMapInfo__Sequence_to_Go(goSlice *[]ProjectedMapInfo, cSlice CProjectedMapInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ProjectedMapInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__msg__ProjectedMapInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__msg__ProjectedMapInfo * uintptr(i)),
		))
		ProjectedMapInfoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ProjectedMapInfo__Sequence_to_C(cSlice *CProjectedMapInfo__Sequence, goSlice []ProjectedMapInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__msg__ProjectedMapInfo)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__msg__ProjectedMapInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__msg__ProjectedMapInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__msg__ProjectedMapInfo * uintptr(i)),
		))
		ProjectedMapInfoTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ProjectedMapInfo__Array_to_Go(goSlice []ProjectedMapInfo, cSlice []CProjectedMapInfo) {
	for i := 0; i < len(cSlice); i++ {
		ProjectedMapInfoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ProjectedMapInfo__Array_to_C(cSlice []CProjectedMapInfo, goSlice []ProjectedMapInfo) {
	for i := 0; i < len(goSlice); i++ {
		ProjectedMapInfoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}