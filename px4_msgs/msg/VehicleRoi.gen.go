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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_roi.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleRoi", VehicleRoiTypeSupport)
}
const (
	VehicleRoi_ROI_NONE uint8 = 0// No region of interest
	VehicleRoi_ROI_WPNEXT uint8 = 1// Point toward next MISSION with optional offset
	VehicleRoi_ROI_WPINDEX uint8 = 2// Point toward given MISSION
	VehicleRoi_ROI_LOCATION uint8 = 3// Point toward fixed location
	VehicleRoi_ROI_TARGET uint8 = 4// Point toward target
	VehicleRoi_ROI_ENUM_END uint8 = 5
)

// Do not create instances of this type directly. Always use NewVehicleRoi
// function instead.
type VehicleRoi struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Mode uint8 `yaml:"mode"`// ROI mode (see above)
	Lat float64 `yaml:"lat"`// Latitude to point to
	Lon float64 `yaml:"lon"`// Longitude to point to
	Alt float32 `yaml:"alt"`// Altitude to point to
	RollOffset float32 `yaml:"roll_offset"`// angle offset in rad. additional angle offsets to next waypoint (only used with ROI_WPNEXT)
	PitchOffset float32 `yaml:"pitch_offset"`// angle offset in rad. additional angle offsets to next waypoint (only used with ROI_WPNEXT)
	YawOffset float32 `yaml:"yaw_offset"`// angle offset in rad. additional angle offsets to next waypoint (only used with ROI_WPNEXT)
}

// NewVehicleRoi creates a new VehicleRoi with default values.
func NewVehicleRoi() *VehicleRoi {
	self := VehicleRoi{}
	self.SetDefaults()
	return &self
}

func (t *VehicleRoi) Clone() *VehicleRoi {
	c := &VehicleRoi{}
	c.Timestamp = t.Timestamp
	c.Mode = t.Mode
	c.Lat = t.Lat
	c.Lon = t.Lon
	c.Alt = t.Alt
	c.RollOffset = t.RollOffset
	c.PitchOffset = t.PitchOffset
	c.YawOffset = t.YawOffset
	return c
}

func (t *VehicleRoi) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleRoi) SetDefaults() {
	t.Timestamp = 0
	t.Mode = 0
	t.Lat = 0
	t.Lon = 0
	t.Alt = 0
	t.RollOffset = 0
	t.PitchOffset = 0
	t.YawOffset = 0
}

// CloneVehicleRoiSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleRoiSlice(dst, src []VehicleRoi) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleRoiTypeSupport types.MessageTypeSupport = _VehicleRoiTypeSupport{}

type _VehicleRoiTypeSupport struct{}

func (t _VehicleRoiTypeSupport) New() types.Message {
	return NewVehicleRoi()
}

func (t _VehicleRoiTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleRoi
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleRoi__create())
}

func (t _VehicleRoiTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleRoi__destroy((*C.px4_msgs__msg__VehicleRoi)(pointer_to_free))
}

func (t _VehicleRoiTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleRoi)
	mem := (*C.px4_msgs__msg__VehicleRoi)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.mode = C.uint8_t(m.Mode)
	mem.lat = C.double(m.Lat)
	mem.lon = C.double(m.Lon)
	mem.alt = C.float(m.Alt)
	mem.roll_offset = C.float(m.RollOffset)
	mem.pitch_offset = C.float(m.PitchOffset)
	mem.yaw_offset = C.float(m.YawOffset)
}

func (t _VehicleRoiTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleRoi)
	mem := (*C.px4_msgs__msg__VehicleRoi)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Mode = uint8(mem.mode)
	m.Lat = float64(mem.lat)
	m.Lon = float64(mem.lon)
	m.Alt = float32(mem.alt)
	m.RollOffset = float32(mem.roll_offset)
	m.PitchOffset = float32(mem.pitch_offset)
	m.YawOffset = float32(mem.yaw_offset)
}

func (t _VehicleRoiTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleRoi())
}

type CVehicleRoi = C.px4_msgs__msg__VehicleRoi
type CVehicleRoi__Sequence = C.px4_msgs__msg__VehicleRoi__Sequence

func VehicleRoi__Sequence_to_Go(goSlice *[]VehicleRoi, cSlice CVehicleRoi__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleRoi, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleRoi__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleRoi * uintptr(i)),
		))
		VehicleRoiTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleRoi__Sequence_to_C(cSlice *CVehicleRoi__Sequence, goSlice []VehicleRoi) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleRoi)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleRoi * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleRoi)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleRoi * uintptr(i)),
		))
		VehicleRoiTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleRoi__Array_to_Go(goSlice []VehicleRoi, cSlice []CVehicleRoi) {
	for i := 0; i < len(cSlice); i++ {
		VehicleRoiTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleRoi__Array_to_C(cSlice []CVehicleRoi, goSlice []VehicleRoi) {
	for i := 0; i < len(goSlice); i++ {
		VehicleRoiTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
