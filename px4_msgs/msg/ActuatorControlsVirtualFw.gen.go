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
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/actuator_controls_virtual_fw.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControlsVirtualFw", ActuatorControlsVirtualFwTypeSupport)
}
const (
	ActuatorControlsVirtualFw_NUM_ACTUATOR_CONTROLS uint8 = 8
	ActuatorControlsVirtualFw_NUM_ACTUATOR_CONTROL_GROUPS uint8 = 6
	ActuatorControlsVirtualFw_INDEX_ROLL uint8 = 0
	ActuatorControlsVirtualFw_INDEX_PITCH uint8 = 1
	ActuatorControlsVirtualFw_INDEX_YAW uint8 = 2
	ActuatorControlsVirtualFw_INDEX_THROTTLE uint8 = 3
	ActuatorControlsVirtualFw_INDEX_FLAPS uint8 = 4
	ActuatorControlsVirtualFw_INDEX_SPOILERS uint8 = 5
	ActuatorControlsVirtualFw_INDEX_AIRBRAKES uint8 = 6
	ActuatorControlsVirtualFw_INDEX_LANDING_GEAR uint8 = 7
	ActuatorControlsVirtualFw_INDEX_GIMBAL_SHUTTER uint8 = 3
	ActuatorControlsVirtualFw_INDEX_CAMERA_ZOOM uint8 = 4
	ActuatorControlsVirtualFw_GROUP_INDEX_ATTITUDE uint8 = 0
	ActuatorControlsVirtualFw_GROUP_INDEX_ATTITUDE_ALTERNATE uint8 = 1
	ActuatorControlsVirtualFw_GROUP_INDEX_GIMBAL uint8 = 2
	ActuatorControlsVirtualFw_GROUP_INDEX_MANUAL_PASSTHROUGH uint8 = 3
	ActuatorControlsVirtualFw_GROUP_INDEX_ALLOCATED_PART1 uint8 = 4
	ActuatorControlsVirtualFw_GROUP_INDEX_ALLOCATED_PART2 uint8 = 5
	ActuatorControlsVirtualFw_GROUP_INDEX_PAYLOAD uint8 = 6
)

// Do not create instances of this type directly. Always use NewActuatorControlsVirtualFw
// function instead.
type ActuatorControlsVirtualFw struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp the data this control response is based on was sampled
	Control [8]float32 `yaml:"control"`
}

// NewActuatorControlsVirtualFw creates a new ActuatorControlsVirtualFw with default values.
func NewActuatorControlsVirtualFw() *ActuatorControlsVirtualFw {
	self := ActuatorControlsVirtualFw{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControlsVirtualFw) Clone() *ActuatorControlsVirtualFw {
	c := &ActuatorControlsVirtualFw{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Control = t.Control
	return c
}

func (t *ActuatorControlsVirtualFw) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControlsVirtualFw) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Control = [8]float32{}
}

// CloneActuatorControlsVirtualFwSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControlsVirtualFwSlice(dst, src []ActuatorControlsVirtualFw) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControlsVirtualFwTypeSupport types.MessageTypeSupport = _ActuatorControlsVirtualFwTypeSupport{}

type _ActuatorControlsVirtualFwTypeSupport struct{}

func (t _ActuatorControlsVirtualFwTypeSupport) New() types.Message {
	return NewActuatorControlsVirtualFw()
}

func (t _ActuatorControlsVirtualFwTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControlsVirtualFw
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControlsVirtualFw__create())
}

func (t _ActuatorControlsVirtualFwTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControlsVirtualFw__destroy((*C.px4_msgs__msg__ActuatorControlsVirtualFw)(pointer_to_free))
}

func (t _ActuatorControlsVirtualFwTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControlsVirtualFw)
	mem := (*C.px4_msgs__msg__ActuatorControlsVirtualFw)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)), m.Control[:])
}

func (t _ActuatorControlsVirtualFwTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControlsVirtualFw)
	mem := (*C.px4_msgs__msg__ActuatorControlsVirtualFw)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_Go(m.Control[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)))
}

func (t _ActuatorControlsVirtualFwTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControlsVirtualFw())
}

type CActuatorControlsVirtualFw = C.px4_msgs__msg__ActuatorControlsVirtualFw
type CActuatorControlsVirtualFw__Sequence = C.px4_msgs__msg__ActuatorControlsVirtualFw__Sequence

func ActuatorControlsVirtualFw__Sequence_to_Go(goSlice *[]ActuatorControlsVirtualFw, cSlice CActuatorControlsVirtualFw__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControlsVirtualFw, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsVirtualFw__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsVirtualFw * uintptr(i)),
		))
		ActuatorControlsVirtualFwTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControlsVirtualFw__Sequence_to_C(cSlice *CActuatorControlsVirtualFw__Sequence, goSlice []ActuatorControlsVirtualFw) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControlsVirtualFw)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControlsVirtualFw * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsVirtualFw)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsVirtualFw * uintptr(i)),
		))
		ActuatorControlsVirtualFwTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControlsVirtualFw__Array_to_Go(goSlice []ActuatorControlsVirtualFw, cSlice []CActuatorControlsVirtualFw) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControlsVirtualFwTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControlsVirtualFw__Array_to_C(cSlice []CActuatorControlsVirtualFw, goSlice []ActuatorControlsVirtualFw) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControlsVirtualFwTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
