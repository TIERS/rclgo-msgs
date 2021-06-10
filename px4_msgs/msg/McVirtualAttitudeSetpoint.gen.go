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

#include <px4_msgs/msg/mc_virtual_attitude_setpoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/McVirtualAttitudeSetpoint", McVirtualAttitudeSetpointTypeSupport)
}
const (
	McVirtualAttitudeSetpoint_FLAPS_OFF uint8 = 0// no flaps
	McVirtualAttitudeSetpoint_FLAPS_LAND uint8 = 1// landing config flaps
	McVirtualAttitudeSetpoint_FLAPS_TAKEOFF uint8 = 2// take-off config flaps
)

// Do not create instances of this type directly. Always use NewMcVirtualAttitudeSetpoint
// function instead.
type McVirtualAttitudeSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	RollBody float32 `yaml:"roll_body"`// body angle in NED frame (can be NaN for FW)
	PitchBody float32 `yaml:"pitch_body"`// body angle in NED frame (can be NaN for FW)
	YawBody float32 `yaml:"yaw_body"`// body angle in NED frame (can be NaN for FW)
	YawSpMoveRate float32 `yaml:"yaw_sp_move_rate"`// rad/s (commanded by user)
	QD [4]float32 `yaml:"q_d"`// Desired quaternion for quaternion control. For quaternion-based attitude control
	ThrustBody [3]float32 `yaml:"thrust_body"`// Normalized thrust command in body NED frame [-1,1]. For clarification: For multicopters thrust_body[0] and thrust[1] are usually 0 and thrust[2] is the negative throttle demand.For fixed wings thrust_x is the throttle demand and thrust_y, thrust_z will usually be zero.
	RollResetIntegral bool `yaml:"roll_reset_integral"`// Reset roll integral part (navigation logic change)
	PitchResetIntegral bool `yaml:"pitch_reset_integral"`// Reset pitch integral part (navigation logic change)
	YawResetIntegral bool `yaml:"yaw_reset_integral"`// Reset yaw integral part (navigation logic change)
	FwControlYaw bool `yaml:"fw_control_yaw"`// control heading with rudder (used for auto takeoff on runway)
	ApplyFlaps uint8 `yaml:"apply_flaps"`// flap config specifier
}

// NewMcVirtualAttitudeSetpoint creates a new McVirtualAttitudeSetpoint with default values.
func NewMcVirtualAttitudeSetpoint() *McVirtualAttitudeSetpoint {
	self := McVirtualAttitudeSetpoint{}
	self.SetDefaults()
	return &self
}

func (t *McVirtualAttitudeSetpoint) Clone() *McVirtualAttitudeSetpoint {
	c := &McVirtualAttitudeSetpoint{}
	c.Timestamp = t.Timestamp
	c.RollBody = t.RollBody
	c.PitchBody = t.PitchBody
	c.YawBody = t.YawBody
	c.YawSpMoveRate = t.YawSpMoveRate
	c.QD = t.QD
	c.ThrustBody = t.ThrustBody
	c.RollResetIntegral = t.RollResetIntegral
	c.PitchResetIntegral = t.PitchResetIntegral
	c.YawResetIntegral = t.YawResetIntegral
	c.FwControlYaw = t.FwControlYaw
	c.ApplyFlaps = t.ApplyFlaps
	return c
}

func (t *McVirtualAttitudeSetpoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *McVirtualAttitudeSetpoint) SetDefaults() {
	t.Timestamp = 0
	t.RollBody = 0
	t.PitchBody = 0
	t.YawBody = 0
	t.YawSpMoveRate = 0
	t.QD = [4]float32{}
	t.ThrustBody = [3]float32{}
	t.RollResetIntegral = false
	t.PitchResetIntegral = false
	t.YawResetIntegral = false
	t.FwControlYaw = false
	t.ApplyFlaps = 0
}

// CloneMcVirtualAttitudeSetpointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMcVirtualAttitudeSetpointSlice(dst, src []McVirtualAttitudeSetpoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var McVirtualAttitudeSetpointTypeSupport types.MessageTypeSupport = _McVirtualAttitudeSetpointTypeSupport{}

