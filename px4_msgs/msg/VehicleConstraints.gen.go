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

#include <px4_msgs/msg/vehicle_constraints.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleConstraints", VehicleConstraintsTypeSupport)
}

// Do not create instances of this type directly. Always use NewVehicleConstraints
// function instead.
type VehicleConstraints struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SpeedXy float32 `yaml:"speed_xy"`// in meters/sec
	SpeedUp float32 `yaml:"speed_up"`// in meters/sec
	SpeedDown float32 `yaml:"speed_down"`// in meters/sec
	WantTakeoff bool `yaml:"want_takeoff"`// tell the controller to initiate takeoff when idling (ignored during flight)
}

// NewVehicleConstraints creates a new VehicleConstraints with default values.
func NewVehicleConstraints() *VehicleConstraints {
	self := VehicleConstraints{}
	self.SetDefaults()
	return &self
}

func (t *VehicleConstraints) Clone() *VehicleConstraints {
	c := &VehicleConstraints{}
	c.Timestamp = t.Timestamp
	c.SpeedXy = t.SpeedXy
	c.SpeedUp = t.SpeedUp
	c.SpeedDown = t.SpeedDown
	c.WantTakeoff = t.WantTakeoff
	return c
}

func (t *VehicleConstraints) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleConstraints) SetDefaults() {
	t.Timestamp = 0
	t.SpeedXy = 0
	t.SpeedUp = 0
	t.SpeedDown = 0
	t.WantTakeoff = false
}

// CloneVehicleConstraintsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleConstraintsSlice(dst, src []VehicleConstraints) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleConstraintsTypeSupport types.MessageTypeSupport = _VehicleConstraintsTypeSupport{}

type _VehicleConstraintsTypeSupport struct{}

func (t _VehicleConstraintsTypeSupport) New() types.Message {
	return NewVehicleConstraints()
}

func (t _VehicleConstraintsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleConstraints
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleConstraints__create())
}

func (t _VehicleConstraintsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleConstraints__destroy((*C.px4_msgs__msg__VehicleConstraints)(pointer_to_free))
}

func (t _VehicleConstraintsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleConstraints)
	mem := (*C.px4_msgs__msg__VehicleConstraints)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.speed_xy = C.float(m.SpeedXy)
	mem.speed_up = C.float(m.SpeedUp)
	mem.speed_down = C.float(m.SpeedDown)
	mem.want_takeoff = C.bool(m.WantTakeoff)
}

func (t _VehicleConstraintsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleConstraints)
	mem := (*C.px4_msgs__msg__VehicleConstraints)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.SpeedXy = float32(mem.speed_xy)
	m.SpeedUp = float32(mem.speed_up)
	m.SpeedDown = float32(mem.speed_down)
	m.WantTakeoff = bool(mem.want_takeoff)
}

func (t _VehicleConstraintsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleConstraints())
}

type CVehicleConstraints = C.px4_msgs__msg__VehicleConstraints
type CVehicleConstraints__Sequence = C.px4_msgs__msg__VehicleConstraints__Sequence

func VehicleConstraints__Sequence_to_Go(goSlice *[]VehicleConstraints, cSlice CVehicleConstraints__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleConstraints, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleConstraints__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleConstraints * uintptr(i)),
		))
		VehicleConstraintsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleConstraints__Sequence_to_C(cSlice *CVehicleConstraints__Sequence, goSlice []VehicleConstraints) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleConstraints)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleConstraints * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleConstraints)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleConstraints * uintptr(i)),
		))
		VehicleConstraintsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleConstraints__Array_to_Go(goSlice []VehicleConstraints, cSlice []CVehicleConstraints) {
	for i := 0; i < len(cSlice); i++ {
		VehicleConstraintsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleConstraints__Array_to_C(cSlice []CVehicleConstraints, goSlice []VehicleConstraints) {
	for i := 0; i < len(goSlice); i++ {
		VehicleConstraintsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
