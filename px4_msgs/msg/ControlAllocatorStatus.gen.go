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

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/control_allocator_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ControlAllocatorStatus", ControlAllocatorStatusTypeSupport)
}
const (
	ControlAllocatorStatus_ACTUATOR_SATURATION_OK int8 = 0// The actuator is not saturated
	ControlAllocatorStatus_ACTUATOR_SATURATION_UPPER_DYN int8 = 1// The actuator is saturated (with a value <= the desired value) because it cannot increase its value faster
	ControlAllocatorStatus_ACTUATOR_SATURATION_UPPER int8 = 2// The actuator is saturated (with a value <= the desired value) because it has reached its maximum value
	ControlAllocatorStatus_ACTUATOR_SATURATION_LOWER_DYN int8 = -1// The actuator is saturated (with a value >= the desired value) because it cannot decrease its value faster
	ControlAllocatorStatus_ACTUATOR_SATURATION_LOWER int8 = -2// The actuator is saturated (with a value >= the desired value) because it has reached its minimum value
)

// Do not create instances of this type directly. Always use NewControlAllocatorStatus
// function instead.
type ControlAllocatorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TorqueSetpointAchieved bool `yaml:"torque_setpoint_achieved"`// Boolean indicating whether the 3D torque setpoint was correctly allocated to actuators. 0 if not achieved, 1 if achieved.
	AllocatedTorque [3]float32 `yaml:"allocated_torque"`// Torque allocated to actuators. Equal to `vehicle_torque_setpoint_s::xyz` if the setpoint was achieved.
	UnallocatedTorque [3]float32 `yaml:"unallocated_torque"`// Unallocated torque. Equal to 0 if the setpoint was achieved.
	ThrustSetpointAchieved bool `yaml:"thrust_setpoint_achieved"`// Boolean indicating whether the 3D thrust setpoint was correctly allocated to actuators. 0 if not achieved, 1 if achieved.
	AllocatedThrust [3]float32 `yaml:"allocated_thrust"`// Thrust allocated to actuators. Equal to `vehicle_thrust_setpoint_s::xyz` if the setpoint was achieved.
	UnallocatedThrust [3]float32 `yaml:"unallocated_thrust"`// Unallocated thrust. Equal to 0 if the setpoint was achieved.
	ActuatorSaturation [16]int8 `yaml:"actuator_saturation"`// Indicates actuator saturation status.
}

// NewControlAllocatorStatus creates a new ControlAllocatorStatus with default values.
func NewControlAllocatorStatus() *ControlAllocatorStatus {
	self := ControlAllocatorStatus{}
	self.SetDefaults()
	return &self
}

func (t *ControlAllocatorStatus) Clone() *ControlAllocatorStatus {
	c := &ControlAllocatorStatus{}
	c.Timestamp = t.Timestamp
	c.TorqueSetpointAchieved = t.TorqueSetpointAchieved
	c.AllocatedTorque = t.AllocatedTorque
	c.UnallocatedTorque = t.UnallocatedTorque
	c.ThrustSetpointAchieved = t.ThrustSetpointAchieved
	c.AllocatedThrust = t.AllocatedThrust
	c.UnallocatedThrust = t.UnallocatedThrust
	c.ActuatorSaturation = t.ActuatorSaturation
	return c
}

func (t *ControlAllocatorStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ControlAllocatorStatus) SetDefaults() {
	t.Timestamp = 0
	t.TorqueSetpointAchieved = false
	t.AllocatedTorque = [3]float32{}
	t.UnallocatedTorque = [3]float32{}
	t.ThrustSetpointAchieved = false
	t.AllocatedThrust = [3]float32{}
	t.UnallocatedThrust = [3]float32{}
	t.ActuatorSaturation = [16]int8{}
}

// CloneControlAllocatorStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneControlAllocatorStatusSlice(dst, src []ControlAllocatorStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ControlAllocatorStatusTypeSupport types.MessageTypeSupport = _ControlAllocatorStatusTypeSupport{}

