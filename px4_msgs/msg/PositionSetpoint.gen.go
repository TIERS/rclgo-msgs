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

#include <px4_msgs/msg/position_setpoint.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/PositionSetpoint", PositionSetpointTypeSupport)
}
const (
	PositionSetpoint_SETPOINT_TYPE_POSITION uint8 = 0// position setpoint
	PositionSetpoint_SETPOINT_TYPE_VELOCITY uint8 = 1// velocity setpoint
	PositionSetpoint_SETPOINT_TYPE_LOITER uint8 = 2// loiter setpoint
	PositionSetpoint_SETPOINT_TYPE_TAKEOFF uint8 = 3// takeoff setpoint
	PositionSetpoint_SETPOINT_TYPE_LAND uint8 = 4// land setpoint, altitude must be ignored, descend until landing
	PositionSetpoint_SETPOINT_TYPE_IDLE uint8 = 5// do nothing, switch off motors or keep at idle speed (MC)
	PositionSetpoint_SETPOINT_TYPE_FOLLOW_TARGET uint8 = 6// setpoint in NED frame (x, y, z, vx, vy, vz) set by follow target
	PositionSetpoint_VELOCITY_FRAME_LOCAL_NED uint8 = 1// MAV_FRAME_LOCAL_NED
	PositionSetpoint_VELOCITY_FRAME_BODY_NED uint8 = 8// MAV_FRAME_BODY_NED
)

// Do not create instances of this type directly. Always use NewPositionSetpoint
// function instead.
type PositionSetpoint struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Valid bool `yaml:"valid"`// true if setpoint is valid
	Type uint8 `yaml:"type"`// setpoint type to adjust behavior of position controller
	Vx float32 `yaml:"vx"`// local velocity setpoint in m/s in NED
	Vy float32 `yaml:"vy"`// local velocity setpoint in m/s in NED
	Vz float32 `yaml:"vz"`// local velocity setpoint in m/s in NED
	VelocityValid bool `yaml:"velocity_valid"`// true if local velocity setpoint valid
	VelocityFrame uint8 `yaml:"velocity_frame"`// to set velocity setpoints in NED or body
	AltValid bool `yaml:"alt_valid"`// do not set for 3D position control. Set to true if you want z-position control while doing vx,vy velocity control.
	Lat float64 `yaml:"lat"`// latitude, in deg
	Lon float64 `yaml:"lon"`// longitude, in deg
	Alt float32 `yaml:"alt"`// altitude AMSL, in m
	Yaw float32 `yaml:"yaw"`// yaw (only for multirotors), in rad [-PI..PI), NaN = hold current yaw
	YawValid bool `yaml:"yaw_valid"`// true if yaw setpoint valid
	Yawspeed float32 `yaml:"yawspeed"`// yawspeed (only for multirotors, in rad/s)
	YawspeedValid bool `yaml:"yawspeed_valid"`// true if yawspeed setpoint valid
	LandingGear int8 `yaml:"landing_gear"`// landing gear: see definition of the states in landing_gear.msg
	LoiterRadius float32 `yaml:"loiter_radius"`// loiter radius (only for fixed wing), in m
	LoiterDirection int8 `yaml:"loiter_direction"`// loiter direction: 1 = CW, -1 = CCW
	AcceptanceRadius float32 `yaml:"acceptance_radius"`// navigation acceptance_radius if we're doing waypoint navigation
	CruisingSpeed float32 `yaml:"cruising_speed"`// the generally desired cruising speed (not a hard constraint)
	CruisingThrottle float32 `yaml:"cruising_throttle"`// the generally desired cruising throttle (not a hard constraint)
	DisableWeatherVane bool `yaml:"disable_weather_vane"`// VTOL: disable (in auto mode) the weather vane feature that turns the nose into the wind
}

// NewPositionSetpoint creates a new PositionSetpoint with default values.
func NewPositionSetpoint() *PositionSetpoint {
	self := PositionSetpoint{}
	self.SetDefaults()
	return &self
}

func (t *PositionSetpoint) Clone() *PositionSetpoint {
	c := &PositionSetpoint{}
	c.Timestamp = t.Timestamp
	c.Valid = t.Valid
	c.Type = t.Type
	c.Vx = t.Vx
	c.Vy = t.Vy
	c.Vz = t.Vz
	c.VelocityValid = t.VelocityValid
	c.VelocityFrame = t.VelocityFrame
	c.AltValid = t.AltValid
	c.Lat = t.Lat
	c.Lon = t.Lon
	c.Alt = t.Alt
	c.Yaw = t.Yaw
	c.YawValid = t.YawValid
	c.Yawspeed = t.Yawspeed
	c.YawspeedValid = t.YawspeedValid
	c.LandingGear = t.LandingGear
	c.LoiterRadius = t.LoiterRadius
	c.LoiterDirection = t.LoiterDirection
	c.AcceptanceRadius = t.AcceptanceRadius
	c.CruisingSpeed = t.CruisingSpeed
	c.CruisingThrottle = t.CruisingThrottle
	c.DisableWeatherVane = t.DisableWeatherVane
	return c
}

