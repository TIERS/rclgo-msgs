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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/TIERS/rclgo-msgs/builtin_interfaces/msg"
	geometry_msgs_msg "github.com/TIERS/rclgo-msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/TIERS/rclgo-msgs/std_msgs/msg"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/image_marker.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/ImageMarker", ImageMarkerTypeSupport)
}
const (
	ImageMarker_CIRCLE int32 = 0
	ImageMarker_LINE_STRIP int32 = 1
	ImageMarker_LINE_LIST int32 = 2
	ImageMarker_POLYGON int32 = 3
	ImageMarker_POINTS int32 = 4
	ImageMarker_ADD int32 = 0
	ImageMarker_REMOVE int32 = 1
)

// Do not create instances of this type directly. Always use NewImageMarker
// function instead.
type ImageMarker struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Ns string `yaml:"ns"`// Namespace which is used with the id to form a unique id.
	Id int32 `yaml:"id"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.
	Type int32 `yaml:"type"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.
	Action int32 `yaml:"action"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.
	Position geometry_msgs_msg.Point `yaml:"position"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.
	Scale float32 `yaml:"scale"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.The scale of the object, e.g. the diameter for a CIRCLE.
	OutlineColor std_msgs_msg.ColorRGBA `yaml:"outline_color"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.The scale of the object, e.g. the diameter for a CIRCLE.The outline color of the marker.
	Filled uint8 `yaml:"filled"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.The scale of the object, e.g. the diameter for a CIRCLE.The outline color of the marker.Whether or not to fill in the shape with color.
	FillColor std_msgs_msg.ColorRGBA `yaml:"fill_color"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.The scale of the object, e.g. the diameter for a CIRCLE.The outline color of the marker.Whether or not to fill in the shape with color.Fill color; in the range: [0.0-1.0]
	Lifetime builtin_interfaces_msg.Duration `yaml:"lifetime"`// Namespace which is used with the id to form a unique id.Unique id within the namespace.One of the above types, e.g. CIRCLE, LINE_STRIP, etc.Either ADD or REMOVE.Two-dimensional coordinate position, in pixel-coordinates.The scale of the object, e.g. the diameter for a CIRCLE.The outline color of the marker.Whether or not to fill in the shape with color.Fill color; in the range: [0.0-1.0]How long the object should last before being automatically deleted.0 indicates forever.
	Points []geometry_msgs_msg.Point `yaml:"points"`// Coordinates in 2D in pixel coords. Used for LINE_STRIP, LINE_LIST, POINTS, etc.
	OutlineColors []std_msgs_msg.ColorRGBA `yaml:"outline_colors"`// Coordinates in 2D in pixel coords. Used for LINE_STRIP, LINE_LIST, POINTS, etc.The color for each line, point, etc. in the points field.
}

// NewImageMarker creates a new ImageMarker with default values.
func NewImageMarker() *ImageMarker {
	self := ImageMarker{}
	self.SetDefaults()
	return &self
}

func (t *ImageMarker) Clone() *ImageMarker {
	c := &ImageMarker{}
	c.Header = *t.Header.Clone()
	c.Ns = t.Ns
	c.Id = t.Id
	c.Type = t.Type
	c.Action = t.Action
	c.Position = *t.Position.Clone()
	c.Scale = t.Scale
	c.OutlineColor = *t.OutlineColor.Clone()
	c.Filled = t.Filled
	c.FillColor = *t.FillColor.Clone()
	c.Lifetime = *t.Lifetime.Clone()
	if t.Points != nil {
		c.Points = make([]geometry_msgs_msg.Point, len(t.Points))
		geometry_msgs_msg.ClonePointSlice(c.Points, t.Points)
	}
	if t.OutlineColors != nil {
		c.OutlineColors = make([]std_msgs_msg.ColorRGBA, len(t.OutlineColors))
		std_msgs_msg.CloneColorRGBASlice(c.OutlineColors, t.OutlineColors)
	}
	return c
}