type _ControlAllocatorStatusTypeSupport struct{}

func (t _ControlAllocatorStatusTypeSupport) New() types.Message {
	return NewControlAllocatorStatus()
}

func (t _ControlAllocatorStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ControlAllocatorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__ControlAllocatorStatus__create())
}

func (t _ControlAllocatorStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ControlAllocatorStatus__destroy((*C.px4_msgs__msg__ControlAllocatorStatus)(pointer_to_free))
}

func (t _ControlAllocatorStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ControlAllocatorStatus)
	mem := (*C.px4_msgs__msg__ControlAllocatorStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.torque_setpoint_achieved = C.bool(m.TorqueSetpointAchieved)
	cSlice_allocated_torque := mem.allocated_torque[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_allocated_torque)), m.AllocatedTorque[:])
	cSlice_unallocated_torque := mem.unallocated_torque[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_unallocated_torque)), m.UnallocatedTorque[:])
	mem.thrust_setpoint_achieved = C.bool(m.ThrustSetpointAchieved)
	cSlice_allocated_thrust := mem.allocated_thrust[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_allocated_thrust)), m.AllocatedThrust[:])
	cSlice_unallocated_thrust := mem.unallocated_thrust[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_unallocated_thrust)), m.UnallocatedThrust[:])
	cSlice_actuator_saturation := mem.actuator_saturation[:]
	primitives.Int8__Array_to_C(*(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_actuator_saturation)), m.ActuatorSaturation[:])
}

func (t _ControlAllocatorStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ControlAllocatorStatus)
	mem := (*C.px4_msgs__msg__ControlAllocatorStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TorqueSetpointAchieved = bool(mem.torque_setpoint_achieved)
	cSlice_allocated_torque := mem.allocated_torque[:]
	primitives.Float32__Array_to_Go(m.AllocatedTorque[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_allocated_torque)))
	cSlice_unallocated_torque := mem.unallocated_torque[:]
	primitives.Float32__Array_to_Go(m.UnallocatedTorque[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_unallocated_torque)))
	m.ThrustSetpointAchieved = bool(mem.thrust_setpoint_achieved)
	cSlice_allocated_thrust := mem.allocated_thrust[:]
	primitives.Float32__Array_to_Go(m.AllocatedThrust[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_allocated_thrust)))
	cSlice_unallocated_thrust := mem.unallocated_thrust[:]
	primitives.Float32__Array_to_Go(m.UnallocatedThrust[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_unallocated_thrust)))
	cSlice_actuator_saturation := mem.actuator_saturation[:]
	primitives.Int8__Array_to_Go(m.ActuatorSaturation[:], *(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_actuator_saturation)))
}

func (t _ControlAllocatorStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ControlAllocatorStatus())
}

type CControlAllocatorStatus = C.px4_msgs__msg__ControlAllocatorStatus
type CControlAllocatorStatus__Sequence = C.px4_msgs__msg__ControlAllocatorStatus__Sequence

func ControlAllocatorStatus__Sequence_to_Go(goSlice *[]ControlAllocatorStatus, cSlice CControlAllocatorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ControlAllocatorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ControlAllocatorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ControlAllocatorStatus * uintptr(i)),
		))
		ControlAllocatorStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ControlAllocatorStatus__Sequence_to_C(cSlice *CControlAllocatorStatus__Sequence, goSlice []ControlAllocatorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ControlAllocatorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ControlAllocatorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ControlAllocatorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ControlAllocatorStatus * uintptr(i)),
		))
		ControlAllocatorStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ControlAllocatorStatus__Array_to_Go(goSlice []ControlAllocatorStatus, cSlice []CControlAllocatorStatus) {
	for i := 0; i < len(cSlice); i++ {
		ControlAllocatorStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ControlAllocatorStatus__Array_to_C(cSlice []CControlAllocatorStatus, goSlice []ControlAllocatorStatus) {
	for i := 0; i < len(goSlice); i++ {
		ControlAllocatorStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
