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

#include <px4_msgs/msg/gimbal_manager_set_manual_control.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/GimbalManagerSetManualControl", GimbalManagerSetManualControlTypeSupport)
}
const (
	GimbalManagerSetManualControl_GIMBAL_MANAGER_FLAGS_RETRACT uint32 = 1
	GimbalManagerSetManualControl_GIMBAL_MANAGER_FLAGS_NEUTRAL uint32 = 2
	GimbalManagerSetManualControl_GIMBAL_MANAGER_FLAGS_ROLL_LOCK uint32 = 4
	GimbalManagerSetManualControl_GIMBAL_MANAGER_FLAGS_PITCH_LOCK uint32 = 8
	GimbalManagerSetManualControl_GIMBAL_MANAGER_FLAGS_YAW_LOCK uint32 = 16
)

// Do not create instances of this type directly. Always use NewGimbalManagerSetManualControl
// function instead.
type GimbalManagerSetManualControl struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	OriginSysid uint8 `yaml:"origin_sysid"`
	OriginCompid uint8 `yaml:"origin_compid"`
	TargetSystem uint8 `yaml:"target_system"`
	TargetComponent uint8 `yaml:"target_component"`
	Flags uint32 `yaml:"flags"`
	GimbalDeviceId uint8 `yaml:"gimbal_device_id"`
	Pitch float32 `yaml:"pitch"`// unitless -1..1, can be NAN
	Yaw float32 `yaml:"yaw"`// unitless -1..1, can be NAN
	PitchRate float32 `yaml:"pitch_rate"`// unitless -1..1, can be NAN
	YawRate float32 `yaml:"yaw_rate"`// unitless -1..1, can be NAN
}

// NewGimbalManagerSetManualControl creates a new GimbalManagerSetManualControl with default values.
func NewGimbalManagerSetManualControl() *GimbalManagerSetManualControl {
	self := GimbalManagerSetManualControl{}
	self.SetDefaults()
	return &self
}

func (t *GimbalManagerSetManualControl) Clone() *GimbalManagerSetManualControl {
	c := &GimbalManagerSetManualControl{}
	c.Timestamp = t.Timestamp
	c.OriginSysid = t.OriginSysid
	c.OriginCompid = t.OriginCompid
	c.TargetSystem = t.TargetSystem
	c.TargetComponent = t.TargetComponent
	c.Flags = t.Flags
	c.GimbalDeviceId = t.GimbalDeviceId
	c.Pitch = t.Pitch
	c.Yaw = t.Yaw
	c.PitchRate = t.PitchRate
	c.YawRate = t.YawRate
	return c
}

func (t *GimbalManagerSetManualControl) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GimbalManagerSetManualControl) SetDefaults() {
	t.Timestamp = 0
	t.OriginSysid = 0
	t.OriginCompid = 0
	t.TargetSystem = 0
	t.TargetComponent = 0
	t.Flags = 0
	t.GimbalDeviceId = 0
	t.Pitch = 0
	t.Yaw = 0
	t.PitchRate = 0
	t.YawRate = 0
}

// CloneGimbalManagerSetManualControlSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGimbalManagerSetManualControlSlice(dst, src []GimbalManagerSetManualControl) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GimbalManagerSetManualControlTypeSupport types.MessageTypeSupport = _GimbalManagerSetManualControlTypeSupport{}

type _GimbalManagerSetManualControlTypeSupport struct{}

func (t _GimbalManagerSetManualControlTypeSupport) New() types.Message {
	return NewGimbalManagerSetManualControl()
}

func (t _GimbalManagerSetManualControlTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GimbalManagerSetManualControl
	return (unsafe.Pointer)(C.px4_msgs__msg__GimbalManagerSetManualControl__create())
}

func (t _GimbalManagerSetManualControlTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GimbalManagerSetManualControl__destroy((*C.px4_msgs__msg__GimbalManagerSetManualControl)(pointer_to_free))
}

func (t _GimbalManagerSetManualControlTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GimbalManagerSetManualControl)
	mem := (*C.px4_msgs__msg__GimbalManagerSetManualControl)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.origin_sysid = C.uint8_t(m.OriginSysid)
	mem.origin_compid = C.uint8_t(m.OriginCompid)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
	mem.flags = C.uint32_t(m.Flags)
	mem.gimbal_device_id = C.uint8_t(m.GimbalDeviceId)
	mem.pitch = C.float(m.Pitch)
	mem.yaw = C.float(m.Yaw)
	mem.pitch_rate = C.float(m.PitchRate)
	mem.yaw_rate = C.float(m.YawRate)
}

func (t _GimbalManagerSetManualControlTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GimbalManagerSetManualControl)
	mem := (*C.px4_msgs__msg__GimbalManagerSetManualControl)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.OriginSysid = uint8(mem.origin_sysid)
	m.OriginCompid = uint8(mem.origin_compid)
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
	m.Flags = uint32(mem.flags)
	m.GimbalDeviceId = uint8(mem.gimbal_device_id)
	m.Pitch = float32(mem.pitch)
	m.Yaw = float32(mem.yaw)
	m.PitchRate = float32(mem.pitch_rate)
	m.YawRate = float32(mem.yaw_rate)
}

func (t _GimbalManagerSetManualControlTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GimbalManagerSetManualControl())
}

type CGimbalManagerSetManualControl = C.px4_msgs__msg__GimbalManagerSetManualControl
type CGimbalManagerSetManualControl__Sequence = C.px4_msgs__msg__GimbalManagerSetManualControl__Sequence

func GimbalManagerSetManualControl__Sequence_to_Go(goSlice *[]GimbalManagerSetManualControl, cSlice CGimbalManagerSetManualControl__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalManagerSetManualControl, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GimbalManagerSetManualControl__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerSetManualControl * uintptr(i)),
		))
		GimbalManagerSetManualControlTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GimbalManagerSetManualControl__Sequence_to_C(cSlice *CGimbalManagerSetManualControl__Sequence, goSlice []GimbalManagerSetManualControl) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GimbalManagerSetManualControl)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GimbalManagerSetManualControl * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GimbalManagerSetManualControl)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerSetManualControl * uintptr(i)),
		))
		GimbalManagerSetManualControlTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GimbalManagerSetManualControl__Array_to_Go(goSlice []GimbalManagerSetManualControl, cSlice []CGimbalManagerSetManualControl) {
	for i := 0; i < len(cSlice); i++ {
		GimbalManagerSetManualControlTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalManagerSetManualControl__Array_to_C(cSlice []CGimbalManagerSetManualControl, goSlice []GimbalManagerSetManualControl) {
	for i := 0; i < len(goSlice); i++ {
		GimbalManagerSetManualControlTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
