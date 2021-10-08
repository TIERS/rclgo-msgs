/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/TIERS/rclgo-msgs/std_msgs/msg"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/msg/navigation_diagnostics.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/NavigationDiagnostics", NavigationDiagnosticsTypeSupport)
}

// Do not create instances of this type directly. Always use NewNavigationDiagnostics
// function instead.
type NavigationDiagnostics struct {
	Header std_msgs_msg.Header `yaml:"header"`
	State string `yaml:"state"`
	WaypointsInBuffer uint8 `yaml:"waypoints_in_buffer"`
	CurrentWaypointId uint8 `yaml:"current_waypoint_id"`
	BumperActive bool `yaml:"bumper_active"`
	LastNavGoal [3]float64 `yaml:"last_nav_goal"`
	CurrentNavGoal [3]float64 `yaml:"current_nav_goal"`
}

// NewNavigationDiagnostics creates a new NavigationDiagnostics with default values.
func NewNavigationDiagnostics() *NavigationDiagnostics {
	self := NavigationDiagnostics{}
	self.SetDefaults()
	return &self
}

func (t *NavigationDiagnostics) Clone() *NavigationDiagnostics {
	c := &NavigationDiagnostics{}
	c.Header = *t.Header.Clone()
	c.State = t.State
	c.WaypointsInBuffer = t.WaypointsInBuffer
	c.CurrentWaypointId = t.CurrentWaypointId
	c.BumperActive = t.BumperActive
	c.LastNavGoal = t.LastNavGoal
	c.CurrentNavGoal = t.CurrentNavGoal
	return c
}

func (t *NavigationDiagnostics) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NavigationDiagnostics) SetDefaults() {
	t.Header.SetDefaults()
	t.State = ""
	t.WaypointsInBuffer = 0
	t.CurrentWaypointId = 0
	t.BumperActive = false
	t.LastNavGoal = [3]float64{}
	t.CurrentNavGoal = [3]float64{}
}

// CloneNavigationDiagnosticsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNavigationDiagnosticsSlice(dst, src []NavigationDiagnostics) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NavigationDiagnosticsTypeSupport types.MessageTypeSupport = _NavigationDiagnosticsTypeSupport{}

type _NavigationDiagnosticsTypeSupport struct{}

func (t _NavigationDiagnosticsTypeSupport) New() types.Message {
	return NewNavigationDiagnostics()
}

func (t _NavigationDiagnosticsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__msg__NavigationDiagnostics
	return (unsafe.Pointer)(C.fog_msgs__msg__NavigationDiagnostics__create())
}

func (t _NavigationDiagnosticsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__msg__NavigationDiagnostics__destroy((*C.fog_msgs__msg__NavigationDiagnostics)(pointer_to_free))
}

func (t _NavigationDiagnosticsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NavigationDiagnostics)
	mem := (*C.fog_msgs__msg__NavigationDiagnostics)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.state), m.State)
	mem.waypoints_in_buffer = C.uint8_t(m.WaypointsInBuffer)
	mem.current_waypoint_id = C.uint8_t(m.CurrentWaypointId)
	mem.bumper_active = C.bool(m.BumperActive)
	cSlice_last_nav_goal := mem.last_nav_goal[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_last_nav_goal)), m.LastNavGoal[:])
	cSlice_current_nav_goal := mem.current_nav_goal[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_current_nav_goal)), m.CurrentNavGoal[:])
}

func (t _NavigationDiagnosticsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NavigationDiagnostics)
	mem := (*C.fog_msgs__msg__NavigationDiagnostics)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.State, unsafe.Pointer(&mem.state))
	m.WaypointsInBuffer = uint8(mem.waypoints_in_buffer)
	m.CurrentWaypointId = uint8(mem.current_waypoint_id)
	m.BumperActive = bool(mem.bumper_active)
	cSlice_last_nav_goal := mem.last_nav_goal[:]
	primitives.Float64__Array_to_Go(m.LastNavGoal[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_last_nav_goal)))
	cSlice_current_nav_goal := mem.current_nav_goal[:]
	primitives.Float64__Array_to_Go(m.CurrentNavGoal[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_current_nav_goal)))
}

func (t _NavigationDiagnosticsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__msg__NavigationDiagnostics())
}

type CNavigationDiagnostics = C.fog_msgs__msg__NavigationDiagnostics
type CNavigationDiagnostics__Sequence = C.fog_msgs__msg__NavigationDiagnostics__Sequence

func NavigationDiagnostics__Sequence_to_Go(goSlice *[]NavigationDiagnostics, cSlice CNavigationDiagnostics__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NavigationDiagnostics, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__msg__NavigationDiagnostics__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__NavigationDiagnostics * uintptr(i)),
		))
		NavigationDiagnosticsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func NavigationDiagnostics__Sequence_to_C(cSlice *CNavigationDiagnostics__Sequence, goSlice []NavigationDiagnostics) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__msg__NavigationDiagnostics)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__msg__NavigationDiagnostics * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__msg__NavigationDiagnostics)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__NavigationDiagnostics * uintptr(i)),
		))
		NavigationDiagnosticsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func NavigationDiagnostics__Array_to_Go(goSlice []NavigationDiagnostics, cSlice []CNavigationDiagnostics) {
	for i := 0; i < len(cSlice); i++ {
		NavigationDiagnosticsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NavigationDiagnostics__Array_to_C(cSlice []CNavigationDiagnostics, goSlice []NavigationDiagnostics) {
	for i := 0; i < len(goSlice); i++ {
		NavigationDiagnosticsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
