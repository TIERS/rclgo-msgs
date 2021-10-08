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

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstatistics_msgs__rosidl_typesupport_c -lstatistics_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <statistics_msgs/msg/statistic_data_type.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("statistics_msgs/StatisticDataType", StatisticDataTypeTypeSupport)
}
const (
	StatisticDataType_STATISTICS_DATA_TYPE_UNINITIALIZED uint8 = 0// Constant for uninitialized
	StatisticDataType_STATISTICS_DATA_TYPE_AVERAGE uint8 = 1// Allowed values
	StatisticDataType_STATISTICS_DATA_TYPE_MINIMUM uint8 = 2// Allowed values
	StatisticDataType_STATISTICS_DATA_TYPE_MAXIMUM uint8 = 3// Allowed values
	StatisticDataType_STATISTICS_DATA_TYPE_STDDEV uint8 = 4// Allowed values
	StatisticDataType_STATISTICS_DATA_TYPE_SAMPLE_COUNT uint8 = 5// Allowed values
)

// Do not create instances of this type directly. Always use NewStatisticDataType
// function instead.
type StatisticDataType struct {
}

// NewStatisticDataType creates a new StatisticDataType with default values.
func NewStatisticDataType() *StatisticDataType {
	self := StatisticDataType{}
	self.SetDefaults()
	return &self
}

func (t *StatisticDataType) Clone() *StatisticDataType {
	c := &StatisticDataType{}
	return c
}

func (t *StatisticDataType) CloneMsg() types.Message {
	return t.Clone()
}

func (t *StatisticDataType) SetDefaults() {
}

// CloneStatisticDataTypeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStatisticDataTypeSlice(dst, src []StatisticDataType) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StatisticDataTypeTypeSupport types.MessageTypeSupport = _StatisticDataTypeTypeSupport{}

type _StatisticDataTypeTypeSupport struct{}

func (t _StatisticDataTypeTypeSupport) New() types.Message {
	return NewStatisticDataType()
}

func (t _StatisticDataTypeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.statistics_msgs__msg__StatisticDataType
	return (unsafe.Pointer)(C.statistics_msgs__msg__StatisticDataType__create())
}

func (t _StatisticDataTypeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.statistics_msgs__msg__StatisticDataType__destroy((*C.statistics_msgs__msg__StatisticDataType)(pointer_to_free))
}

func (t _StatisticDataTypeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _StatisticDataTypeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _StatisticDataTypeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__statistics_msgs__msg__StatisticDataType())
}

type CStatisticDataType = C.statistics_msgs__msg__StatisticDataType
type CStatisticDataType__Sequence = C.statistics_msgs__msg__StatisticDataType__Sequence

func StatisticDataType__Sequence_to_Go(goSlice *[]StatisticDataType, cSlice CStatisticDataType__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]StatisticDataType, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.statistics_msgs__msg__StatisticDataType__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_statistics_msgs__msg__StatisticDataType * uintptr(i)),
		))
		StatisticDataTypeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func StatisticDataType__Sequence_to_C(cSlice *CStatisticDataType__Sequence, goSlice []StatisticDataType) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.statistics_msgs__msg__StatisticDataType)(C.malloc((C.size_t)(C.sizeof_struct_statistics_msgs__msg__StatisticDataType * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.statistics_msgs__msg__StatisticDataType)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_statistics_msgs__msg__StatisticDataType * uintptr(i)),
		))
		StatisticDataTypeTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func StatisticDataType__Array_to_Go(goSlice []StatisticDataType, cSlice []CStatisticDataType) {
	for i := 0; i < len(cSlice); i++ {
		StatisticDataTypeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func StatisticDataType__Array_to_C(cSlice []CStatisticDataType, goSlice []StatisticDataType) {
	for i := 0; i < len(goSlice); i++ {
		StatisticDataTypeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
