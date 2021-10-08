/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/geofence_result.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/GeofenceResult", GeofenceResultTypeSupport)
}
const (
	GeofenceResult_GF_ACTION_NONE uint8 = 0// no action on geofence violation
	GeofenceResult_GF_ACTION_WARN uint8 = 1// critical mavlink message
	GeofenceResult_GF_ACTION_LOITER uint8 = 2// switch to AUTO|LOITER
	GeofenceResult_GF_ACTION_RTL uint8 = 3// switch to AUTO|RTL
	GeofenceResult_GF_ACTION_TERMINATE uint8 = 4// flight termination
	GeofenceResult_GF_ACTION_LAND uint8 = 5// switch to AUTO|LAND
)

// Do not create instances of this type directly. Always use NewGeofenceResult
// function instead.
type GeofenceResult struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	GeofenceViolated bool `yaml:"geofence_violated"`// true if the geofence is violated
	GeofenceAction uint8 `yaml:"geofence_action"`// action to take when geofence is violated
	HomeRequired bool `yaml:"home_required"`// true if the geofence requires a valid home position
}

// NewGeofenceResult creates a new GeofenceResult with default values.
func NewGeofenceResult() *GeofenceResult {
	self := GeofenceResult{}
	self.SetDefaults()
	return &self
}

func (t *GeofenceResult) Clone() *GeofenceResult {
	c := &GeofenceResult{}
	c.Timestamp = t.Timestamp
	c.GeofenceViolated = t.GeofenceViolated
	c.GeofenceAction = t.GeofenceAction
	c.HomeRequired = t.HomeRequired
	return c
}

func (t *GeofenceResult) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GeofenceResult) SetDefaults() {
	t.Timestamp = 0
	t.GeofenceViolated = false
	t.GeofenceAction = 0
	t.HomeRequired = false
}

// CloneGeofenceResultSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGeofenceResultSlice(dst, src []GeofenceResult) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GeofenceResultTypeSupport types.MessageTypeSupport = _GeofenceResultTypeSupport{}

type _GeofenceResultTypeSupport struct{}

func (t _GeofenceResultTypeSupport) New() types.Message {
	return NewGeofenceResult()
}

func (t _GeofenceResultTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GeofenceResult
	return (unsafe.Pointer)(C.px4_msgs__msg__GeofenceResult__create())
}

func (t _GeofenceResultTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GeofenceResult__destroy((*C.px4_msgs__msg__GeofenceResult)(pointer_to_free))
}

func (t _GeofenceResultTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GeofenceResult)
	mem := (*C.px4_msgs__msg__GeofenceResult)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.geofence_violated = C.bool(m.GeofenceViolated)
	mem.geofence_action = C.uint8_t(m.GeofenceAction)
	mem.home_required = C.bool(m.HomeRequired)
}

func (t _GeofenceResultTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GeofenceResult)
	mem := (*C.px4_msgs__msg__GeofenceResult)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.GeofenceViolated = bool(mem.geofence_violated)
	m.GeofenceAction = uint8(mem.geofence_action)
	m.HomeRequired = bool(mem.home_required)
}

func (t _GeofenceResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GeofenceResult())
}

type CGeofenceResult = C.px4_msgs__msg__GeofenceResult
type CGeofenceResult__Sequence = C.px4_msgs__msg__GeofenceResult__Sequence

func GeofenceResult__Sequence_to_Go(goSlice *[]GeofenceResult, cSlice CGeofenceResult__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GeofenceResult, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GeofenceResult__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GeofenceResult * uintptr(i)),
		))
		GeofenceResultTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GeofenceResult__Sequence_to_C(cSlice *CGeofenceResult__Sequence, goSlice []GeofenceResult) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GeofenceResult)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GeofenceResult * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GeofenceResult)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GeofenceResult * uintptr(i)),
		))
		GeofenceResultTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GeofenceResult__Array_to_Go(goSlice []GeofenceResult, cSlice []CGeofenceResult) {
	for i := 0; i < len(cSlice); i++ {
		GeofenceResultTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GeofenceResult__Array_to_C(cSlice []CGeofenceResult, goSlice []GeofenceResult) {
	for i := 0; i < len(goSlice); i++ {
		GeofenceResultTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
