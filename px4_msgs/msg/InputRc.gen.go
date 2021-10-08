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

#include <px4_msgs/msg/input_rc.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/InputRc", InputRcTypeSupport)
}
const (
	InputRc_RC_INPUT_SOURCE_UNKNOWN uint8 = 0
	InputRc_RC_INPUT_SOURCE_PX4FMU_PPM uint8 = 1
	InputRc_RC_INPUT_SOURCE_PX4IO_PPM uint8 = 2
	InputRc_RC_INPUT_SOURCE_PX4IO_SPEKTRUM uint8 = 3
	InputRc_RC_INPUT_SOURCE_PX4IO_SBUS uint8 = 4
	InputRc_RC_INPUT_SOURCE_PX4IO_ST24 uint8 = 5
	InputRc_RC_INPUT_SOURCE_MAVLINK uint8 = 6
	InputRc_RC_INPUT_SOURCE_QURT uint8 = 7
	InputRc_RC_INPUT_SOURCE_PX4FMU_SPEKTRUM uint8 = 8
	InputRc_RC_INPUT_SOURCE_PX4FMU_SBUS uint8 = 9
	InputRc_RC_INPUT_SOURCE_PX4FMU_ST24 uint8 = 10
	InputRc_RC_INPUT_SOURCE_PX4FMU_SUMD uint8 = 11
	InputRc_RC_INPUT_SOURCE_PX4FMU_DSM uint8 = 12
	InputRc_RC_INPUT_SOURCE_PX4IO_SUMD uint8 = 13
	InputRc_RC_INPUT_SOURCE_PX4FMU_CRSF uint8 = 14
	InputRc_RC_INPUT_SOURCE_PX4FMU_GHST uint8 = 15
	InputRc_RC_INPUT_MAX_CHANNELS uint8 = 18// Maximum number of R/C input channels in the system. S.Bus has up to 18 channels.
)

// Do not create instances of this type directly. Always use NewInputRc
// function instead.
type InputRc struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampLastSignal uint64 `yaml:"timestamp_last_signal"`// last valid reception time
	ChannelCount uint8 `yaml:"channel_count"`// number of channels actually being seen
	Rssi int32 `yaml:"rssi"`// receive signal strength indicator (RSSI): < 0: Undefined, 0: no signal, 100: full reception
	RcFailsafe bool `yaml:"rc_failsafe"`// explicit failsafe flag: true on TX failure or TX out of range , false otherwise. Only the true state is reliable, as there are some (PPM) receivers on the market going into failsafe without telling us explicitly.
	RcLost bool `yaml:"rc_lost"`// RC receiver connection status: True,if no frame has arrived in the expected time, false otherwise. True usually means that the receiver has been disconnected, but can also indicate a radio link loss on "stupid" systems. Will remain false, if a RX with failsafe option continues to transmit frames after a link loss.
	RcLostFrameCount uint16 `yaml:"rc_lost_frame_count"`// Number of lost RC frames. Note: intended purpose: observe the radio link quality if RSSI is not available. This value must not be used to trigger any failsafe-alike funtionality.
	RcTotalFrameCount uint16 `yaml:"rc_total_frame_count"`// Number of total RC frames. Note: intended purpose: observe the radio link quality if RSSI is not available. This value must not be used to trigger any failsafe-alike funtionality.
	RcPpmFrameLength uint16 `yaml:"rc_ppm_frame_length"`// Length of a single PPM frame. Zero for non-PPM systems
	InputSource uint8 `yaml:"input_source"`// Input source
	Values [18]uint16 `yaml:"values"`// measured pulse widths for each of the supported channels
}

// NewInputRc creates a new InputRc with default values.
func NewInputRc() *InputRc {
	self := InputRc{}
	self.SetDefaults()
	return &self
}

