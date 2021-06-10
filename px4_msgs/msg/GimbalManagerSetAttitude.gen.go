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

#include <px4_msgs/msg/gimbal_manager_set_attitude.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/GimbalManagerSetAttitude", GimbalManagerSetAttitudeTypeSupport)
}
const (
	GimbalManagerSetAttitude_GIMBAL_MANAGER_FLAGS_RETRACT uint32 = 1
	GimbalManagerSetAttitude_GIMBAL_MANAGER_FLAGS_NEUTRAL uint32 = 2
	GimbalManagerSetAttitude_GIMBAL_MANAGER_FLAGS_ROLL_LOCK uint32 = 4
	GimbalManagerSetAttitude_GIMBAL_MANAGER_FLAGS_PITCH_LOCK uint32 = 8
	GimbalManagerSetAttitude_GIMBAL_MANAGER_FLAGS_YAW_LOCK uint32 = 16
)

// Do not create instances of this type directly. Always use NewGimbalManagerSetAttitude
// function instead.
type GimbalManagerSetAttitude struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	OriginSysid uint8 `yaml:"origin_sysid"`
	OriginCompid uint8 `yaml:"origin_compid"`
	TargetSystem uint8 `yaml:"target_system"`
	TargetComponent uint8 `yaml:"target_component"`
	Flags uint32 `yaml:"flags"`
	GimbalDeviceId uint8 `yaml:"gimbal_device_id"`
	Q [4]float32 `yaml:"q"`
	AngularVelocityX float32 `yaml:"angular_velocity_x"`
	AngularVelocityY float32 `yaml:"angular_velocity_y"`
	AngularVelocityZ float32 `yaml:"angular_velocity_z"`
}

// NewGimbalManagerSetAttitude creates a new GimbalManagerSetAttitude with default values.
func NewGimbalManagerSetAttitude() *GimbalManagerSetAttitude {
	self := GimbalManagerSetAttitude{}
	self.SetDefaults()
	return &self
}

func (t *GimbalManagerSetAttitude) Clone() *GimbalManagerSetAttitude {
	c := &GimbalManagerSetAttitude{}
	c.Timestamp = t.Timestamp
	c.OriginSysid = t.OriginSysid
	c.OriginCompid = t.OriginCompid
	c.TargetSystem = t.TargetSystem
	c.TargetComponent = t.TargetComponent
	c.Flags = t.Flags
	c.GimbalDeviceId = t.GimbalDeviceId
	c.Q = t.Q
	c.AngularVelocityX = t.AngularVelocityX
	c.AngularVelocityY = t.AngularVelocityY
	c.AngularVelocityZ = t.AngularVelocityZ
	return c
}

func (t *GimbalManagerSetAttitude) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GimbalManagerSetAttitude) SetDefaults() {
	t.Timestamp = 0
	t.OriginSysid = 0
	t.OriginCompid = 0
	t.TargetSystem = 0
	t.TargetComponent = 0
	t.Flags = 0
	t.GimbalDeviceId = 0
	t.Q = [4]float32{}
	t.AngularVelocityX = 0
	t.AngularVelocityY = 0
	t.AngularVelocityZ = 0
}

// CloneGimbalManagerSetAttitudeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGimbalManagerSetAttitudeSlice(dst, src []GimbalManagerSetAttitude) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GimbalManagerSetAttitudeTypeSupport types.MessageTypeSupport = _GimbalManagerSetAttitudeTypeSupport{}

type _GimbalManagerSetAttitudeTypeSupport struct{}

func (t _GimbalManagerSetAttitudeTypeSupport) New() types.Message {
	return NewGimbalManagerSetAttitude()
}

func (t _GimbalManagerSetAttitudeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__GimbalManagerSetAttitude
	return (unsafe.Pointer)(C.px4_msgs__msg__GimbalManagerSetAttitude__create())
}

func (t _GimbalManagerSetAttitudeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__GimbalManagerSetAttitude__destroy((*C.px4_msgs__msg__GimbalManagerSetAttitude)(pointer_to_free))
}

func (t _GimbalManagerSetAttitudeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GimbalManagerSetAttitude)
	mem := (*C.px4_msgs__msg__GimbalManagerSetAttitude)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.origin_sysid = C.uint8_t(m.OriginSysid)
	mem.origin_compid = C.uint8_t(m.OriginCompid)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
	mem.flags = C.uint32_t(m.Flags)
	mem.gimbal_device_id = C.uint8_t(m.GimbalDeviceId)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)), m.Q[:])
	mem.angular_velocity_x = C.float(m.AngularVelocityX)
	mem.angular_velocity_y = C.float(m.AngularVelocityY)
	mem.angular_velocity_z = C.float(m.AngularVelocityZ)
}

func (t _GimbalManagerSetAttitudeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GimbalManagerSetAttitude)
	mem := (*C.px4_msgs__msg__GimbalManagerSetAttitude)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.OriginSysid = uint8(mem.origin_sysid)
	m.OriginCompid = uint8(mem.origin_compid)
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
	m.Flags = uint32(mem.flags)
	m.GimbalDeviceId = uint8(mem.gimbal_device_id)
	cSlice_q := mem.q[:]
	primitives.Float32__Array_to_Go(m.Q[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_q)))
	m.AngularVelocityX = float32(mem.angular_velocity_x)
	m.AngularVelocityY = float32(mem.angular_velocity_y)
	m.AngularVelocityZ = float32(mem.angular_velocity_z)
}

func (t _GimbalManagerSetAttitudeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__GimbalManagerSetAttitude())
}

type CGimbalManagerSetAttitude = C.px4_msgs__msg__GimbalManagerSetAttitude
type CGimbalManagerSetAttitude__Sequence = C.px4_msgs__msg__GimbalManagerSetAttitude__Sequence

func GimbalManagerSetAttitude__Sequence_to_Go(goSlice *[]GimbalManagerSetAttitude, cSlice CGimbalManagerSetAttitude__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalManagerSetAttitude, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__GimbalManagerSetAttitude__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerSetAttitude * uintptr(i)),
		))
		GimbalManagerSetAttitudeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GimbalManagerSetAttitude__Sequence_to_C(cSlice *CGimbalManagerSetAttitude__Sequence, goSlice []GimbalManagerSetAttitude) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__GimbalManagerSetAttitude)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__GimbalManagerSetAttitude * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__GimbalManagerSetAttitude)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__GimbalManagerSetAttitude * uintptr(i)),
		))
		GimbalManagerSetAttitudeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GimbalManagerSetAttitude__Array_to_Go(goSlice []GimbalManagerSetAttitude, cSlice []CGimbalManagerSetAttitude) {
	for i := 0; i < len(cSlice); i++ {
		GimbalManagerSetAttitudeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalManagerSetAttitude__Array_to_C(cSlice []CGimbalManagerSetAttitude, goSlice []GimbalManagerSetAttitude) {
	for i := 0; i < len(goSlice); i++ {
		GimbalManagerSetAttitudeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
