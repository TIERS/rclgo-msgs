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

#include <px4_msgs/msg/rc_parameter_map.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/RcParameterMap", RcParameterMapTypeSupport)
}
const (
	RcParameterMap_RC_PARAM_MAP_NCHAN uint8 = 3// This limit is also hardcoded in the enum RC_CHANNELS_FUNCTION in rc_channels.h
	RcParameterMap_PARAM_ID_LEN uint8 = 16// corresponds to MAVLINK_MSG_PARAM_VALUE_FIELD_PARAM_ID_LEN
)

// Do not create instances of this type directly. Always use NewRcParameterMap
// function instead.
type RcParameterMap struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Valid [3]bool `yaml:"valid"`// true for RC-Param channels which are mapped to a param
	ParamIndex [3]int32 `yaml:"param_index"`// corresponding param index, this field is ignored if set to -1, in this case param_id will be used
	ParamId [51]byte `yaml:"param_id"`// MAP_NCHAN * (ID_LEN + 1) chars, corresponding param id, null terminated
	Scale [3]float32 `yaml:"scale"`// scale to map the RC input [-1, 1] to a parameter value
	Value0 [3]float32 `yaml:"value0"`// initial value around which the parameter value is changed
	ValueMin [3]float32 `yaml:"value_min"`// minimal parameter value
	ValueMax [3]float32 `yaml:"value_max"`// minimal parameter value
}

// NewRcParameterMap creates a new RcParameterMap with default values.
func NewRcParameterMap() *RcParameterMap {
	self := RcParameterMap{}
	self.SetDefaults()
	return &self
}

func (t *RcParameterMap) Clone() *RcParameterMap {
	c := &RcParameterMap{}
	c.Timestamp = t.Timestamp
	c.Valid = t.Valid
	c.ParamIndex = t.ParamIndex
	c.ParamId = t.ParamId
	c.Scale = t.Scale
	c.Value0 = t.Value0
	c.ValueMin = t.ValueMin
	c.ValueMax = t.ValueMax
	return c
}

func (t *RcParameterMap) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RcParameterMap) SetDefaults() {
	t.Timestamp = 0
	t.Valid = [3]bool{}
	t.ParamIndex = [3]int32{}
	t.ParamId = [51]byte{}
	t.Scale = [3]float32{}
	t.Value0 = [3]float32{}
	t.ValueMin = [3]float32{}
	t.ValueMax = [3]float32{}
}

// CloneRcParameterMapSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRcParameterMapSlice(dst, src []RcParameterMap) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RcParameterMapTypeSupport types.MessageTypeSupport = _RcParameterMapTypeSupport{}

type _RcParameterMapTypeSupport struct{}

func (t _RcParameterMapTypeSupport) New() types.Message {
	return NewRcParameterMap()
}

func (t _RcParameterMapTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__RcParameterMap
	return (unsafe.Pointer)(C.px4_msgs__msg__RcParameterMap__create())
}

func (t _RcParameterMapTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__RcParameterMap__destroy((*C.px4_msgs__msg__RcParameterMap)(pointer_to_free))
}

func (t _RcParameterMapTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RcParameterMap)
	mem := (*C.px4_msgs__msg__RcParameterMap)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_valid := mem.valid[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_valid)), m.Valid[:])
	cSlice_param_index := mem.param_index[:]
	primitives.Int32__Array_to_C(*(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_param_index)), m.ParamIndex[:])
	cSlice_param_id := mem.param_id[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_param_id)), m.ParamId[:])
	cSlice_scale := mem.scale[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_scale)), m.Scale[:])
	cSlice_value0 := mem.value0[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value0)), m.Value0[:])
	cSlice_value_min := mem.value_min[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value_min)), m.ValueMin[:])
	cSlice_value_max := mem.value_max[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value_max)), m.ValueMax[:])
}

func (t _RcParameterMapTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RcParameterMap)
	mem := (*C.px4_msgs__msg__RcParameterMap)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_valid := mem.valid[:]
	primitives.Bool__Array_to_Go(m.Valid[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_valid)))
	cSlice_param_index := mem.param_index[:]
	primitives.Int32__Array_to_Go(m.ParamIndex[:], *(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_param_index)))
	cSlice_param_id := mem.param_id[:]
	primitives.Char__Array_to_Go(m.ParamId[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_param_id)))
	cSlice_scale := mem.scale[:]
	primitives.Float32__Array_to_Go(m.Scale[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_scale)))
	cSlice_value0 := mem.value0[:]
	primitives.Float32__Array_to_Go(m.Value0[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value0)))
	cSlice_value_min := mem.value_min[:]
	primitives.Float32__Array_to_Go(m.ValueMin[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value_min)))
	cSlice_value_max := mem.value_max[:]
	primitives.Float32__Array_to_Go(m.ValueMax[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_value_max)))
}

func (t _RcParameterMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__RcParameterMap())
}

type CRcParameterMap = C.px4_msgs__msg__RcParameterMap
type CRcParameterMap__Sequence = C.px4_msgs__msg__RcParameterMap__Sequence

func RcParameterMap__Sequence_to_Go(goSlice *[]RcParameterMap, cSlice CRcParameterMap__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RcParameterMap, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__RcParameterMap__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RcParameterMap * uintptr(i)),
		))
		RcParameterMapTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func RcParameterMap__Sequence_to_C(cSlice *CRcParameterMap__Sequence, goSlice []RcParameterMap) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__RcParameterMap)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__RcParameterMap * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__RcParameterMap)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__RcParameterMap * uintptr(i)),
		))
		RcParameterMapTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func RcParameterMap__Array_to_Go(goSlice []RcParameterMap, cSlice []CRcParameterMap) {
	for i := 0; i < len(cSlice); i++ {
		RcParameterMapTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RcParameterMap__Array_to_C(cSlice []CRcParameterMap, goSlice []RcParameterMap) {
	for i := 0; i < len(goSlice); i++ {
		RcParameterMapTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
