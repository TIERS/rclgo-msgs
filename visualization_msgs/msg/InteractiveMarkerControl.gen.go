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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/interactive_marker_control.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/InteractiveMarkerControl", InteractiveMarkerControlTypeSupport)
}
const (
	InteractiveMarkerControl_INHERIT uint8 = 0// Orientation mode: controls how orientation changes.INHERIT: Follow orientation of interactive markerFIXED: Keep orientation fixed at initial stateVIEW_FACING: Align y-z plane with screen (x: forward, y:left, z:up).
	InteractiveMarkerControl_FIXED uint8 = 1// Orientation mode: controls how orientation changes.INHERIT: Follow orientation of interactive markerFIXED: Keep orientation fixed at initial stateVIEW_FACING: Align y-z plane with screen (x: forward, y:left, z:up).
	InteractiveMarkerControl_VIEW_FACING uint8 = 2// Orientation mode: controls how orientation changes.INHERIT: Follow orientation of interactive markerFIXED: Keep orientation fixed at initial stateVIEW_FACING: Align y-z plane with screen (x: forward, y:left, z:up).
	InteractiveMarkerControl_NONE uint8 = 0// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_MENU uint8 = 1// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_BUTTON uint8 = 2// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_MOVE_AXIS uint8 = 3// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_MOVE_PLANE uint8 = 4// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_ROTATE_AXIS uint8 = 5// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_MOVE_ROTATE uint8 = 6// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
	InteractiveMarkerControl_MOVE_3D uint8 = 7// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS."3D" interaction modes work with the mouse+SHIFT+CTRL or with 3D cursors.MOVE_3D: Translate freely in 3D space.ROTATE_3D: Rotate freely in 3D space about the origin of parent frame.MOVE_ROTATE_3D: Full 6-DOF freedom of translation and rotation about the cursor origin.
	InteractiveMarkerControl_ROTATE_3D uint8 = 8// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS."3D" interaction modes work with the mouse+SHIFT+CTRL or with 3D cursors.MOVE_3D: Translate freely in 3D space.ROTATE_3D: Rotate freely in 3D space about the origin of parent frame.MOVE_ROTATE_3D: Full 6-DOF freedom of translation and rotation about the cursor origin.
	InteractiveMarkerControl_MOVE_ROTATE_3D uint8 = 9// Interaction mode for this controlNONE: This control is only meant for visualization; no context menu.MENU: Like NONE, but right-click menu is active.BUTTON: Element can be left-clicked.MOVE_AXIS: Translate along local x-axis.MOVE_PLANE: Translate in local y-z plane.ROTATE_AXIS: Rotate around local x-axis.MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS."3D" interaction modes work with the mouse+SHIFT+CTRL or with 3D cursors.MOVE_3D: Translate freely in 3D space.ROTATE_3D: Rotate freely in 3D space about the origin of parent frame.MOVE_ROTATE_3D: Full 6-DOF freedom of translation and rotation about the cursor origin.
)

// Do not create instances of this type directly. Always use NewInteractiveMarkerControl
// function instead.
type InteractiveMarkerControl struct {
	Name string `yaml:"name"`// Identifying string for this control.You need to assign a unique value to this to receive feedback from the GUIon what actions the user performs on this control (e.g. a button click).
	Orientation geometry_msgs_msg.Quaternion `yaml:"orientation"`// Defines the local coordinate frame (relative to the pose of the parentinteractive marker) in which is being rotated and translated.Default: Identity
	OrientationMode uint8 `yaml:"orientation_mode"`
	InteractionMode uint8 `yaml:"interaction_mode"`
	AlwaysVisible bool `yaml:"always_visible"`// If true, the contained markers will also be visiblewhen the gui is not in interactive mode.
	Markers []Marker `yaml:"markers"`// Markers to be displayed as custom visual representation.Leave this empty to use the default control handles.Note:- The markers can be defined in an arbitrary coordinate frame,but will be transformed into the local frame of the interactive marker.- If the header of a marker is empty, its pose will be interpreted asrelative to the pose of the parent interactive marker.
	IndependentMarkerOrientation bool `yaml:"independent_marker_orientation"`// In VIEW_FACING mode, set this to true if you don't want the markersto be aligned with the camera view point. The markers will show upas in INHERIT mode.
	Description string `yaml:"description"`// Short description (< 40 characters) of what this control does,e.g. "Move the robot".Default: A generic description based on the interaction mode
}