func (t *PositionSetpoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PositionSetpoint) SetDefaults() {
	t.Timestamp = 0
	t.Valid = false
	t.Type = 0
	t.Vx = 0
	t.Vy = 0
	t.Vz = 0
	t.VelocityValid = false
	t.VelocityFrame = 0
	t.AltValid = false
	t.Lat = 0
	t.Lon = 0
	t.Alt = 0
	t.Yaw = 0
	t.YawValid = false
	t.Yawspeed = 0
	t.YawspeedValid = false
	t.LandingGear = 0
	t.LoiterRadius = 0
	t.LoiterDirection = 0
	t.AcceptanceRadius = 0
	t.CruisingSpeed = 0
	t.CruisingThrottle = 0
	t.DisableWeatherVane = false
}

// ClonePositionSetpointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePositionSetpointSlice(dst, src []PositionSetpoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PositionSetpointTypeSupport types.MessageTypeSupport = _PositionSetpointTypeSupport{}

type _PositionSetpointTypeSupport struct{}

func (t _PositionSetpointTypeSupport) New() types.Message {
	return NewPositionSetpoint()
}

func (t _PositionSetpointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__PositionSetpoint
	return (unsafe.Pointer)(C.px4_msgs__msg__PositionSetpoint__create())
}

func (t _PositionSetpointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__PositionSetpoint__destroy((*C.px4_msgs__msg__PositionSetpoint)(pointer_to_free))
}

func (t _PositionSetpointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PositionSetpoint)
	mem := (*C.px4_msgs__msg__PositionSetpoint)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.valid = C.bool(m.Valid)
	mem._type = C.uint8_t(m.Type)
	mem.vx = C.float(m.Vx)
	mem.vy = C.float(m.Vy)
	mem.vz = C.float(m.Vz)
	mem.velocity_valid = C.bool(m.VelocityValid)
	mem.velocity_frame = C.uint8_t(m.VelocityFrame)
	mem.alt_valid = C.bool(m.AltValid)
	mem.lat = C.double(m.Lat)
	mem.lon = C.double(m.Lon)
	mem.alt = C.float(m.Alt)
	mem.yaw = C.float(m.Yaw)
	mem.yaw_valid = C.bool(m.YawValid)
	mem.yawspeed = C.float(m.Yawspeed)
	mem.yawspeed_valid = C.bool(m.YawspeedValid)
	mem.landing_gear = C.int8_t(m.LandingGear)
	mem.loiter_radius = C.float(m.LoiterRadius)
	mem.loiter_direction = C.int8_t(m.LoiterDirection)
	mem.acceptance_radius = C.float(m.AcceptanceRadius)
	mem.cruising_speed = C.float(m.CruisingSpeed)
	mem.cruising_throttle = C.float(m.CruisingThrottle)
	mem.disable_weather_vane = C.bool(m.DisableWeatherVane)
}

func (t _PositionSetpointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PositionSetpoint)
	mem := (*C.px4_msgs__msg__PositionSetpoint)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Valid = bool(mem.valid)
	m.Type = uint8(mem._type)
	m.Vx = float32(mem.vx)
	m.Vy = float32(mem.vy)
	m.Vz = float32(mem.vz)
	m.VelocityValid = bool(mem.velocity_valid)
	m.VelocityFrame = uint8(mem.velocity_frame)
	m.AltValid = bool(mem.alt_valid)
	m.Lat = float64(mem.lat)
	m.Lon = float64(mem.lon)
	m.Alt = float32(mem.alt)
	m.Yaw = float32(mem.yaw)
	m.YawValid = bool(mem.yaw_valid)
	m.Yawspeed = float32(mem.yawspeed)
	m.YawspeedValid = bool(mem.yawspeed_valid)
	m.LandingGear = int8(mem.landing_gear)
	m.LoiterRadius = float32(mem.loiter_radius)
	m.LoiterDirection = int8(mem.loiter_direction)
	m.AcceptanceRadius = float32(mem.acceptance_radius)
	m.CruisingSpeed = float32(mem.cruising_speed)
	m.CruisingThrottle = float32(mem.cruising_throttle)
	m.DisableWeatherVane = bool(mem.disable_weather_vane)
}

func (t _PositionSetpointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__PositionSetpoint())
}

type CPositionSetpoint = C.px4_msgs__msg__PositionSetpoint
type CPositionSetpoint__Sequence = C.px4_msgs__msg__PositionSetpoint__Sequence

func PositionSetpoint__Sequence_to_Go(goSlice *[]PositionSetpoint, cSlice CPositionSetpoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PositionSetpoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__PositionSetpoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionSetpoint * uintptr(i)),
		))
		PositionSetpointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PositionSetpoint__Sequence_to_C(cSlice *CPositionSetpoint__Sequence, goSlice []PositionSetpoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__PositionSetpoint)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__PositionSetpoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__PositionSetpoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionSetpoint * uintptr(i)),
		))
		PositionSetpointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PositionSetpoint__Array_to_Go(goSlice []PositionSetpoint, cSlice []CPositionSetpoint) {
	for i := 0; i < len(cSlice); i++ {
		PositionSetpointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PositionSetpoint__Array_to_C(cSlice []CPositionSetpoint, goSlice []PositionSetpoint) {
	for i := 0; i < len(goSlice); i++ {
		PositionSetpointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
