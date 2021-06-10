/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/nav_sat_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/NavSatStatus", NavSatStatusTypeSupport)
}
const (
	NavSatStatus_STATUS_NO_FIX int8 = -1// unable to fix position
	NavSatStatus_STATUS_FIX int8 = 0// unaugmented fix
	NavSatStatus_STATUS_SBAS_FIX int8 = 1// with satellite-based augmentation
	NavSatStatus_STATUS_GBAS_FIX int8 = 2// with ground-based augmentation
	NavSatStatus_SERVICE_GPS uint16 = 1
	NavSatStatus_SERVICE_GLONASS uint16 = 2
	NavSatStatus_SERVICE_COMPASS uint16 = 4// includes BeiDou.
	NavSatStatus_SERVICE_GALILEO uint16 = 8
)

// Do not create instances of this type directly. Always use NewNavSatStatus
// function instead.
type NavSatStatus struct {
	Status int8 `yaml:"status"`
	Service uint16 `yaml:"service"`
}

// NewNavSatStatus creates a new NavSatStatus with default values.
func NewNavSatStatus() *NavSatStatus {
	self := NavSatStatus{}
	self.SetDefaults()
	return &self
}

func (t *NavSatStatus) Clone() *NavSatStatus {
	c := &NavSatStatus{}
	c.Status = t.Status
	c.Service = t.Service
	return c
}

func (t *NavSatStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NavSatStatus) SetDefaults() {
	t.Status = 0
	t.Service = 0
}

// CloneNavSatStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNavSatStatusSlice(dst, src []NavSatStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NavSatStatusTypeSupport types.MessageTypeSupport = _NavSatStatusTypeSupport{}

type _NavSatStatusTypeSupport struct{}

func (t _NavSatStatusTypeSupport) New() types.Message {
	return NewNavSatStatus()
}

func (t _NavSatStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__NavSatStatus
	return (unsafe.Pointer)(C.sensor_msgs__msg__NavSatStatus__create())
}

func (t _NavSatStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__NavSatStatus__destroy((*C.sensor_msgs__msg__NavSatStatus)(pointer_to_free))
}

func (t _NavSatStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NavSatStatus)
	mem := (*C.sensor_msgs__msg__NavSatStatus)(dst)
	mem.status = C.int8_t(m.Status)
	mem.service = C.uint16_t(m.Service)
}

func (t _NavSatStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NavSatStatus)
	mem := (*C.sensor_msgs__msg__NavSatStatus)(ros2_message_buffer)
	m.Status = int8(mem.status)
	m.Service = uint16(mem.service)
}

func (t _NavSatStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__NavSatStatus())
}

type CNavSatStatus = C.sensor_msgs__msg__NavSatStatus
type CNavSatStatus__Sequence = C.sensor_msgs__msg__NavSatStatus__Sequence

func NavSatStatus__Sequence_to_Go(goSlice *[]NavSatStatus, cSlice CNavSatStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NavSatStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__NavSatStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__NavSatStatus * uintptr(i)),
		))
		NavSatStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func NavSatStatus__Sequence_to_C(cSlice *CNavSatStatus__Sequence, goSlice []NavSatStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__NavSatStatus)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__NavSatStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__NavSatStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__NavSatStatus * uintptr(i)),
		))
		NavSatStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func NavSatStatus__Array_to_Go(goSlice []NavSatStatus, cSlice []CNavSatStatus) {
	for i := 0; i < len(cSlice); i++ {
		NavSatStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NavSatStatus__Array_to_C(cSlice []CNavSatStatus, goSlice []NavSatStatus) {
	for i := 0; i < len(goSlice); i++ {
		NavSatStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
