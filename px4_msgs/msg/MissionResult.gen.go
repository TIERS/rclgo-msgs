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

#include <px4_msgs/msg/mission_result.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/MissionResult", MissionResultTypeSupport)
}
const (
	MissionResult_MISSION_EXECUTION_MODE_NORMAL uint8 = 0// Execute the mission according to the planned items
	MissionResult_MISSION_EXECUTION_MODE_REVERSE uint8 = 1// Execute the mission in reverse order, ignoring commands and converting all waypoints to normal ones
	MissionResult_MISSION_EXECUTION_MODE_FAST_FORWARD uint8 = 2// Execute the mission as fast as possible, for example converting loiter waypoints to normal ones
)

// Do not create instances of this type directly. Always use NewMissionResult
// function instead.
type MissionResult struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	InstanceCount uint32 `yaml:"instance_count"`// Instance count of this mission. Increments monotonically whenever the mission is modified
	SeqReached int32 `yaml:"seq_reached"`// Sequence of the mission item which has been reached, default -1
	SeqCurrent uint16 `yaml:"seq_current"`// Sequence of the current mission item
	SeqTotal uint16 `yaml:"seq_total"`// Total number of mission items
	Valid bool `yaml:"valid"`// true if mission is valid
	Warning bool `yaml:"warning"`// true if mission is valid, but has potentially problematic items leading to safety warnings
	Finished bool `yaml:"finished"`// true if mission has been completed
	Failure bool `yaml:"failure"`// true if the mission cannot continue or be completed for some reason
	StayInFailsafe bool `yaml:"stay_in_failsafe"`// true if the commander should not switch out of the failsafe mode
	FlightTermination bool `yaml:"flight_termination"`// true if the navigator demands a flight termination from the commander app
	ItemDoJumpChanged bool `yaml:"item_do_jump_changed"`// true if the number of do jumps remaining has changed
	ItemChangedIndex uint16 `yaml:"item_changed_index"`// indicate which item has changed
	ItemDoJumpRemaining uint16 `yaml:"item_do_jump_remaining"`// set to the number of do jumps remaining for that item
	ExecutionMode uint8 `yaml:"execution_mode"`// indicates the mode in which the mission is executed
}

// NewMissionResult creates a new MissionResult with default values.
func NewMissionResult() *MissionResult {
	self := MissionResult{}
	self.SetDefaults()
	return &self
}

func (t *MissionResult) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *MissionResult) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var MissionResultTypeSupport types.MessageTypeSupport = _MissionResultTypeSupport{}

type _MissionResultTypeSupport struct{}

func (t _MissionResultTypeSupport) New() types.Message {
	return NewMissionResult()
}

func (t _MissionResultTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__MissionResult
	return (unsafe.Pointer)(C.px4_msgs__msg__MissionResult__create())
}

func (t _MissionResultTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__MissionResult__destroy((*C.px4_msgs__msg__MissionResult)(pointer_to_free))
}

func (t _MissionResultTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MissionResult)
	mem := (*C.px4_msgs__msg__MissionResult)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.instance_count = C.uint32_t(m.InstanceCount)
	mem.seq_reached = C.int32_t(m.SeqReached)
	mem.seq_current = C.uint16_t(m.SeqCurrent)
	mem.seq_total = C.uint16_t(m.SeqTotal)
	mem.valid = C.bool(m.Valid)
	mem.warning = C.bool(m.Warning)
	mem.finished = C.bool(m.Finished)
	mem.failure = C.bool(m.Failure)
	mem.stay_in_failsafe = C.bool(m.StayInFailsafe)
	mem.flight_termination = C.bool(m.FlightTermination)
	mem.item_do_jump_changed = C.bool(m.ItemDoJumpChanged)
	mem.item_changed_index = C.uint16_t(m.ItemChangedIndex)
	mem.item_do_jump_remaining = C.uint16_t(m.ItemDoJumpRemaining)
	mem.execution_mode = C.uint8_t(m.ExecutionMode)
}

func (t _MissionResultTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MissionResult)
	mem := (*C.px4_msgs__msg__MissionResult)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.InstanceCount = uint32(mem.instance_count)
	m.SeqReached = int32(mem.seq_reached)
	m.SeqCurrent = uint16(mem.seq_current)
	m.SeqTotal = uint16(mem.seq_total)
	m.Valid = bool(mem.valid)
	m.Warning = bool(mem.warning)
	m.Finished = bool(mem.finished)
	m.Failure = bool(mem.failure)
	m.StayInFailsafe = bool(mem.stay_in_failsafe)
	m.FlightTermination = bool(mem.flight_termination)
	m.ItemDoJumpChanged = bool(mem.item_do_jump_changed)
	m.ItemChangedIndex = uint16(mem.item_changed_index)
	m.ItemDoJumpRemaining = uint16(mem.item_do_jump_remaining)
	m.ExecutionMode = uint8(mem.execution_mode)
}

func (t _MissionResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__MissionResult())
}

type CMissionResult = C.px4_msgs__msg__MissionResult
type CMissionResult__Sequence = C.px4_msgs__msg__MissionResult__Sequence

func MissionResult__Sequence_to_Go(goSlice *[]MissionResult, cSlice CMissionResult__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MissionResult, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__MissionResult__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(i)),
		))
		MissionResultTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MissionResult__Sequence_to_C(cSlice *CMissionResult__Sequence, goSlice []MissionResult) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__MissionResult)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__MissionResult)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__MissionResult * uintptr(i)),
		))
		MissionResultTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MissionResult__Array_to_Go(goSlice []MissionResult, cSlice []CMissionResult) {
	for i := 0; i < len(cSlice); i++ {
		MissionResultTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MissionResult__Array_to_C(cSlice []CMissionResult, goSlice []MissionResult) {
	for i := 0; i < len(goSlice); i++ {
		MissionResultTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}