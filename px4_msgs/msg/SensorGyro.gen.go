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

#include <px4_msgs/msg/sensor_gyro.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorGyro", SensorGyroTypeSupport)
}
const (
	SensorGyro_ORB_QUEUE_LENGTH uint8 = 8
)

// Do not create instances of this type directly. Always use NewSensorGyro
// function instead.
type SensorGyro struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	X float32 `yaml:"x"`// angular velocity in the FRD board frame X-axis in rad/s
	Y float32 `yaml:"y"`// angular velocity in the FRD board frame Y-axis in rad/s
	Z float32 `yaml:"z"`// angular velocity in the FRD board frame Z-axis in rad/s
	Temperature float32 `yaml:"temperature"`// temperature in degrees Celsius
	ErrorCount uint32 `yaml:"error_count"`
	Samples uint8 `yaml:"samples"`// number of raw samples that went into this message
}

// NewSensorGyro creates a new SensorGyro with default values.
func NewSensorGyro() *SensorGyro {
	self := SensorGyro{}
	self.SetDefaults()
	return &self
}

func (t *SensorGyro) Clone() *SensorGyro {
	c := &SensorGyro{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.DeviceId = t.DeviceId
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.Temperature = t.Temperature
	c.ErrorCount = t.ErrorCount
	c.Samples = t.Samples
	return c
}

func (t *SensorGyro) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorGyro) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.DeviceId = 0
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.Temperature = 0
	t.ErrorCount = 0
	t.Samples = 0
}

// CloneSensorGyroSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorGyroSlice(dst, src []SensorGyro) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorGyroTypeSupport types.MessageTypeSupport = _SensorGyroTypeSupport{}

type _SensorGyroTypeSupport struct{}

func (t _SensorGyroTypeSupport) New() types.Message {
	return NewSensorGyro()
}

func (t _SensorGyroTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorGyro
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorGyro__create())
}

func (t _SensorGyroTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorGyro__destroy((*C.px4_msgs__msg__SensorGyro)(pointer_to_free))
}

func (t _SensorGyroTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorGyro)
	mem := (*C.px4_msgs__msg__SensorGyro)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.device_id = C.uint32_t(m.DeviceId)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	mem.temperature = C.float(m.Temperature)
	mem.error_count = C.uint32_t(m.ErrorCount)
	mem.samples = C.uint8_t(m.Samples)
}

func (t _SensorGyroTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorGyro)
	mem := (*C.px4_msgs__msg__SensorGyro)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.DeviceId = uint32(mem.device_id)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	m.Temperature = float32(mem.temperature)
	m.ErrorCount = uint32(mem.error_count)
	m.Samples = uint8(mem.samples)
}

func (t _SensorGyroTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorGyro())
}

type CSensorGyro = C.px4_msgs__msg__SensorGyro
type CSensorGyro__Sequence = C.px4_msgs__msg__SensorGyro__Sequence

func SensorGyro__Sequence_to_Go(goSlice *[]SensorGyro, cSlice CSensorGyro__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorGyro, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorGyro__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyro * uintptr(i)),
		))
		SensorGyroTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorGyro__Sequence_to_C(cSlice *CSensorGyro__Sequence, goSlice []SensorGyro) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorGyro)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorGyro * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorGyro)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyro * uintptr(i)),
		))
		SensorGyroTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorGyro__Array_to_Go(goSlice []SensorGyro, cSlice []CSensorGyro) {
	for i := 0; i < len(cSlice); i++ {
		SensorGyroTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorGyro__Array_to_C(cSlice []CSensorGyro, goSlice []SensorGyro) {
	for i := 0; i < len(goSlice); i++ {
		SensorGyroTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
