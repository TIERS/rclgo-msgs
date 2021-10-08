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

#include <px4_msgs/msg/vehicle_angular_velocity_groundtruth.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleAngularVelocityGroundtruth", VehicleAngularVelocityGroundtruthTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleAngularVelocityGroundtruth
// function instead.
type VehicleAngularVelocityGroundtruth struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// timestamp of the data sample on which this message is based (microseconds)
	Xyz [3]float32 `yaml:"xyz"`// Bias corrected angular velocity about the FRD body frame XYZ-axis in rad/s
}

// NewVehicleAngularVelocityGroundtruth creates a new VehicleAngularVelocityGroundtruth with default values.
func NewVehicleAngularVelocityGroundtruth() *VehicleAngularVelocityGroundtruth {
	self := VehicleAngularVelocityGroundtruth{}
	self.SetDefaults()
	return &self
}

func (t *VehicleAngularVelocityGroundtruth) Clone() *VehicleAngularVelocityGroundtruth {
	c := &VehicleAngularVelocityGroundtruth{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Xyz = t.Xyz
	return c
}

func (t *VehicleAngularVelocityGroundtruth) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleAngularVelocityGroundtruth) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Xyz = [3]float32{}
}

// CloneVehicleAngularVelocityGroundtruthSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleAngularVelocityGroundtruthSlice(dst, src []VehicleAngularVelocityGroundtruth) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleAngularVelocityGroundtruthTypeSupport types.MessageTypeSupport = _VehicleAngularVelocityGroundtruthTypeSupport{}

type _VehicleAngularVelocityGroundtruthTypeSupport struct{}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) New() types.Message {
	return NewVehicleAngularVelocityGroundtruth()
}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleAngularVelocityGroundtruth
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleAngularVelocityGroundtruth__create())
}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleAngularVelocityGroundtruth__destroy((*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth)(pointer_to_free))
}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleAngularVelocityGroundtruth)
	mem := (*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)), m.Xyz[:])
}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleAngularVelocityGroundtruth)
	mem := (*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_xyz := mem.xyz[:]
	primitives.Float32__Array_to_Go(m.Xyz[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_xyz)))
}

func (t _VehicleAngularVelocityGroundtruthTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleAngularVelocityGroundtruth())
}

type CVehicleAngularVelocityGroundtruth = C.px4_msgs__msg__VehicleAngularVelocityGroundtruth
type CVehicleAngularVelocityGroundtruth__Sequence = C.px4_msgs__msg__VehicleAngularVelocityGroundtruth__Sequence

func VehicleAngularVelocityGroundtruth__Sequence_to_Go(goSlice *[]VehicleAngularVelocityGroundtruth, cSlice CVehicleAngularVelocityGroundtruth__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleAngularVelocityGroundtruth, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocityGroundtruth * uintptr(i)),
		))
		VehicleAngularVelocityGroundtruthTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleAngularVelocityGroundtruth__Sequence_to_C(cSlice *CVehicleAngularVelocityGroundtruth__Sequence, goSlice []VehicleAngularVelocityGroundtruth) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocityGroundtruth * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleAngularVelocityGroundtruth)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleAngularVelocityGroundtruth * uintptr(i)),
		))
		VehicleAngularVelocityGroundtruthTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleAngularVelocityGroundtruth__Array_to_Go(goSlice []VehicleAngularVelocityGroundtruth, cSlice []CVehicleAngularVelocityGroundtruth) {
	for i := 0; i < len(cSlice); i++ {
		VehicleAngularVelocityGroundtruthTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleAngularVelocityGroundtruth__Array_to_C(cSlice []CVehicleAngularVelocityGroundtruth, goSlice []VehicleAngularVelocityGroundtruth) {
	for i := 0; i < len(goSlice); i++ {
		VehicleAngularVelocityGroundtruthTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