func (t *InputRc) Clone() *InputRc {
	c := &InputRc{}
	c.Timestamp = t.Timestamp
	c.TimestampLastSignal = t.TimestampLastSignal
	c.ChannelCount = t.ChannelCount
	c.Rssi = t.Rssi
	c.RcFailsafe = t.RcFailsafe
	c.RcLost = t.RcLost
	c.RcLostFrameCount = t.RcLostFrameCount
	c.RcTotalFrameCount = t.RcTotalFrameCount
	c.RcPpmFrameLength = t.RcPpmFrameLength
	c.InputSource = t.InputSource
	c.Values = t.Values
	return c
}

func (t *InputRc) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InputRc) SetDefaults() {
	t.Timestamp = 0
	t.TimestampLastSignal = 0
	t.ChannelCount = 0
	t.Rssi = 0
	t.RcFailsafe = false
	t.RcLost = false
	t.RcLostFrameCount = 0
	t.RcTotalFrameCount = 0
	t.RcPpmFrameLength = 0
	t.InputSource = 0
	t.Values = [18]uint16{}
}

// CloneInputRcSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInputRcSlice(dst, src []InputRc) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InputRcTypeSupport types.MessageTypeSupport = _InputRcTypeSupport{}

type _InputRcTypeSupport struct{}

func (t _InputRcTypeSupport) New() types.Message {
	return NewInputRc()
}

func (t _InputRcTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__InputRc
	return (unsafe.Pointer)(C.px4_msgs__msg__InputRc__create())
}

func (t _InputRcTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__InputRc__destroy((*C.px4_msgs__msg__InputRc)(pointer_to_free))
}

func (t _InputRcTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InputRc)
	mem := (*C.px4_msgs__msg__InputRc)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_last_signal = C.uint64_t(m.TimestampLastSignal)
	mem.channel_count = C.uint8_t(m.ChannelCount)
	mem.rssi = C.int32_t(m.Rssi)
	mem.rc_failsafe = C.bool(m.RcFailsafe)
	mem.rc_lost = C.bool(m.RcLost)
	mem.rc_lost_frame_count = C.uint16_t(m.RcLostFrameCount)
	mem.rc_total_frame_count = C.uint16_t(m.RcTotalFrameCount)
	mem.rc_ppm_frame_length = C.uint16_t(m.RcPpmFrameLength)
	mem.input_source = C.uint8_t(m.InputSource)
	cSlice_values := mem.values[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_values)), m.Values[:])
}

func (t _InputRcTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InputRc)
	mem := (*C.px4_msgs__msg__InputRc)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampLastSignal = uint64(mem.timestamp_last_signal)
	m.ChannelCount = uint8(mem.channel_count)
	m.Rssi = int32(mem.rssi)
	m.RcFailsafe = bool(mem.rc_failsafe)
	m.RcLost = bool(mem.rc_lost)
	m.RcLostFrameCount = uint16(mem.rc_lost_frame_count)
	m.RcTotalFrameCount = uint16(mem.rc_total_frame_count)
	m.RcPpmFrameLength = uint16(mem.rc_ppm_frame_length)
	m.InputSource = uint8(mem.input_source)
	cSlice_values := mem.values[:]
	primitives.Uint16__Array_to_Go(m.Values[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_values)))
}

func (t _InputRcTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__InputRc())
}

type CInputRc = C.px4_msgs__msg__InputRc
type CInputRc__Sequence = C.px4_msgs__msg__InputRc__Sequence

func InputRc__Sequence_to_Go(goSlice *[]InputRc, cSlice CInputRc__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InputRc, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__InputRc__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__InputRc * uintptr(i)),
		))
		InputRcTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InputRc__Sequence_to_C(cSlice *CInputRc__Sequence, goSlice []InputRc) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__InputRc)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__InputRc * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__InputRc)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__InputRc * uintptr(i)),
		))
		InputRcTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InputRc__Array_to_Go(goSlice []InputRc, cSlice []CInputRc) {
	for i := 0; i < len(cSlice); i++ {
		InputRcTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InputRc__Array_to_C(cSlice []CInputRc, goSlice []InputRc) {
	for i := 0; i < len(goSlice); i++ {
		InputRcTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
