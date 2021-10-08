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

#include <px4_msgs/msg/sensor_gyro_fft.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorGyroFft", SensorGyroFftTypeSupport)
}

// Do not create instances of this type directly. Always use NewSensorGyroFft
// function instead.
type SensorGyroFft struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`
	DeviceId uint32 `yaml:"device_id"`// unique device ID for the sensor that does not change between power cycles
	SensorSampleRateHz float32 `yaml:"sensor_sample_rate_hz"`
	ResolutionHz float32 `yaml:"resolution_hz"`
	PeakFrequenciesX [6]float32 `yaml:"peak_frequencies_x"`// x axis peak frequencies
	PeakFrequenciesY [6]float32 `yaml:"peak_frequencies_y"`// y axis peak frequencies
	PeakFrequenciesZ [6]float32 `yaml:"peak_frequencies_z"`// z axis peak frequencies
	PeakMagnitudeX [6]uint32 `yaml:"peak_magnitude_x"`// x axis peak frequencies magnitude
	PeakMagnitudeY [6]uint32 `yaml:"peak_magnitude_y"`// y axis peak frequencies magnitude
	PeakMagnitudeZ [6]uint32 `yaml:"peak_magnitude_z"`// z axis peak frequencies magnitude
}

// NewSensorGyroFft creates a new SensorGyroFft with default values.
func NewSensorGyroFft() *SensorGyroFft {
	self := SensorGyroFft{}
	self.SetDefaults()
	return &self
}

func (t *SensorGyroFft) Clone() *SensorGyroFft {
	c := &SensorGyroFft{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.DeviceId = t.DeviceId
	c.SensorSampleRateHz = t.SensorSampleRateHz
	c.ResolutionHz = t.ResolutionHz
	c.PeakFrequenciesX = t.PeakFrequenciesX
	c.PeakFrequenciesY = t.PeakFrequenciesY
	c.PeakFrequenciesZ = t.PeakFrequenciesZ
	c.PeakMagnitudeX = t.PeakMagnitudeX
	c.PeakMagnitudeY = t.PeakMagnitudeY
	c.PeakMagnitudeZ = t.PeakMagnitudeZ
	return c
}

func (t *SensorGyroFft) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorGyroFft) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.DeviceId = 0
	t.SensorSampleRateHz = 0
	t.ResolutionHz = 0
	t.PeakFrequenciesX = [6]float32{}
	t.PeakFrequenciesY = [6]float32{}
	t.PeakFrequenciesZ = [6]float32{}
	t.PeakMagnitudeX = [6]uint32{}
	t.PeakMagnitudeY = [6]uint32{}
	t.PeakMagnitudeZ = [6]uint32{}
}

// CloneSensorGyroFftSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorGyroFftSlice(dst, src []SensorGyroFft) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorGyroFftTypeSupport types.MessageTypeSupport = _SensorGyroFftTypeSupport{}

type _SensorGyroFftTypeSupport struct{}

func (t _SensorGyroFftTypeSupport) New() types.Message {
	return NewSensorGyroFft()
}

func (t _SensorGyroFftTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorGyroFft
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorGyroFft__create())
}

func (t _SensorGyroFftTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorGyroFft__destroy((*C.px4_msgs__msg__SensorGyroFft)(pointer_to_free))
}

func (t _SensorGyroFftTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorGyroFft)
	mem := (*C.px4_msgs__msg__SensorGyroFft)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.device_id = C.uint32_t(m.DeviceId)
	mem.sensor_sample_rate_hz = C.float(m.SensorSampleRateHz)
	mem.resolution_hz = C.float(m.ResolutionHz)
	cSlice_peak_frequencies_x := mem.peak_frequencies_x[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_x)), m.PeakFrequenciesX[:])
	cSlice_peak_frequencies_y := mem.peak_frequencies_y[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_y)), m.PeakFrequenciesY[:])
	cSlice_peak_frequencies_z := mem.peak_frequencies_z[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_z)), m.PeakFrequenciesZ[:])
	cSlice_peak_magnitude_x := mem.peak_magnitude_x[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_x)), m.PeakMagnitudeX[:])
	cSlice_peak_magnitude_y := mem.peak_magnitude_y[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_y)), m.PeakMagnitudeY[:])
	cSlice_peak_magnitude_z := mem.peak_magnitude_z[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_z)), m.PeakMagnitudeZ[:])
}

func (t _SensorGyroFftTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorGyroFft)
	mem := (*C.px4_msgs__msg__SensorGyroFft)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.DeviceId = uint32(mem.device_id)
	m.SensorSampleRateHz = float32(mem.sensor_sample_rate_hz)
	m.ResolutionHz = float32(mem.resolution_hz)
	cSlice_peak_frequencies_x := mem.peak_frequencies_x[:]
	primitives.Float32__Array_to_Go(m.PeakFrequenciesX[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_x)))
	cSlice_peak_frequencies_y := mem.peak_frequencies_y[:]
	primitives.Float32__Array_to_Go(m.PeakFrequenciesY[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_y)))
	cSlice_peak_frequencies_z := mem.peak_frequencies_z[:]
	primitives.Float32__Array_to_Go(m.PeakFrequenciesZ[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_peak_frequencies_z)))
	cSlice_peak_magnitude_x := mem.peak_magnitude_x[:]
	primitives.Uint32__Array_to_Go(m.PeakMagnitudeX[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_x)))
	cSlice_peak_magnitude_y := mem.peak_magnitude_y[:]
	primitives.Uint32__Array_to_Go(m.PeakMagnitudeY[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_y)))
	cSlice_peak_magnitude_z := mem.peak_magnitude_z[:]
	primitives.Uint32__Array_to_Go(m.PeakMagnitudeZ[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_peak_magnitude_z)))
}

func (t _SensorGyroFftTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorGyroFft())
}

type CSensorGyroFft = C.px4_msgs__msg__SensorGyroFft
type CSensorGyroFft__Sequence = C.px4_msgs__msg__SensorGyroFft__Sequence

func SensorGyroFft__Sequence_to_Go(goSlice *[]SensorGyroFft, cSlice CSensorGyroFft__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorGyroFft, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorGyroFft__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(i)),
		))
		SensorGyroFftTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorGyroFft__Sequence_to_C(cSlice *CSensorGyroFft__Sequence, goSlice []SensorGyroFft) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorGyroFft)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorGyroFft)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorGyroFft * uintptr(i)),
		))
		SensorGyroFftTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorGyroFft__Array_to_Go(goSlice []SensorGyroFft, cSlice []CSensorGyroFft) {
	for i := 0; i < len(cSlice); i++ {
		SensorGyroFftTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorGyroFft__Array_to_C(cSlice []CSensorGyroFft, goSlice []SensorGyroFft) {
	for i := 0; i < len(goSlice); i++ {
		SensorGyroFftTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