// NewInteractiveMarkerControl creates a new InteractiveMarkerControl with default values.
func NewInteractiveMarkerControl() *InteractiveMarkerControl {
	self := InteractiveMarkerControl{}
	self.SetDefaults()
	return &self
}

func (t *InteractiveMarkerControl) Clone() *InteractiveMarkerControl {
	c := &InteractiveMarkerControl{}
	c.Name = t.Name
	c.Orientation = *t.Orientation.Clone()
	c.OrientationMode = t.OrientationMode
	c.InteractionMode = t.InteractionMode
	c.AlwaysVisible = t.AlwaysVisible
	if t.Markers != nil {
		c.Markers = make([]Marker, len(t.Markers))
		CloneMarkerSlice(c.Markers, t.Markers)
	}
	c.IndependentMarkerOrientation = t.IndependentMarkerOrientation
	c.Description = t.Description
	return c
}

func (t *InteractiveMarkerControl) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InteractiveMarkerControl) SetDefaults() {
	t.Name = ""
	t.Orientation.SetDefaults()
	t.OrientationMode = 0
	t.InteractionMode = 0
	t.AlwaysVisible = false
	t.Markers = nil
	t.IndependentMarkerOrientation = false
	t.Description = ""
}

// CloneInteractiveMarkerControlSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInteractiveMarkerControlSlice(dst, src []InteractiveMarkerControl) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InteractiveMarkerControlTypeSupport types.MessageTypeSupport = _InteractiveMarkerControlTypeSupport{}

type _InteractiveMarkerControlTypeSupport struct{}

func (t _InteractiveMarkerControlTypeSupport) New() types.Message {
	return NewInteractiveMarkerControl()
}

func (t _InteractiveMarkerControlTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerControl
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerControl__create())
}

func (t _InteractiveMarkerControlTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerControl__destroy((*C.visualization_msgs__msg__InteractiveMarkerControl)(pointer_to_free))
}

func (t _InteractiveMarkerControlTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InteractiveMarkerControl)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerControl)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	geometry_msgs_msg.QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.orientation), &m.Orientation)
	mem.orientation_mode = C.uint8_t(m.OrientationMode)
	mem.interaction_mode = C.uint8_t(m.InteractionMode)
	mem.always_visible = C.bool(m.AlwaysVisible)
	Marker__Sequence_to_C(&mem.markers, m.Markers)
	mem.independent_marker_orientation = C.bool(m.IndependentMarkerOrientation)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.description), m.Description)
}

func (t _InteractiveMarkerControlTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InteractiveMarkerControl)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerControl)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	geometry_msgs_msg.QuaternionTypeSupport.AsGoStruct(&m.Orientation, unsafe.Pointer(&mem.orientation))
	m.OrientationMode = uint8(mem.orientation_mode)
	m.InteractionMode = uint8(mem.interaction_mode)
	m.AlwaysVisible = bool(mem.always_visible)
	Marker__Sequence_to_Go(&m.Markers, mem.markers)
	m.IndependentMarkerOrientation = bool(mem.independent_marker_orientation)
	primitives.StringAsGoStruct(&m.Description, unsafe.Pointer(&mem.description))
}

func (t _InteractiveMarkerControlTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerControl())
}

type CInteractiveMarkerControl = C.visualization_msgs__msg__InteractiveMarkerControl
type CInteractiveMarkerControl__Sequence = C.visualization_msgs__msg__InteractiveMarkerControl__Sequence

func InteractiveMarkerControl__Sequence_to_Go(goSlice *[]InteractiveMarkerControl, cSlice CInteractiveMarkerControl__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerControl, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerControl__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerControl * uintptr(i)),
		))
		InteractiveMarkerControlTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InteractiveMarkerControl__Sequence_to_C(cSlice *CInteractiveMarkerControl__Sequence, goSlice []InteractiveMarkerControl) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerControl)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerControl * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerControl)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerControl * uintptr(i)),
		))
		InteractiveMarkerControlTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InteractiveMarkerControl__Array_to_Go(goSlice []InteractiveMarkerControl, cSlice []CInteractiveMarkerControl) {
	for i := 0; i < len(cSlice); i++ {
		InteractiveMarkerControlTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerControl__Array_to_C(cSlice []CInteractiveMarkerControl, goSlice []InteractiveMarkerControl) {
	for i := 0; i < len(goSlice); i++ {
		InteractiveMarkerControlTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