func (t *ImageMarker) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ImageMarker) SetDefaults() {
	t.Header.SetDefaults()
	t.Ns = ""
	t.Id = 0
	t.Type = 0
	t.Action = 0
	t.Position.SetDefaults()
	t.Scale = 0
	t.OutlineColor.SetDefaults()
	t.Filled = 0
	t.FillColor.SetDefaults()
	t.Lifetime.SetDefaults()
	t.Points = nil
	t.OutlineColors = nil
}

// CloneImageMarkerSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneImageMarkerSlice(dst, src []ImageMarker) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ImageMarkerTypeSupport types.MessageTypeSupport = _ImageMarkerTypeSupport{}

type _ImageMarkerTypeSupport struct{}

func (t _ImageMarkerTypeSupport) New() types.Message {
	return NewImageMarker()
}

func (t _ImageMarkerTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__ImageMarker
	return (unsafe.Pointer)(C.visualization_msgs__msg__ImageMarker__create())
}

func (t _ImageMarkerTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__ImageMarker__destroy((*C.visualization_msgs__msg__ImageMarker)(pointer_to_free))
}

func (t _ImageMarkerTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ImageMarker)
	mem := (*C.visualization_msgs__msg__ImageMarker)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.ns), m.Ns)
	mem.id = C.int32_t(m.Id)
	mem._type = C.int32_t(m.Type)
	mem.action = C.int32_t(m.Action)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.position), &m.Position)
	mem.scale = C.float(m.Scale)
	std_msgs_msg.ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&mem.outline_color), &m.OutlineColor)
	mem.filled = C.uint8_t(m.Filled)
	std_msgs_msg.ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&mem.fill_color), &m.FillColor)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.lifetime), &m.Lifetime)
	geometry_msgs_msg.Point__Sequence_to_C((*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)), m.Points)
	std_msgs_msg.ColorRGBA__Sequence_to_C((*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.outline_colors)), m.OutlineColors)
}

func (t _ImageMarkerTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ImageMarker)
	mem := (*C.visualization_msgs__msg__ImageMarker)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.Ns, unsafe.Pointer(&mem.ns))
	m.Id = int32(mem.id)
	m.Type = int32(mem._type)
	m.Action = int32(mem.action)
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.Position, unsafe.Pointer(&mem.position))
	m.Scale = float32(mem.scale)
	std_msgs_msg.ColorRGBATypeSupport.AsGoStruct(&m.OutlineColor, unsafe.Pointer(&mem.outline_color))
	m.Filled = uint8(mem.filled)
	std_msgs_msg.ColorRGBATypeSupport.AsGoStruct(&m.FillColor, unsafe.Pointer(&mem.fill_color))
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Lifetime, unsafe.Pointer(&mem.lifetime))
	geometry_msgs_msg.Point__Sequence_to_Go(&m.Points, *(*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)))
	std_msgs_msg.ColorRGBA__Sequence_to_Go(&m.OutlineColors, *(*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.outline_colors)))
}

func (t _ImageMarkerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__ImageMarker())
}

type CImageMarker = C.visualization_msgs__msg__ImageMarker
type CImageMarker__Sequence = C.visualization_msgs__msg__ImageMarker__Sequence

func ImageMarker__Sequence_to_Go(goSlice *[]ImageMarker, cSlice CImageMarker__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ImageMarker, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__ImageMarker__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__ImageMarker * uintptr(i)),
		))
		ImageMarkerTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ImageMarker__Sequence_to_C(cSlice *CImageMarker__Sequence, goSlice []ImageMarker) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__ImageMarker)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__ImageMarker * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__ImageMarker)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__ImageMarker * uintptr(i)),
		))
		ImageMarkerTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ImageMarker__Array_to_Go(goSlice []ImageMarker, cSlice []CImageMarker) {
	for i := 0; i < len(cSlice); i++ {
		ImageMarkerTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ImageMarker__Array_to_C(cSlice []CImageMarker, goSlice []ImageMarker) {
	for i := 0; i < len(goSlice); i++ {
		ImageMarkerTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