type _McVirtualAttitudeSetpointTypeSupport struct{}

func (t _McVirtualAttitudeSetpointTypeSupport) New() types.Message {
	return NewMcVirtualAttitudeSetpoint()
}

func (t _McVirtualAttitudeSetpointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__McVirtualAttitudeSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__McVirtualAttitudeSetpoint__create())
}

func (t _McVirtualAttitudeSetpointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__McVirtualAttitudeSetpoint__destroy((*C.px4_msgs__msg__McVirtualAttitudeSetpoint)(pointer_to_free))
}

func (t _McVirtualAttitudeSetpointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*McVirtualAttitudeSetpoint)
	mem := (*C.px4_msgs__msg__McVirtualAttitudeSetpoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.roll_body = C.float(m.RollBody)
	mem.pitch_body = C.float(m.PitchBody)
	mem.yaw_body = C.float(m.YawBody)
	mem.yaw_sp_move_rate = C.float(m.YawSpMoveRate)
	cSlice_q_d := mem.q_d[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q_d)), m.QD[:])
	cSlice_thrust_body := mem.thrust_body[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)), m.ThrustBody[:])
	mem.roll_reset_integral = C.bool(m.RollResetIntegral)
	mem.pitch_reset_integral = C.bool(m.PitchResetIntegral)
	mem.yaw_reset_integral = C.bool(m.YawResetIntegral)
	mem.fw_control_yaw = C.bool(m.FwControlYaw)
	mem.apply_flaps = C.uint8_t(m.ApplyFlaps)
}

func (t _McVirtualAttitudeSetpointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*McVirtualAttitudeSetpoint)
	mem := (*C.px4_msgs__msg__McVirtualAttitudeSetpoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.RollBody = float32(mem.roll_body)
	m.PitchBody = float32(mem.pitch_body)
	m.YawBody = float32(mem.yaw_body)
	m.YawSpMoveRate = float32(mem.yaw_sp_move_rate)
	cSlice_q_d := mem.q_d[:]
	primitives.Float32__Array_to_Go(m.QD[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q_d)))
	cSlice_thrust_body := mem.thrust_body[:]
	primitives.Float32__Array_to_Go(m.ThrustBody[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_thrust_body)))
	m.RollResetIntegral = bool(mem.roll_reset_integral)
	m.PitchResetIntegral = bool(mem.pitch_reset_integral)
	m.YawResetIntegral = bool(mem.yaw_reset_integral)
	m.FwControlYaw = bool(mem.fw_control_yaw)
	m.ApplyFlaps = uint8(mem.apply_flaps)
}

func (t _McVirtualAttitudeSetpointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__McVirtualAttitudeSetpoint())
}

type CMcVirtualAttitudeSetpoint = C.px4_msgs__msg__McVirtualAttitudeSetpoint
type CMcVirtualAttitudeSetpoint__Sequence = C.px4_msgs__msg__McVirtualAttitudeSetpoint__Sequence

func McVirtualAttitudeSetpoint__Sequence_to_Go(goSlice *[]McVirtualAttitudeSetpoint, cSlice CMcVirtualAttitudeSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]McVirtualAttitudeSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__McVirtualAttitudeSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__McVirtualAttitudeSetpoint * uintptr(i)),
		))
		McVirtualAttitudeSetpointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func McVirtualAttitudeSetpoint__Sequence_to_C(cSlice *CMcVirtualAttitudeSetpoint__Sequence, goSlice []McVirtualAttitudeSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__McVirtualAttitudeSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__McVirtualAttitudeSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__McVirtualAttitudeSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__McVirtualAttitudeSetpoint * uintptr(i)),
		))
		McVirtualAttitudeSetpointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func McVirtualAttitudeSetpoint__Array_to_Go(goSlice []McVirtualAttitudeSetpoint, cSlice []CMcVirtualAttitudeSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		McVirtualAttitudeSetpointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func McVirtualAttitudeSetpoint__Array_to_C(cSlice []CMcVirtualAttitudeSetpoint, goSlice []McVirtualAttitudeSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		McVirtualAttitudeSetpointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
