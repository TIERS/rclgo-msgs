/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo-msgs/builtin_interfaces/msg"
	geometry_msgs_msg "github.com/tiiuae/rclgo-msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/marker.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/Marker", MarkerTypeSupport)
}
const (
	Marker_ARROW int32 = 0
	Marker_CUBE int32 = 1
	Marker_SPHERE int32 = 2
	Marker_CYLINDER int32 = 3
	Marker_LINE_STRIP int32 = 4
	Marker_LINE_LIST int32 = 5
	Marker_CUBE_LIST int32 = 6
	Marker_SPHERE_LIST int32 = 7
	Marker_POINTS int32 = 8
	Marker_TEXT_VIEW_FACING int32 = 9
	Marker_MESH_RESOURCE int32 = 10
	Marker_TRIANGLE_LIST int32 = 11
	Marker_ADD int32 = 0
	Marker_MODIFY int32 = 0
	Marker_DELETE int32 = 2
	Marker_DELETEALL int32 = 3
)

// Do not create instances of this type directly. Always use NewMarker
// function instead.
type Marker struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header for timestamp and frame id.
	Ns string `yaml:"ns"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.
	Id int32 `yaml:"id"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.
	Type int32 `yaml:"type"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.
	Action int32 `yaml:"action"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objects
	Pose geometry_msgs_msg.Pose `yaml:"pose"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objectsPose of the object with respect the frame_id specified in the header.
	Scale geometry_msgs_msg.Vector3 `yaml:"scale"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objectsPose of the object with respect the frame_id specified in the header.Scale of the object; 1,1,1 means default (usually 1 meter square).
	Color std_msgs_msg.ColorRGBA `yaml:"color"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objectsPose of the object with respect the frame_id specified in the header.Scale of the object; 1,1,1 means default (usually 1 meter square).Color of the object; in the range: [0.0-1.0]
	Lifetime builtin_interfaces_msg.Duration `yaml:"lifetime"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objectsPose of the object with respect the frame_id specified in the header.Scale of the object; 1,1,1 means default (usually 1 meter square).Color of the object; in the range: [0.0-1.0]How long the object should last before being automatically deleted.0 indicates forever.
	FrameLocked bool `yaml:"frame_locked"`// Header for timestamp and frame id.Namespace in which to place the object.Used in conjunction with id to create a unique name for the object.Object ID used in conjunction with the namespace for manipulating and deleting the object later.Type of object.Action to take; one of:- 0 add/modify an object- 1 (deprecated)- 2 deletes an object- 3 deletes all objectsPose of the object with respect the frame_id specified in the header.Scale of the object; 1,1,1 means default (usually 1 meter square).Color of the object; in the range: [0.0-1.0]How long the object should last before being automatically deleted.0 indicates forever.If this marker should be frame-locked, i.e. retransformed into its frame every timestep.
	Points []geometry_msgs_msg.Point `yaml:"points"`// Only used if the type specified has some use for them (eg. POINTS, LINE_STRIP, etc.)
	Colors []std_msgs_msg.ColorRGBA `yaml:"colors"`// Only used if the type specified has some use for them (eg. POINTS, LINE_STRIP, etc.)Only used if the type specified has some use for them (eg. POINTS, LINE_STRIP, etc.)The number of colors provided must either be 0 or equal to the number of points provided.NOTE: alpha is not yet used
	Text string `yaml:"text"`// Only used for text markers
	MeshResource string `yaml:"mesh_resource"`// Only used for MESH_RESOURCE markers.
	MeshUseEmbeddedMaterials bool `yaml:"mesh_use_embedded_materials"`// Only used for MESH_RESOURCE markers.
}

// NewMarker creates a new Marker with default values.
func NewMarker() *Marker {
	self := Marker{}
	self.SetDefaults()
	return &self
}

func (t *Marker) Clone() *Marker {
	c := &Marker{}
	c.Header = *t.Header.Clone()
	c.Ns = t.Ns
	c.Id = t.Id
	c.Type = t.Type
	c.Action = t.Action
	c.Pose = *t.Pose.Clone()
	c.Scale = *t.Scale.Clone()
	c.Color = *t.Color.Clone()
	c.Lifetime = *t.Lifetime.Clone()
	c.FrameLocked = t.FrameLocked
	if t.Points != nil {
		c.Points = make([]geometry_msgs_msg.Point, len(t.Points))
		geometry_msgs_msg.ClonePointSlice(c.Points, t.Points)
	}
	if t.Colors != nil {
		c.Colors = make([]std_msgs_msg.ColorRGBA, len(t.Colors))
		std_msgs_msg.CloneColorRGBASlice(c.Colors, t.Colors)
	}
	c.Text = t.Text
	c.MeshResource = t.MeshResource
	c.MeshUseEmbeddedMaterials = t.MeshUseEmbeddedMaterials
	return c
}

