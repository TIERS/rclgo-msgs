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
	geometry_msgs_msg "github.com/tiiuae/rclgo-msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/interactive_marker.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/InteractiveMarker", InteractiveMarkerTypeSupport)
}

// Do not create instances of this type directly. Always use NewInteractiveMarker
// function instead.
type InteractiveMarker struct {
	Header std_msgs_msg.Header `yaml:"header"`// Time/frame info.If header.time is set to 0, the marker will be retransformed intoits frame on each timestep. You will receive the pose feedbackin the same frame.Otherwise, you might receive feedback in a different frame.For rviz, this will be the current 'fixed frame' set by the user.
	Pose geometry_msgs_msg.Pose `yaml:"pose"`// Initial pose. Also, defines the pivot point for rotations.
	Name string `yaml:"name"`// Identifying string. Must be globally unique inthe topic that this message is sent through.
	Description string `yaml:"description"`// Short description (< 40 characters).
	Scale float32 `yaml:"scale"`// Scale to be used for default controls (default=1).
	MenuEntries []MenuEntry `yaml:"menu_entries"`// All menu and submenu entries associated with this marker.
	Controls []InteractiveMarkerControl `yaml:"controls"`// List of controls displayed for this marker.
}

// NewInteractiveMarker creates a new InteractiveMarker with default values.
func NewInteractiveMarker() *InteractiveMarker {
	self := InteractiveMarker{}
	self.SetDefaults()
	return &self
}

func (t *InteractiveMarker) Clone() *InteractiveMarker {
	c := &InteractiveMarker{}
	c.Header = *t.Header.Clone()
	c.Pose = *t.Pose.Clone()
	c.Name = t.Name
	c.Description = t.Description
	c.Scale = t.Scale
	if t.MenuEntries != nil {
		c.MenuEntries = make([]MenuEntry, len(t.MenuEntries))
		CloneMenuEntrySlice(c.MenuEntries, t.MenuEntries)
	}
	if t.Controls != nil {
		c.Controls = make([]InteractiveMarkerControl, len(t.Controls))
		CloneInteractiveMarkerControlSlice(c.Controls, t.Controls)
	}
	return c
}

func (t *InteractiveMarker) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InteractiveMarker) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
	t.Name = ""
	t.Description = ""
	t.Scale = 0
	t.MenuEntries = nil
	t.Controls = nil
}

// CloneInteractiveMarkerSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInteractiveMarkerSlice(dst, src []InteractiveMarker) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InteractiveMarkerTypeSupport types.MessageTypeSupport = _InteractiveMarkerTypeSupport{}

type _InteractiveMarkerTypeSupport struct{}

func (t _InteractiveMarkerTypeSupport) New() types.Message {
	return NewInteractiveMarker()
}

func (t _InteractiveMarkerTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarker
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarker__create())
}

func (t _InteractiveMarkerTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarker__destroy((*C.visualization_msgs__msg__InteractiveMarker)(pointer_to_free))
}

func (t _InteractiveMarkerTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InteractiveMarker)
	mem := (*C.visualization_msgs__msg__InteractiveMarker)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	geometry_msgs_msg.PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.description), m.Description)
	mem.scale = C.float(m.Scale)
	MenuEntry__Sequence_to_C(&mem.menu_entries, m.MenuEntries)
	InteractiveMarkerControl__Sequence_to_C(&mem.controls, m.Controls)
}

func (t _InteractiveMarkerTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InteractiveMarker)
	mem := (*C.visualization_msgs__msg__InteractiveMarker)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	geometry_msgs_msg.PoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	primitives.StringAsGoStruct(&m.Description, unsafe.Pointer(&mem.description))
	m.Scale = float32(mem.scale)
	MenuEntry__Sequence_to_Go(&m.MenuEntries, mem.menu_entries)
	InteractiveMarkerControl__Sequence_to_Go(&m.Controls, mem.controls)
}

func (t _InteractiveMarkerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarker())
}

type CInteractiveMarker = C.visualization_msgs__msg__InteractiveMarker
type CInteractiveMarker__Sequence = C.visualization_msgs__msg__InteractiveMarker__Sequence

func InteractiveMarker__Sequence_to_Go(goSlice *[]InteractiveMarker, cSlice CInteractiveMarker__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarker, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarker__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(i)),
		))
		InteractiveMarkerTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InteractiveMarker__Sequence_to_C(cSlice *CInteractiveMarker__Sequence, goSlice []InteractiveMarker) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarker)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarker)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(i)),
		))
		InteractiveMarkerTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InteractiveMarker__Array_to_Go(goSlice []InteractiveMarker, cSlice []CInteractiveMarker) {
	for i := 0; i < len(cSlice); i++ {
		InteractiveMarkerTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarker__Array_to_C(cSlice []CInteractiveMarker, goSlice []InteractiveMarker) {
	for i := 0; i < len(goSlice); i++ {
		InteractiveMarkerTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
