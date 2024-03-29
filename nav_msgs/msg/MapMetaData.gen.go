/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/TIERS/rclgo-msgs/builtin_interfaces/msg"
	geometry_msgs_msg "github.com/TIERS/rclgo-msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <nav_msgs/msg/map_meta_data.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("nav_msgs/MapMetaData", MapMetaDataTypeSupport)
}

// Do not create instances of this type directly. Always use NewMapMetaData
// function instead.
type MapMetaData struct {
	MapLoadTime builtin_interfaces_msg.Time `yaml:"map_load_time"`// The time at which the map was loaded
	Resolution float32 `yaml:"resolution"`// The map resolution [m/cell]
	Width uint32 `yaml:"width"`// Map width [cells]
	Height uint32 `yaml:"height"`// Map height [cells]
	Origin geometry_msgs_msg.Pose `yaml:"origin"`// The origin of the map [m, m, rad].  This is the real-world pose of thebottom left corner of cell (0,0) in the map.
}

// NewMapMetaData creates a new MapMetaData with default values.
func NewMapMetaData() *MapMetaData {
	self := MapMetaData{}
	self.SetDefaults()
	return &self
}

func (t *MapMetaData) Clone() *MapMetaData {
	c := &MapMetaData{}
	c.MapLoadTime = *t.MapLoadTime.Clone()
	c.Resolution = t.Resolution
	c.Width = t.Width
	c.Height = t.Height
	c.Origin = *t.Origin.Clone()
	return c
}

func (t *MapMetaData) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MapMetaData) SetDefaults() {
	t.MapLoadTime.SetDefaults()
	t.Resolution = 0
	t.Width = 0
	t.Height = 0
	t.Origin.SetDefaults()
}

// CloneMapMetaDataSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMapMetaDataSlice(dst, src []MapMetaData) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MapMetaDataTypeSupport types.MessageTypeSupport = _MapMetaDataTypeSupport{}

type _MapMetaDataTypeSupport struct{}

func (t _MapMetaDataTypeSupport) New() types.Message {
	return NewMapMetaData()
}

func (t _MapMetaDataTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.nav_msgs__msg__MapMetaData
	return (unsafe.Pointer)(C.nav_msgs__msg__MapMetaData__create())
}

func (t _MapMetaDataTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.nav_msgs__msg__MapMetaData__destroy((*C.nav_msgs__msg__MapMetaData)(pointer_to_free))
}

func (t _MapMetaDataTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MapMetaData)
	mem := (*C.nav_msgs__msg__MapMetaData)(dst)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.map_load_time), &m.MapLoadTime)
	mem.resolution = C.float(m.Resolution)
	mem.width = C.uint32_t(m.Width)
	mem.height = C.uint32_t(m.Height)
	geometry_msgs_msg.PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.origin), &m.Origin)
}

func (t _MapMetaDataTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MapMetaData)
	mem := (*C.nav_msgs__msg__MapMetaData)(ros2_message_buffer)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.MapLoadTime, unsafe.Pointer(&mem.map_load_time))
	m.Resolution = float32(mem.resolution)
	m.Width = uint32(mem.width)
	m.Height = uint32(mem.height)
	geometry_msgs_msg.PoseTypeSupport.AsGoStruct(&m.Origin, unsafe.Pointer(&mem.origin))
}

func (t _MapMetaDataTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__nav_msgs__msg__MapMetaData())
}

type CMapMetaData = C.nav_msgs__msg__MapMetaData
type CMapMetaData__Sequence = C.nav_msgs__msg__MapMetaData__Sequence

func MapMetaData__Sequence_to_Go(goSlice *[]MapMetaData, cSlice CMapMetaData__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MapMetaData, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.nav_msgs__msg__MapMetaData__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__msg__MapMetaData * uintptr(i)),
		))
		MapMetaDataTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MapMetaData__Sequence_to_C(cSlice *CMapMetaData__Sequence, goSlice []MapMetaData) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.nav_msgs__msg__MapMetaData)(C.malloc((C.size_t)(C.sizeof_struct_nav_msgs__msg__MapMetaData * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.nav_msgs__msg__MapMetaData)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_nav_msgs__msg__MapMetaData * uintptr(i)),
		))
		MapMetaDataTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MapMetaData__Array_to_Go(goSlice []MapMetaData, cSlice []CMapMetaData) {
	for i := 0; i < len(cSlice); i++ {
		MapMetaDataTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MapMetaData__Array_to_C(cSlice []CMapMetaData, goSlice []MapMetaData) {
	for i := 0; i < len(goSlice); i++ {
		MapMetaDataTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
