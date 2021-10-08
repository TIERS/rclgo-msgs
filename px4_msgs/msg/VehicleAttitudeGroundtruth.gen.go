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

#include <px4_msgs/msg/vehicle_attitude_groundtruth.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleAttitudeGroundtruth", VehicleAttitudeGroundtruthTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleAttitudeGroundtruth
// function instead.
type VehicleAttitudeGroundtruth struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	Q [4]float32 `yaml:"q"`// Quaternion rotation from the FRD body frame to the NED earth frame
	DeltaQReset [4]float32 `yaml:"delta_q_reset"`// Amount by which quaternion has changed during last reset
	QuatResetCounter uint8 `yaml:"quat_reset_counter"`// Quaternion reset counter
}

// NewVehicleAttitudeGroundtruth creates a new VehicleAttitudeGroundtruth with default values.
func NewVehicleAttitudeGroundtruth() *VehicleAttitudeGroundtruth {
	self := VehicleAttitudeGroundtruth{}
	self.SetDefaults()
	return &self
}

func (t *VehicleAttitudeGroundtruth) Clone() *VehicleAttitudeGroundtruth {
	c := &VehicleAttitudeGroundtruth{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Q = t.Q
	c.DeltaQReset = t.DeltaQReset
	c.QuatResetCounter = t.QuatResetCounter
	return c
}

func (t *VehicleAttitudeGroundtruth) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleAttitudeGroundtruth) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Q = [4]float32{}
	t.DeltaQReset = [4]float32{}
	t.QuatResetCounter = 0
}

// CloneVehicleAttitudeGroundtruthSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleAttitudeGroundtruthSlice(dst, src []VehicleAttitudeGroundtruth) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleAttitudeGroundtruthTypeSupport types.MessageTypeSupport = _VehicleAttitudeGroundtruthTypeSupport{}

type _VehicleAttitudeGroundtruthTypeSupport struct{}

func (t _VehicleAttitudeGroundtruthTypeSupport) New() types.Message {
	return NewVehicleAttitudeGroundtruth()
}

func (t _VehicleAttitudeGroundtruthTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAttitudeGroundtruth
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAttitudeGroundtruth__create())
}

func (t _VehicleAttitudeGroundtruthTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAttitudeGroundtruth__destroy((*C.px4_msgs__msg__VehicleAttitudeGroundtruth)(pointer_to_free))
}

func (t _VehicleAttitudeGroundtruthTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleAttitudeGroundtruth)
	mem := (*C.px4_msgs__msg__VehicleAttitudeGroundtruth)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)), m.DeltaQReset[:])
	mem.quat_reset_counter = C.uint8_t(m.QuatResetCounter)
}

func (t _VehicleAttitudeGroundtruthTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleAttitudeGroundtruth)
	mem := (*C.px4_msgs__msg__VehicleAttitudeGroundtruth)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	cSlice_delta_q_reset := mem.delta_q_reset[:]
	primitives.Float32__Array_to_Go(m.DeltaQReset[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_q_reset)))
	m.QuatResetCounter = uint8(mem.quat_reset_counter)
}

func (t _VehicleAttitudeGroundtruthTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAttitudeGroundtruth())
}

type CVehicleAttitudeGroundtruth = C.px4_msgs__msg__VehicleAttitudeGroundtruth
type CVehicleAttitudeGroundtruth__Sequence = C.px4_msgs__msg__VehicleAttitudeGroundtruth__Sequence

func VehicleAttitudeGroundtruth__Sequence_to_Go(goSlice *[]VehicleAttitudeGroundtruth, cSlice CVehicleAttitudeGroundtruth__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAttitudeGroundtruth, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAttitudeGroundtruth__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitudeGroundtruth * uintptr(i)),
		))
		VehicleAttitudeGroundtruthTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleAttitudeGroundtruth__Sequence_to_C(cSlice *CVehicleAttitudeGroundtruth__Sequence, goSlice []VehicleAttitudeGroundtruth) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAttitudeGroundtruth)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAttitudeGroundtruth * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAttitudeGroundtruth)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAttitudeGroundtruth * uintptr(i)),
		))
		VehicleAttitudeGroundtruthTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleAttitudeGroundtruth__Array_to_Go(goSlice []VehicleAttitudeGroundtruth, cSlice []CVehicleAttitudeGroundtruth) {
	for i := 0; i < len(cSlice); i++ {
		VehicleAttitudeGroundtruthTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAttitudeGroundtruth__Array_to_C(cSlice []CVehicleAttitudeGroundtruth, goSlice []VehicleAttitudeGroundtruth) {
	for i := 0; i < len(goSlice); i++ {
		VehicleAttitudeGroundtruthTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
