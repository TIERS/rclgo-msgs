/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package pendulum_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/TIERS/rclgo-msgs/builtin_interfaces/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpendulum_msgs__rosidl_typesupport_c -lpendulum_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pendulum_msgs/msg/rttest_results.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pendulum_msgs/RttestResults", RttestResultsTypeSupport)
}

// Do not create instances of this type directly. Always use NewRttestResults
// function instead.
type RttestResults struct {
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`
	Command JointCommand `yaml:"command"`
	State JointState `yaml:"state"`
	CurLatency uint64 `yaml:"cur_latency"`
	MeanLatency float64 `yaml:"mean_latency"`
	MinLatency uint64 `yaml:"min_latency"`
	MaxLatency uint64 `yaml:"max_latency"`
	MinorPagefaults uint64 `yaml:"minor_pagefaults"`
	MajorPagefaults uint64 `yaml:"major_pagefaults"`
}

// NewRttestResults creates a new RttestResults with default values.
func NewRttestResults() *RttestResults {
	self := RttestResults{}
	self.SetDefaults()
	return &self
}

func (t *RttestResults) Clone() *RttestResults {
	c := &RttestResults{}
	c.Stamp = *t.Stamp.Clone()
	c.Command = *t.Command.Clone()
	c.State = *t.State.Clone()
	c.CurLatency = t.CurLatency
	c.MeanLatency = t.MeanLatency
	c.MinLatency = t.MinLatency
	c.MaxLatency = t.MaxLatency
	c.MinorPagefaults = t.MinorPagefaults
	c.MajorPagefaults = t.MajorPagefaults
	return c
}

func (t *RttestResults) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RttestResults) SetDefaults() {
	t.Stamp.SetDefaults()
	t.Command.SetDefaults()
	t.State.SetDefaults()
	t.CurLatency = 0
	t.MeanLatency = 0
	t.MinLatency = 0
	t.MaxLatency = 0
	t.MinorPagefaults = 0
	t.MajorPagefaults = 0
}

// CloneRttestResultsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRttestResultsSlice(dst, src []RttestResults) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RttestResultsTypeSupport types.MessageTypeSupport = _RttestResultsTypeSupport{}

type _RttestResultsTypeSupport struct{}

func (t _RttestResultsTypeSupport) New() types.Message {
	return NewRttestResults()
}

func (t _RttestResultsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pendulum_msgs__msg__RttestResults
	return (unsafe.Pointer)(C.pendulum_msgs__msg__RttestResults__create())
}

func (t _RttestResultsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pendulum_msgs__msg__RttestResults__destroy((*C.pendulum_msgs__msg__RttestResults)(pointer_to_free))
}

func (t _RttestResultsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RttestResults)
	mem := (*C.pendulum_msgs__msg__RttestResults)(dst)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
	JointCommandTypeSupport.AsCStruct(unsafe.Pointer(&mem.command), &m.Command)
	JointStateTypeSupport.AsCStruct(unsafe.Pointer(&mem.state), &m.State)
	mem.cur_latency = C.uint64_t(m.CurLatency)
	mem.mean_latency = C.double(m.MeanLatency)
	mem.min_latency = C.uint64_t(m.MinLatency)
	mem.max_latency = C.uint64_t(m.MaxLatency)
	mem.minor_pagefaults = C.uint64_t(m.MinorPagefaults)
	mem.major_pagefaults = C.uint64_t(m.MajorPagefaults)
}

func (t _RttestResultsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RttestResults)
	mem := (*C.pendulum_msgs__msg__RttestResults)(ros2_message_buffer)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
	JointCommandTypeSupport.AsGoStruct(&m.Command, unsafe.Pointer(&mem.command))
	JointStateTypeSupport.AsGoStruct(&m.State, unsafe.Pointer(&mem.state))
	m.CurLatency = uint64(mem.cur_latency)
	m.MeanLatency = float64(mem.mean_latency)
	m.MinLatency = uint64(mem.min_latency)
	m.MaxLatency = uint64(mem.max_latency)
	m.MinorPagefaults = uint64(mem.minor_pagefaults)
	m.MajorPagefaults = uint64(mem.major_pagefaults)
}

func (t _RttestResultsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pendulum_msgs__msg__RttestResults())
}

type CRttestResults = C.pendulum_msgs__msg__RttestResults
type CRttestResults__Sequence = C.pendulum_msgs__msg__RttestResults__Sequence

func RttestResults__Sequence_to_Go(goSlice *[]RttestResults, cSlice CRttestResults__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RttestResults, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.pendulum_msgs__msg__RttestResults__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(i)),
		))
		RttestResultsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func RttestResults__Sequence_to_C(cSlice *CRttestResults__Sequence, goSlice []RttestResults) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.pendulum_msgs__msg__RttestResults)(C.malloc((C.size_t)(C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.pendulum_msgs__msg__RttestResults)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_pendulum_msgs__msg__RttestResults * uintptr(i)),
		))
		RttestResultsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func RttestResults__Array_to_Go(goSlice []RttestResults, cSlice []CRttestResults) {
	for i := 0; i < len(cSlice); i++ {
		RttestResultsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RttestResults__Array_to_C(cSlice []CRttestResults, goSlice []RttestResults) {
	for i := 0; i < len(goSlice); i++ {
		RttestResultsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
