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
	sensor_msgs_msg "github.com/TIERS/rclgo-msgs/sensor_msgs/msg"
	std_msgs_msg "github.com/TIERS/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpcl_msgs__rosidl_typesupport_c -lpcl_msgs__rosidl_generator_c
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pcl_msgs/msg/polygon_mesh.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pcl_msgs/PolygonMesh", PolygonMeshTypeSupport)
}

// Do not create instances of this type directly. Always use NewPolygonMesh
// function instead.
type PolygonMesh struct {
	Header std_msgs_msg.Header `yaml:"header"`// Separate header for the polygonal surface
	Cloud sensor_msgs_msg.PointCloud2 `yaml:"cloud"`// Separate header for the polygonal surfaceVertices of the mesh as a point cloud
	Polygons []Vertices `yaml:"polygons"`// Separate header for the polygonal surfaceVertices of the mesh as a point cloudList of polygons
}

// NewPolygonMesh creates a new PolygonMesh with default values.
func NewPolygonMesh() *PolygonMesh {
	self := PolygonMesh{}
	self.SetDefaults()
	return &self
}

func (t *PolygonMesh) Clone() *PolygonMesh {
	c := &PolygonMesh{}
	c.Header = *t.Header.Clone()
	c.Cloud = *t.Cloud.Clone()
	if t.Polygons != nil {
		c.Polygons = make([]Vertices, len(t.Polygons))
		CloneVerticesSlice(c.Polygons, t.Polygons)
	}
	return c
}

func (t *PolygonMesh) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PolygonMesh) SetDefaults() {
	t.Header.SetDefaults()
	t.Cloud.SetDefaults()
	t.Polygons = nil
}

// ClonePolygonMeshSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePolygonMeshSlice(dst, src []PolygonMesh) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PolygonMeshTypeSupport types.MessageTypeSupport = _PolygonMeshTypeSupport{}

type _PolygonMeshTypeSupport struct{}

func (t _PolygonMeshTypeSupport) New() types.Message {
	return NewPolygonMesh()
}

func (t _PolygonMeshTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__PolygonMesh
	return (unsafe.Pointer)(C.pcl_msgs__msg__PolygonMesh__create())
}

func (t _PolygonMeshTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__PolygonMesh__destroy((*C.pcl_msgs__msg__PolygonMesh)(pointer_to_free))
}

func (t _PolygonMeshTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PolygonMesh)
	mem := (*C.pcl_msgs__msg__PolygonMesh)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	sensor_msgs_msg.PointCloud2TypeSupport.AsCStruct(unsafe.Pointer(&mem.cloud), &m.Cloud)
	Vertices__Sequence_to_C(&mem.polygons, m.Polygons)
}

func (t _PolygonMeshTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PolygonMesh)
	mem := (*C.pcl_msgs__msg__PolygonMesh)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	sensor_msgs_msg.PointCloud2TypeSupport.AsGoStruct(&m.Cloud, unsafe.Pointer(&mem.cloud))
	Vertices__Sequence_to_Go(&m.Polygons, mem.polygons)
}

func (t _PolygonMeshTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__PolygonMesh())
}

type CPolygonMesh = C.pcl_msgs__msg__PolygonMesh
type CPolygonMesh__Sequence = C.pcl_msgs__msg__PolygonMesh__Sequence

func PolygonMesh__Sequence_to_Go(goSlice *[]PolygonMesh, cSlice CPolygonMesh__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PolygonMesh, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pcl_msgs__msg__PolygonMesh__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__PolygonMesh * uintptr(i)),
		))
		PolygonMeshTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PolygonMesh__Sequence_to_C(cSlice *CPolygonMesh__Sequence, goSlice []PolygonMesh) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__PolygonMesh)(C.malloc((C.size_t)(C.sizeof_struct_pcl_msgs__msg__PolygonMesh * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pcl_msgs__msg__PolygonMesh)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pcl_msgs__msg__PolygonMesh * uintptr(i)),
		))
		PolygonMeshTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PolygonMesh__Array_to_Go(goSlice []PolygonMesh, cSlice []CPolygonMesh) {
	for i := 0; i < len(cSlice); i++ {
		PolygonMeshTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PolygonMesh__Array_to_C(cSlice []CPolygonMesh, goSlice []PolygonMesh) {
	for i := 0; i < len(goSlice); i++ {
		PolygonMeshTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
