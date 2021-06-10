/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package statistics_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstatistics_msgs__rosidl_typesupport_c -lstatistics_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <statistics_msgs/msg/statistic_data_point.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("statistics_msgs/StatisticDataPoint", StatisticDataPointTypeSupport)
}

// Do not create instances of this type directly. Always use NewStatisticDataPoint
// function instead.
type StatisticDataPoint struct {
	DataType uint8 `yaml:"data_type"`// The statistic type of this data point, defined in StatisticDataType.msgDefault value should be StatisticDataType.STATISTICS_DATA_TYPE_UNINITIALIZED (0).
	Data float64 `yaml:"data"`// The value of the data point
}

// NewStatisticDataPoint creates a new StatisticDataPoint with default values.
func NewStatisticDataPoint() *StatisticDataPoint {
	self := StatisticDataPoint{}
	self.SetDefaults()
	return &self
}

func (t *StatisticDataPoint) Clone() *StatisticDataPoint {
	c := &StatisticDataPoint{}
	c.DataType = t.DataType
	c.Data = t.Data
	return c
}

func (t *StatisticDataPoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *StatisticDataPoint) SetDefaults() {
	t.DataType = 0
	t.Data = 0
}

// CloneStatisticDataPointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStatisticDataPointSlice(dst, src []StatisticDataPoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StatisticDataPointTypeSupport types.MessageTypeSupport = _StatisticDataPointTypeSupport{}

type _StatisticDataPointTypeSupport struct{}

func (t _StatisticDataPointTypeSupport) New() types.Message {
	return NewStatisticDataPoint()
}

func (t _StatisticDataPointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.statistics_msgs__msg__StatisticDataPoint
	return (unsafe.Pointer)(C.statistics_msgs__msg__StatisticDataPoint__create())
}

func (t _StatisticDataPointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.statistics_msgs__msg__StatisticDataPoint__destroy((*C.statistics_msgs__msg__StatisticDataPoint)(pointer_to_free))
}

func (t _StatisticDataPointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*StatisticDataPoint)
	mem := (*C.statistics_msgs__msg__StatisticDataPoint)(dst)
	mem.data_type = C.uint8_t(m.DataType)
	mem.data = C.double(m.Data)
}

func (t _StatisticDataPointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*StatisticDataPoint)
	mem := (*C.statistics_msgs__msg__StatisticDataPoint)(ros2_message_buffer)
	m.DataType = uint8(mem.data_type)
	m.Data = float64(mem.data)
}

func (t _StatisticDataPointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__statistics_msgs__msg__StatisticDataPoint())
}

type CStatisticDataPoint = C.statistics_msgs__msg__StatisticDataPoint
type CStatisticDataPoint__Sequence = C.statistics_msgs__msg__StatisticDataPoint__Sequence

func StatisticDataPoint__Sequence_to_Go(goSlice *[]StatisticDataPoint, cSlice CStatisticDataPoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]StatisticDataPoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.statistics_msgs__msg__StatisticDataPoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_statistics_msgs__msg__StatisticDataPoint * uintptr(i)),
		))
		StatisticDataPointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func StatisticDataPoint__Sequence_to_C(cSlice *CStatisticDataPoint__Sequence, goSlice []StatisticDataPoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.statistics_msgs__msg__StatisticDataPoint)(C.malloc((C.size_t)(C.sizeof_struct_statistics_msgs__msg__StatisticDataPoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.statistics_msgs__msg__StatisticDataPoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_statistics_msgs__msg__StatisticDataPoint * uintptr(i)),
		))
		StatisticDataPointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func StatisticDataPoint__Array_to_Go(goSlice []StatisticDataPoint, cSlice []CStatisticDataPoint) {
	for i := 0; i < len(cSlice); i++ {
		StatisticDataPointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func StatisticDataPoint__Array_to_C(cSlice []CStatisticDataPoint, goSlice []StatisticDataPoint) {
	for i := 0; i < len(goSlice); i++ {
		StatisticDataPointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
