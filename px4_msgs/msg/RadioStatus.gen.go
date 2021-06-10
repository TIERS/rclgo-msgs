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

#include <px4_msgs/msg/radio_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/RadioStatus", RadioStatusTypeSupport)
}

// Do not create instances of this type directly. Always use NewRadioStatus
// function instead.
type RadioStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Rssi uint8 `yaml:"rssi"`// local signal strength
	RemoteRssi uint8 `yaml:"remote_rssi"`// remote signal strength
	Txbuf uint8 `yaml:"txbuf"`// how full the tx buffer is as a percentage
	Noise uint8 `yaml:"noise"`// background noise level
	RemoteNoise uint8 `yaml:"remote_noise"`// remote background noise level
	Rxerrors uint16 `yaml:"rxerrors"`// receive errors
	Fix uint16 `yaml:"fix"`// count of error corrected packets
}

// NewRadioStatus creates a new RadioStatus with default values.
func NewRadioStatus() *RadioStatus {
	self := RadioStatus{}
	self.SetDefaults()
	return &self
}

func (t *RadioStatus) Clone() *RadioStatus {
	c := &RadioStatus{}
	c.Timestamp = t.Timestamp
	c.Rssi = t.Rssi
	c.RemoteRssi = t.RemoteRssi
	c.Txbuf = t.Txbuf
	c.Noise = t.Noise
	c.RemoteNoise = t.RemoteNoise
	c.Rxerrors = t.Rxerrors
	c.Fix = t.Fix
	return c
}

func (t *RadioStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RadioStatus) SetDefaults() {
	t.Timestamp = 0
	t.Rssi = 0
	t.RemoteRssi = 0
	t.Txbuf = 0
	t.Noise = 0
	t.RemoteNoise = 0
	t.Rxerrors = 0
	t.Fix = 0
}

// CloneRadioStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRadioStatusSlice(dst, src []RadioStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RadioStatusTypeSupport types.MessageTypeSupport = _RadioStatusTypeSupport{}

type _RadioStatusTypeSupport struct{}

func (t _RadioStatusTypeSupport) New() types.Message {
	return NewRadioStatus()
}

func (t _RadioStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__RadioStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__RadioStatus__create())
}

func (t _RadioStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__RadioStatus__destroy((*C.px4_msgs__msg__RadioStatus)(pointer_to_free))
}

func (t _RadioStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RadioStatus)
	mem := (*C.px4_msgs__msg__RadioStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.rssi = C.uint8_t(m.Rssi)
	mem.remote_rssi = C.uint8_t(m.RemoteRssi)
	mem.txbuf = C.uint8_t(m.Txbuf)
	mem.noise = C.uint8_t(m.Noise)
	mem.remote_noise = C.uint8_t(m.RemoteNoise)
	mem.rxerrors = C.uint16_t(m.Rxerrors)
	mem.fix = C.uint16_t(m.Fix)
}

func (t _RadioStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RadioStatus)
	mem := (*C.px4_msgs__msg__RadioStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Rssi = uint8(mem.rssi)
	m.RemoteRssi = uint8(mem.remote_rssi)
	m.Txbuf = uint8(mem.txbuf)
	m.Noise = uint8(mem.noise)
	m.RemoteNoise = uint8(mem.remote_noise)
	m.Rxerrors = uint16(mem.rxerrors)
	m.Fix = uint16(mem.fix)
}

func (t _RadioStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__RadioStatus())
}

type CRadioStatus = C.px4_msgs__msg__RadioStatus
type CRadioStatus__Sequence = C.px4_msgs__msg__RadioStatus__Sequence

func RadioStatus__Sequence_to_Go(goSlice *[]RadioStatus, cSlice CRadioStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RadioStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__RadioStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RadioStatus * uintptr(i)),
		))
		RadioStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func RadioStatus__Sequence_to_C(cSlice *CRadioStatus__Sequence, goSlice []RadioStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__RadioStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__RadioStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__RadioStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RadioStatus * uintptr(i)),
		))
		RadioStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func RadioStatus__Array_to_Go(goSlice []RadioStatus, cSlice []CRadioStatus) {
	for i := 0; i < len(cSlice); i++ {
		RadioStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RadioStatus__Array_to_C(cSlice []CRadioStatus, goSlice []RadioStatus) {
	for i := 0; i < len(goSlice); i++ {
		RadioStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