func (t *Marker) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Marker) SetDefaults() {
	t.Header.SetDefaults()
	t.Ns = ""
	t.Id = 0
	t.Type = 0
	t.Action = 0
	t.Pose.SetDefaults()
	t.Scale.SetDefaults()
	t.Color.SetDefaults()
	t.Lifetime.SetDefaults()
	t.FrameLocked = false
	t.Points = nil
	t.Colors = nil
	t.Text = ""
	t.MeshResource = ""
	t.MeshUseEmbeddedMaterials = false
}

// CloneMarkerSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMarkerSlice(dst, src []Marker) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MarkerTypeSupport types.MessageTypeSupport = _MarkerTypeSupport{}

type _MarkerTypeSupport struct{}

func (t _MarkerTypeSupport) New() types.Message {
	return NewMarker()
}

func (t _MarkerTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__Marker
	return (unsafe.Pointer)(C.visualization_msgs__msg__Marker__create())
}

func (t _MarkerTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__Marker__destroy((*C.visualization_msgs__msg__Marker)(pointer_to_free))
}

func (t _MarkerTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Marker)
	mem := (*C.visualization_msgs__msg__Marker)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.ns), m.Ns)
	mem.id = C.int32_t(m.Id)
	mem._type = C.int32_t(m.Type)
	mem.action = C.int32_t(m.Action)
	geometry_msgs_msg.PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.scale), &m.Scale)
	std_msgs_msg.ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&mem.color), &m.Color)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.lifetime), &m.Lifetime)
	mem.frame_locked = C.bool(m.FrameLocked)
	geometry_msgs_msg.Point__Sequence_to_C((*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)), m.Points)
	std_msgs_msg.ColorRGBA__Sequence_to_C((*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.colors)), m.Colors)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.text), m.Text)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.mesh_resource), m.MeshResource)
	mem.mesh_use_embedded_materials = C.bool(m.MeshUseEmbeddedMaterials)
}

func (t _MarkerTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Marker)
	mem := (*C.visualization_msgs__msg__Marker)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.Ns, unsafe.Pointer(&mem.ns))
	m.Id = int32(mem.id)
	m.Type = int32(mem._type)
	m.Action = int32(mem.action)
	geometry_msgs_msg.PoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.Scale, unsafe.Pointer(&mem.scale))
	std_msgs_msg.ColorRGBATypeSupport.AsGoStruct(&m.Color, unsafe.Pointer(&mem.color))
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Lifetime, unsafe.Pointer(&mem.lifetime))
	m.FrameLocked = bool(mem.frame_locked)
	geometry_msgs_msg.Point__Sequence_to_Go(&m.Points, *(*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)))
	std_msgs_msg.ColorRGBA__Sequence_to_Go(&m.Colors, *(*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.colors)))
	primitives.StringAsGoStruct(&m.Text, unsafe.Pointer(&mem.text))
	primitives.StringAsGoStruct(&m.MeshResource, unsafe.Pointer(&mem.mesh_resource))
	m.MeshUseEmbeddedMaterials = bool(mem.mesh_use_embedded_materials)
}

func (t _MarkerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__Marker())
}

type CMarker = C.visualization_msgs__msg__Marker
type CMarker__Sequence = C.visualization_msgs__msg__Marker__Sequence

func Marker__Sequence_to_Go(goSlice *[]Marker, cSlice CMarker__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Marker, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__Marker__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__Marker * uintptr(i)),
		))
		MarkerTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Marker__Sequence_to_C(cSlice *CMarker__Sequence, goSlice []Marker) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__Marker)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__Marker * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__Marker)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__Marker * uintptr(i)),
		))
		MarkerTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Marker__Array_to_Go(goSlice []Marker, cSlice []CMarker) {
	for i := 0; i < len(cSlice); i++ {
		MarkerTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Marker__Array_to_C(cSlice []CMarker, goSlice []Marker) {
	for i := 0; i < len(goSlice); i++ {
		MarkerTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
